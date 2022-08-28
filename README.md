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
  - channel: mychannel
    chaincode: general2
    public:
      publicKeyX: 0ce2fa6e66521155f780573beb0e5f18d0aeea6b9a145f54e5c8c442efd15ecf
      publicKeyY: fa332850bffd6e06cbbd6e29ac851fe12da302c74550c3d75e24db54a2a1fdd7
    private:
      privateKey: 0b15a775077e438bce6ebcb7b30c3e61d9909ee861568723661d4728ee701068
    organizationsDir: organizations
    walletLabel: general2
```

# 部署
```shell
// 编译文件
CGO_ENABLED=0 build -o generalapp
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
| :---: | :---: | :---: | :---: | :---: |
|chaincodeName|ture|string|调用链码名称|notary|
|docType|true|string|链码名称sm2加密16进制字符串|048502008185906f51be4d64c75db8424e30ea680ad875e8fa11e53588549ad07c1bc6146e0d79f12d99fd49707d608757f4b625ea6f708ce8adcb7530be550eb921f14a37bf3001eb85a9c61860fbb2cfdcb59ae74fc46f56c74ad320b15378e132ee0b968a783bcf7f13e41650ea3cb1e5a93bf8246e02|

###### 返回字段
|返回字段|字段类型|说明                              |
|:---:   |:---:|:---:|
|code   |int    |请求状态 0：失败；1：成功   |
|msg  |string | 请求信息                      |
|data |string |链码返回数据                         |

###### 接口示例
``` json
{
    "code": 1,
    "msg": "success",
    "data": null
}
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
| :---: | :---: | :---: | :---: | :---: |
|chaincodeName|ture|string|调用链码名称|notary|
|keys|true|string|链码参数sm2加密16进制字符串|048502008185906f51be4d64c75db8424e30ea680ad875e8fa11e53588549ad07c1bc6146e0d79f12d99fd49707d608757f4b625ea6f708ce8adcb7530be550eb921f14a37bf3001eb85a9c61860fbb2cfdcb59ae74fc46f56c74ad320b15378e132ee0b968a783bcf7f13e41650ea3cb1e5a93bf8246e02|
|content|true|string|链码数据sm2加密16进制字符串|04a347d2dcd6c2157c478e3a5c87b7b380c8dc52cd37c2dd5bd4f72613ff8ca3238f2bf6f24a58121de0d8809eee2bdc4035b0ba56e67996df3dc3065b999ef61a716dc830f61d38a740c475aa2b23e80b6ee1a38c4c5a213ac4cbfd33594b1a94dea8f87dfeb91c51fb7d1a3110f18effbd457f2c1afbf0e9183486cfea3e411d28698f72b7e825a9f380d5092a4c190e2001c3bd764c1a6345cf5c7623b48d9017b7869e04df8ddb3d6fd2d5f2320fac411a82f57466334ee70567ac1ccc2fc7bbf0eebfaa|

###### 返回字段
|返回字段|字段类型|说明                              |
|:---:   |:---:|:---:|
|code   |int    |请求状态 0：失败；1：成功   |
|msg  |string | 请求信息                      |
|data |string |链码返回数据                         |

###### 接口示例
``` json
{
    "code": 1,
    "msg": "success",
    "data": null
}
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
| :---: | :---: | :---: | :---: | :---: |
|chaincodeName|ture|string|调用链码名称|notary|
|keys|true|string|链码参数sm2加密16进制字符串|048502008185906f51be4d64c75db8424e30ea680ad875e8fa11e53588549ad07c1bc6146e0d79f12d99fd49707d608757f4b625ea6f708ce8adcb7530be550eb921f14a37bf3001eb85a9c61860fbb2cfdcb59ae74fc46f56c74ad320b15378e132ee0b968a783bcf7f13e41650ea3cb1e5a93bf8246e02|
|content|true|string|链码数据sm2加密16进制字符串|0485038cc2ee0104fb98c19b986ddc25fe9945c428289b9776afdad58d0c368d91500e73066e2c26fd19160afb5bea0d7fbd64346a68af2568d2925bb108ac955a5af51a452e2076b2e00a661bbce5d3229353fd3eb7b0206c0c7342e4b8173be79982a2aae8a4b8e2e3d52f2cff5b54b4bd411e5db1ecda57fed569ff8b3328c92bce01410d1cfe534e4a3423d1eb00bf15fa691eaa14140d7d60da02e2349b1b7ce4249fe371b23b1f2cbde7de2634718eca92a42fcd9d347622e8e7f0cbd8c53f5684308f|

