package main

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"strconv"
)

func listTransferContract(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	prefix := prefixTransferContract
	resultsIterator, err := stub.GetStateByPartialCompositeKey(prefix, []string{})
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	var results []interface{}
	for resultsIterator.HasNext() {
		kvResult, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}

		c := transferContract{}
		err = json.Unmarshal(kvResult.Value, &c)
		if err != nil {
			return shim.Error(err.Error())
		}

		results = append(results, c)
	}

	resultsAsBytes, err := json.Marshal(results)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(resultsAsBytes)
}

func listTransferContractByBuyer(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("TRIT_ERROR_3")
	}

	dto := struct {
		Username string `json:"username"`
	}{}

	err := json.Unmarshal([]byte(args[0]), &dto)
	if err != nil {
		return shim.Error(err.Error())
	}

	prefix := prefixTransferContract
	resultsIterator, err := stub.GetStateByPartialCompositeKey(prefix, []string{})
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	var results []interface{}
	for resultsIterator.HasNext() {
		kvResult, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}

		c := transferContract{}
		err = json.Unmarshal(kvResult.Value, &c)
		if err != nil {
			return shim.Error(err.Error())
		}

		if c.Buyer == dto.Username {
			results = append(results, c)
		}
	}

	resultsAsBytes, err := json.Marshal(results)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(resultsAsBytes)
}

func listTransferContractBySeller(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("TRIT_ERROR_3")
	}

	dto := struct {
		Username string `json:"username"`
	}{}

	err := json.Unmarshal([]byte(args[0]), &dto)
	if err != nil {
		return shim.Error(err.Error())
	}

	prefix := prefixTransferContract
	resultsIterator, err := stub.GetStateByPartialCompositeKey(prefix, []string{})
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	var results []interface{}
	for resultsIterator.HasNext() {
		kvResult, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}

		c := transferContract{}
		err = json.Unmarshal(kvResult.Value, &c)
		if err != nil {
			return shim.Error(err.Error())
		}

		if c.Seller == dto.Username {
			results = append(results, c)
		}
	}

	resultsAsBytes, err := json.Marshal(results)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(resultsAsBytes)
}

func listChangePriceContract(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	prefix := prefixChangePriceContract
	resultsIterator, err := stub.GetStateByPartialCompositeKey(prefix, []string{})
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	var results []interface{}
	for resultsIterator.HasNext() {
		kvResult, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}

		c := changePriceContract{}
		err = json.Unmarshal(kvResult.Value, &c)
		if err != nil {
			return shim.Error(err.Error())
		}

		results = append(results, c)
	}

	resultsAsBytes, err := json.Marshal(results)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(resultsAsBytes)
}

func listPublishContract(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	prefix := prefixPublishContract
	resultsIterator, err := stub.GetStateByPartialCompositeKey(prefix, []string{})
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	var results []interface{}
	for resultsIterator.HasNext() {
		kvResult, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}

		c := publishContract{}
		err = json.Unmarshal(kvResult.Value, &c)
		if err != nil {
			return shim.Error(err.Error())
		}

		results = append(results, c)
	}

	resultsAsBytes, err := json.Marshal(results)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(resultsAsBytes)
}

func listPublishContractByPublisher(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("TRIT_ERROR_3")
	}

	dto := struct {
		Username string `json:"username"`
	}{}

	err := json.Unmarshal([]byte(args[0]), &dto)
	if err != nil {
		return shim.Error(err.Error())
	}

	prefix := prefixPublishContract
	resultsIterator, err := stub.GetStateByPartialCompositeKey(prefix, []string{})
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	var results []interface{}
	for resultsIterator.HasNext() {
		kvResult, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}

		c := publishContract{}
		err = json.Unmarshal(kvResult.Value, &c)
		if err != nil {
			return shim.Error(err.Error())
		}

		if c.Publisher == dto.Username {
			results = append(results, c)
		}

	}

	resultsAsBytes, err := json.Marshal(results)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(resultsAsBytes)
}

