docker exec -it cli bash
peer chaincode invoke -n chaincode_advance_tutorial_lbh -v 0  -c '{"Args":["putCompose","lbh","123","456","789"]}' -C myc