const Web3 = require("web3");
const utils = require('web3-utils');

async function test() {
  console.log(`get balance`)
}
test().then(async function () {
  const url = 'https://nodetest.3fchain.org';
  const web3 = new Web3(url)

  console.info(utils.sha3("Approval(address,address,uint256)"))

/*

  web3.sha3("Approval(address,address,uint256)").then(console.log)
*/

  web3.eth.getBalance("0x417Ef196A4ecEe6a971e8144633123030B468B14").then(console.log);
})