func listSellTritAdvertisingContract(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	prefix := prefixSellTritAdvertisingContract
	resultsIterator, err := stub.GetStateByPartialCompositeKey(prefix, []string{})
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	var results []interface{}
	for resultsIterator.HasNext() {
		kvResult, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}

		c := sellTritAdvertisingContract{}
		err = json.Unmarshal(kvResult.Value, &c)

		if err != nil {
			return shim.Error(err.Error())
		}

		c.Amount ,err = getHT(stub, []string{prefix + c.UUID, strconv.Itoa(c.Amount)})

		results = append(results, c)
	}

	resultsAsBytes, err := json.Marshal(results)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(resultsAsBytes)
}

func createPublishContract(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("TRIT_ERROR_3")
	}

	creatorOrg, creatorCertIssuer, creatorUserName, err := getTxCreatorInfo(stub)
	if err != nil {
		logger.Error(fmt.Sprintf("Error extracting creator identity info: %s\n", err.Error()))
		return shim.Error("TRIT_ERROR_4")
	}

	if !authenticateRegulatorOrg(creatorOrg, creatorCertIssuer) {
		return shim.Error("TRIT_ERROR_1")
	}

	dto := struct {
		UUID   string `json:"uuid"`
		Id     string `json:"trit_id"` // mã của bất động sản - ví dụ CHA, CHB
		Amount int    `json:"amount"`  // số lượng chứng chỉ quỹ được bán ra
		Time   string  `json:"time"`

	}{}
	err = json.Unmarshal([]byte(args[0]), &dto)
	if err != nil {
		return shim.Error(err.Error())
	}

	//validate amount
	if dto.Amount <= 0 {
		return shim.Error("TRIT_ERROR_2")
	}

	//validate UUID
	if len(dto.UUID) == 0 {
		return shim.Error("TRIT_ERROR)7")
	}

	//validate Id
	if len(dto.Id) == 0 {
		return shim.Error("TRIT_ERROR_8")
	}

	//validate real estate state

	estate, err := getRealEstateState(stub, dto.Id)
	if err != nil {
		return shim.Error(err.Error())
	}

	if len(estate.Shareholders) > 0 {
		return shim.Error("TRIT_ERROR_9")
	}

	//validate new contract id
	err = validatePublishContractId(stub, dto.UUID)
	if err != nil {
		return shim.Error(err.Error())
	}

	//get publisher
	p, err := getPublisherState(stub, creatorUserName)
	if err != nil {
		return shim.Error(err.Error())
	}

	//get owner
	owner, err := getShareholderState(stub, estate.OwnerId)
	if err != nil {
		return shim.Error(err.Error())
	}

	/** EVERY THING OK **/

	err = estate.publishTrit(stub, dto.Amount, &owner)
	if err != nil {
		return shim.Error(err.Error())
	}
	err = owner.saveShareHolderState(stub)
	if err != nil {
		return shim.Error(err.Error())
	}

	//add trit id to published Trit array
	p.setTritIdToPublisherPublishedTrits(dto.Id)
	err = p.savePublisherState(stub)
	if err != nil {
		return shim.Error(err.Error())
	}

	contract := publishContract{
		UUID:      dto.UUID,
		Publisher: creatorUserName,
		Amount:    dto.Amount,
		TritId:    dto.Id,
		Time:      dto.Time,
	}

	contractKey, err := getPublishContractKey(stub, dto.UUID)
	if err != nil {
		return shim.Error(err.Error())
	}
	contractAsBytes, err := json.Marshal(contract)
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState(contractKey, contractAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Return success, if the new contract has been created
	return shim.Success(nil)
}

