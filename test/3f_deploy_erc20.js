const Web3 = require("web3");
const base58   = require('base58-js');

async function test() {
  console.log(`3f deploying erc20`)
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
          "internalType": "string",
          "name": "issuerAuthorities",
          "type": "string"
        },
        {
          "internalType": "string",
          "name": "issuerName",
          "type": "string"
        },
        {
          "internalType": "string",
          "name": "description",
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
      "inputs": [],
      "name": "description",
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
      "inputs": [],
      "name": "issuerAuthorities",
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
      "name": "issuerName",
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
  const bytecode = `60806040526040516200104838038062001048833981016040819052620000269162000247565b86868686866003620000398682620003d2565b506004620000488582620003d2565b506005620000578482620003d2565b506006620000668382620003d2565b506007620000758282620003d2565b5050505050506200008d82826200009a60201b60201c565b50505050505050620004c6565b6001600160a01b038216620000f55760405162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f206164647265737300604482015260640160405180910390fd5b80600260008282546200010991906200049e565b90915550506001600160a01b038216600081815260208181526040808320805486019055518481527fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a35050565b505050565b634e487b7160e01b600052604160045260246000fd5b600082601f8301126200018d57600080fd5b81516001600160401b0380821115620001aa57620001aa62000165565b604051601f8301601f19908116603f01168101908282118183101715620001d557620001d562000165565b81604052838152602092508683858801011115620001f257600080fd5b600091505b83821015620002165785820183015181830184015290820190620001f7565b600093810190920192909252949350505050565b80516001600160a01b03811681146200024257600080fd5b919050565b600080600080600080600060e0888a0312156200026357600080fd5b87516001600160401b03808211156200027b57600080fd5b620002898b838c016200017b565b985060208a0151915080821115620002a057600080fd5b620002ae8b838c016200017b565b975060408a0151915080821115620002c557600080fd5b620002d38b838c016200017b565b965060608a0151915080821115620002ea57600080fd5b620002f88b838c016200017b565b955060808a01519150808211156200030f57600080fd5b506200031e8a828b016200017b565b9350506200032f60a089016200022a565b915060c0880151905092959891949750929550565b600181811c908216806200035957607f821691505b6020821081036200037a57634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156200016057600081815260208120601f850160051c81016020861015620003a95750805b601f850160051c820191505b81811015620003ca57828155600101620003b5565b505050505050565b81516001600160401b03811115620003ee57620003ee62000165565b6200040681620003ff845462000344565b8462000380565b602080601f8311600181146200043e5760008415620004255750858301515b600019600386901b1c1916600185901b178555620003ca565b600085815260208120601f198616915b828110156200046f578886015182559484019460019091019084016200044e565b50858210156200048e5787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b80820180821115620004c057634e487b7160e01b600052601160045260246000fd5b92915050565b610b7280620004d66000396000f3fe608060405234801561001057600080fd5b50600436106101165760003560e01c806340c10f19116100a257806395d89b411161007157806395d89b411461021f5780639dc29fac14610227578063a457c2d71461023a578063a9059cbb1461024d578063dd62ed3e1461026057600080fd5b806340c10f19146101c857806356189cb4146101db57806370a08231146101ee5780637284e4161461021757600080fd5b806323b872dd116100e957806323b872dd146101835780632e9868f81461019657806330d2b1721461019e578063313ce567146101a657806339509351146101b557600080fd5b806306fdde031461011b578063095ea7b31461013957806318160ddd1461015c578063222f5be01461016e575b600080fd5b610123610273565b60405161013091906109bc565b60405180910390f35b61014c610147366004610a26565b610305565b6040519015158152602001610130565b6002545b604051908152602001610130565b61018161017c366004610a50565b61031f565b005b61014c610191366004610a50565b61032f565b610123610353565b610123610362565b60405160128152602001610130565b61014c6101c3366004610a26565b610371565b6101816101d6366004610a26565b610393565b6101816101e9366004610a50565b6103a1565b6101606101fc366004610a8c565b6001600160a01b031660009081526020819052604090205490565b6101236103ac565b6101236103bb565b610181610235366004610a26565b6103ca565b61014c610248366004610a26565b6103d4565b61014c61025b366004610a26565b610454565b61016061026e366004610aae565b610462565b60606003805461028290610ae1565b80601f01602080910402602001604051908101604052809291908181526020018280546102ae90610ae1565b80156102fb5780601f106102d0576101008083540402835291602001916102fb565b820191906000526020600020905b8154815290600101906020018083116102de57829003601f168201915b5050505050905090565b60003361031381858561048d565b60019150505b92915050565b61032a8383836105b1565b505050565b60003361033d858285610757565b6103488585856105b1565b506001949350505050565b60606006805461028290610ae1565b60606005805461028290610ae1565b6000336103138185856103848383610462565b61038e9190610b1b565b61048d565b61039d82826107cb565b5050565b61032a83838361048d565b60606007805461028290610ae1565b60606004805461028290610ae1565b61039d828261088a565b600033816103e28286610462565b9050838110156104475760405162461bcd60e51b815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f77604482015264207a65726f60d81b60648201526084015b60405180910390fd5b610348828686840361048d565b6000336103138185856105b1565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b6001600160a01b0383166104ef5760405162461bcd60e51b8152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f206164646044820152637265737360e01b606482015260840161043e565b6001600160a01b0382166105505760405162461bcd60e51b815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f206164647265604482015261737360f01b606482015260840161043e565b6001600160a01b0383811660008181526001602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925910160405180910390a3505050565b6001600160a01b0383166106155760405162461bcd60e51b815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f206164604482015264647265737360d81b606482015260840161043e565b6001600160a01b0382166106775760405162461bcd60e51b815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201526265737360e81b606482015260840161043e565b6001600160a01b038316600090815260208190526040902054818110156106ef5760405162461bcd60e51b815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e7420657863656564732062604482015265616c616e636560d01b606482015260840161043e565b6001600160a01b03848116600081815260208181526040808320878703905593871680835291849020805487019055925185815290927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a35b50505050565b60006107638484610462565b9050600019811461075157818110156107be5760405162461bcd60e51b815260206004820152601d60248201527f45524332303a20696e73756666696369656e7420616c6c6f77616e6365000000604482015260640161043e565b610751848484840361048d565b6001600160a01b0382166108215760405162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f206164647265737300604482015260640161043e565b80600260008282546108339190610b1b565b90915550506001600160a01b038216600081815260208181526040808320805486019055518481527fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a35050565b6001600160a01b0382166108ea5760405162461bcd60e51b815260206004820152602160248201527f45524332303a206275726e2066726f6d20746865207a65726f206164647265736044820152607360f81b606482015260840161043e565b6001600160a01b0382166000908152602081905260409020548181101561095e5760405162461bcd60e51b815260206004820152602260248201527f45524332303a206275726e20616d6f756e7420657863656564732062616c616e604482015261636560f01b606482015260840161043e565b6001600160a01b0383166000818152602081815260408083208686039055600280548790039055518581529192917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a3505050565b600060208083528351808285015260005b818110156109e9578581018301518582016040015282016109cd565b506000604082860101526040601f19601f8301168501019250505092915050565b80356001600160a01b0381168114610a2157600080fd5b919050565b60008060408385031215610a3957600080fd5b610a4283610a0a565b946020939093013593505050565b600080600060608486031215610a6557600080fd5b610a6e84610a0a565b9250610a7c60208501610a0a565b9150604084013590509250925092565b600060208284031215610a9e57600080fd5b610aa782610a0a565b9392505050565b60008060408385031215610ac157600080fd5b610aca83610a0a565b9150610ad860208401610a0a565b90509250929050565b600181811c90821680610af557607f821691505b602082108103610b1557634e487b7160e01b600052602260045260246000fd5b50919050565b8082018082111561031957634e487b7160e01b600052601160045260246000fdfea2646970667358221220c9cce2b2d34d217eb9ce4ec61b3f085011eeb64b40f5bf5d79df834d3c11bcc864736f6c63430008110033`
  const privKey = '0954ab7f6009370daf31a57c364b3251760dc9779058a40a532884043444f8ce';
  const address = 'FFF3nyQJTh2PohN4Z1gBHXZ8vTLaoJcVJRyteoYEC4ZUToAov8KWPGgbxV';

  const args = ['Charity’s Crypto Currency', 'NGO', 'Finance Future Factory GmbH', 'OneWorld Humanity Foundation (NGO)','Gold , Gems , Jewellery, Real Estate, Mining, Nature Resources,Human Resources  and All kind value assets', to0xAddress(address), '10000000000000000000000000000'];

  const web3 = new Web3('https://node.3fchain.org');
  const incrementer = new web3.eth.Contract(abi);
  const incrementerTx = incrementer.deploy({ data: bytecode, arguments: args, });

  const createTransaction = await web3.eth.accounts.signTransaction({
    from: to0xAddress(address),
    data: incrementerTx.encodeABI(),
    gas: '1200000',
  },privKey);

  const createReceipt = await web3.eth.sendSignedTransaction(
    createTransaction.rawTransaction
  );

  console.log('Contract deployed at 0x address', to0xAddress(createReceipt.contractAddress));
  console.log('Contract deployed at 3f address', createReceipt.contractAddress);
})

