# Test Fabcar Chaincode in Dev-mode (typescript)

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

```
docker exec -it chaincode bash

cd /opt/gopath/src/chaincode/fabcar/typescript
npm install
npm run build
$(npm bin)/fabric-chaincode-node start --peer.address peer:7052 --chaincode-id-name "mycontract:v0"
```

## terminal3

```
docker exec -it cli bash

peer chaincode install -p /opt/gopath/src/chaincodedev/chaincode/fabcar/typescript -n mycontract -v v0 -l node
peer chaincode instantiate -n mycontract -v v0 -c '{"Args":[]}' -C myc

peer chaincode invoke -n mycontract -c '{"function":"initLedger","Args":[]}' -C myc
peer chaincode query -n mycontract -c '{"function":"queryAllCars","Args":[]}' -C myc
peer chaincode query -n mycontract -c '{"function":"queryCar","Args":["CAR1"]}' -C myc
peer chaincode invoke -n mycontract -c '{"function":"changeCarOwner","Args":["CAR1", "Chris"]}' -C myc
peer chaincode query -n mycontract -c '{"function":"queryCar","Args":["CAR1"]}' -C myc
```
