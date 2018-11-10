docker exec -it cli bash
# story1
peer chaincode invoke -n chaincode_basic_tutorial_lbh -v 0  -c '{"Args":["setSampleAsset","abc","123"]}' -C myc
peer chaincode invoke -n chaincode_basic_tutorial_lbh -v 0  -c '{"Args":["getSampleAsset","abc"]}' -C myc
peer chaincode invoke -n chaincode_basic_tutorial_lbh -v 0  -c '{"Args":["setSampleAsset","abc","456"]}' -C myc
peer chaincode invoke -n chaincode_basic_tutorial_lbh -v 0  -c '{"Args":["getHistoryForSample","abc"]}' -C myc

# story2
# simulate data
peer chaincode invoke -n chaincode_basic_tutorial_lbh -v 0  -c '{"Args":["simulateTA"]}' -C myc
peer chaincode invoke -n chaincode_basic_tutorial_lbh -v 0  -c '{"Args":["getTA","1"]}' -C myc
peer chaincode invoke -n chaincode_basic_tutorial_lbh -v 0  -c '{"Args":["putTA","4","20181111","lbh","500","singleValentine"]}' -C myc
peer chaincode invoke -n chaincode_basic_tutorial_lbh -v 0  -c '{"Args":["getTAAveragePrice","1","4"]}' -C myc
peer chaincode invoke -n chaincode_basic_tutorial_lbh -v 0  -c '{"Args":["changeTaHolder","4","amanda"]}' -C myc
peer chaincode invoke -n chaincode_basic_tutorial_lbh -v 0  -c '{"Args":["getTaHistoryHolder","4"]}' -C myc
