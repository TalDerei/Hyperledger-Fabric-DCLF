// SPDX-License-Identifier: MIT

pragma solidity >=0.7.0 <0.9.0;

contract agreement {
    address private coordinator = 0x5B38Da6a701c568545dCfcB03FcB875f56beddC4;
    address payable party;
    bool status = false;

    modifier isCoordinator() {
        require(msg.sender == coordinator, "Only the coordinator can call this function.");
        _;
    }
    
    string link = "some link to a predetermined agreement template";
    string agreement_content;
    

    constructor(string memory _title, string memory _entity){
        agreement_content = string(abi.encodePacked("I ",_entity, " hereby agree with the agreement on", link, " regarding my origional creative content ", _title, "."));
        party = payable(address(msg.sender));
    }

    function approve() public isCoordinator returns(bool){
        if(status == false){
            status = true;
            return true;
        }
        return false;
    }

    function deny() public isCoordinator{
        selfdestruct(party);
    }

    function getAgreement() public view returns(string memory) {
        return agreement_content;
    }

    function getAddress() public view returns(address){
        return address(this);
    }
}