# GeneralApp

通用链码调用sdk服务

# 配置文件编写示例
```yaml
# 系统配置
system:
# 这里修改端口的话需要修改dockerfile expose配置
  addr: 0.0.0.0:8001

# 链码sdk初始化配置
chaincodes:
# 每一项代表链接池中的一个链码sdk链接,channel表示链码所在通道名称
  - channel: mychannel
#   链码名称
    chaincode: general
    # 链码公钥x,y坐标16进制字符串
    public:
      publicKeyX: 0ce2fa6e66521155f780573beb0e5f18d0aeea6b9a145f54e5c8c442efd15ecf
      publicKeyY: fa332850bffd6e06cbbd6e29ac851fe12da302c74550c3d75e24db54a2a1fdd7
    # 链码私钥16进制字符串
    private:
      privateKey: 0b15a775077e438bce6ebcb7b30c3e61d9909ee861568723661d4728ee701068
    # 组织身份信息文件夹
    organizationsDir: organizations
    # 钱包id可以链码名称配置一致
    walletLabel: general
    # connect链接文件目录
    ccpPath: organizations/peerOrganizations/org1.example.com/connection-org1.yaml
    # msp身份信息目录
    cerdPath: organizations/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
  - channel: mychannel
    chaincode: general2
    public:
      publicKeyX: 0ce2fa6e66521155f780573beb0e5f18d0aeea6b9a145f54e5c8c442efd15ecf
      publicKeyY: fa332850bffd6e06cbbd6e29ac851fe12da302c74550c3d75e24db54a2a1fdd7
    private:
      privateKey: 0b15a775077e438bce6ebcb7b30c3e61d9909ee861568723661d4728ee701068
    organizationsDir: organizations
    walletLabel: general2
    # connect链接文件目录
    ccpPath: organizations/peerOrganizations/org1.example.com/connection-org1.yaml
    # msp身份信息目录
    cerdPath: organizations/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
```

# 部署
```shell
// 编译文件
CGO_ENABLED=0 go build -o generalapp .
// 编译docker镜像
docker build -t generalapp:v0.0.1 .

// 启动
// 启动时端口映射需要与dockerfile中对应
// 当有多个organizations身份信息文件夹时需要在启动时映射进容器,并且要与config.yaml文件中organizationsDir配置路径名称一致
docker run -itd --restart=always --name genralapp --net=host -p 8001:8001 -v $(pwd)/organizations:/home/generalapp/organizations -v $(pwd)/config.yaml:/home/generalapp/config.yaml generalapp:v0.0.1
```

# 接口调用示例

---

**1\. initChaincode**
###### 接口功能
> 初始化链码

###### URL
> http://127.0.0.1:8001/initChaincode

###### 支持格式
> JSON

###### HTTP请求方式
> POST

###### 请求参数
|参数|必选|类型|说明|example|
| :--- | :--- | :--- | :--- | :--- |
|chaincodeName|ture|string|调用链码名称|notary|
|data|true|string|链码请求数据sm2加密16进制字符串|048502008185906f51be4d64c75db8424e30ea680ad875e8fa11e53588549ad07c1bc6146e0d79f12d99fd49707d608757f4b625ea6f708ce8adcb7530be550eb921f14a37bf3001eb85a9c61860fbb2cfdcb59ae74fc46f56c74ad320b15378e132ee0b968a783bcf7f13e41650ea3cb1e5a93bf8246e02|

###### 链码请求json数据
|参数|必选|类型|说明|example|
| :--- | :--- | :--- | :--- | :--- |
|docType|true|string|couchdb文档类型|notary|

###### 链码请求数据未加密前示例
```json
{
  "docType": "notary"
}
```

###### 返回字段解密后结构
|返回字段|字段类型|说明                              |
|:---:   |:---:|:---:|
|code   |int    |请求状态 0：失败；1：成功   |
|msg  |string | 请求信息                      |
|data |string |链码返回数据                         |

