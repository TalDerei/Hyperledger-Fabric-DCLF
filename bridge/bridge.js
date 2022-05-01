const path = require('path');
const abiDecoder = require('abi-decoder');
// modules for fabric chaincode env
const { Fabric } = require('fabric-contract-api')
const shim = require('fabric-shim')
const { Gateway, Wallets } = require('fabric-network');
const FabricCAServices = require('fabric-ca-client');

//modules for chaincode application
const { buildCAClient, registerAndEnrollUser, enrollAdmin } = require('./utils/CAUtil.js');
const { buildCCPOrg1, buildWallet } = require('./utils/AppUtil.js');

// modules for EVM
const Web3 = require('web3')

const abi = require('./abi.json') // Copy the ABI json object from the contract into abi.json file
abiDecoder.addABI(abi);
const rpcURL = 'ws://127.0.0.1:7545' // RPC URL in ws (HTTP does not support subscribe)
const web3 = new Web3(rpcURL)
//For event LedgerCall(uint256 indexed id, address proposer, uint256[] values, bytes[] calldatas)
//Topic is: 0x7099c1be026061e9e1bede1b3fc6e4ef1f778a59c21d71358947920c16a401ea
const ContractAddress = '0x7aC3f2615C66135cC5796B6A11422B67eebDe839';
// In case there is a need to call the contract
// const contract = new web3.eth.Contract(abi, ContractAddress)

// Chaincode specific variables
// const channelName = 'mychannel';
// const chaincodeName = 'copyright'; // Change this to the current name of the deployed contract

var subscription = web3.eth.subscribe('logs',{
    address: ContractAddress
	    //topics: ['0x845a1778394c4eba39d9dcfe84e423a7eca678740f2ec16fea33c4d9fa6a171b']
}, function(error, result){
    if(error) console.log(error);
}).on("data", function(trxData){
	const result_data = trxData.data;
    console.log("Event received", result_data);
	console.log(result_data.length/64);
	var calldata = new Array(Math.floor((result_data.length/64)));
	for(let i = 0; i < Math.floor((result_data.length/64)); i++){
		calldata[i] = result_data.substring(2+i*64,2+(i+1)*64);
	}
	console.log(calldata);
	// event parameter check (LedgerCall from GovernorContract will have 4 parameters)

	//console.log("ABI decoded", abiDecoder.decodeLogs(trxData));
    // ************ Begin Fabric Application ************
    // try{
    //     //setting up test auth, need to rework this
	// 	const ccp = buildCCPOrg1();
	// 	const caClient = buildCAClient(FabricCAServices, ccp, 'ca.org1.example.com');
	// 	const wallet = await buildWallet(Wallets, walletPath);
	// 	await enrollAdmin(caClient, wallet, mspOrg1);
	// 	await registerAndEnrollUser(caClient, wallet, mspOrg1, org1UserId, 'org1.department1');
	// 	const gateway = new Gateway();
    //     try{
    //         await gateway.connect(ccp, {
	// 			wallet,
	// 			identity: org1UserId,
	// 			discovery: { enabled: true, asLocalhost: true } // using asLocalhost as this gateway is using a fabric network deployed locally
	// 		});
    //         // The actual process here with the gateway established

	// 		// Build a network instance based on the channel where the smart contract is deployed
	// 		const network = await gateway.getNetwork(channelName);
	// 		// Get the contract from the network.
	// 		const contract = network.getContract(chaincodeName);

    //         //TODO: Start calling the actual contract (which is yet to be implemented)

    //     }finally {
	// 		// Disconnect from the gateway when the application is closing
	// 		// This will close all connections to the network
	// 		gateway.disconnect();
	// 	}
    // } catch (error) {
	// 	console.error(`******** FAILED to run the application: ${error}`);
	// }
    // ************ End Fabric Application ************
})