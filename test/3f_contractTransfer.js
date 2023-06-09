﻿const Web3 = require("web3");
const base58 = require('base58-js');

async function test() {
  console.log(`3f contract erc20 transfer`)
}
test().then(async function () {
  const abi=[
    {
      "inputs": [
        {
          "internalType": "string",
          "name": "name",
          "type": "string"
        },
        {
          "internalType": "string",
          "name": "symbol",
          "type": "string"
        },
        {
          "internalType": "address",
          "name": "initialAccount",
          "type": "address"
        },
        {
          "internalType": "uint256",
          "name": "initialBalance",
          "type": "uint256"
        }
      ],
      "stateMutability": "payable",
      "type": "constructor"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": true,
          "internalType": "address",
          "name": "owner",
          "type": "address"
        },
        {
          "indexed": true,
          "internalType": "address",
          "name": "spender",
          "type": "address"
        },
        {
          "indexed": false,
          "internalType": "uint256",
          "name": "value",
          "type": "uint256"
        }
      ],
      "name": "Approval",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": true,
          "internalType": "address",
          "name": "from",
          "type": "address"
        },
        {
          "indexed": true,
          "internalType": "address",
          "name": "to",
          "type": "address"
        },
        {
          "indexed": false,
          "internalType": "uint256",
          "name": "value",
          "type": "uint256"
        }
      ],
      "name": "Transfer",
      "type": "event"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "owner",
          "type": "address"
        },
        {
          "internalType": "address",
          "name": "spender",
          "type": "address"
        }
      ],
      "name": "allowance",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "spender",
          "type": "address"
        },
        {
          "internalType": "uint256",
          "name": "amount",
          "type": "uint256"
        }
      ],
      "name": "approve",
      "outputs": [
        {
          "internalType": "bool",
          "name": "",
          "type": "bool"
        }
      ],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "owner",
          "type": "address"
        },
        {
          "internalType": "address",
          "name": "spender",
          "type": "address"
        },
        {
          "internalType": "uint256",
          "name": "value",
          "type": "uint256"
        }
      ],
      "name": "approveInternal",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "account",
          "type": "address"
        }
      ],
      "name": "balanceOf",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "account",
          "type": "address"
        },
        {
          "internalType": "uint256",
          "name": "amount",
          "type": "uint256"
        }
      ],
      "name": "burn",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "decimals",
      "outputs": [
        {
          "internalType": "uint8",
          "name": "",
          "type": "uint8"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "spender",
          "type": "address"
        },
        {
          "internalType": "uint256",
          "name": "subtractedValue",
          "type": "uint256"
        }
      ],
      "name": "decreaseAllowance",
      "outputs": [
        {
          "internalType": "bool",
          "name": "",
          "type": "bool"
        }
      ],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "spender",
          "type": "address"
        },
        {
          "internalType": "uint256",
          "name": "addedValue",
          "type": "uint256"
        }
      ],
      "name": "increaseAllowance",
      "outputs": [
        {
          "internalType": "bool",
          "name": "",
          "type": "bool"
        }
      ],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "account",
          "type": "address"
        },
        {
          "internalType": "uint256",
          "name": "amount",
          "type": "uint256"
        }
      ],
      "name": "mint",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "name",
      "outputs": [
        {
          "internalType": "string",
          "name": "",
          "type": "string"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "symbol",
      "outputs": [
        {
          "internalType": "string",
          "name": "",
          "type": "string"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "totalSupply",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "to",
          "type": "address"
        },
        {
          "internalType": "uint256",
          "name": "amount",
          "type": "uint256"
        }
      ],
      "name": "transfer",
      "outputs": [
        {
          "internalType": "bool",
          "name": "",
          "type": "bool"
        }
      ],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "from",
          "type": "address"
        },
        {
          "internalType": "address",
          "name": "to",
          "type": "address"
        },
        {
          "internalType": "uint256",
          "name": "amount",
          "type": "uint256"
        }
      ],
      "name": "transferFrom",
      "outputs": [
        {
          "internalType": "bool",
          "name": "",
          "type": "bool"
        }
      ],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "from",
          "type": "address"
        },
        {
          "internalType": "address",
          "name": "to",
          "type": "address"
        },
        {
          "internalType": "uint256",
          "name": "value",
          "type": "uint256"
        }
      ],
      "name": "transferInternal",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    }
  ]
  const bytecode = `608060405260405162000f2538038062000f258339810160408190526200002691620001f4565b8383600362000036838262000315565b50600462000045828262000315565b5050506200005a82826200006460201b60201c565b5050505062000409565b6001600160a01b038216620000bf5760405162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f206164647265737300604482015260640160405180910390fd5b8060026000828254620000d39190620003e1565b90915550506001600160a01b038216600081815260208181526040808320805486019055518481527fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a35050565b505050565b634e487b7160e01b600052604160045260246000fd5b600082601f8301126200015757600080fd5b81516001600160401b03808211156200017457620001746200012f565b604051601f8301601f19908116603f011681019082821181831017156200019f576200019f6200012f565b81604052838152602092508683858801011115620001bc57600080fd5b600091505b83821015620001e05785820183015181830184015290820190620001c1565b600093810190920192909252949350505050565b600080600080608085870312156200020b57600080fd5b84516001600160401b03808211156200022357600080fd5b620002318883890162000145565b955060208701519150808211156200024857600080fd5b50620002578782880162000145565b604087015190945090506001600160a01b03811681146200027757600080fd5b6060959095015193969295505050565b600181811c908216806200029c57607f821691505b602082108103620002bd57634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156200012a57600081815260208120601f850160051c81016020861015620002ec5750805b601f850160051c820191505b818110156200030d57828155600101620002f8565b505050505050565b81516001600160401b038111156200033157620003316200012f565b620003498162000342845462000287565b84620002c3565b602080601f831160018114620003815760008415620003685750858301515b600019600386901b1c1916600185901b1785556200030d565b600085815260208120601f198616915b82811015620003b25788860151825594840194600190910190840162000391565b5085821015620003d15787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b808201808211156200040357634e487b7160e01b600052601160045260246000fd5b92915050565b610b0c80620004196000396000f3fe608060405234801561001057600080fd5b50600436106100f55760003560e01c806340c10f19116100975780639dc29fac116100665780639dc29fac146101ee578063a457c2d714610201578063a9059cbb14610214578063dd62ed3e1461022757600080fd5b806340c10f191461019757806356189cb4146101aa57806370a08231146101bd57806395d89b41146101e657600080fd5b8063222f5be0116100d3578063222f5be01461014d57806323b872dd14610162578063313ce56714610175578063395093511461018457600080fd5b806306fdde03146100fa578063095ea7b31461011857806318160ddd1461013b575b600080fd5b61010261023a565b60405161010f9190610956565b60405180910390f35b61012b6101263660046109c0565b6102cc565b604051901515815260200161010f565b6002545b60405190815260200161010f565b61016061015b3660046109ea565b6102e6565b005b61012b6101703660046109ea565b6102f6565b6040516012815260200161010f565b61012b6101923660046109c0565b61031a565b6101606101a53660046109c0565b61033c565b6101606101b83660046109ea565b61034a565b61013f6101cb366004610a26565b6001600160a01b031660009081526020819052604090205490565b610102610355565b6101606101fc3660046109c0565b610364565b61012b61020f3660046109c0565b61036e565b61012b6102223660046109c0565b6103ee565b61013f610235366004610a48565b6103fc565b60606003805461024990610a7b565b80601f016020809104026020016040519081016040528092919081815260200182805461027590610a7b565b80156102c25780601f10610297576101008083540402835291602001916102c2565b820191906000526020600020905b8154815290600101906020018083116102a557829003601f168201915b5050505050905090565b6000336102da818585610427565b60019150505b92915050565b6102f183838361054b565b505050565b6000336103048582856106f1565b61030f85858561054b565b506001949350505050565b6000336102da81858561032d83836103fc565b6103379190610ab5565b610427565b6103468282610765565b5050565b6102f1838383610427565b60606004805461024990610a7b565b6103468282610824565b6000338161037c82866103fc565b9050838110156103e15760405162461bcd60e51b815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f77604482015264207a65726f60d81b60648201526084015b60405180910390fd5b61030f8286868403610427565b6000336102da81858561054b565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b6001600160a01b0383166104895760405162461bcd60e51b8152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f206164646044820152637265737360e01b60648201526084016103d8565b6001600160a01b0382166104ea5760405162461bcd60e51b815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f206164647265604482015261737360f01b60648201526084016103d8565b6001600160a01b0383811660008181526001602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925910160405180910390a3505050565b6001600160a01b0383166105af5760405162461bcd60e51b815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f206164604482015264647265737360d81b60648201526084016103d8565b6001600160a01b0382166106115760405162461bcd60e51b815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201526265737360e81b60648201526084016103d8565b6001600160a01b038316600090815260208190526040902054818110156106895760405162461bcd60e51b815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e7420657863656564732062604482015265616c616e636560d01b60648201526084016103d8565b6001600160a01b03848116600081815260208181526040808320878703905593871680835291849020805487019055925185815290927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a35b50505050565b60006106fd84846103fc565b905060001981146106eb57818110156107585760405162461bcd60e51b815260206004820152601d60248201527f45524332303a20696e73756666696369656e7420616c6c6f77616e636500000060448201526064016103d8565b6106eb8484848403610427565b6001600160a01b0382166107bb5760405162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f20616464726573730060448201526064016103d8565b80600260008282546107cd9190610ab5565b90915550506001600160a01b038216600081815260208181526040808320805486019055518481527fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a35050565b6001600160a01b0382166108845760405162461bcd60e51b815260206004820152602160248201527f45524332303a206275726e2066726f6d20746865207a65726f206164647265736044820152607360f81b60648201526084016103d8565b6001600160a01b038216600090815260208190526040902054818110156108f85760405162461bcd60e51b815260206004820152602260248201527f45524332303a206275726e20616d6f756e7420657863656564732062616c616e604482015261636560f01b60648201526084016103d8565b6001600160a01b0383166000818152602081815260408083208686039055600280548790039055518581529192917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a3505050565b600060208083528351808285015260005b8181101561098357858101830151858201604001528201610967565b506000604082860101526040601f19601f8301168501019250505092915050565b80356001600160a01b03811681146109bb57600080fd5b919050565b600080604083850312156109d357600080fd5b6109dc836109a4565b946020939093013593505050565b6000806000606084860312156109ff57600080fd5b610a08846109a4565b9250610a16602085016109a4565b9150604084013590509250925092565b600060208284031215610a3857600080fd5b610a41826109a4565b9392505050565b60008060408385031215610a5b57600080fd5b610a64836109a4565b9150610a72602084016109a4565b90509250929050565b600181811c90821680610a8f57607f821691505b602082108103610aaf57634e487b7160e01b600052602260045260246000fd5b50919050565b808201808211156102e057634e487b7160e01b600052601160045260246000fdfea26469706673582212207ca09df96f6e3e119d33ae846077fc6ed8cc06a335725909cab9d2df5aecfbaa64736f6c63430008110033`
  const privKey = '0954ab7f6009370daf31a57c364b3251760dc9779058a40a532884043444f8ce';
  const address = 'FFF3nyQJTh2PohN4Z1gBHXZ8vTLaoJcVJRyteoYEC4ZUToAov8KWPGgbxV';


 /* const privKey = 'c862dfb65d03054a5d98cd1cb02d0697d132effe7d5b5046fb8416072cb68554';
  const address = 'FFF5ykma29oAcKrrXmQwJuaW2bfEynuaupEDzcsL6jn1JYXESVStzjzSa1';*/


 // const web3 = new Web3('https://nodetest.3fchain.org')
  const web3 = new Web3('https://node.3fchain.org')

  const contractAddress = 'FFF3SqyhwnE2CHYvQiRJNDoaBHHpUTUbfWnZmmSs7y8ebLHtykm5k1rnL9'
  const incrementer = new web3.eth.Contract(abi, to0xAddress(contractAddress));

 // const address = "FFF3bQsF3vtgptUfMwAbxUSU3RYB9MvK83CQ11yRPtPhTzGnRd54d3ceUB"
  //const toAddress = "FFF5ymJWkMYJzLpPFw5gyW14kHaEn61bDdGgfei2eUDX7FFMVCcv8biMjw"
  const toAddress = "FFF65VmNYJyy3p968Nz7DvktwBTpKTK8yzpHnp857a8uAHnzakDJYo6vxu"
  const encoded = incrementer.methods.transfer(to0xAddress(toAddress), '10000000000000000000').encodeABI();
  const createTransaction = await web3.eth.accounts.signTransaction(
    {
      from: to0xAddress(address),
      to: to0xAddress(contractAddress),
      data: encoded,
      gas: '1200000',
    },
    privKey
  );
  const createReceipt = await web3.eth.sendSignedTransaction(
    createTransaction.rawTransaction
  );
  console.log(`Tx successfull with hash: ${createReceipt.transactionHash}`);

})

