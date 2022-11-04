const web3 = require("web3")
const RLP = require('rlp');

// Configure
const validators = [
  
   {
     "consensusAddr": "0x10a2c15925873171a329b6cedb6905e4669e8fb2",
     "feeAddr": "0x10a2c15925873171a329b6cedb6905e4669e8fb2",
     "bscFeeAddr": "0x10a2c15925873171a329b6cedb6905e4669e8fb2",
     "votingPower": 0x0000000000000064,
   },
   {
     "consensusAddr": "0x0a86be0121df45de1a82094255cafc34f1119f53",
     "feeAddr": "0x0a86be0121df45de1a82094255cafc34f1119f53",
     "bscFeeAddr": "0x0a86be0121df45de1a82094255cafc34f1119f53",
     "votingPower": 0x0000000000000064,
   },
   {
     "consensusAddr": "0x1b6469567b4b9e642b164952830c0a294dde607e",
     "feeAddr": "0x1b6469567b4b9e642b164952830c0a294dde607e",
     "bscFeeAddr": "0x1b6469567b4b9e642b164952830c0a294dde607e",
     "votingPower": 0x0000000000000064,
   },
   {
     "consensusAddr": "0x3bd1fc7a8caf8a9a44794a8b8b7743cec53a98ca",
     "feeAddr": "0x3bd1fc7a8caf8a9a44794a8b8b7743cec53a98ca",
     "bscFeeAddr": "0x3bd1fc7a8caf8a9a44794a8b8b7743cec53a98ca",
     "votingPower": 0x0000000000000064,
   },
   {
     "consensusAddr": "0xc73087bab0ba1d327ba37ac8fdc8a12d3deb824a",
     "feeAddr": "0xc73087bab0ba1d327ba37ac8fdc8a12d3deb824a",
     "bscFeeAddr": "0xc73087bab0ba1d327ba37ac8fdc8a12d3deb824a",
     "votingPower": 0x0000000000000064,
   },
];

// ===============  Do not edit below ====
function generateExtradata(validators) {
  let extraVanity =Buffer.alloc(32);
  let validatorsBytes = extraDataSerialize(validators);
  let extraSeal =Buffer.alloc(65);
  return Buffer.concat([extraVanity,validatorsBytes,extraSeal]);
}

function extraDataSerialize(validators) {
  let n = validators.length;
  let arr = [];
  for (let i = 0;i<n;i++) {
    let validator = validators[i];
    arr.push(Buffer.from(web3.utils.hexToBytes(validator.consensusAddr)));
  }
  return Buffer.concat(arr);
}

function validatorUpdateRlpEncode(validators) {
  let n = validators.length;
  let vals = [];
  for (let i = 0;i<n;i++) {
    vals.push([
      validators[i].consensusAddr,
      validators[i].bscFeeAddr,
      validators[i].feeAddr,
      validators[i].votingPower,
    ]);
  }
  let pkg = [0x00, vals];
  return web3.utils.bytesToHex(RLP.encode(pkg));
}

extraValidatorBytes = generateExtradata(validators);
validatorSetBytes = validatorUpdateRlpEncode(validators);

exports = module.exports = {
  extraValidatorBytes: extraValidatorBytes,
  validatorSetBytes: validatorSetBytes,
}