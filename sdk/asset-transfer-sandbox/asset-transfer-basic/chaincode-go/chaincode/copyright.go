package chaincode

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

// SmartContract provides functions for managing an Asset
type SmartContract struct {
	contractapi.Contract
}

type TrackInfo struct {
	Title   string   `json:"title"`
	Album   string   `json:"album"`
	Genres  []string `json:"genres"`
	Runtime int      `json:"runtime"`
	Authors []string `json:"authors"`
}

// Asset describes basic details of what makes up a simple asset
type Copyright struct {
	ID                   string     `json:"ID"`
	RegistrationNumber   string     `json:"registrationNumber"`
	RegistrationDate     string     `json:"registrationDate"`
	Owner                string     `json:"owner"`
	Track                *TrackInfo `json:"track`
	LegalContractURL     string     `json:"legalContractUrl"`
	AlternativeSourceURL string     `json:"alternativeSourceURL"`
}

// InitLedger adds a base set of assets to the ledger
func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	copyrights := []Copyright{
		{
			ID:                 "1",
			RegistrationNumber: "SR0000620204",
			RegistrationDate:   "2008-12-11",
			Owner:              "Sample_Owner",
			Track: &TrackInfo{
				Title: "Heartless",
				Album: "808s & Heartbreak",
				Genres: []string{
					"Hip-Hop/Rap",
					"Rythm and Blues",
				},
				Runtime: 211,
				Authors: []string{
					"Kanye West",
				},
			},
			LegalContractURL:     `https://cocatalog.loc.gov/cgi-bin/Pwebrecon.cgi?v1=1&ti=1,1&Search%5FArg=kanye%20west%20heartless&Search%5FCode=FT%2A&CNT=25&PID=emPBh99LW_3hUod-sKRk3rFLAV08y&SEQ=20220213164148&SID=5`,
			AlternativeSourceURL: "",
		},
	}

	for _, asset := range copyrights {
		assetJSON, err := json.Marshal(asset)
		if err != nil {
			return err
		}

		err = ctx.GetStub().PutState(asset.ID, assetJSON)
		if err != nil {
			return fmt.Errorf("failed to put to world state. %v", err)
		}
	}

	return nil
}

// Invoke routes incoming requests to this smart contract to the proper function
func (s *SmartContract) Invoke(APIStub shim.ChaincodeStubInterface) sc.Response {
	function, args := APIStub.GetFunctionAndParameters()
	switch function {
	case "Mint":
		return s.Mint(args[0], args[1], args[2])
	case "ReadAsset":
		res, _ := s.ReadAsset(args[0], args[1])
		return res
	case "UpdateRegistrationNumber":
		return s.UpdateRegistrationNumber(args[0], args[1], args[2])
	case "UpdateRegistrationDate":
		return s.UpdateRegistrationDate(args[0], args[1], args[2])
	case "AddAuthor":
		return s.AddAuthor(args[0], args[1], args[2])
	case "UpdateLegalContractURL":
		return s.UpdateLegalContractURL(args[0], args[1], args[2])
	case "UpdateAlternativeSourceURL":
		return s.UpdateAlternativeSourceURL(args[0], args[1], args[2])
	case "DeleteAsset":
		return s.DeleteAsset(args[0], args[1])
	case "AssetExists":
		res, _ := s.AssetExists(args[0], args[1])
		return res
	case "TransferAsset":
		return s.TransferAsset(args[0], args[1], args[2], args[3])
	case "GetAllAssets":
		res, _ := s.GetAllAssets(args[0])
		return res
	}
	return "Invalid function call"
}

// Mint creates a new copyright and stores it in world state with given id
func (s *SmartContract) Mint(ctx contractapi.TransactionContextInterface, id string, description string) error {
	exists, err := s.AssetExists(ctx, id)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("an asset with ID: %s already exists", id)
	}

	var asset Copyright
	err = json.Unmarshal([]byte(description), &asset)
	if err != nil {
		return fmt.Errorf("asset description format incorrect")
	}

	assetJSON, err := json.Marshal(asset)
	return ctx.GetStub().PutState(id, assetJSON)
}

