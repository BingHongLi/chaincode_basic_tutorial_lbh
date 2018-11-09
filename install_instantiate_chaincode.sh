docker exec -it cli bash
cd chaincode/chaincode_basic_tutorial_lbh
peer chaincode install -p chaincodedev/chaincode/chaincode_basic_tutorial_lbh -n chaincode_basic_tutorial_lbh -v 0
peer chaincode instantiate -n chaincode_basic_tutorial_lbh -v 0  -c '{"Args":[]}' -C myc