func createTransferContractForBuyer(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("TRIT_ERROR_3")
	}

	creatorOrg, creatorCertIssuer, creatorUserName, err := getTxCreatorInfo(stub)
	if err != nil {
		logger.Error(fmt.Sprintf("Error extracting creator identity info: %s\n", err.Error()))
		return shim.Error("TRIT_ERROR_4")
	}

	if !authenticateShareholderOrg(creatorOrg, creatorCertIssuer) {
		return shim.Error("TRIT_ERROR_5")
	}

	dto := struct {
		UUID                          string `json:"uuid"`
		Amount                        int    `json:"amount"`
		SellTritAdvertisingContractId string `json:"sell_trit_advertising_contract_id"`
		CreateTime 					  string `json:"create_time"`
	}{}

	err = json.Unmarshal([]byte(args[0]), &dto)
	if err != nil {
		return shim.Error(err.Error())
	}

	//validate UUID
	if len(dto.UUID) == 0 {
		return shim.Error("UUID is empty")
	}

	//get buyer
	buyer, err := getShareholderState(stub, creatorUserName)
	if err != nil {
		return shim.Error(err.Error())
	}

	//validate Sell Trit Advertising Contract state
	advertisingContractAsBytes, err := validateSellTritAdvertisingContractState(stub, dto.SellTritAdvertisingContractId)
	if err != nil {
		return shim.Error(err.Error())
	}

	adContract := sellTritAdvertisingContract{}
	err = json.Unmarshal(advertisingContractAsBytes, &adContract)
	if err != nil {
		return shim.Error(err.Error())
	}

	//get seller
	seller, err := getShareholderState(stub, adContract.Seller)
	if err != nil {
		return shim.Error(err.Error())
	}

	//check buyer coin
	if !buyer.checkEnoughCoin(adContract.Price * dto.Amount) {
		return shim.Error("TRIT_ERROR_10")
	}

	//validate get estate state
	estate, err := getRealEstateState(stub, adContract.TritId)
	if err != nil {
		return shim.Error(err.Error())
	}

	if len(estate.Shareholders) == 0 {
		return shim.Error("TRIT_ERROR_11")
	}

	//validate amount
	if dto.Amount <= 0 {
		return shim.Error("TRIT_ERROR_2")
	}

	if dto.Amount > adContract.Amount {
		return shim.Error("TRIT_ERROR_12")
	}

	//validate new transfer contract id
	err = validateTransferContractId(stub, dto.UUID)
	if err != nil {
		return shim.Error(err.Error())
	}

	/** Every Thing OK **/

	//change amount of sell trit advertising contract
	err = adContract.changeAmount(stub, dto.Amount, &seller)
	if err != nil {
		return shim.Error(err.Error())
	}

	//transfer trit
	err = estate.transferTrit(stub, &buyer, &seller, dto.Amount)
	if err != nil {
		return shim.Error(err.Error())
	}

	//transfer coin
	buyer.Balance -= adContract.Price * dto.Amount
	seller.Balance += adContract.Price * dto.Amount

	err = buyer.saveShareHolderState(stub)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = seller.saveShareHolderState(stub)

	if err != nil {
		return shim.Error(err.Error())
	}

	//create transfer contract
	contract := transferContract{
		UUID:      dto.UUID,
		Buyer:     creatorUserName,
		Seller:    adContract.Seller,
		TritId:    adContract.TritId,
		TritPrice: adContract.Price,
		Amount:    dto.Amount,
		Time: 	   dto.CreateTime,
	}

	contractKey, err := getTransferContractKey(stub, dto.UUID)
	if err != nil {
		return shim.Error(err.Error())
	}
	contractAsBytes, err := json.Marshal(contract)
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState(contractKey, contractAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Return success, if the new contract has been created
	return shim.Success(nil)

}

func createChangePriceContract(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("TRIT_ERROR_3")
	}

	creatorOrg, creatorCertIssuer, creatorUserName, err := getTxCreatorInfo(stub)
	if err != nil {

		logger.Error(fmt.Sprintf("Error extracting creator identity info: %s\n", err.Error()))
		return shim.Error("TRIT_ERROR_4")
	}

	if !authenticateRegulatorOrg(creatorOrg, creatorCertIssuer) {
		return shim.Error("TRIT_ERROR_1")
	}

	dto := struct {
		UUID       string `json:"uuid"`
		Id    	   string `json:"trit_id"` // mã của bất động sản - ví dụ CHA, CHB
		Price      int    `json:"price"`   // gia cua bds được bán ra
		createTime string `json:"create_time"`
	}{}

	err = json.Unmarshal([]byte(args[0]), &dto)
	if err != nil {
		return shim.Error(err.Error())
	}

	//validate UUID
	if len(dto.UUID) == 0 {
		return shim.Error("TRIT_ERROR_7")
	}

	//validate real estate state
	estate, err := getRealEstateState(stub, dto.Id)
	if err != nil {
		return shim.Error(err.Error())
	}

	//validate price
	if dto.Price <= 0 {
		return shim.Error("TRIT_ERROR_13")
	}

	//validate new contract id
	err = validateChangePriceContractId(stub, dto.UUID)
	if err != nil {
		return shim.Error(err.Error())
	}

	/** Every Thing Ok **/
	err = estate.changePrice(stub, dto.Price)
	if err != nil {
		return shim.Error(err.Error())
	}

	contract := changePriceContract{
		UUID:      dto.UUID,
		Publisher: creatorUserName,
		Price:     dto.Price,
		TritId:    dto.Id,
		Time:      dto.createTime,
	}

	contractKey, err := getChangePriceContractKey(stub, dto.UUID)
	if err != nil {
		return shim.Error(err.Error())
	}
	contractAsBytes, err := json.Marshal(contract)
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState(contractKey, contractAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Return success, if the new contract has been created
	return shim.Success(nil)
}