// ReadAsset returns the asset stored in the world state with given id.
func (s *SmartContract) ReadAsset(ctx contractapi.TransactionContextInterface, id string) (*Copyright, error) {
	assetJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if assetJSON == nil {
		return nil, fmt.Errorf("the asset %s does not exist", id)
	}

	var asset Copyright
	err = json.Unmarshal(assetJSON, &asset)
	if err != nil {
		return nil, err
	}

	return &asset, nil
}

//SETTERS FOR COPYRIGHT FILEDS. TO BE REMOVED LATER
func (s *SmartContract) UpdateRegistrationNumber(ctx contractapi.TransactionContextInterface, id string, registrationNumber string) error {
	exists, err := s.AssetExists(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the asset %s does not exist", id)
	}

	asset, err := s.ReadAsset(ctx, id)
	if err != nil {
		return err
	}

	asset.RegistrationNumber = registrationNumber
	assetJSON, err := json.Marshal(asset)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, assetJSON)
}

func (s *SmartContract) UpdateRegistrationDate(ctx contractapi.TransactionContextInterface, id string, registrationDate string) error {
	exists, err := s.AssetExists(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the asset %s does not exist", id)
	}

	asset, err := s.ReadAsset(ctx, id)
	if err != nil {
		return err
	}

	asset.RegistrationDate = registrationDate
	assetJSON, err := json.Marshal(asset)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, assetJSON)
}

func (s *SmartContract) AddAuthor(ctx contractapi.TransactionContextInterface, id string, newAuthor string) error {
	exists, err := s.AssetExists(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the asset %s does not exist", id)
	}

	asset, err := s.ReadAsset(ctx, id)
	if err != nil {
		return err
	}

	asset.Track.Authors = append(asset.Track.Authors, newAuthor)
	assetJSON, err := json.Marshal(asset)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, assetJSON)
}

func (s *SmartContract) UpdateLegalContractURL(ctx contractapi.TransactionContextInterface, id string, legalContractURL string) error {
	exists, err := s.AssetExists(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the asset %s does not exist", id)
	}

	asset, err := s.ReadAsset(ctx, id)
	if err != nil {
		return err
	}

	asset.LegalContractURL = legalContractURL
	assetJSON, err := json.Marshal(asset)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, assetJSON)
}

func (s *SmartContract) UpdateAlternativeSourceURL(ctx contractapi.TransactionContextInterface, id string, alternativeSourceURL string) error {
	exists, err := s.AssetExists(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the asset %s does not exist", id)
	}

	asset, err := s.ReadAsset(ctx, id)
	if err != nil {
		return err
	}

	asset.AlternativeSourceURL = alternativeSourceURL
	assetJSON, err := json.Marshal(asset)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, assetJSON)
}

// DeleteAsset deletes an given asset from the world state.
func (s *SmartContract) DeleteAsset(ctx contractapi.TransactionContextInterface, id string) error {
	exists, err := s.AssetExists(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the asset %s does not exist", id)
	}

	return ctx.GetStub().DelState(id)
}

// AssetExists returns true when asset with given ID exists in world state
func (s *SmartContract) AssetExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
	assetJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}

	return assetJSON != nil, nil
}

// TransferAsset updates the owner field of asset with given id in world state.
func (s *SmartContract) TransferAsset(ctx contractapi.TransactionContextInterface, id string, newOwner string, callerAddress string) error {
	asset, err := s.ReadAsset(ctx, id)
	if err != nil {
		return err
	}

	if asset.Owner != callerAddress {
		return fmt.Errorf("Only the owner of this asset can transfer it.")
	}
	asset.Owner = newOwner
	assetJSON, err := json.Marshal(asset)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, assetJSON)
}

// GetAllAssets returns all assets found in world state
func (s *SmartContract) GetAllAssets(ctx contractapi.TransactionContextInterface) ([]*Copyright, error) {
	// range query with empty string for startKey and endKey does an
	// open-ended query of all assets in the chaincode namespace.
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var assets []*Copyright
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var asset Copyright
		err = json.Unmarshal(queryResponse.Value, &asset)
		if err != nil {
			return nil, err
		}
		assets = append(assets, &asset)
	}

	return assets, nil
}