var to0xAddress = function(value){
  if (value.length === 40 ||value.length === 42 ){
     return value
  }
  return '0x'+byteToString(base58.base58_to_binary(value.substring(3)))
}

var toFFFAddress =function(value){
  if (value.length === 58 || value.length === 60 ){
    return value
 }
  return 'FFF'+base58.binary_to_base58(stringToByte(value.substring(2)))
}

var byteToString= function(arr) {
  if(typeof arr === 'string') {
      return arr;
  }
  var str = '',
      _arr = arr;
  for(var i = 0; i < _arr.length; i++) {
      var one = _arr[i].toString(2),
          v = one.match(/^1+?(?=0)/);
      if(v && one.length == 8) {
          var bytesLength = v[0].length;
          var store = _arr[i].toString(2).slice(7 - bytesLength);
          for(var st = 1; st < bytesLength; st++) {
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

var stringToByte= function(str) {
  const bytes = new Array();
  let len, c;
  // eslint-disable-next-line prefer-const
  len = str.length;
  for (let i = 0; i < len; i++) {
    c = str.charCodeAt(i);
    if(c >= 0x010000 && c <= 0x10FFFF) {
      bytes.push(((c >> 18) & 0x07) | 0xF0);
      bytes.push(((c >> 12) & 0x3F) | 0x80);
      bytes.push(((c >> 6) & 0x3F) | 0x80);
      bytes.push((c & 0x3F) | 0x80);
    } else if(c >= 0x000800 && c <= 0x00FFFF) {
      bytes.push(((c >> 12) & 0x0F) | 0xE0);
      bytes.push(((c >> 6) & 0x3F) | 0x80);
      bytes.push((c & 0x3F) | 0x80);
    } else if(c >= 0x000080 && c <= 0x0007FF) {
      bytes.push(((c >> 6) & 0x1F) | 0xC0);
      bytes.push((c & 0x3F) | 0x80);
    } else {
      bytes.push(c & 0xFF);
    }
  }
  return bytes;
}
