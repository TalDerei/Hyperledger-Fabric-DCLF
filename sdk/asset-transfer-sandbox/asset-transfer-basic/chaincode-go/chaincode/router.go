package chaincode

import (
	"crypto"
	"crypto/rsa"
	"crypto/x509"
	"strconv"
	"strings"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func (s *SmartContract) ProcessMessage(ctx contractapi.TransactionContextInterface, msg string, sig []byte, msgHashSum []uint8, marshaledPublicKey []byte) string {
	// Unmarshal public key
	unmarshaledPublicKey, err := x509.ParsePKCS1PublicKey(marshaledPublicKey)

	// verify that the message was signed using the private key corresponding to the public key
	err = VerifyMessageSignature(msg, msgHashSum, sig, unmarshaledPublicKey)
	if err != nil {
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

func VerifyMessageSignature(msg string, msgHashSum []uint8, sig []byte, publicKey *rsa.PublicKey) error {
	return rsa.VerifyPSS(publicKey, crypto.SHA256, msgHashSum, sig, nil)
}
