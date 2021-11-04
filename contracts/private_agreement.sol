// SPDX-License-Identifier: MIT

pragma solidity >=0.7.0 <0.9.0;

contract agreement {
    address private coordinator = 0x5B38Da6a701c568545dCfcB03FcB875f56beddC4;
    address payable party;
    bool status = false;
    string contact;
    bool key = false;
    uint256 aes_key;

    modifier isCoordinator() {
        require(msg.sender == coordinator, "Only the coordinator can call this function.");
        _;
    }

    modifier isParty() {
        require(msg.sender == party, "Only the party can call this function.");
        _;
    }

    modifier isInvolved() {
        require(msg.sender == party || msg.sender == coordinator, "Only the party can call this function.");
        _;
    }
    
    string link = "some link to a predetermined agreement template";
    string agreement_content;
    bool uploaded = false;
    
    //provide the coordinator with a contact number so an AES key can be sent to you
    //will need to make the input more uniform by providing a secure way to communicate in API/Web/Mobile App
    constructor(string memory _contact){
        contact = _contact;
        party = payable(address(msg.sender));
    }

    function getContact() public view isCoordinator returns(string memory){
        return contact;
    }

    function keySent() public isCoordinator{
        key = true;
    }

    function keyStatus() public view returns(bool){
        return key;
    }

    function uploadAgreement(string memory _agreement_content) public isParty{
        agreement_content = _agreement_content;
        uploaded = true;
    }

    function approve() public isCoordinator returns(bool){
        if(status == false){
            if(uploaded){
                status = true;
                return true;
            }
        }
        return false;
    }

    //will not allow the coordinator to destruct the agreement once approved
    function deny() public isCoordinator returns(bool){
        if(status == false){
            selfdestruct(party);
            return true;
        }
        return false;
    }

    //the party or the coordinator can choose to make their contract public by sharing the encryption key
    function makePublic(uint256 _aes_key) public isInvolved{
        aes_key = _aes_key;
    }

    function getAESKey() public view returns(uint256){
        return aes_key;
    }

    function getAgreement() public view returns(string memory) {
        return agreement_content;
    }

    function getAddress() public view returns(address){
        return address(this);
    }
}