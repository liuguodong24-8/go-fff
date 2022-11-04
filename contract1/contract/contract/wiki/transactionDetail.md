## 获取区块信息

### 1) 请求地址

>api/transactions/:hash

### 2) 调用方式：HTTP GET

### 3) 接口描述：

* 获取区块历史列表

### 4) 请求参数:

#### GET参数:
| 字段名称      | 字段说明 | 类型     | 必填  | 备注     |
|-----------| ----  |--------|-----|--------|

### 5) 请求返回结果:

```
{
    "code": 0,
    "data": [
        {
            "hash": "0x7f07b92874b596d7d627f9895ea459fd5a68c5a0671851437ab48238edd4e25c",
            "nonce": 5,
            "to": "0x5d41c57b1c492bda3a256f0e48570d12dd626727",
            "value": 0,
            "gasPrice": 0,
            "gas": 21000,
            "accessList": null,
            "chainId": 0,
            "cost": 0,
            "data": "",
            "gasFeeCap": 0,
            "gasTipCap": 0,
            "protected": true
        },
        {
            "hash": "0x7f07b92874b596d7d627f9895ea459fd5a68c5a0671851437ab48238edd4e25c",
            "nonce": 5,
            "to": "0x5d41c57b1c492bda3a256f0e48570d12dd626727",
            "value": 0,
            "gasPrice": 0,
            "gas": 21000,
            "accessList": null,
            "chainId": 0,
            "cost": 0,
            "data": "",
            "gasFeeCap": 0,
            "gasTipCap": 0,
            "protected": true
        }
    ],
    "page": {
        "total": 3,
        "page": 1,
        "perPage": 2,
        "prev": 0,
        "next": 2,
        "totalPage": 2
    }
}
```

### 6) 请求返回结果参数说明:
| 字段名称 | 字段说明     | 类型    | 备注 |
|------|----------|-------|----  |
| code | 0成功 其他失败 | int   | - | 
| data |          | array | - |
