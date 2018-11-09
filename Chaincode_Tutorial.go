package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"fmt"
)

type SampleChaincode struct {}


func (sc *SampleChaincode) Init( stub shim.ChaincodeStubInterface) peer.Response {

	return shim.Success(nil)

}

func (sc *SampleChaincode) Invoke( stub shim.ChaincodeStubInterface) peer.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := stub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger
	if function == "setSampleAsset" {
		return sc.putSampleAsset(stub, args)
	} else if function == "getSampleAsset" {
		return sc.getSampleAsset(stub, args)
	} else if function == "getHistoryForSample" {
		return sc.getHistorySampleAsset(stub, args)
	}

	return shim.Error("Invalid Smart Contract function name.")

}


func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SampleChaincode))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}

}