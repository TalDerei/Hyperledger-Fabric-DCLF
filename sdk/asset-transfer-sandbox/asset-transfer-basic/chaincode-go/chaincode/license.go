package dclf

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type License struct {
	ID                        string     `json:"ID"`
	Copyright                 *Copyright `json:"copyright"`
	Owner                     string     `json:"owner"`
	Created                   int        `json:"created"`
	Duration                  int        `json:"duration"`
	Rules                     []string   `json:"rules"`
	LegalContractURL          string     `json:"legalContractUrl"`
	EarlyTerminationInitiator string     `json:"earlyTerminationInitiator"`
}

// Mint creates a new license and stores it in world state with given id
func (s *SmartContract) MintLicense(ctx contractapi.TransactionContextInterface, copyrightID string, to string, id string, description string) error {
	copyright, err := s.ReadCopyright(ctx, copyrightID)
	if err != nil {
		return fmt.Errorf("Copyright with ID: %v does not exist", copyrightID)
	}

	var asset License
	err = json.Unmarshal([]byte(description), &asset)
	if err != nil {
		return fmt.Errorf("asset description format incorrect")
	}

	asset.ID = id
	asset.Owner = to
	asset.Copyright = copyright

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

// InitiateTermination allows either the owner of the license or referenced copyright to propose the agreement be terminated
func (s *SmartContract) InitiateTermination(ctx contractapi.TransactionContextInterface, tokenId string, caller string) error {
	assetJSON, err := ctx.GetStub().GetState(tokenId)
	if err != nil {
		return fmt.Errorf("failed to read from world state: %v", err)
	}
	if assetJSON == nil {
		return fmt.Errorf("the asset %s does not exist", tokenId)
	}

	var asset License
	err = json.Unmarshal(assetJSON, &asset)
	if err != nil {
		return err
	}
	if asset.EarlyTerminationInitiator != "" {
		return fmt.Errorf("the asset %s has already had a termination initiated", tokenId)
	}
	copyright := *asset.Copyright
	if caller != asset.Owner || caller != copyright.Owner {
		return fmt.Errorf("user %s is not approved to modify this asset", caller)
	}

	asset.EarlyTerminationInitiator = caller
	assetJSON, err = json.Marshal(asset)
	if err != nil {
		return err
	}
	return ctx.GetStub().PutState(tokenId, assetJSON)
}

// ApproveTermination completes an approved termination and deletes the license
func (s *SmartContract) ApproveTermination(ctx contractapi.TransactionContextInterface, tokenId string, caller string) error {
	assetJSON, err := ctx.GetStub().GetState(tokenId)
	if err != nil {
		return fmt.Errorf("failed to read from world state: %v", err)
	}
	if assetJSON == nil {
		return fmt.Errorf("the asset %s does not exist", tokenId)
	}

	var asset License
	err = json.Unmarshal(assetJSON, &asset)
	if err != nil {
		return err
	}
	if asset.EarlyTerminationInitiator == "" {
		return fmt.Errorf("the asset %s has not had a termination initiated", tokenId)
	}
	copyright := *asset.Copyright
	if caller != asset.Owner || caller != copyright.Owner {
		return fmt.Errorf("user %s is not approved to modify this asset", caller)
	}
	if (caller == asset.Owner && asset.EarlyTerminationInitiator == asset.Owner) || (caller == copyright.Owner && asset.EarlyTerminationInitiator == copyright.Owner) {
		return fmt.Errorf("user %s initiated the termination", caller)
	}

	return ctx.GetStub().DelState(tokenId)
}

// SetURL acts as a one-time setter to bind a license to a legal contract
func (s *SmartContract) SetURL(ctx contractapi.TransactionContextInterface, tokenId string, url string) error {
	assetJSON, err := ctx.GetStub().GetState(tokenId)
	if err != nil {
		return fmt.Errorf("failed to read from world state: %v", err)
	}
	if assetJSON == nil {
		return fmt.Errorf("the asset %s does not exist", tokenId)
	}

	var asset License
	err = json.Unmarshal(assetJSON, &asset)
	if err != nil {
		return err
	}
	if asset.LegalContractURL != "" {
		return fmt.Errorf("License with ID: %v already has contract URL: %s", err, asset.LegalContractURL)
	}

	asset.LegalContractURL = url
	assetJSON, err = json.Marshal(asset)
	if err != nil {
		return err
	}
	return ctx.GetStub().PutState(tokenId, assetJSON)
}

// ReadAsset returns the asset stored in the world state with given id.
func (s *SmartContract) ReadLicense(ctx contractapi.TransactionContextInterface, id string) (*License, error) {
	assetJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if assetJSON == nil {
		return nil, fmt.Errorf("the asset %s does not exist", id)
	}

	var asset License
	err = json.Unmarshal(assetJSON, &asset)
	if err != nil {
		return nil, err
	}

	return &asset, nil
}

// AssetExists returns true when asset with given ID exists in world state
func (s *SmartContract) LicenseExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
	assetJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}

	return assetJSON != nil, nil
}
