package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

type SampleAsset struct {}

/*


*/
func (sc *SampleChaincode) putSampleAsset(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	// 將用戶的第一個參數當作key, 第二的參數當作內容，並轉換成byte，存入區塊鏈
	stub.PutState(args[0],[]byte(args[1]))

	// 回傳成功塞值
	return shim.Success([]byte("insert success"))

}

/*


*/
func (sc *SampleChaincode) getSampleAsset(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	// 依照用戶輸入的參數向區塊鏈取得內容，並存回result變數
	result, _ := stub.GetState(args[0])

	// 將結果傳回
	return shim.Success(result)

}


func (sc *SampleChaincode) getHistorySampleAsset(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	valueIter,_ := stub.GetHistoryForKey(args[0])

	var value string
	for valueIter.HasNext() {
		result, _ := valueIter.Next()
		value +=  string(result.Value) + "||"
	}

	return shim.Success( []byte(value) )

}