func createSellTritAdvertisingContract(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("TRIT_ERROR_3")
	}

	creatorOrg, creatorCertIssuer, creatorUserName, err := getTxCreatorInfo(stub)
	if err != nil {
		logger.Error(fmt.Sprintf("Error extracting creator identity info: %s\n", err.Error()))
		return shim.Error("TRIT_ERROR_4")
	}

	if !authenticateShareholderOrg(creatorOrg, creatorCertIssuer) {
		return shim.Error("TRIT_ERROR_6")
	}

	dto := struct {
		UUID       string  `json:"uuid"`
		TritId     string  `json:"trit_id"` // id chứng chỉ quỹ
		Amount     int     `json:"amount"`  // tổng số lượng chỉnh chỉ quỹ trao đổi
		Price      int     `json:"price"`   // giá bán
		CreateTime string  `json:"create_time"`
	}{}

	err = json.Unmarshal([]byte(args[0]), &dto)
	if err != nil {
		return shim.Error(err.Error())
	}

	//validate UUID
	if len(dto.UUID) == 0 {
		return shim.Error("TRIT_ERROR_7")
	}

	//validate price
	if dto.Price <= 0 {
		return shim.Error("TRIT_ERROR_15")
	}

	//validate amount
	if dto.Amount <= 0 {
		return shim.Error("TRIT_ERROR_2")
	}

	seller, err := getShareholderState(stub, creatorUserName)
	if err != nil {
		return shim.Error(err.Error())
	}

	totalAdTrit, err := seller.getTotalAdvertisingTritOf(stub, dto.TritId)
	if err != nil {
		return shim.Error(err.Error())
	}

	if tritOfseller, ok := seller.TritList[dto.TritId]; ok {
		if tritOfseller < dto.Amount {
			return shim.Error("TRIT_ERROR_14")
		}

		if dto.Amount > tritOfseller-totalAdTrit {
			return shim.Error(fmt.Sprintf("TRIT_ERROR_16"))
		}
	} else {
		return shim.Error("TRIT_ERROR_17")
	}

	//validate real estate state
	_, err = validateRealEstateState(stub, dto.TritId)
	if err != nil {
		return shim.Error(err.Error())
	}

	//validate new contract id
	err = validateSellTritAdvertisingContractId(stub, dto.UUID)
	if err != nil {
		return shim.Error(err.Error())
	}

	/** EVERY THING OK **/

	contractKey, err := getSellTritAdvertisingContractKey(stub, dto.UUID)
	contract := sellTritAdvertisingContract{
		UUID:   dto.UUID,
		Seller: creatorUserName,
		TritId: dto.TritId,
		Amount: dto.Amount,
		Price:  dto.Price,
		Time: 	dto.CreateTime,
	}

	contractAsBytes, err := json.Marshal(contract)
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState(contractKey, contractAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	seller.setIdToShareholderSellTritAdvertisingContractList(contract.TritId, contract.UUID)
	err = seller.saveShareHolderState(stub)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Return success, if the new contract has been created
	return shim.Success(nil)
}

