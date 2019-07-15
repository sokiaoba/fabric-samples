# Test Fabcar Chaincode in Dev-mode

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

cd /opt/gopath/src/chaincode/copyright
go build -o copyright
CORE_CHAINCODE_LOGLEVEL=debug CORE_PEER_ADDRESS=peer:7052 CORE_CHAINCODE_ID_NAME=mycc:v0 ./copyright
```

## terminal3

* go

```
docker exec -it cli bash

peer chaincode install -p chaincodedev/chaincode/copyright -n mycc -v v0
peer chaincode instantiate -n mycc -v v0 -c '{"Args":[]}' -C myc

peer chaincode invoke -n mycc -c '{"function":"initLedger","Args":[]}' -C myc

peer chaincode invoke -n mycc -c '{"function":"createRightHolder","Args":["rightHolder1","John Lennon", "2019-07-15 00:00:00"]}' -C myc
peer chaincode invoke -n mycc -c '{"function":"createRightHolder","Args":["rightHolder2","Paul McCartney", "2019-07-15 00:00:00"]}' -C myc

peer chaincode invoke -n mycc -c '{"function":"createSong","Args":["song1", "Hey Jude", "hash value of Hey Jude", "2019-07-15 00:00:00"]}' -C myc
peer chaincode invoke -n mycc -c '{"function":"createCopyright","Args":["copyright1","song1","rightHolder1","50","1","2019-07-15 00:00:00"]}' -C myc
peer chaincode invoke -n mycc -c '{"function":"createCopyright","Args":["copyright2","song1","rightHolder2","50","1","2019-07-15 00:00:00"]}' -C myc

peer chaincode invoke -n mycc -c '{"function":"createSong","Args":["song2", "Strawberry Fields Forever", "hash value of Strawberry Fields Forever", "2019-07-15 00:00:00"]}' -C myc
peer chaincode invoke -n mycc -c '{"function":"createCopyright","Args":["copyright1","song2","rightHolder1","100","1","2019-07-15 00:00:00"]}' -C myc

peer chaincode invoke -n mycc -c '{"function":"querySong","Args":["song1"]}' -C myc
peer chaincode invoke -n mycc -c '{"function":"querySong","Args":["song2"]}' -C myc

peer chaincode invoke -n mycc -c '{"function":"querySong","Args":["rightHolder1"]}' -C myc
```
