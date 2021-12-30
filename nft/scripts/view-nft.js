require('dotenv').config();
const API_URL = process.env.API_URL;
const PUBLIC_KEY = process.env.PUBLIC_KEY;
const PRIVATE_KEY = process.env.PRIVATE_KEY;

const { createAlchemyWeb3 } = require("@alch/alchemy-web3");
const web3 = createAlchemyWeb3(API_URL);

const contract = require("../artifacts/contracts/MyNFT.sol/MyNFT.json");
const contractAddress = "0xaC76Ac9995eb52AaDA526F53333942E49eC9A206";
const nftContract = new web3.eth.Contract(contract.abi, contractAddress);

async function viewNFT() {
	const ret = await nftContract.methods.totalSupply().call();
	
	//const ret = await web3.eth.call(tx);
	console.log(ret);

	const ret2 = await nftContract.methods.tokenURI(1).call();
	console.log(ret2);
}

viewNFT();
