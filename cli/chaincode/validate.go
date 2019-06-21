package main

import (
	"errors"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func validatePublisherState(stub shim.ChaincodeStubInterface, id string) ([]byte, error) {
	key, err := getPublisherKey(stub, id)
	if err != nil {
		return nil, err
	}

	publisherAsBytes, err := stub.GetState(key)
	if err != nil {
		return nil, err
	}
	if publisherAsBytes == nil {
		return nil, errors.New(fmt.Sprintf("TRIT_ERROR_33"))
	}

	return publisherAsBytes, nil
}

func validateShareholderState(stub shim.ChaincodeStubInterface, id string) ([]byte, error) {
	key, err := getShareholderKey(stub, id)
	if err != nil {
		return nil, err
	}

	shareholderAsBytes, err := stub.GetState(key)
	if err != nil {
		return nil, err
	}
	if shareholderAsBytes == nil {
		return nil, errors.New(fmt.Sprintf("TRIT_ERROR_34"))
	}

	return shareholderAsBytes, nil
}

func validateRealEstateState(stub shim.ChaincodeStubInterface, id string) ([]byte, error) {
	key, err := getRealEstateKey(stub, id)
	if err != nil {
		return nil, err
	}

	realEstateAsBytes, err := stub.GetState(key)
	if err != nil {
		return nil, err
	}
	if realEstateAsBytes == nil {
		return nil, errors.New(fmt.Sprintf("TRIT_ERROR_35"))
	}

	return realEstateAsBytes, nil
}

func validateSellTritAdvertisingContractState(stub shim.ChaincodeStubInterface, id string) ([]byte, error) {
	key, err := getSellTritAdvertisingContractKey(stub, id)
	if err != nil {
		return nil, err
	}

	contractAsBytes, err := stub.GetState(key)
	if err != nil {
		return nil, err
	}
	if contractAsBytes == nil {
		return nil, errors.New(fmt.Sprintf("TRIT_ERROR_36"))
	}

	return contractAsBytes, nil
}

func validateRealEstateId(stub shim.ChaincodeStubInterface, id string) error {
	realEstateKey, err := getRealEstateKey(stub, id)
	if err != nil {
		return err
	}

	realEstateAsBytes, err := stub.GetState(realEstateKey)
	if realEstateAsBytes != nil {
		return errors.New("TRIT_ERROR_37")
	}
	return nil
}

func validateSellTritAdvertisingContractId(stub shim.ChaincodeStubInterface, id string) error {
	contractKey, err := getSellTritAdvertisingContractKey(stub, id)
	if err != nil {
		return err
	}

	contractAsBytes, err := stub.GetState(contractKey)
	if err != nil {
		return err
	}
	if contractAsBytes != nil {
		return errors.New(fmt.Sprintf("TRIT_ERROR_38"))
	}

	return nil
}

func validateTransferContractId(stub shim.ChaincodeStubInterface, id string) error {
	contractKey, err := getTransferContractKey(stub, id)
	if err != nil {
		return err
	}

	contractAsBytes, err := stub.GetState(contractKey)
	if err != nil {
		return err
	}
	if contractAsBytes != nil {
		return errors.New(fmt.Sprintf("TRIT_ERROR_39"))
	}

	return nil
}

func validateChangePriceContractId(stub shim.ChaincodeStubInterface, id string) error {
	contractKey, err := getChangePriceContractKey(stub, id)
	if err != nil {
		return err
	}

	contractAsBytes, err := stub.GetState(contractKey)
	if err != nil {
		return err
	}

	if contractAsBytes != nil {
		return errors.New(fmt.Sprintf("TRIT_ERORR_40"))
	}

	return nil
}

func validatePublishContractId(stub shim.ChaincodeStubInterface, id string) error {
	contractKey, err := getPublishContractKey(stub, id)
	if err != nil {
		return err
	}

	contractAsBytes, err := stub.GetState(contractKey)
	if err != nil {
		return err
	}
	if contractAsBytes != nil {
		return errors.New(fmt.Sprintf("TRIT_ERROR_41"))
	}

	return nil
}

func validateShareholderId(stub shim.ChaincodeStubInterface, id string) error {
	key, err := getShareholderKey(stub, id)
	if err != nil {
		return err
	}

	userAsBytes, err := stub.GetState(key)
	if userAsBytes != nil {
		return errors.New("TRIT_ERROR_42")
	}
	return nil
}

func validatePublisherId(stub shim.ChaincodeStubInterface, id string) error {
	key, err := getPublisherKey(stub, id)
	if err != nil {
		return err
	}

	userAsBytes, err := stub.GetState(key)
	if userAsBytes != nil {
		return errors.New("TRIT_ERROR_43")
	}
	return nil
}