var to0xAddress = function (value) {
  if (value.length === 40 || value.length === 42) {
    return value
  }
  return '0x' + byteToString(base58.base58_to_binary(value.substring(3)))
}

var toFFFAddress = function (value) {
  if (value.length === 58 || value.length === 58) {
    return value
  }
  return 'FFF' + base58.binary_to_base58(stringToByte(value.substring(2)))
}

var byteToString = function (arr) {
  if (typeof arr === 'string') {
    return arr;
  }
  var str = '',
    _arr = arr;
  for (var i = 0; i < _arr.length; i++) {
    var one = _arr[i].toString(2),
      v = one.match(/^1+?(?=0)/);
    if (v && one.length == 8) {
      var bytesLength = v[0].length;
      var store = _arr[i].toString(2).slice(7 - bytesLength);
      for (var st = 1; st < bytesLength; st++) {
        store += _arr[st + i].toString(2).slice(2);
      }
      str += String.fromCharCode(parseInt(store, 2));
      i += bytesLength - 1;
    } else {
      str += String.fromCharCode(_arr[i]);
    }
  }
  return str;
}

var stringToByte = function (str) {
  const bytes = new Array();
  let len, c;
  // eslint-disable-next-line prefer-const
  len = str.length;
  for (let i = 0; i < len; i++) {
    c = str.charCodeAt(i);
    if (c >= 0x010000 && c <= 0x10FFFF) {
      bytes.push(((c >> 18) & 0x07) | 0xF0);
      bytes.push(((c >> 12) & 0x3F) | 0x80);
      bytes.push(((c >> 6) & 0x3F) | 0x80);
      bytes.push((c & 0x3F) | 0x80);
    } else if (c >= 0x000800 && c <= 0x00FFFF) {
      bytes.push(((c >> 12) & 0x0F) | 0xE0);
      bytes.push(((c >> 6) & 0x3F) | 0x80);
      bytes.push((c & 0x3F) | 0x80);
    } else if (c >= 0x000080 && c <= 0x0007FF) {
      bytes.push(((c >> 6) & 0x1F) | 0xC0);
      bytes.push((c & 0x3F) | 0x80);
    } else {
      bytes.push(c & 0xFF);
    }
  }
  return bytes;
}