func createAndPublishRealEstate(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("TRIT_ERROR_3")
	}

	creatorOrg, creatorCertIssuer, creatorUserName, err := getTxCreatorInfo(stub)
	if err != nil {
		logger.Error(fmt.Sprintf("Error extracting creator identity info: %s\n", err.Error()))
		return shim.Error("TRIT_ERROR_4")
	}

	if !authenticateRegulatorOrg(creatorOrg, creatorCertIssuer) {
		return shim.Error("TRIT_ERROR_1")
	}

	dto := struct {
		UUID        string  `json:"uuid"`
		Id          string  `json:"id"`           // Mã của bất động sản - ví dụ CHA, CHB
		Price       int     `json:"price"`        // Giá trị
		SquareMeter float64 `json:"square_meter"` // Diện tích
		Address     string  `json:"address"`      // Địa chỉ
		OwnerId     string  `json:"owner_id"`     // Id của chủ sở hữu bất động sản
		Amount      int     `json:"amount"`       // Số lượng chứng chỉ quỹ được bán ra
		Description string  `json:"description"`  // Miêu tả
		CreateTime	string	`json:"create_time"`

	}{}

	err = json.Unmarshal([]byte(args[0]), &dto)
	if err != nil {
		return shim.Error(err.Error())
	}

	//validate id
	if len(dto.Id) == 0 {
		return shim.Error("TRIT_ERROR_8")
	}

	//validate uuid
	if len(dto.UUID) == 0 {
		return shim.Error("TRIT_ERROR_7")
	}

	//validate square meter
	if dto.SquareMeter <= 0 {
		return shim.Error("TRIT_ERROR_18")
	}

	//validate amount
	if dto.Amount <= 0 {
		return shim.Error("TRIT_ERROR_2")
	}

	//validate price
	if dto.Price <= 0 {
		return shim.Error("TRIT_ERROR_19")
	}

	//validate description
	if len(dto.Description) == 0 {
		return shim.Error("TRIT_ERROR_20")
	}

	//validate address
	if len(dto.Address) == 0 {
		return shim.Error("TRIT_ERROR_21")
	}

	//validate new real estate id
	err = validateRealEstateId(stub, dto.Id)
	if err != nil {
		return shim.Error(err.Error())
	}

	//get publisher
	p, err := getPublisherState(stub, creatorUserName)
	if err != nil {
		return shim.Error(err.Error())
	}

	// get owner
	owner, err := getShareholderState(stub, dto.OwnerId)
	if err != nil {
		return shim.Error(err.Error())
	}

	//every thing OK

	estate := realEstate{
		Id:           dto.Id,
		Shareholders: nil,
		Price:        dto.Price,
		SquareMeter:  dto.SquareMeter,
		Address:      dto.Address,
		OwnerId:      dto.OwnerId,

	}

	err = estate.publishTrit(stub, dto.Amount, &owner)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = owner.saveShareHolderState(stub)
	if err != nil {
		return shim.Error(err.Error())
	}

	//add trit id to publisher's published trit array
	p.setTritIdToPublisherPublishedTrits(dto.Id)
	err = p.savePublisherState(stub)
	if err != nil {
		return shim.Error(err.Error())
	}

	//validate new contract id
	err = validatePublishContractId(stub, dto.UUID)
	if err != nil {
		return shim.Error(err.Error())
	}

	//save publish contract to ledger
	contract := publishContract{
		UUID:      dto.UUID,
		Publisher: creatorUserName,
		Amount:    dto.Amount,
		TritId:    dto.Id,
		Time: 		  dto.CreateTime,
	}

	contractAsBytes, err := json.Marshal(contract)
	if err != nil {
		return shim.Error(err.Error())
	}

	contractKey, err := getPublishContractKey(stub, dto.UUID)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(contractKey, contractAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	logger.Info("==== createAndPublishRealEstate success ====");

	// Return success, if the new contract has been created
	return shim.Success(nil)

}

