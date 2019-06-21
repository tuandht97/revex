package main

import (
	"encoding/json"
	"errors"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func (p *publisher) setTritIdToPublisherPublishedTrits(tritId string) {

	if tritId == "" {
		return
	}

	if len(p.PublishedTrits) == 0 {
		p.PublishedTrits = []string{}
	}

	p.PublishedTrits = append(p.PublishedTrits, tritId)

}

func getPublisherState(stub shim.ChaincodeStubInterface, id string) (publisher, error) {
	if len(id) == 0 {
		return publisher{}, errors.New("TRIT_ERROR_8")
	}

	userKey, err := getPublisherKey(stub, id)
	if err != nil {
		return publisher{}, err
	}
	userBytes, _ := stub.GetState(userKey)
	if len(userBytes) == 0 {
		return publisher{}, errors.New("TRIT_ERROR_27")
	}

	response := publisher{}
	err = json.Unmarshal(userBytes, &response)
	return response, err
}

func (p publisher) savePublisherState(stub shim.ChaincodeStubInterface) error {

	userKey, err := getPublisherKey(stub, p.Username)
	if err != nil {
		return err
	}
	userAsBytes, err := json.Marshal(p)
	if err != nil {
		return err
	}
	err = stub.PutState(userKey, userAsBytes)
	return err
}
