package chaincode

import (
	"strconv"
	"strings"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"

	hexutil "github.com/quan8/go-ethereum/common/hexutil"
	crypto "github.com/quan8/go-ethereum/crypto"
)

func (s *SmartContract) ProcessMessage(ctx contractapi.TransactionContextInterface, msg string, sig string, pubKey string) string {
	hash := crypto.Keccak256Hash(msg)

	signature, err := hexutil.Decode(sig)
	if err != nil {
		return "Error decoding signature"
	}

	sigPublicKey, err := hexutil.Decode(pubKey)
	if err != nil {
		return "Error decoding public key"
	}

	// verify that the message was signed using the private key corresponding to the public key
	signatureNoRecoverID := signature[:len(signature)-1] // remove recovery id
	verified := crypto.VerifySignature(sigPublicKey, hash.Bytes(), signatureNoRecoverID)
	if !verified {
		return "Invalid message signature"
	}

	// parse message
	splitMsg := strings.Split(msg, "_")
	chaincodeName := splitMsg[0]
	funcName := splitMsg[1]
	numParams, _ := strconv.ParseInt(splitMsg[2], 0, 64)
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
