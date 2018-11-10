package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"strconv"
	"encoding/json"
	"fmt"
)

/*
	定義資產
*/
type TraceAsset struct {
	AssetId int `json:"assetId"`
	Mfd string `json:"mfd"`
	Holder string `json:"holder"`
	TransPrice int  `json:"transPrice"`
	Other string `json:"other"`
}

/*
	新增or更新資產
	預計輸入參數
	"4","20181111","lbh","500","singleValentine"
*/
func (sc *SampleChaincode) putTraceAssetIntoBC(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	//將用戶輸入轉換所需變數
	assetIdIsString, mfd, holder, transPriceIsString,other :=  args[0], args[1], args[2], args[3],args[4]

	// 將用戶id與交易價格從字串轉成數字
	assetId,_ := strconv.Atoi(assetIdIsString)
	transPrice,_ := strconv.Atoi(transPriceIsString)

	// 將變數組合成資產
	traceAsset := TraceAsset{assetId, mfd, holder, transPrice,other}

	// 將資產轉成區塊鏈可懂的二進制資料
	traceAssetAsbyte, _ := json.Marshal(traceAsset)

	// 放入資產
	err := stub.PutState(assetIdIsString,traceAssetAsbyte)

	// 若出錯，則回報錯誤
	if err !=nil {
		shim.Error(err.Error())
	}

	// 成功插入，則將資產傳回
	return shim.Success(traceAssetAsbyte)
}

/*
	取得資產
	預計輸入參數
	"4"
*/
func (sc *SampleChaincode) getTraceAssetFromBC(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	// 將用戶輸入轉換所需變數
	assetIdIsString := args[0]

	// 使用assetIdIsString向區塊鏈取得資料
	transAssetAsByte, _ := stub.GetState(assetIdIsString)

	// 若沒取得資料，則回應系統錯誤
	if transAssetAsByte == nil {
		return shim.Error("Could not locate transAsset")
	}

	// 若取得資料，將完整資料傳回
	return shim.Success(transAssetAsByte)

}

/*
	資產所有權轉讓
	預計輸入參數
	"4","amanda"
*/
func (sc *SampleChaincode) changeTraceAssetIntoBC(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	// 將用戶輸入轉換所需變數
	assetIdIsString := args[0]
	holder := args[1]
	// 使用assetIdIsString向區塊鏈取得二進制資料
	transAssetAsByteFromBC, _ := stub.GetState(assetIdIsString)

	// 將此二進制資料轉換成可操作的TraceAsset
	tranAsset := TraceAsset{}
	json.Unmarshal(transAssetAsByteFromBC, &tranAsset)

	// 將資產持有人更換成  holder變數之內容
	tranAsset.Holder = holder

	// 將資產轉回二進制資料
	transAssetAsByte, _ := json.Marshal(tranAsset)

	// 將二進制資料放回區塊鏈中
	err := stub.PutState(assetIdIsString, transAssetAsByte)

	// 若有異常，則回傳系統錯誤
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to change tranAsset holder: %s", args[0]))
	}

	// 若正常，則將資料傳回
	return shim.Success(transAssetAsByte)

}

/*
	市場資產整體平均價格
	預計輸入參數
	"1","4"
*/
func (sc *SampleChaincode) getAverageAssetPrice(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	// 將用戶輸入轉換所需變數，起始key與結束key
	startKey := args[0]
	endKey := args[1]

	// 取得批量資料
	resultsIterator,_ := stub.GetStateByRange(startKey,endKey)
	defer resultsIterator.Close()

	// 轉換成TransAsset的Slice陣列
	var transAssetPriceSlice []int
	for resultsIterator.HasNext(){

		transAssetAsByte,_ := resultsIterator.Next()
		tranAsset := TraceAsset{}
		json.Unmarshal(transAssetAsByte.Value, &tranAsset)

		transAssetPriceSlice = append(transAssetPriceSlice, tranAsset.TransPrice)

	}

	// 計算平均價格
	var sumPrice int
	for _, element := range transAssetPriceSlice {
		sumPrice += element
	}
	averagePrice := sumPrice/len(transAssetPriceSlice)

	// 傳回平均價格
	return shim.Success([]byte(strconv.Itoa(averagePrice)))

}


/*
	特定資產的持有人歷史變化
	預計輸入參數
	"4"
*/
func (sc *SampleChaincode) getAssetHolderHistory(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	// 將用戶輸入轉換所需變數
	assetIdIsString := args[0]

	// 取得該資產的歷史紀錄
	assetIter, _ := stub.GetHistoryForKey(assetIdIsString)
	defer assetIter.Close()

	// 取出歷史持有人
	var holderHistory string
	for assetIter.HasNext() {
		result, _ := assetIter.Next()
		tranAsset := TraceAsset{}
		json.Unmarshal(result.Value, &tranAsset)

		holderHistory +=  tranAsset.Holder + "||"
	}
	// 將結果傳回
	return shim.Success([]byte(holderHistory))
}

// 提供模擬資料
func (sc *SampleChaincode) simulateData(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	traceAssets := []TraceAsset{
		TraceAsset{AssetId:1,Mfd:"20181109",Holder:"lbh",TransPrice:100, Other:""},
		TraceAsset{AssetId:2,Mfd:"20181111",Holder:"lbh",TransPrice:200, Other:""},
		TraceAsset{AssetId:3,Mfd:"20181113",Holder:"lbh",TransPrice:300, Other:""},
	}

	for _, traceAsset := range traceAssets {
		traceAssetAsBytes, _ := json.Marshal(traceAsset)
		stub.PutState(strconv.Itoa(traceAsset.AssetId),traceAssetAsBytes)
		fmt.Println("Added", traceAsset)
	}

	return shim.Success(nil)
}