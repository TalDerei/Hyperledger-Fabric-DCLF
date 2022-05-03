package dclf

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type PermissionedAddress struct {
	Address      string   `json:"address"`
	OwnedContent []string `json:"ownedContent"`
}

// AddPermissionedAddress adds a new newly-permissioned address to the world state
func (s *SmartContract) AddPermissionedAddress(ctx contractapi.TransactionContextInterface, addr string) error {
	address := PermissionedAddress{
		Address: addr,
	}

	addressJSON, err := json.Marshal(address)
	if err != nil {
		return err
	}

	err = ctx.GetStub().PutState(string(address.Address), addressJSON)
	if err != nil {
		return fmt.Errorf("failed to put to world state. %v", err)
	}

	return nil
}

// DeletePermissionedAddress removes a previously permissioned address from the world state (removing persmissions)
func (s *SmartContract) DeletePermissionedAddress(ctx contractapi.TransactionContextInterface, addr string) error {
	exists, err := s.PermissionedAddressExists(ctx, addr)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the address %s does not exist", addr)
	}

	return ctx.GetStub().DelState(string(addr))
}

// GetPermissionedAddress returns a permissioned address from world state
func (s *SmartContract) GetPermissionedAddress(ctx contractapi.TransactionContextInterface, addr string) (bool, error) {
	addressJSON, err := ctx.GetStub().GetState(addr)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}

	return addressJSON != nil, nil
}

// IsPermissionedAddress returns true if the passed address exists in world state, false if not
func (s *SmartContract) IsPermissionedAddress(ctx contractapi.TransactionContextInterface, addr string) bool {
	exists, _ := s.PermissionedAddressExists(ctx, addr)
	return exists
}

func (s *SmartContract) PermissionedAddressExists(ctx contractapi.TransactionContextInterface, addr string) (bool, error) {
	assetJSON, err := ctx.GetStub().GetState(addr)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}

	return assetJSON != nil, nil
}
