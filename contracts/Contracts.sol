// SPDX-License-Identifier: MIT

pragma solidity >=0.7.0 <0.9.0;

contract Copyright {
    //store the address of the coordinator
    address private coordinator;
    
    //check if the message sender is the coordinator
    modifier isCoordinator() {
        require(msg.sender == coordinator, "Only the coordinator can call this function.");
        _;
    }
    
    //information pertaining to the copyright claim
    string registrationNumber;
    uint registrationDate;
    string title;
    string[] claimants;
    string[] authors;
    string contact;
    
    //information pertaining to the work itself
    struct track {
        string title;
        string album;
        string genre;
        uint runtime;
    }
    
    track tr;
    
    constructor(string memory _registrationNumber, uint _registrationDate, string memory _title, string[] memory _claimants, string[] memory _authors, string memory _contact,
                    string memory _album, string memory _genre, uint _runtime) {
        coordinator = msg.sender;
        
        registrationNumber = _registrationNumber;
        registrationDate = _registrationDate;
        title = _title;
        claimants = _claimants;
        authors = _authors;
        contact = _contact;
        
        tr = track({title: _title, album: _album, genre: _genre, runtime: _runtime});
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
    
    function getClaimants() public view returns(string[] memory) {
        return claimants;
    }
    
    function addClaimant(string memory _claimant) public isCoordinator {
        claimants[claimants.length] = _claimant;
    }
    
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
    
    function getAuthors() public view returns(string[] memory) {
        return authors;
    }
    
    function getContact() public view returns(string memory) {
        return contact;
    }
    
    function getTrack() public view returns(track memory) {
        return tr;
    }
}