###### 返回字段
|返回字段|字段类型|说明                              |
|:---:   |:---:|:---:|
|code   |int    |请求状态 0：失败；1：成功   |
|msg  |string | 请求信息                      |
|data |string |链码返回数据                         |

###### 接口示例
``` json
{
    "code": 1,
    "msg": "success",
    "data": null
}
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
| :---: | :---: | :---: | :---: | :---: |
|chaincodeName|ture|string|调用链码名称|notary|
|keys|true|string|链码参数sm2加密16进制字符串|048502008185906f51be4d64c75db8424e30ea680ad875e8fa11e53588549ad07c1bc6146e0d79f12d99fd49707d608757f4b625ea6f708ce8adcb7530be550eb921f14a37bf3001eb85a9c61860fbb2cfdcb59ae74fc46f56c74ad320b15378e132ee0b968a783bcf7f13e41650ea3cb1e5a93bf8246e02|

###### 返回字段
|返回字段|字段类型|说明                              |
|:---:   |:---:|:---:|
|code   |int    |请求状态 0：失败；1：成功   |
|msg  |string | 请求信息                      |
|data |string |链码返回数据                         |

###### 接口示例
``` json
{
    "code": 1,
    "msg": "success",
    "data": null
}
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
| :---: | :---: | :---: | :---: | :---: |
|chaincodeName|ture|string|调用链码名称|notary|
|keys|true|string|链码参数sm2加密16进制字符串|048502008185906f51be4d64c75db8424e30ea680ad875e8fa11e53588549ad07c1bc6146e0d79f12d99fd49707d608757f4b625ea6f708ce8adcb7530be550eb921f14a37bf3001eb85a9c61860fbb2cfdcb59ae74fc46f56c74ad320b15378e132ee0b968a783bcf7f13e41650ea3cb1e5a93bf8246e02|

###### 返回字段
|返回字段|字段类型|说明                              |
|:---:   |:---:|:---:|
|code   |int    |请求状态 0：失败；1：成功   |
|msg  |string | 请求信息                      |
|data |string |链码返回数据(16进制字符串需要sm2解密处理)|

