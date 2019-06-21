package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

type revexCCQ struct {
	ObjectType string `json:"docType"`   // field for couchdb
	Symbol     string `json:"symbol"`    // mã
	Total      int    `json:"total"`     // tổng số mã
	Price      int    `json:"price"`     // giá bán
	TranType   int    `json:"tranType"`  // loại giao dịch, 0: Mua - 1: Bán
	Timestamp  int    `json:"timestamp"` // ngày giao dịch dạng Timestamp
	UserA      string `json:"userA"`     // nhà đầu tư
	UserB      string `json:"userB"`     // sàn
	Org        string `json:"org"`       // tổ chức CCQ
}

type revexBDS struct {
	ObjectType     string `json:"docType"`        // field for couchdb
	Symbol         string `json:"symbol"`         // mã
	Price          int    `json:"price"`          // giá bán
	Address        string `json:"address"`        // địa chỉ
	Area           int    `json:"area"`           // diện tích
	NumberFloor    int    `json:"numberFloor"`    // số tầng
	NumberBedroom  int    `json:"numberBedroom"`  // số phòng ngủ
	NumberRestroom int    `json:"numberRestroom"` // số phòng vệ sinh
	TranType       int    `json:"tranType"`       // loại giao dịch, 0: Mua - 1: Bán
	Timestamp      int    `json:"timestamp"`      // ngày giao dịch dạng Timestamp
	UserA          string `json:"userA"`          // nhà đầu tư
	UserB          string `json:"userB"`          // sàn
	Org            string `json:"org"`            // tổ chức BDS
}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()
	fmt.Println("invoke is running " + function)

	if function == "queryRevexs" {
		return t.queryRevexs(stub, args)
	} else if function == "getAll" {
		return t.getAll(stub)
	} else if function == "importBDS" {
		return t.importBDS(stub, args)
	} else if function == "importCCQ" {
		return t.importCCQ(stub, args)
	}

	fmt.Println("invoke did not find func: " + function) //error
	return shim.Error("Received unknown function invocation " + function)
}

func constructQueryResponseFromIterator(resultsIterator shim.StateQueryIteratorInterface) (*bytes.Buffer, error) {
	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString(string(queryResponse.Value))
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	return &buffer, nil
}

func (t *SimpleChaincode) queryRevexs(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	//   0
	// "queryString"
	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	queryString := args[0]

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)
}

func getQueryResultForQueryString(stub shim.ChaincodeStubInterface, queryString string) ([]byte, error) {

	fmt.Printf("- getQueryResultForQueryString queryString:\n%s\n", queryString)

	resultsIterator, err := stub.GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	buffer, err := constructQueryResponseFromIterator(resultsIterator)
	if err != nil {
		return nil, err
	}

	fmt.Printf("- getQueryResultForQueryString queryResult:\n%s\n", buffer.String())

	return buffer.Bytes(), nil
}

func (t *SimpleChaincode) getAll(stub shim.ChaincodeStubInterface) pb.Response {

	resultsIterator, err := stub.GetStateByRange("", "")
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	buffer, err := constructQueryResponseFromIterator(resultsIterator)
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Printf("- getMarblesByRange queryResult:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

func (t *SimpleChaincode) importBDS(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var revexs []revexBDS
	err := json.Unmarshal([]byte(args[0]), &revexs)
	if err != nil {
		panic(err)
		return shim.Error(err.Error())
	}

	fmt.Print(revexs)

	i := 0
	for i < len(revexs) {
		revexs[i].ObjectType = "revexBDS"
		revexs[i].Org = "BDS"
		revexKey := strings.Join([]string{revexs[i].Symbol, revexs[i].UserA, revexs[i].UserB, strconv.Itoa(revexs[i].TranType), strconv.Itoa(revexs[i].Timestamp), revexs[i].Org}, "-")
		revexJSONasBytes, _ := json.Marshal(revexs[i])
		stub.PutState(revexKey, revexJSONasBytes)
		i = i + 1
	}

	return shim.Success(nil)
}

func (t *SimpleChaincode) importCCQ(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var revexs []revexCCQ
	err := json.Unmarshal([]byte(args[0]), &revexs)
	if err != nil {
		panic(err)
		return shim.Error(err.Error())
	}

	i := 0
	for i < len(revexs) {
		revexs[i].ObjectType = "revexCCQ"
		revexs[i].Org = "CCQ"
		revexKey := strings.Join([]string{revexs[i].Symbol, revexs[i].UserA, revexs[i].UserB, strconv.Itoa(revexs[i].TranType), strconv.Itoa(revexs[i].Timestamp), revexs[i].Org}, "-")
		revexJSONasBytes, _ := json.Marshal(revexs[i])
		stub.PutState(revexKey, revexJSONasBytes)
		i = i + 1
	}

	return shim.Success(nil)
}
