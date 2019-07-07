# init

```
docker rm -f $(docker ps -aq)  
docker rmi -f $(docker images)
```

# terminal1 

```
cd ~/Dev/fabric-samples/chaincode-docker-devmode/  
docker-compose -f docker-compose-simple.yaml up
```

# terminal2

```
docker exec -it chaincode bash

cd playground  
npm install  
npm run build  
CORE_CHAINCODE_ID_NAME="mycc:v0" node ./ --peer.address peer:7052
```

# terminal3

```
docker exec -it cli bash

peer chaincode install -p /opt/gopath/src/chaincodedev/chaincode/playground -n mycc -v v0 -l node  
peer chaincode instantiate -n mycc -v v0 -c '{"Args":[]}' -C myc   
peer chaincode invoke -n mycc -c '{"Args":["invoke"]}' -C myc
```