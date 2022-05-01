package chaincode

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type Retriever struct {
	CurveParams *elliptic.CurveParams `json:"Curve"`
	MyX         string                `json:"X"`
	MyY         string                `json:"Y"`
}

func (s *SmartContract) ProcessMessage(ctx contractapi.TransactionContextInterface, msg string, sig []byte, msgHashSum []byte, marshaledPublicKey []byte) string {
	// Hash message and compare to hash passed from client
	testMsgHash := sha256.Sum256([]byte(msg))
	if bytes.Compare(msgHashSum, testMsgHash[:]) != 0 {
		return "Invalid message hash."
	}

	// Use retriever struct to unmarshal json, then copy fields to EC public key struct
	retriever := new(Retriever)

	err := json.Unmarshal(marshaledPublicKey, &retriever)
	if err != nil {
		return "Unmarshal failed."
	}

	var unmarshaledECPublicKey ecdsa.PublicKey
	unmarshaledECPublicKey.Curve = retriever.CurveParams
	newX := new(big.Int)
	newX, ok := newX.SetString(retriever.MyX, 10)
	if !ok {
		fmt.Println("SetString X failed")
	}
	newY := new(big.Int)
	newY, ok = newY.SetString(retriever.MyY, 10)
	if !ok {
		fmt.Println("SetString Y failed")
	}
	unmarshaledECPublicKey.X = newX
	unmarshaledECPublicKey.Y = newY

	// verify that the message was signed using the private key corresponding to the public key
	valid := ecdsa.VerifyASN1(&unmarshaledECPublicKey, msgHashSum[:], sig)
	if !valid {
		return "Message singature is not valid."
	}

	// parse message
	splitMsg := strings.Split(msg, "_")
	chaincodeName := splitMsg[0]
	funcName := splitMsg[1]
	numParams, _ := strconv.ParseInt(splitMsg[2])
	params := splitMsg[3:]

	// prepare params as byte slices
	paramsBytes := make([][]byte, numParams+1)
	paramsBytes = append(paramsBytes, []byte(funcName))
	for i := 0; i < int(numParams); i++ {
		paramsBytes = append(paramsBytes, []byte(params[i]))
	}

	// call chaincode
	res := ctx.GetStub().InvokeChaincode(chaincodeName, paramsBytes, ctx.GetStub().GetChannelID())
	return res.GetMessage()
}
