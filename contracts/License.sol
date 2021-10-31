// SPDX-License-Identifier: MIT

pragma solidity >=0.7.0 <0.9.0;

contract License {
    address licensee;
    string licenseeName;
    address copyright;
    uint validFrom;
    uint validTo;
    
    constructor(address _licensee, string memory _licenseeName, address _copyright, uint _validFrom, uint _validTo) {
        licensee = _licensee;
        licenseeName = _licenseeName;
        copyright = _copyright;
        validFrom = _validFrom;
        validTo = _validTo;
    }
    
    function getLicensee() public view returns(address) {
        return licensee;
    }
    
    function getLicenseeName() public view returns(string memory) {
        return licenseeName;
    }
    
    function getCopyright() public view returns(address) {
        return copyright;
    }
    
    function getValidFrom() public view returns(uint) {
        return validFrom;
    }
    
    function getValidTo() public view returns(uint) {
        return validTo;
    }
   
}