package main

import (
	"encoding/json"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
	"strconv"
)

type SmartContract struct {
}

type Song struct {
	ObjectType string      `json:"docType"`
	Name       string      `json:"name"`
	Hash       string      `json:"hash"`
	Copyrights []Copyright `json:"copyrights"`
	Timestamp  string      `json:"timestamp"`
}

type RightHolder struct {
	ObjectType string      `json:"docType"`
	Name       string      `json:"name"`
	Copyrights []Copyright `json:"copyrights"`
	Timestamp  string      `json:"timestamp"`
}

type Copyright struct {
	ObjectType  string `json:"docType"`
	Song        string `json:"song"`        // song id
	RightHolder string `json:"rightHolder"` // right holder id
	Percentage  int    `json:"percentage"`
	Type        int    `json:"type"` // 1. song writing, 2. lyric writing, 2. secondary use
	Timestamp   string `json:"timestamp"`
}

func (s *SmartContract) Init(stub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

func (s *SmartContract) Invoke(stub shim.ChaincodeStubInterface) sc.Response {
	function, args := stub.GetFunctionAndParameters()

	if function == "initLedger" {
		return s.initLedger(stub)
	} else if function == "querySong" {
		return s.querySong(stub, args)
	} else if function == "queryRightHolder" {
		return s.queryRightHolder(stub, args)
	} else if function == "createSong" {
		return s.createSong(stub, args)
	} else if function == "createRightHolder" {
		return s.createRightHolder(stub, args)
	} else if function == "createCopyright" {
		return s.createCopyright(stub, args)
	}

	return shim.Error("Invalid function name.")
}

func (s *SmartContract) initLedger(stub shim.ChaincodeStubInterface) sc.Response {
	// Todo
	return shim.Success(nil)
}

func (s *SmartContract) createSong(stub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}

	var (
		songId    = args[0]
		name      = args[1]
		hash      = args[2]
		timestamp = args[3]
	)

	song := Song{
		ObjectType: "Song",
		Name:       name,
		Hash:       hash,
		Copyrights: []string{},
		Timestamp:  timestamp,
	}

	songAsBytes, _ := json.Marshal(song)
	stub.PutState(songId, songAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) createRightHolder(stub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	var (
		rightHolderId = args[0]
		name          = args[1]
		timestamp     = args[2]
	)

	rightHolder := RightHolder{
		ObjectType: "RightHolder",
		Name:       name,
		Copyrights: []string{},
		Timestamp:  timestamp,
	}

	rightHolderAsBytes, _ := json.Marshal(rightHolder)
	stub.PutState(rightHolderId, rightHolderAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) createCopyright(stub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 6 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	var (
		copyrightId   = args[0]
		songId        = args[1]
		rightHolderId = args[2]
		timestamp     = args[5]
	)

	songAsBytes, err := stub.GetState(songId)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get Song for " + songId + "\"}"
		return shim.Error(jsonResp)
	}

	rightHolderAsBytes, err := stub.GetState(rightHolderId)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get RightHolder for " + rightHolderId + "\"}"
		return shim.Error(jsonResp)
	}

	percentage, err := strconv.Atoi(args[3])
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get Percentage for " + args[3] + "\"}"
		return shim.Error(jsonResp)
	}

	_type, err := strconv.Atoi(args[4])
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get Type for " + args[4] + "\"}"
		return shim.Error(jsonResp)
	}

	copyright := Copyright{
		ObjectType:  "CopyRight",
		Song:        songId,
		RightHolder: rightHolderId,
		Percentage:  percentage,
		Type:        _type,
		Timestamp:   timestamp,
	}

	copyrightAsBytes, _ := json.Marshal(copyright)
	stub.PutState(copyrightId, copyrightAsBytes)

	song := Song{}
	json.Unmarshal(songAsBytes, &song)
	song.Copyrights = append(song.Copyrights, copyright)

	songAsBytes, _ = json.Marshal(song)
	stub.PutState(songId, songAsBytes)

	rightHolder := RightHolder{}
	json.Unmarshal(rightHolderAsBytes, &rightHolder)
	rightHolder.Copyrights = append(rightHolder.Copyrights, copyright)

	rightHolderAsBytes, _ = json.Marshal(rightHolder)
	stub.PutState(rightHolderId, rightHolderAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) querySong(stub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	var (
		songId = args[0]
	)

	songAsBytes, err := stub.GetState(songId)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get Song for " + songId + "\"}"
		return shim.Error(jsonResp)
	}

	return shim.Success(songAsBytes)
}

func (s *SmartContract) queryRightHolder(stub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	var (
		rightHolderId = args[0]
	)

	rightHolderAsBytes, err := stub.GetState(rightHolderId)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get RightHolder for " + rightHolderId + "\"}"
		return shim.Error(jsonResp)
	}

	return shim.Success(songAsBytes)
}

func main() {
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error starting SmartContract: %s", err)
	}
}
