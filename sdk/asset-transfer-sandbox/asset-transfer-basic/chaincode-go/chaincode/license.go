package chaincode

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type SmartContract struct {
	contractapi.Contract
}

type License struct {
	ID               string   `json:"ID"`
	Owner            string   `json:"owner"`
	Created          int      `json:"created"`
	Duration         int      `json:"duration"`
	Rules            []string `json:"rules"`
	LegalContractURL string   `json:"legalContractUrl"`
}

// Mint creates a new license and stores it in world state with given id
func (s *SmartContract) Mint(ctx contractapi.TransactionContextInterface, copyrightID string, to string, id string, description string) error {
	exists, err := s.AssetExists(ctx, copyrightID)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("Copyright with ID: %s does not exist", copyrightID)
	}

	var asset License
	err = json.Unmarshal([]byte(description), &asset)
	if err != nil {
		return fmt.Errorf("asset description format incorrect")
	}

	assetJSON, err := json.Marshal(asset)
	return ctx.GetStub().PutState(id, assetJSON)
}

// OwnerOf returns the holder of the license with given id
func (s *SmartContract) OwnerOf(ctx contractapi.TransactionContextInterface, id string) (string, error) {
	assetJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return "", fmt.Errorf("failed to read from world state: %v", err)
	}
	if assetJSON == nil {
		return "", fmt.Errorf("the asset %s does not exist", id)
	}

	var asset License
	err = json.Unmarshal(assetJSON, &asset)
	if err != nil {
		return "", err
	}

	return asset.Owner, nil
}

// AssetExists returns true when asset with given ID exists in world state
func (s *SmartContract) AssetExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
	assetJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}

	return assetJSON != nil, nil
}
