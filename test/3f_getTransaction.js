const Web3 = require("web3");
const base58   = require('base58-js');

async function test() {
  console.log(`3f get Transaction`)
}
test().then(async function () {
  const url = 'http://87.118.86.2:8488';
  const web3 = new Web3(url)
  new web3.eth.getTransaction('0x240968a62e361fb727e7b2aea7f70a1c20349e7ffeaf1f2831ba8d11fa55ecb4').then(console.log);
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