###### 接口示例
``` json
{
    "code": 1,
    "msg": "success",
    "data": "0452856340c7d6764e6672c211188f9ad46ba66b2c6b8002c05a1a415e4317b44464c361393e3831cc267c404f5f9cbc6499c1364c714241980fe1952f0f87eb6d8e4ae99c91c3966bfa0750fd5d4cf9f646fa6e3ce755f5bb98cc0aca66f00d1a39ef3c20c243105476ee1bb3083a8e399c53f1acb7bf118a6cd0388068d38d67e69168dacd3348c4e5942208609b5a78cf633fa677036e9d4697a0a9825e744a30e2d6b3d0c87934d95bd9f37815a3e3cccd58ee8e03112cf378715224540abcc5ac22cb54"
}
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
| :---: | :---: | :---: | :---: | :---: |
|chaincodeName|ture|string|调用链码名称|notary|
|keys|true|string|链码参数sm2加密16进制字符串|048502008185906f51be4d64c75db8424e30ea680ad875e8fa11e53588549ad07c1bc6146e0d79f12d99fd49707d608757f4b625ea6f708ce8adcb7530be550eb921f14a37bf3001eb85a9c61860fbb2cfdcb59ae74fc46f56c74ad320b15378e132ee0b968a783bcf7f13e41650ea3cb1e5a93bf8246e02|

###### 返回字段
|返回字段|字段类型|说明                              |
|:---:   |:---:|:---:|
|code   |int    |请求状态 0：失败；1：成功   |
|msg  |string | 请求信息                      |
|data |string |链码返回数据(16进制字符串需要sm2解密处理)|

###### 接口示例
``` json
{
    "code": 1,
    "msg": "success",
    "data": "046bf9702922fccc7b5b0b493efdff705e605d887a93c93f85e293f51409122491b1b48c715087656cf6c990a164ac83949ac6a1c08c3f8310af1a68ca14535097019909302d3fba7de6cc50f5ca99c641598866d2acf853a592a9f927ab8587d5a3d0de19d33818b0611c48f31d6c3ea8a1d33c11e9590afdbbe490a4810f30919d3838c670c47cde29140a7765ff5e124d65143b7b368c78acde75258a1bc3e4ae549c7692ff332e925b0259cae068125d5fbe43da7bb22ffad82c4adba2e382aec88163307dc9"
}
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
| :---: | :---: | :---: | :---: | :---: |
|chaincodeName|ture|string|调用链码名称|notary|
|keys|true|string|链码参数sm2加密16进制字符串|048502008185906f51be4d64c75db8424e30ea680ad875e8fa11e53588549ad07c1bc6146e0d79f12d99fd49707d608757f4b625ea6f708ce8adcb7530be550eb921f14a37bf3001eb85a9c61860fbb2cfdcb59ae74fc46f56c74ad320b15378e132ee0b968a783bcf7f13e41650ea3cb1e5a93bf8246e02|

###### 返回字段
|返回字段|字段类型|说明                              |
|:---:   |:---:|:---:|
|code   |int    |请求状态 0：失败；1：成功   |
|msg  |string | 请求信息                      |
|data |string |链码返回数据(16进制字符串需要sm2解密处理)|

