docker exec -it cli bash
# story1
# 輸入基本數據 -> 取得基本數據 -> 更新基本數據 -> 檢視數據變換歷史
peer chaincode invoke -n chaincode_basic_tutorial_lbh -v 0  -c '{"Args":["setSampleAsset","abc","123"]}' -C myc
peer chaincode invoke -n chaincode_basic_tutorial_lbh -v 0  -c '{"Args":["getSampleAsset","abc"]}' -C myc
peer chaincode invoke -n chaincode_basic_tutorial_lbh -v 0  -c '{"Args":["setSampleAsset","abc","456"]}' -C myc
peer chaincode invoke -n chaincode_basic_tutorial_lbh -v 0  -c '{"Args":["getHistoryForSample","abc"]}' -C myc

# story2
# simulate data
# 輸入模擬資產 -> 取得資產 -> 追加資產
# -> 取得平均價格 -> 計算平均價格 -> 變更資產持有人 -> 檢視資產持有人歷史
peer chaincode invoke -n chaincode_basic_tutorial_lbh -v 0  -c '{"Args":["simulateTA"]}' -C myc
peer chaincode invoke -n chaincode_basic_tutorial_lbh -v 0  -c '{"Args":["getTA","1"]}' -C myc
peer chaincode invoke -n chaincode_basic_tutorial_lbh -v 0  -c '{"Args":["putTA","4","20181111","lbh","500","singleValentine"]}' -C myc
peer chaincode invoke -n chaincode_basic_tutorial_lbh -v 0  -c '{"Args":["getTAAveragePrice","1","5"]}' -C myc
peer chaincode invoke -n chaincode_basic_tutorial_lbh -v 0  -c '{"Args":["changeTaHolder","4","amanda"]}' -C myc
peer chaincode invoke -n chaincode_basic_tutorial_lbh -v 0  -c '{"Args":["getTaHistoryHolder","4"]}' -C myc
