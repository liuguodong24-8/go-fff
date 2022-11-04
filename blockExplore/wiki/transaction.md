## 获取交易列表

### 1) 请求地址

>api/transactions

### 2) 调用方式：HTTP GET

### 3) 接口描述：

* 获取交易列表

### 4) 请求参数:

#### GET参数:
| 字段名称      | 字段说明 | 类型     | 必填  | 备注     |
|-----------| ----  |--------|-----|--------|
| page      |  | int    | Y   | 当前页    | 
| page_size |  | int    | Y   | 分页大小   |
| to        |  | string | N   | 区块交易地址 |

### 5) 请求返回结果:

```
{
    {
    "code": 0,
    "data": [
        {
            "hash": "0xb21b1de8ea834b2a7fe3c35755f8d5664125f106da9cdf9a25fac3eea3099cc9",
            "nonce": 0,
            "to": "0x0000000000000000000000000000000000001000",
            "value": 0,
            "gasPrice": 0,
            "gas": 9223372036854775807,
            "blockHash": "0x4998d50f31deb10242937088b2f36269645504b98cfc15b186ad9ff77ef365b3",
            "blockNumber": 1,
            "input": "0xe1c7392a",
            "from": "0x85245585de1768e6d40b46c5b4d0c17785ab7a9e",
            "transactionIndex": 0
        },
        {
            "hash": "0x66caf29268bb5a5aa36bac358f8a7eaaf69b4c41a0a743d0e3819512c411a5d5",
            "nonce": 1,
            "to": "0x917a3fea65fc150ad7a5ed7cb0ee455eab7bf163",
            "value": 319535557742690304,
            "gasPrice": 1000000000,
            "gas": 21000,
            "blockHash": "0x398f4f3e5a2b84534413225e8cfdafcc5416c7ac9dc8262ad04931a6601789e2",
            "blockNumber": 419,
            "input": "0x",
            "from": "0x85245585de1768e6d40b46c5b4d0c17785ab7a9e",
            "transactionIndex": 0
        },
        {
            "hash": "0x4184a0d0e54d95ea351f0de6d46c83b2797f34fae5faf3fa837e06d36dc76541",
            "nonce": 2,
            "to": "0xab9a49c82642de2515fcf83e505e54f3e41a2767",
            "value": 6060067626849927168,
            "gasPrice": 1000000000,
            "gas": 21000,
            "blockHash": "0x3cf6e58252bed385a33b06971268db92bcf7ed62454d64a2e1fa1082b4ada900",
            "blockNumber": 440,
            "input": "0x",
            "from": "0x85245585de1768e6d40b46c5b4d0c17785ab7a9e",
            "transactionIndex": 0
        }
    ],
    "page": {
        "total": 10,
        "page": 1,
        "perPage": 3,
        "totalPage": 4
    }
}
```

### 6) 请求返回结果参数说明:
|  字段名称   | 字段说明 | 类型     | 备注 |
|  ----  | ----  |--------|----  |
| hash  | 交易哈希 | string | - | 
| nonce  |  本次交易之前发送方已经生成的交易数量| int    | - |
| blockHash  | 交易所在块的哈希，对于挂起块，该值为null | string | - |
| blockNumber  | 交易所在块的编号，对于挂起块，该值为null | int    | - |
| transactionIndex  |交易在块中的索引位置，挂起块该值为null  | int    | - |
| from  | 交易发送方地址 | string | - |
| to  |  交易接收方地址，对于合约创建交易，该值为null| int    | - |
| value  | 发送的以太数量，单位：wei | int    | - |
| gasPrice  | 发送方提供的gas价格，单位：wei | array  | - |
| gas  | 发送方提供的gas可用量 | int    | - |
| input  | 随交易发送的数据 | int    | - |

