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
    //format: index {index} name {name} transfer_type {transfer_type} location {location} share {share}
    string[] claimants;
    //format: index {index} name {name} domicile {domicile} citizenship {count} {citizenship} authorship {count} {authorship}
    string[] authors;
    //format: index {index} name {name} cell {cell} location {location}
    string[] contacts;
    string[] link;
    
    //information pertaining to the work itself
    struct track {
        string title;
        string album;
        string genre;
        uint runtime;
    }
    track tr;
    
    constructor(string memory _registrationNumber, uint _registrationDate, string memory _title, string[] memory _claimants, string[] memory _authors, string[] memory _contact,
                    string memory _album, string memory _genre, uint _runtime, string[] memory _link) {
        //coordinator = msg.sender;
        coordinator = msg.sender;
        registrationNumber = _registrationNumber;
        registrationDate = _registrationDate;
        title = _title;
        claimants = _claimants;
        authors = _authors;
        contacts = _contact;
        link = _link;
        
        tr = track({title: _title, album: _album, genre: _genre, runtime: _runtime});
    }
    
    //TODO: Is it cheaper to slice the data and be able to return individual content in claimants, authors, contacts here or return whole thing (computation vs data size)

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
        claimants.push(_claimant);
        //claimants[claimants.length] = _claimant;
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

    //the commented line has some "payable" issue
    function addAuthors(string memory _authors) public isCoordinator {
        authors.push(_authors);
        //claimants[claimants.length] = _claimant;
    }
    
    function dropAuthors(uint _index) public isCoordinator returns(bool) {
        if (_index >= authors.length) {
            return false;
        }
        
        for (uint i = _index; i < authors.length - 1; i++) {
            authors[i] = authors[i + 1];
        }
        delete authors[authors.length - 1];
        return true;
    }


    function getContact() public view returns(string[] memory) {
        return contacts;
    }

    function addContact(string memory _contacts) public isCoordinator {
        contacts.push(_contacts);
        //claimants[claimants.length] = _claimant;
    }
    
    function dropContact(uint _index) public isCoordinator returns(bool) {
        if (_index >= contacts.length) {
            return false;
        }
        
        for (uint i = _index; i < contacts.length - 1; i++) {
            contacts[i] = contacts[i + 1];
        }
        delete contacts[contacts.length - 1];
        return true;
    }


    function getTrack() public view returns(track memory) {
        return tr;
    }


    function getLink() public view returns(string[] memory){
        return link;
    }

    function addLink(string memory _link) public isCoordinator {
        link.push(_link);
        //claimants[claimants.length] = _claimant;
    }
    
    function dropLink(uint _index) public isCoordinator returns(bool) {
        if (_index >= link.length) {
            return false;
        }
        
        for (uint i = _index; i < link.length - 1; i++) {
            link[i] = link[i + 1];
        }
        delete link[link.length - 1];
        return true;
    }

}