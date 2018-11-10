package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
	"encoding/json"
	"strconv"
	"fmt"
)

type TraceAsset struct {
	assetId int `json:"assetId"`
	mfd string `json:"mfd"`
	holder string `json:"holder"`
	transPrice int  `json:"transPrice"`
	other string `json:"other"`
}

// 新增or更新資產
// 預計輸入參數
// "4","20181111","lbh","500","singleValentine"
func (sc *SampleChaincode) putTraceAssetIntoBC(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	assetIdIsString, mfd, holder, transPriceIsString,other :=  args[0], args[1], args[2], args[3],args[4]

	assetId,_ := strconv.Atoi(assetIdIsString)
	transPrice,_ := strconv.Atoi(transPriceIsString)

	traceAsset := TraceAsset{assetId, mfd, holder, transPrice,other}

	traceAssetAsbyte, _ := json.Marshal(traceAsset)

	err := stub.PutState(assetIdIsString,traceAssetAsbyte)

	if err !=nil {
		shim.Error(err.Error())
	}

	return shim.Success(traceAssetAsbyte)
}

// 取得資產
// 預計輸入參數
// "4"
func (sc *SampleChaincode) getTraceAssetFromBC(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	assetIdIsString := args[0]

	transAssetAsByte, _ := stub.GetState(assetIdIsString)

	if transAssetAsByte == nil {
		return shim.Error("Could not locate transAsset")
	}

	return shim.Success(transAssetAsByte)

}

// 資產所有權轉讓
// 預計輸入參數
// "4","amanda"
func (sc *SampleChaincode) changeTraceAssetIntoBC(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	assetIdIsString := args[0]

	transAssetAsByteFromBC, _ := stub.GetState(assetIdIsString)
	tranAsset := TraceAsset{}
	json.Unmarshal(transAssetAsByteFromBC, &tranAsset)

	tranAsset.holder = args[1]

	transAssetAsByte, _ := json.Marshal(tranAsset)

	err := stub.PutState(assetIdIsString, transAssetAsByte)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to change tranAsset holder: %s", args[0]))
	}

	return shim.Success(transAssetAsByte)

}

// 市場資產整體平均價格
// 預計輸入參數
// "1","4"
func (sc *SampleChaincode) getAverageAssetPrice(stub shim.ChaincodeStubInterface, args []string) sc.Response {


	startKey := args[0]
	endKey := args[1]

	resultsIterator,_ := stub.GetStateByRange(startKey,endKey)
	defer resultsIterator.Close()

	var transAssetPriceSlice []int
	for resultsIterator.HasNext(){

		transAssetAsByte,_ := resultsIterator.Next()
		tranAsset := TraceAsset{}
		json.Unmarshal(transAssetAsByte.Value, &tranAsset)

		transAssetPriceSlice = append(transAssetPriceSlice, tranAsset.transPrice)

	}

	// average price
	var sumPrice int
	for _, element := range transAssetPriceSlice {
		sumPrice += element
	}
	averagePrice := sumPrice/len(transAssetPriceSlice)

	return shim.Success([]byte(strconv.Itoa(averagePrice)))

}


// 特定資產的持有人歷史變化
// "4"
func (sc *SampleChaincode) getAssetHolderHistory(stub shim.ChaincodeStubInterface, args []string) sc.Response {


	assetIdIsString := args[0]

	assetIter, _ := stub.GetHistoryForKey(assetIdIsString)
	defer assetIter.Close()

	var holderHistory string
	for assetIter.HasNext() {
		result, _ := assetIter.Next()
		tranAsset := TraceAsset{}
		json.Unmarshal(result.Value, &tranAsset)

		holderHistory +=  tranAsset.holder + "||"
	}
	return shim.Success([]byte(holderHistory))
}

// 提供模擬資料
func (sc *SampleChaincode) simulateData(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	traceAssets := []TraceAsset{
		TraceAsset{assetId:1,mfd:"20181109",holder:"lbh",transPrice:100, other:""},
		TraceAsset{assetId:2,mfd:"20181111",holder:"lbh",transPrice:200, other:""},
		TraceAsset{assetId:3,mfd:"20181113",holder:"lbh",transPrice:300, other:""},
	}

	for _, traceAsset := range traceAssets {
		traceAssetAsBytes, _ := json.Marshal(traceAsset)
		stub.PutState(strconv.Itoa(traceAsset.assetId),traceAssetAsBytes)
		fmt.Println("Added", traceAsset)
	}

	return shim.Success(nil)
}