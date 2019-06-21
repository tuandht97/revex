package main

import (
	"errors"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"strconv"
)

func deleteHT (APIstub shim.ChaincodeStubInterface, args []string) error {

	// Retrieve the variable name
	name := args[0]

	// Delete all delta rows
	deltaResultsIterator, deltaErr := APIstub.GetStateByPartialCompositeKey("varName~op~value~txID", []string{name})
	if deltaErr != nil {
		return errors.New(fmt.Sprintf("Could not retrieve delta rows for %s: %s", name, deltaErr.Error()))
	}
	defer deltaResultsIterator.Close()

	// Ensure the variable exists
	if !deltaResultsIterator.HasNext() {
		return errors.New(fmt.Sprintf("No variable by the name %s exists", name))
	}

	// Iterate through result set and delete all indices
	var i int
	for i = 0; deltaResultsIterator.HasNext(); i++ {
		responseRange, nextErr := deltaResultsIterator.Next()
		if nextErr != nil {
			return errors.New(fmt.Sprintf("Could not retrieve next delta row: %s", nextErr.Error()))
		}

		deltaRowDelErr := APIstub.DelState(responseRange.Key)
		if deltaRowDelErr != nil {
			return errors.New(fmt.Sprintf("Could not delete delta row: %s", deltaRowDelErr.Error()))
		}
	}

	return nil
}

func getHT(APIstub shim.ChaincodeStubInterface, args []string)(int, error) {
	// Check we have a valid number of args

	name := args[0]
	startValue, err := strconv.Atoi(args[1])

	if err != nil {
		return 0, err
	}
	// Get all deltas for the variable
	deltaResultsIterator, deltaErr := APIstub.GetStateByPartialCompositeKey("varName~op~value~txID", []string{name})
	if deltaErr != nil {
		return -1, errors.New(fmt.Sprintf("Could not retrieve value for %s: %s", name, deltaErr.Error()))
	}
	defer deltaResultsIterator.Close()

	// Check the variable existed
	if !deltaResultsIterator.HasNext() {
		return startValue, nil
	}

	// Iterate through result set and compute final value
	var finalVal int
	finalVal += startValue

	var i int64
	for i = 0; deltaResultsIterator.HasNext(); i++ {
		// Get the next row
		responseRange, nextErr := deltaResultsIterator.Next()
		if nextErr != nil {
			return -1, errors.New(nextErr.Error())
		}

		// Split the composite key into its component parts
		_, keyParts, splitKeyErr := APIstub.SplitCompositeKey(responseRange.Key)
		if splitKeyErr != nil {
			return -1, errors.New(splitKeyErr.Error())
		}

		// Retrieve the delta value and operation
		operation := keyParts[1]
		valueStr := keyParts[2]

		// Convert the value string and perform the operation
		value, convErr := strconv.Atoi(valueStr)
		if convErr != nil {
			return -1, errors.New(convErr.Error())
		}

		switch operation {
		case "+":
			finalVal += value
		case "-":
			finalVal -= value
		default:
			return 1, errors.New(fmt.Sprintf("Unrecognized operation %s", operation))
		}
	}

	return finalVal, nil
}

func updateHT(APIstub shim.ChaincodeStubInterface, args []string) error {
	// Check we have a valid number of args


	// Extract the args
	name := args[0]
	op := args[2]
	_, err := strconv.Atoi(args[1])
	if err != nil {
		return errors.New("Provided value was not a number" + err.Error())
	}

	// Make sure a valid operator is provided
	if op != "+" && op != "-" {
		return errors.New(fmt.Sprintf("Operator %s is unrecognized", op))
	}

	// Retrieve info needed for the update procedure
	txid := APIstub.GetTxID()
	compositeIndexName := "varName~op~value~txID"

	// Create the composite key that will allow us to query for all deltas on a particular variable
	compositeKey, compositeErr := APIstub.CreateCompositeKey(compositeIndexName, []string{name, op, args[1], txid})
	if compositeErr != nil {
		return errors.New(fmt.Sprintf("Could not create a composite key for %s: %s", name, compositeErr.Error()))
	}

	// Save the composite key index
	compositePutErr := APIstub.PutState(compositeKey, []byte{0x00})
	if compositePutErr != nil {
		return errors.New(fmt.Sprintf("Could not put operation for %s in the ledger: %s", name, compositePutErr.Error()))
	}

	return nil
}