func changePriceSellTritAdvertisingContract(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("TRIT_ERROR_3")
	}

	creatorOrg, creatorCertIssuer, creatorUserName, err := getTxCreatorInfo(stub)
	if err != nil {
		logger.Error(fmt.Sprintf("Error extracting creator identity info: %s\n", err.Error()))
		return shim.Error("TRIT_ERROR_4")
	}

	if !authenticateShareholderOrg(creatorOrg, creatorCertIssuer) {
		return shim.Error("TRIT_ERROR_5")
	}

	dto := struct {
		Price                         int    `json:"price"`
		SellTritAdvertisingContractId string `json:"sell_trit_advertising_contract_id"`
	}{}

	err = json.Unmarshal([]byte(args[0]), &dto)
	if err != nil {
		return shim.Error(err.Error())
	}

	//Validate Price
	if dto.Price <= 0 {
		return shim.Error("TRIT_ERROR_15")
	}

	//validate Sell Trit Advertising Contract state

	advertisingContractAsBytes, err := validateSellTritAdvertisingContractState(stub, dto.SellTritAdvertisingContractId)
	if err != nil {
		return shim.Error(err.Error())
	}

	adContract := sellTritAdvertisingContract{}
	err = json.Unmarshal(advertisingContractAsBytes, &adContract)
	if err != nil {
		return shim.Error(err.Error())
	}

	//validate seller
	if adContract.Seller != creatorUserName {
		logger.Error(creatorUserName + "doesn't has permission to change Advertising contract with id =  " + dto.SellTritAdvertisingContractId)
		return shim.Error("TRIT_ERROR_6")
	}

	//change price of advertising contract
	err = adContract.changePrice(stub, dto.Price)

	if err != nil {
		return shim.Error(err.Error())
	}

	// Return success, if the contract has been changed
	return shim.Success(nil)

}

func deleteSellTritAdvertisingContract(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("TRIT_ERROR_3")
	}

	creatorOrg, creatorCertIssuer, creatorUserName, err := getTxCreatorInfo(stub)
	if err != nil {
		logger.Error(fmt.Sprintf("Error extracting creator identity info: %s\n", err.Error()))
		return shim.Error("TRIT_ERROR_4")
	}

	if !authenticateShareholderOrg(creatorOrg, creatorCertIssuer) {
		return shim.Error("TRIT_ERROR_5")
	}

	dto := struct {
		Id string `json:"id"`
	}{}

	err = json.Unmarshal([]byte(args[0]), &dto)
	if err != nil {
		return shim.Error(err.Error())
	}

	contractAsBytes, err := validateSellTritAdvertisingContractState(stub, dto.Id)
	if err != nil {
		return shim.Error(err.Error())
	}

	contract := sellTritAdvertisingContract{}
	err = json.Unmarshal(contractAsBytes, &contract)
	if err != nil {
		return shim.Error(err.Error())
	}

	if contract.Seller != creatorUserName {
		return shim.Error("TRIT_ERROR_22")
	}
	seller, err := getShareholderState(stub, contract.Seller)
	if err != nil {
		return shim.Error(err.Error())
	}

	key, err := getSellTritAdvertisingContractKey(stub, dto.Id)
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.DelState(key)
	if err != nil {
		return shim.Error(err.Error())
	}

	seller.deleteIdToShareholderSellTritAdvertisingContractList(contract.TritId, contract.UUID)
	err = seller.saveShareHolderState(stub)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}

func listSellTritAdvertisingContractByUser(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("TRIT_ERROR_3")
	}
	dto := struct {
		Username string `json:"username"`
	}{}
	err := json.Unmarshal([]byte(args[0]), &dto)
	if err != nil {
		return shim.Error(err.Error())
	}

	prefix := prefixSellTritAdvertisingContract
	resultsIterator, err := stub.GetStateByPartialCompositeKey(prefix, []string{})
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	var results []interface{}
	for resultsIterator.HasNext() {
		kvResult, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}

		c := sellTritAdvertisingContract{}
		err = json.Unmarshal(kvResult.Value, &c)
		if err != nil {
			return shim.Error(err.Error())
		}

		if err != nil {
			return shim.Error(err.Error())
		}

		if c.Seller == dto.Username {
			c.Amount ,err = getHT(stub, []string{prefix + c.UUID, strconv.Itoa(c.Amount)})
			results = append(results, c)
		}



	}

	resultsAsBytes, err := json.Marshal(results)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(resultsAsBytes)

}
