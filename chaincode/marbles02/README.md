# Test Marbles Chaincode in Dev-mode

## init

```
docker rm -f $(docker ps -aq)
docker rmi -f $(docker images)
```

## terminal1 

```
cd ~/Dev/fabric-samples/chaincode-docker-devmode/
docker-compose -f docker-compose-simple.yaml up
```

## terminal2

* go

```
docker exec -it chaincode bash

cd /opt/gopath/src/chaincode/marbles02/go
go build -o marbles_chaincode
CORE_CHAINCODE_LOGLEVEL=debug CORE_PEER_ADDRESS=peer:7052 CORE_CHAINCODE_ID_NAME=marbles:v0 ./marbles_chaincode
```

## terminal3

* go

```
docker exec -it cli bash

peer chaincode install -p chaincodedev/chaincode/marbles02/go -n marbles -v v0
peer chaincode instantiate -n marbles -v v0 -c '{"Args":[]}' -C myc

peer chaincode invoke -C myc -n marbles -c '{"Args":["initMarble","marble1","blue","35","tom"]}'
peer chaincode invoke -C myc -n marbles -c '{"Args":["initMarble","marble2","red","50","tom"]}'
peer chaincode invoke -C myc -n marbles -c '{"Args":["initMarble","marble3","blue","70","tom"]}'
peer chaincode invoke -C myc -n marbles -c '{"Args":["transferMarble","marble2","jerry"]}'
peer chaincode invoke -C myc -n marbles -c '{"Args":["transferMarblesBasedOnColor","blue","jerry"]}'
peer chaincode invoke -C myc -n marbles -c '{"Args":["delete","marble1"]}'

peer chaincode query -C myc -n marbles -c '{"Args":["readMarble","marble1"]}'
peer chaincode query -C myc -n marbles -c '{"Args":["getMarblesByRange","marble1","marble3"]}'
peer chaincode query -C myc -n marbles -c '{"Args":["getHistoryForMarble","marble1"]}'
```
