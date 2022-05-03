package dclf

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/hyperledger/fabric-samples/asset-transfer-basic/chaincode-go/chaincode/mocks"

	hexutil "github.com/quan8/go-ethereum/common/hexutil"
	crypto "github.com/quan8/go-ethereum/crypto"
)

// SmartContract provides functions for managing an Asset
type SmartContract struct {
	contractapi.Contract
}

type transactionContext interface {
	contractapi.TransactionContextInterface
}

func (s *SmartContract) ProcessMessage(ctx contractapi.TransactionContextInterface, msg string, sig string, pubKey string) (string, error) {
	hash := crypto.Keccak256Hash([]byte(msg))

	signature, err := hexutil.Decode(sig)
	if err != nil {
		return "", err
	}

	sigPublicKey, err := hexutil.Decode(pubKey)
	if err != nil {
		return "", err
	}

	// verify that the message was signed using the private key corresponding to the public key
	signatureNoRecoverID := signature[:len(signature)-1] // remove recovery id
	verified := crypto.VerifySignature(sigPublicKey, hash.Bytes(), signatureNoRecoverID)
	if !verified {
		return "", err
	}

	// derive address from public key
	publicKeyAbbrevBytes, err := hexutil.Decode("0x" + pubKey[4:])
	if err != nil {
		return "", err
	}
	publicKeyHash := crypto.Keccak256(publicKeyAbbrevBytes)
	address := "0x" + hexutil.Encode(publicKeyHash)[26:]

	// parse message
	splitMsg := strings.Split(msg, "__")
	funcName := splitMsg[0]
	params := splitMsg[1:]

	// prepare params as byte slices
	// paramsBytes := make([][]byte, numParams+1)
	// paramsBytes = append(paramsBytes, []byte(funcName))
	// for i := 0; i < int(numParams); i++ {
	// 	paramsBytes = append(paramsBytes, []byte(params[i]))
	// }

	// call target functiom
	// res := ctx.GetStub().InvokeChaincode(chaincodeName, paramsBytes, ctx.GetStub().GetChannelID())
	// return res.GetMessage()

	return s.CallSmartContract(funcName, address, params)

}

func (s *SmartContract) CallSmartContract(address string, funcName string, params []string) (string, error) {
	ctx := &mocks.TransactionContext{}
	switch funcName {
	case "MintCopyright":
		err := s.MintCopyright(ctx, params[0], params[1])
		if err != nil {
			return "", err
		}
		return "Minted new copyright with ID: " + params[0], nil
	case "ReadCopyright":
		res, err := s.ReadCopyright(ctx, params[0])
		if err != nil {
			return "", err
		}

		resJSON, err := json.Marshal(res)
		if err != nil {
			return "", err
		}
		return string(resJSON), nil
	case "DeleteCopyright":
		err := s.DeleteCopyright(ctx, params[0])
		if err != nil {
			return "", err
		}
		return "Deleted copyright with ID: " + params[0], nil
	case "TransferCopyright":
		err := s.TransferCopyright(ctx, params[0], params[1], address)
		if err != nil {
			return "", err
		}
		return "Transfered copyright " + params[0] + " ownership from " + " address to " + params[1], nil
	case "MintLicense":
		err := s.MintLicense(ctx, params[0], address, params[1], params[2])
		if err != nil {
			return "", err
		}
		return "Minted new license with ID: " + params[1], nil
	case "OwnerOf":
		res, err := s.OwnerOf(ctx, params[0])
		if err != nil {
			return "", err
		}
		return res, nil
	case "InitiateTermination":
		err := s.InitiateTermination(ctx, params[0], address)
		if err != nil {
			return "", err
		}
		return "Initiated termination of license with ID: " + params[0], nil
	case "ApproveTermination":
		err := s.ApproveTermination(ctx, params[0], address)
		if err != nil {
			return "", err
		}
		return "Terminated license with ID: " + params[0], nil
	case "SetURL":
		err := s.SetURL(ctx, params[0], params[1])
		if err != nil {
			return "", err
		}
		return "Set contract URL on license with ID: " + params[0], nil
	case "ReadLicense":
		res, err := s.ReadLicense(ctx, params[0])
		if err != nil {
			return "", err
		}

		resJSON, err := json.Marshal(res)
		if err != nil {
			return "", err
		}

		return string(resJSON), nil
	}
	return "", fmt.Errorf("Invalid function name: %v", funcName)
}
