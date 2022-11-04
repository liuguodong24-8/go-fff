## 上传nft

### 1) 请求地址

>api/mint-nft

### 2) 调用方式：HTTP POST

### 3) 接口描述：

* 上传nft

### 4) 请求参数:

#### GET参数:
| 字段名称          | 字段说明                                    | 类型     | 必填 | 备注   |
|---------------|-----------------------------------------|--------|----  |------|
| id            | nf的id                                   | int    | Y |  |
| nft_name      | NFT名字                                   | string | Y |  |
| nft_content   | NFT介绍                                   | string | Y |  |
| nft_image_url | NFT图片地址                                 | string | Y |  |
| nft_cert_list | nft证书地址                                 | array  | Y |  |
| type          | 类型  FFF = 1 ETH = 2 BSC = 3 | int    | Y |  |


### 5) 请求返回结果:

```
{
    "code": 0,
    "msg": "创建NFT哈希已创建"
}
```

### 6) 请求返回结果参数说明:
| 字段名称   | 字段说明                | 类型     | 备注 |
|--------|---------------------|--------|----  |
| code   | 0成功 1参数错误 2数据为空 3失败 | int    | - | 
| msg |                     | string | - |