###### 接口示例
``` json
041ced976b5cf99f5482bd3d7d34ada224fe930bd07537e5001b66a6dad845ce0158b070401a64f05daf972c8169611690b02d41b291a2ebfac4537bd04a8d6a8e71985c4e812b989502d1f035e806ad4e00474612dc326851509e60e9aafa0b5cd556050e742a55a30cf369ddb86e213d31c6764b337e5d17de8705a8664611a95545d293fb5f
```

---

**2\. add**
###### 接口功能
> 添加数据

###### URL
> http://127.0.0.1:8001/add

###### 支持格式
> JSON

###### HTTP请求方式
> POST

###### 请求参数
|参数|必选|类型|说明|example|
| :--- | :--- | :--- | :--- | :--- |
|chaincodeName|ture|string|调用链码名称|notary|
|data|true|string|链码请求数据sm2加密16进制字符串|048502008185906f51be4d64c75db8424e30ea680ad875e8fa11e53588549ad07c1bc6146e0d79f12d99fd49707d608757f4b625ea6f708ce8adcb7530be550eb921f14a37bf3001eb85a9c61860fbb2cfdcb59ae74fc46f56c74ad320b15378e132ee0b968a783bcf7f13e41650ea3cb1e5a93bf8246e02|

###### 链码请求json数据
|参数|必选|类型|说明|example|
| :--- | :--- | :--- | :--- | :--- |
|keys|true|[]string|参数字符串数组|["144","C402022072640"]|
|content|true|string|数据json字符串|{\"notaryOfficeId\":40,\"serviceType\":3,\"notarizationNumber\":\"KFGZS003\",\"payTime\":\"2021-12-12 10:17:07\"}|

###### 链码请求数据未加密前示例
```json
{
  "keys": "[\"144\",\"C402022072640\"]",
  "content": "{\"notaryOfficeId\":40,\"serviceType\":3,\"notarizationNumber\":\"KFGZS003\",\"payTime\":\"2021-12-12 10:17:07\"}"
}
```

###### 返回字段解密后结构
|返回字段|字段类型|说明                              |
|:---:   |:---:|:---:|
|code   |int    |请求状态 0：失败；1：成功   |
|msg  |string | 请求信息                      |
|data |string |链码返回txId                       |

###### 接口示例
``` json
0406a2573d7424d1118278d03711fd6874a4459ff2ad349d1a5939b9e5654b296b920b4408aef5506b07e35873b08226aa77062e1e7f0211d01de093b999965e67c10f3ff1e1b93afe44bcfe6b2379cc41e46e270b7023d0e40b333ff8d6fb398a6bfe266d4787f5e30a91c78af9ac03dc9e5e92b3d2dda6a36bf1394cc16e44cab3cf40524ea3
```

---

**3\. update**
###### 接口功能
> 更新数据

###### URL
> http://127.0.0.1:8001/update

###### 支持格式
> JSON

###### HTTP请求方式
> POST

###### 请求参数
|参数|必选|类型|说明|example|
| :--- | :--- | :--- | :--- | :--- |
|chaincodeName|ture|string|调用链码名称|notary|
|data|true|string|链码请求数据sm2加密16进制字符串|048502008185906f51be4d64c75db8424e30ea680ad875e8fa11e53588549ad07c1bc6146e0d79f12d99fd49707d608757f4b625ea6f708ce8adcb7530be550eb921f14a37bf3001eb85a9c61860fbb2cfdcb59ae74fc46f56c74ad320b15378e132ee0b968a783bcf7f13e41650ea3cb1e5a93bf8246e02|

###### 链码请求json数据
|参数|必选|类型|说明|example|
| :--- | :--- | :--- | :--- | :--- |
|keys|true|[]string|参数字符串数组|["144","C402022072640"]|
|content|true|string|数据json字符串|{\"notaryOfficeId\":40,\"serviceType\":3,\"notarizationNumber\":\"KFGZS003\",\"payTime\":\"2021-12-12 10:17:07\"}|

###### 链码请求数据未加密前示例
```json
{
  "keys": "[\"144\",\"C402022072640\"]",
  "content": "{\"notaryOfficeId\":40,\"serviceType\":12312,\"notarizationNumber\":\"KFGZS003\",\"payTime\":\"2021-12-12 10:17:07\"}"
}
```

###### 返回字段解密后结构
|返回字段|字段类型|说明                              |
|:---:   |:---:|:---:|
|code   |int    |请求状态 0：失败；1：成功   |
|msg  |string | 请求信息                      |
|data |string |链码返回txId                         |

###### 接口示例
``` json
0406a2573d7424d1118278d03711fd6874a4459ff2ad349d1a5939b9e5654b296b920b4408aef5506b07e35873b08226aa77062e1e7f0211d01de093b999965e67c10f3ff1e1b93afe44bcfe6b2379cc41e46e270b7023d0e40b333ff8d6fb398a6bfe266d4787f5e30a91c78af9ac03dc9e5e92b3d2dda6a36bf1394cc16e44cab3cf40524ea3
```

---

**4\. delete**
###### 接口功能
> 删除数据

###### URL
> http://127.0.0.1:8001/delete

###### 支持格式
> JSON

###### HTTP请求方式
> POST

###### 请求参数
|参数|必选|类型|说明|example|
| :--- | :--- | :--- | :--- | :--- |
|chaincodeName|ture|string|调用链码名称|notary|
|data|true|string|链码请求数据sm2加密16进制字符串|048502008185906f51be4d64c75db8424e30ea680ad875e8fa11e53588549ad07c1bc6146e0d79f12d99fd49707d608757f4b625ea6f708ce8adcb7530be550eb921f14a37bf3001eb85a9c61860fbb2cfdcb59ae74fc46f56c74ad320b15378e132ee0b968a783bcf7f13e41650ea3cb1e5a93bf8246e02|

###### 链码请求json数据
|参数|必选|类型|说明|example|
| :--- | :--- | :--- | :--- | :--- |
|keys|true|[]string|参数字符串数组|["144","C402022072640"]|

###### 链码请求数据未加密前示例
```json
{
  "keys": "[\"144\",\"C402022072640\"]"
}
```

###### 返回字段解密后解构
|返回字段|字段类型|说明                              |
|:---:   |:---:|:---:|
|code   |int    |请求状态 0：失败；1：成功   |
|msg  |string | 请求信息                      |
|data |string |链码返回txId                       |

###### 接口示例
``` json
0406a2573d7424d1118278d03711fd6874a4459ff2ad349d1a5939b9e5654b296b920b4408aef5506b07e35873b08226aa77062e1e7f0211d01de093b999965e67c10f3ff1e1b93afe44bcfe6b2379cc41e46e270b7023d0e40b333ff8d6fb398a6bfe266d4787f5e30a91c78af9ac03dc9e5e92b3d2dda6a36bf1394cc16e44cab3cf40524ea3
```

---

**5\. query**
###### 接口功能
> 根据keys查询单条数据

###### URL
> http://127.0.0.1:8001/query

###### 支持格式
> JSON

###### HTTP请求方式
> POST

###### 请求参数
|参数|必选|类型|说明|example|
| :--- | :--- | :--- | :--- | :--- |
|chaincodeName|ture|string|调用链码名称|notary|
|data|true|string|链码请求数据sm2加密16进制字符串|048502008185906f51be4d64c75db8424e30ea680ad875e8fa11e53588549ad07c1bc6146e0d79f12d99fd49707d608757f4b625ea6f708ce8adcb7530be550eb921f14a37bf3001eb85a9c61860fbb2cfdcb59ae74fc46f56c74ad320b15378e132ee0b968a783bcf7f13e41650ea3cb1e5a93bf8246e02|

###### 链码请求json数据
|参数|必选|类型|说明|example|
| :--- | :--- | :--- | :--- | :--- |
|keys|true|[]string|参数字符串数组|["144","C402022072640"]|

###### 链码请求数据未加密前示例
```json
{
  "keys": "[\"144\",\"C402022072640\"]"
}
```

###### 返回字段解密后解构
|返回字段|字段类型|说明                              |
|:---:   |:---:|:---:|
|code   |int    |请求状态 0：失败；1：成功   |
|msg  |string | 请求信息                      |
|data |string |链码返回数据|

###### 接口示例
``` json
044d92d3ed19bfdccc77ee2275a441a94c05b5beeaf73faa31ceafc083428709ed583507fea0d3d3a7aa00e1bedd6b200dff0a61f83d9415f5076a92a98f5617786255f0913d6248d5383cb67be11710bb7bdb6169630605d65548f6a38883c320df0704ed41c6fd9d860391c8d55f70d01cdb4f9f1c5c84ff03d3fd94e28886e1dee3491f6cdcb559a9402ffc482de89670878a84033da5639853a0dad34c69f5366674a0b0f4cc3878f03f82e9b2b4947f90b57e48d1544416e58b352da8d8439d86214471ccaa994dc6e528b7166f628ebf5545a99eef1ad8fb60a6ab7aed5c70d9811391f21490db8b1732f5f6915d91b88254cb
```

---

**6\. queryAll**
###### 接口功能
> 根据keys查询所有数据

###### URL
> http://127.0.0.1:8001/queryAll

###### 支持格式
> JSON

###### HTTP请求方式
> POST

###### 请求参数
|参数|必选|类型|说明|example|
| :--- | :--- | :--- | :--- | :--- |
|chaincodeName|ture|string|调用链码名称|notary|
|data|true|string|链码请求数据sm2加密16进制字符串|048502008185906f51be4d64c75db8424e30ea680ad875e8fa11e53588549ad07c1bc6146e0d79f12d99fd49707d608757f4b625ea6f708ce8adcb7530be550eb921f14a37bf3001eb85a9c61860fbb2cfdcb59ae74fc46f56c74ad320b15378e132ee0b968a783bcf7f13e41650ea3cb1e5a93bf8246e02|

###### 链码请求json数据
|参数|必选|类型|说明|example|
| :--- | :--- | :--- | :--- | :--- |
|keys|true|[]string|参数字符串数组|["144","C402022072640"]|

###### 链码请求数据未加密前示例
```json
{
  "keys": "[]"
}
```

###### 返回字段解密后格式
|返回字段|字段类型|说明                              |
|:---:   |:---:|:---:|
|code   |int    |请求状态 0：失败；1：成功   |
|msg  |string | 请求信息                      |
|data |string |链码返回数据|

###### 接口示例
``` json
0482f84b3f6361eaf0330002a7a8e642b1849e8896930ef16c7d84f977be95e1d613a082033034a2639b54fd58c4d89f67585ce7c9224543bd69f593383811a0c32688ff94a60d89c08a312f56f6fe33d19783668b22374b42393532dbe3eab42d3ff45bda3f306697d9dd186a08df293f75b4aeb8844fb7fc0c4b392034401acce99cc5ba1bf6aff74fb698a484fb7806e71b5ab2e3427762da6946eaa3a41fc24e9e84d25e47f8bfbfdf7c8a0c8f25274b1cd46233728b1e30ebdfb8ad641325c417661ec5ab9c29ded7c5adc20a5ec838f1357e704dc92abd2f7a118d024338e785fe8135a1e16969d7dbbeee422c711f0daf7e4cee5a677b4e208310e1867db226a0b03e7b4c7aa065547340b411e94f2a1afa2807739d9a0c9260a9aca6ddb2f302fbff1e9b5a511e4a4acc935119e08c807cd800de6aff5d34b03f6e0bf3f58f21ce3147071e433fa51002fd51921e492ac2567c6a7e85da43c59c3cf72a3fd7f62941c0d6334f7cfd87f285fd1962d021bbf04cd2d0f762225d19691fbff439fb5cefb56dd11bc03fc6f6671e24c000f823f2758bb79cfe7a15532fa9d2820829d4dda6bd28143615433165823388e2f4262d87e30842194614787b09e0fa1364906655c3259b5f5e00c20af0e94e6f52c21422baf94c639524b17c9af3100249f5c41be88d30c0b023735a7d311f27009bacf37501a4bdd40a8a0d3872637b85f64e4a40b016fe4826942870653b6c29296f91c618b040fb0bd9ceca65c2de6362e0a808dcb2b314ca7f8afe918904e7c43b0a97be667e33605bbb49d6f5f8535d5805fb42d68acb1a70f8a521946fe9335cbfaf86814f99d64e347939bf9f4220ef2c9471e06750916b3f4cb74c857fcb30610d0ad75a1c3c870eb67c5242a68b5187feeccbe9bc7eb22e9996ac534d7e9b274c99a231e0d360e8d7d076c677515e652cd0d9a8b3656c9dcf5481dafc2d5611d9865f9c1da3ac7bfa7f70fd8c3877f34a61942d34840b30624f932f2f1cb3964b0f54092ab7a5fdfb5eaf502025a5292a49fbf48b128c81ea0d72787ea838e980501d2b58f6f9963a5b5d43622e9f87403eacc5b056070668fdf39d03b43c48bd2a7c29da9e3edc1f345d253bac10cfdefa18b541758cfc880866e66c323575df57c61bb43ab723c8caf2742b6484e58b899f1a86c8e02a696f5fa23aa3a09080f525753cd82f1049833fd86c819876fd9a85af4b7fedf595c2a574e41f7acc5778ab47ebe8bf6bff821c0757021372ec6a379e42c279161526e4495d9b8ed4da097b6d162579a015a8a30f4bdcd21fee787bf9d613facc4ce28e34b00e180eea1307da39423b5698d90c198c1b170539fca918f3b2f44ca2bf913bccce7d94c33155657f35df170b9f3e00ced660
```

---

**7\. queryLog**
###### 接口功能
> 根据keys查询log

###### URL
> http://127.0.0.1:8001/queryLog

###### 支持格式
> JSON

###### HTTP请求方式
> POST

###### 请求参数
|参数|必选|类型|说明|example|
| :--- | :--- | :--- | :--- | :--- |
|chaincodeName|ture|string|调用链码名称|notary|
|data|true|string|链码请求数据sm2加密16进制字符串|048502008185906f51be4d64c75db8424e30ea680ad875e8fa11e53588549ad07c1bc6146e0d79f12d99fd49707d608757f4b625ea6f708ce8adcb7530be550eb921f14a37bf3001eb85a9c61860fbb2cfdcb59ae74fc46f56c74ad320b15378e132ee0b968a783bcf7f13e41650ea3cb1e5a93bf8246e02|

###### 链码请求json数据
|参数|必选|类型|说明|example|
| :--- | :--- | :--- | :--- | :--- |
|keys|true|[]string|参数字符串数组|["144","C402022072640"]|

###### 链码请求数据未加密前示例
```json
{
  "keys": "[\"144\",\"C402022072640\"]"
}
```

###### 返回字段解密后结构
|返回字段|字段类型|说明                              |
|:---:   |:---:|:---:|
|code   |int    |请求状态 0：失败；1：成功   |
|msg  |string | 请求信息                      |
|data |string |链码返回数据|

###### 接口示例
``` json
044b281a218bf3c2bdcc0b5f255b26aab64d816e39cd33ffeeada033e97df5498f407c7e679a2cca761edf0600bd43225feb74a8ee8c460b532207d28a213eba0c7384b1d3b7fb36cfe4acd8367dce8e3c69ddaf86ac904d9b55881690908329e0eff5fa72bd881950c7fbe9e867ef11d8293b7987b442a25db0194eb2d13bf2841ce1767e88ab08941fd5ab110d8d8727ec6e3b4aa9b6d3f7d33ec666d63101af6ad57d41509e0087c8a40172c4ab17b3e19877a142fdcbebc1e92e1c47d8e5abaa4db790d963010fd1fe007d5e74b403203b7ed03bbef9ac2546f42bef69ce9bd787fc21d7c879f560e8fbb542a2b13a6e99f2f843c9adab9dc705fde87b057239332541f9057ce647413310b3de77dbac8bef45bc5b95d232594882b6b1316168dc96055a212a2b3713227a1eb2d0c377c1133cc86a60dc4bdb6c22b552c5b81ba7a567484c39d3c35dc95d8a8311e777e8e4955330975f9086992fdb4ac88bd1fc1a31a3f68db93f42133e72bd7424af6121c8bbd5d13897cbe28767d9c08a264f98f4af0356bd3cb717057fc10a29ff27b9cd2186f422d411f71629be2840150a05397fe14f7cf4abd5111fc07d72d847
```

---

**8\. querysByPagination**
###### 接口功能
> 分页查询

###### URL
> http://127.0.0.1:8001/querysByPagination

###### 支持格式
> JSON

###### HTTP请求方式
> POST

###### 请求参数
|参数|必选|类型|说明|example|
| :--- | :--- | :--- | :--- | :--- |
|chaincodeName|ture|string|调用链码名称|notary|
|data|true|string|链码请求数据sm2加密16进制字符串|048502008185906f51be4d64c75db8424e30ea680ad875e8fa11e53588549ad07c1bc6146e0d79f12d99fd49707d608757f4b625ea6f708ce8adcb7530be550eb921f14a37bf3001eb85a9c61860fbb2cfdcb59ae74fc46f56c74ad320b15378e132ee0b968a783bcf7f13e41650ea3cb1e5a93bf8246e02|

###### 链码请求json数据
|参数|必选|类型|说明|example|
| :--- | :--- | :--- | :--- | :--- |
|keys|true|[]string|参数字符串数组|["144","C402022072640"]|
|pagesize|true|string|分页大小|"12"|
|nextmark|true|string|链码返回mark标志|""|

###### 链码请求数据未加密前示例
```json
{
  "keys": "[\"144\",\"C402022072640\"]",
  "pagesize": "12",
  "nextmark": ""
}
```

###### 返回字段解密后结构
|返回字段|字段类型|说明                              |
|:---:   |:---:|:---:|
|code   |int    |请求状态 0：失败；1：成功   |
|msg  |string | 请求信息                      |
|data |string |链码返回数据|

###### 接口示例
``` json
0451feeaf93749b9a61544436fdd239ab4b1b8c1eb7e8c4443949bf3b2bf1581ac5cbeae3e74979ce364c5a02c705dcb00a727287fa85133d75104f1b43369c54172c827c4cdfa70324db3cfc2c2ade9de2d278965633919c3ce1463091c8549c013f730675bcba7b569da7bf9149547021db8c54b24780ef64e349f8a6a45301d23e444e1c44bbc983efc1fb25fc49afebe342b062ebcdfa900a68c446e13d8153fb9ccd4b4562503819d1549621c90cde2fb07d2b65a6941896aa41505ae39202f2c527ef8222e38854b9fc212d6459fe332786be72cb7853b8743912c6a86bf065d9f00fcfcce25180cf1639aad7a9ddc176365536dea4093b1cb252bb80887c7e8b336cba4abd36cf99ff12e4293c94e251206d24f349a1ddec10e1690ba95de8e40a60e28ee57fbb8f5c1d4f7502801f6e7ac2e14ee0c3502ddfbde3036cbaaa9a89c83be8220
```

---

**8\. check**
###### 接口功能
> 验证数据

###### URL
> http://127.0.0.1:8001/check

###### 支持格式
> JSON

###### HTTP请求方式
> POST

###### 请求参数
|参数|必选|类型|说明|example|
| :--- | :--- | :--- | :--- | :--- |
|chaincodeName|ture|string|调用链码名称|notary|
|data|true|string|链码请求数据sm2加密16进制字符串|044e52e7fd422ed4a34a5a02a949268d8c593f0ac00dafd524f439d4808eeea11b86b9dc8be4cdf01b6249d24e3bd7e71fef490cc3ce5bab7eb2d46fe46061e0d4612f65361c6b2fd4140d743421263a670acad7fa9b01d659f7d6f971be7225535d819663da497b4d4199800fab2fddb8c9dbb4022da7e0975a20f75fc7a03aeed9472859e8e013dc5f17732f0bc23b6ab5cb8c8d35061d8b6f5f73431fb6b1839b4052367edbb4a815c51edb8fcae81cf03ca3e7b4fc2ab0b61b6c2a59ecdb4fec4f912a8f6e2d0d3f825517e7102157d8d38881f1a6fb8201e64dff272e09b8bb581c4ffc7f24722ae14674f70f9811603547a19e8a37fd36847a992ce5f78454c010c91f56fb23adcfb3eccf6f8496d0e1513b8d46df85aa54ddfbd7848add623f57023c80101175a96f32f276|

###### 链码请求json数据
|参数|必选|类型|说明|example|
| :--- | :--- | :--- | :--- | :--- |
|keys|true|[]string|参数字符串数组|["121","C402","12312"]|
|compares|true|map[string]string|比较字段键值对,验证数据字段名为key,链上记录字段名为value|{"notaryOfficeId":"notaryOfficeId"}|
|content|true|string|数据json字符串|"{\"notaryOfficeId\":40,\"serviceType\":3,\"notarizationNumber\":\"KFGZS003\",\"payTime\":\"2021-12-12 10:17:07\"}"|
|checkType|true|string|验证类型|"test"|

###### 链码请求数据未加密前示例
```json
{
  "keys": ["144","C402022072640"],
  "compares": {"notaryOfficeId":"notaryOfficeId"},
  "content": "{\"notaryOfficeId\":40,\"serviceType\":12312,\"notarizationNumber\":\"KFGZS003\",\"payTime\":\"2021-12-12 10:17:07\"}",
  "checkType": "test"
}
```

###### 返回字段解密后结构
|返回字段|字段类型|说明                              |
|:---:   |:---:|:---:|
|code   |int    |请求状态 0：失败；1：成功   |
|msg  |string | 请求信息                      |
|data |string |checkResult:链码返回验证结果0失败；1成功,txId:验证记录交易ID|

###### 接口示例
``` json
0451feeaf93749b9a61544436fdd239ab4b1b8c1eb7e8c4443949bf3b2bf1581ac5cbeae3e74979ce364c5a02c705dcb00a727287fa85133d75104f1b43369c54172c827c4cdfa70324db3cfc204d72d7d5d511f843f6cf22ab73e110af6ad21ba47b7e0289c59a60cbf6d6977c47137ef494db504a2bb08ea5edf939e41c20c858f2a80f651c41217df9be9cd4d0e4d9b4d3ff64c1e7a7c9753b0f715ed06fabe66b4d6e01d337c43ffb6287704425c9c56b6c65c2756a1f3f5b6c14ccfc0e5e7b99957685f538eb546224254b18674be79a790a86075956f762a4f3fb2b594817cbb5988de17b63d7431fdf528e8a20779c36aa5599e7b8ef917902e754ef133f9b6b70b52a22fd039e4fb040d0ea52ca2
```