###### 接口示例
``` json
{
    "code": 1,
    "msg": "success",
    "data": "04cae7237e9c62d46814f9ac2deea7a287708338b55acf6b5a08130a7881ef1347fc82fa91f8c714de7767dacb34f0761fb323fb6f3e10ac30419f3c14eff0620773986f3896f88199b73c5a10b5597b8e64d55189f0ff4adb73912849c01e3fcf2f7598600f54e07835558a7aeb4666d0235c2540b08f54b5aba7d4b2481dbfd6d3ae93323b7b7f7b258d4ef4cb7467b2410f1cfcd5501029d04f7cc1dc8a908f03837cbba67726713d0426139236ba4a0aa86648c2c5862062001536321709dc18c77224ba014e32f699c1c44a5e6b5070081269619b3e7c04c48cdacac96bf0c65b49d8efc73ba3cb8e9659d52bf99d0f5eee7d3a9c93bdc4ab7e285777d87665890f3a29e6b99f794ec2ccf69275a28b32dadc48766dd0351eabcc23215254a0a13c5d6fce5b956f0006a02f03493d405723074934d5cf620d0b1b4dd43000b2ee6e4cc9e434adf7feefcc1a43cf610572245363efb567b4bd723d037276b0d1e57b5a5e893fb955c939e38e6fa3db3985d5633238f40b8df08997e63f2deac5eb65c906f32a25a242aa9a030a8735480cca63c07c47263019e311ded95359393b4efd6517d8f97384f8c03dcabf782063da482c4bb5e1c0b2ca231304de83986e3775b1f7e7c8f6b0189b19a678894e072e84fb287fa895451e4bf98d6b8186886b2afb7fa3e39094c5af06075a6d53b7e84e05395928620fdfca0e496443f1f0d9fdf5ef2d7ab6185e6cc5eab6b18d1eec249acc7243f788918d83ed1d596bc0e298e2446af552d7d5dcd6cc8f7012b24710b7dfaa3df60772ba8f780b932316c04b86f464aad7bba6ba8365476745c229bc7da07704aa9b485963761e44637249bf007f99e8f8430a92c36ed1f294fadad0d7615b5e5325c144133951a24704d6e9d349f7d837562f5a1c8ada4dcae621f47ecbf4c7197f6ad3c8fde77ce14482d6e4413ed3cb6e83f76e38835917d70b71c9c2a2954d34ad8247cc492bf05f73066b270a84c939a8a367c2e9e8d6a31193249dc0ae451e6e01b7750cb4ac0db22c63516decb2548dc186edce91f1e0e3512f85393661f3d2d251a8c8e1b9fa1e07b3bf80bd7b87fec686a16f7693ecd555b101a6250d1e9cb92e2e0a655ee8bbf0c76369e5e427a860e8c5e4f539a08f51a8c8b6aa24cb3482960d71b39907dd4764948b9839697dde23f9828f6a6db0b07dbba41eca821211472b33c046a09ef71c1ff5f1a94ea99e0fb306e0173b6e7e1a6cf1074030a19848a076d6969e2c4dd0a6d8b40ba7912d9bc13f2349fd5362be04ec33405b5f0b3243bea5dfe474e1dde6cb1fd1ae505cde07a1418406687e07633eec77a5c7ea9d159906381fa62538c54fbe53543d3148c16d3db5ce8e2fda84759045c28f24bb860d9e61a4d2d227e41c967d94b717021c919bf5fd701690a4952247c8161b22667c2568841e8c89b0fda10d769e"
}
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
| :---: | :---: | :---: | :---: | :---: |
|chaincodeName|ture|string|调用链码名称|notary|
|keys|true|string|链码参数sm2加密16进制字符串|04320e2938d399a34b76dfd99bcb8ca5fd6695a9bb4dc7ca02715a9e4225ce6a06e71ea903dd9c117f1ee7dadba76793a075a228c4f90d767dedbe147a5f7cd03c06631a7d1c3facae432046197e89af45d822ba5f5739373d2fce1c3ad2153fd982b5|
|pageSize|true|string|分页大小整数类型sm2加密16进制字符串|04c5472ce881dda7c3792fbad7b18acd906c6cae43b8b2a17bf29180dad7c97be73961532d598eb1b6a43131025d2a6c216f8feb7b10db35f91a33251f6309e004a908d42715df3e29a2f6554f47532e5e1ac4449301dada80f0f0f45ce12571f7c5af|
|nextmark|false|string|链码返回bookmark标签sm2加密16进制字符串||

###### 返回字段
|返回字段|字段类型|说明                              |
|:---:   |:---:|:---:|
|code   |int    |请求状态 0：失败；1：成功   |
|msg  |string | 请求信息                      |
|data |string |链码返回数据(16进制字符串需要sm2解密处理)|

###### 接口示例
``` json
{
    "code": 1,
    "msg": "success",
    "data": "04728e49f6d4b695d5ff07415a16565eb07a5c5af7d55e452c79816883844f3f56eb003a87a54040c2da0e88ca196aecaa585585ffed50c4451dfcd01e11749fea90d9e294839e301036978b05f0466802dcebfc14beb9cb8f9bcaec3c097574e12ff7f18823ac02a3fbb6e8e57b8e5faeab1f9234eaa43d5f47764cdf774a86fcaca43f0d914a95f9cca4cc971eae6705e23d68a59958aa53a4088706d5a8e09e390b6e68b8c2c8942dcae11522cb05c6bf0190c8844536bbee6b46bbd846b31ba54750157eee0a437b110079222129ab91a52621bfbb05ec8339ea40321fa883e0c7c64462f99e933208ac2342eb0c5110ac201654e47bbc574727d9be13cf75"
}
```