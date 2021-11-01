// SPDX-License-Identifier: MIT

pragma solidity >=0.7.0 <0.9.0;
// This is allow me to return struct after the smart contract is called
pragma experimental ABIEncoderV2;

contract Copyright {
    //store the address of the coordinator
    address private coordinator = 0xd27fbEC844d2c6dA0f14Ab3f4B95fbB48fa7bFE1;
    
    //check if the message sender is the coordinator
    //will this not allow anyone not the coordinator to call get functions? We might want to allow those kind of operations
    modifier isCoordinator() {
        require(msg.sender == coordinator, "Only the coordinator can call this function.");
        _;
    }
    
    //information pertaining to the copyright claim
    struct claimant{
        string name;
        string transfer_type;
        string location;
        int share;
    }

    struct author{
        string name;
        string domicile;
        string[] citizenship;
        string[] authorship;
    }
    struct contact{
        string name;
        string cell;
        string location;
    }
    string registrationNumber;
    uint registrationDate;
    uint registrationBlockTime;
    string title;
    claimant[] claimants;
    author[] authors;
    contact[] contacts;
    
    //information pertaining to the work itself
    struct track {
        string title;
        string album;
        string genre;
        uint runtime;
    }
    
    track tr;
    
    //I found out about ABIEncoderV2 after I finished this part, should I change the construtor input to struct, too?
    constructor(string memory _registrationNumber, uint _registrationDate, string memory _title, 
                    string[] memory climant_names, string[] memory climant_transfer_types, 
                    string[] memory climant_locations, int[] memory climant_shares, string[] memory author_names, 
                    string[] memory domicile, string[][] memory citzenships, string[][] memory authorships,
                    string[] memory contact_names, string[] memory contact_cells, string[] memory contact_location,
                    string memory _album, string memory _genre, uint _runtime) {
        coordinator = msg.sender;
        registrationNumber = _registrationNumber;
        registrationDate = _registrationDate;
        registrationBlockTime = block.timestamp;
        title = _title;
        if(climant_names.length == climant_transfer_types.length 
            && climant_names.length == climant_locations.length 
            && climant_names.length == climant_shares.length){
            if(author_names.length == domicile.length
                && author_names.length == citzenships.length
                && author_names.length == authorships.length){
                    if(contact_names.length == contact_cells.length
                        && contact_names.length == contact_location.length){
                        // only perform the action when all length match. Is there a better way to match this?
                        for(uint i = 0; i < climant_names.length; i++){
                            claimants[i] = claimant(climant_names[i], climant_transfer_types[i], climant_locations[i], climant_shares[i]);
                        }
                        for(uint i = 0; i < author_names.length; i++){
                            authors[i] = author(author_names[i], domicile[i], citzenships[i], authorships[i]);
                        }
                        for(uint i = 0; i < contact_names.length; i++){
                            contacts[i] = contact(contact_names[i], contact_cells[i], contact_location[i]);
                        }
                        tr = track({title: _title, album: _album, genre: _genre, runtime: _runtime});
                    }

            }
        }
    }
    
    function getRegistrationNumber() public view returns(string memory) {
        return registrationNumber;
    }
    
    function getRegistrationDate() public view returns(uint) {
        return registrationDate;
    }
    
    function getTitle() public view returns(string memory) {
        return title;
    }

    function getClaimants() public view returns(claimant[] memory) {
        return claimants;
    }
    
    function addClaimant(claimant memory _claimant) public isCoordinator {
        claimants[claimants.length] = _claimant;
    }
    
    // why should we drop by index? We might want to change it so that it drops by name
    function dropClaimant(uint _index) public isCoordinator returns(bool) {
        if (_index >= claimants.length) {
            return false;
        }
        
        for (uint i = _index; i < claimants.length - 1; i++) {
            claimants[i] = claimants[i + 1];
        }
        delete claimants[claimants.length - 1];
        return true;
    }
    
    function getAuthors() public view returns(author[] memory) {
        return authors;
    }
    
    function getContact() public view returns(contact[] memory) {
        return contacts;
    }
    
    function getTrack() public view returns(track memory) {
        return tr;
    }

}