package test

import (
	"fmt"
	"github.com/fff-chain/go-fff/crypto"
	"github.com/fff-chain/go-fff/global_config/wallet"
	"testing"
)

func TestTransfer(t *testing.T) {
	str := "Transfer(address,address,uint256)"
	fmt.Println(crypto.Keccak256Hash([]byte(str)))
	return

	wl := wallet.Init("https://nodetest.3fchain.org/", "http://127.0.0.1:8545")

	keyJson := `{"address":"FFF3WCZBKvZYxxPVKxzMLqNJWN6MwaZEhsaBUY72TT9GoBXdNZmZ3rRmzz","crypto":{"cipher":"aes-128-ctr","ciphertext":"abc1acba2f5933b53f17c16ba83df1bbf33d9541275b49e405d11b16001da297","cipherparams":{"iv":"774250d3fbc384786cb710cd93a9031b"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"fbee1d148b04eced5de6b4cffa340e8c3aa95fe7348a072ec4df34f540e0bfb8"},"mac":"8f82ae1f4fd144c1d1bb7b79be00cd4990b0c9152862bfe830a1cc12ddd22fc1"},"id":"01c87816-4b8b-40b9-89f7-ab02731f334c","version":3}`
	fmt.Println(wl.KeystoreToPrivateKey(keyJson, ""))

	return

	// 合约事件签名哈希值，对事件的字符做keccak散列运算
	/*str := "Transfer(address,address,uint256)"
	fmt.Println(crypto.Keccak256Hash([]byte(str)))
	return*/

	//wl := wallet.Init("http://87.118.86.85:28488", "http://127.0.0.1:8545")
	//wl := wallet.Init("https://nodetest.3fchain.org/", "http://127.0.0.1:8545")
	/*	wl := wallet.Init("https://nodetest.3fchain.org/", "http://127.0.0.1:8545")
		wl := wallet.Init("https://nodetest.3fchain.org/", "http://127.0.0.1:8545")*/

	/*h, _ := wl.CurrClient.BlockByNumber(context.TODO(), new(big.Int).SetInt64(1436054))

	fmt.Println(h.Transactions())*/

	/*hash := wl.Transfer("c862dfb65d03054a5d98cd1cb02d0697d132effe7d5b5046fb8416072cb68554", "FFF3nwSsLMMPgoQWHqWd5LNsR21x15iTnBCJv9MpGJcFB5WTnMch3xrv9P", global_config.EthToWei(1000), "", 0)
	log.Println(hash)*/

	//正式链
	/*wl := wallet.Init("https://node.3fchain.org/", "http://127.0.0.1:8545")

	hash := wl.Transfer("0954ab7f6009370daf31a57c364b3251760dc9779058a40a532884043444f8ce", "FFF65VmNYJyy3p968Nz7DvktwBTpKTK8yzpHnp857a8uAHnzakDJYo6vxu", global_config.EthToWei(10), "", 0)

	fmt.Println(hash)
	*/
	/*wl := wallet.Init("https://node.3fchain.org/", "http://127.0.0.1:8545")
	keyJson := `{"address":"FFF3nyQJTh2PohN4Z1gBHXZ8vTLaoJcVJRyteoYEC4ZUToAov8KWPGgbxV","crypto":{"cipher":"aes-128-ctr","ciphertext":"f56e388a255a6b01dc532dabe8c957b9d4d258b8d9c7b83fba45f1335328abbe","cipherparams":{"iv":"205abf014f4bb770804d5ce877c7e9ca"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"b36aab6b4193831daab5f8e44402c576a6751da2838c7ed4402ef3409e71e036"},"mac":"3ef890a81dfd5dc7502e17b041441dcefaa61ed58f6dc375b01fee9c382ddd61"},"id":"aa2d9c34-8ebc-4b64-bac3-ca239779c245","version":3}`
	fmt.Println(wl.KeystoreToPrivateKey(keyJson, "Cq9q6il8NCdl0Fnh"))*/

	//"0954ab7f6009370daf31a57c364b3251760dc9779058a40a532884043444f8ce"

	//c862dfb65d03054a5d98cd1cb02d0697d132effe7d5b5046fb8416072cb68554
	/*wallet.GetPublicAddressFromKey("6bcaedfb46c3ad907ac9eec52c1d9414bfb0e743b82adbe6c6752d0baab60920")
	fmt.Println(wallet.GetPublicAddressFromKey("6bcaedfb46c3ad907ac9eec52c1d9414bfb0e743b82adbe6c6752d0baab60920"))
	*/
}
