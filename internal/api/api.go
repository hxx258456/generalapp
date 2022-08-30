// 请求流程:解析参数=>解密参数=>请求链码=>加密请求结果=>返回加密结果
//
package api

import (
	"encoding/hex"
	"encoding/json"
	"generalapp/internal/model"
	"generalapp/pkg/sdkpool"
	"generalapp/pkg/utils"
	"log"

	"github.com/gin-gonic/gin"
)

// 初始化链码
func InitChaincode(c *gin.Context) {
	// 参数解析
	param := model.Request{}
	if err := c.ShouldBindJSON(&param); err != nil {
		resp := model.Response{
			Code: 0,
			Msg:  "参数解析错误: " + err.Error(),
			Data: nil,
		}
		c.JSON(200, resp)
		return
	}

	// 检查链码sdk是否在连接池中初始化
	if _, ok := sdkpool.SdkPoll[param.ChaincodeName]; !ok {
		resp := model.Response{
			Code: 0,
			Msg:  param.ChaincodeName + " sdk未初始化",
			Data: nil,
		}
		c.JSON(200, resp)
		return
	}

	// 16进制获取加密结果
	hexByte, err := hex.DecodeString(param.Data)
	if err != nil {
		resp := model.Response{
			Code: 0,
			Msg:  err.Error(),
			Data: nil,
		}
		respByte, _ := json.Marshal(&resp)
		encResp, _ := utils.Encrypt(sdkpool.SdkPoll[param.ChaincodeName].Public, string(respByte))

		c.String(200, hex.EncodeToString(encResp))
		return
	}

	// sm2解密请求数据
	decData, err := utils.Decrypt(sdkpool.SdkPoll[param.ChaincodeName].Private, hexByte)
	if err != nil {
		resp := model.Response{
			Code: 0,
			Msg:  err.Error(),
			Data: nil,
		}
		respByte, _ := json.Marshal(&resp)
		encResp, _ := utils.Encrypt(sdkpool.SdkPoll[param.ChaincodeName].Public, string(respByte))

		c.String(200, hex.EncodeToString(encResp))
		return
	}
	log.Println("解密结果[data]: ", decData)

	data := new(model.InitData)
	if err := json.Unmarshal([]byte(decData), data); err != nil {
		resp := model.Response{
			Code: 0,
			Msg:  err.Error(),
			Data: nil,
		}
		respByte, _ := json.Marshal(&resp)
		encResp, _ := utils.Encrypt(sdkpool.SdkPoll[param.ChaincodeName].Public, string(respByte))

		c.String(200, hex.EncodeToString(encResp))
		return
	}

	// sdk发起交易
	var result = []byte{}
	result, err = sdkpool.SdkPoll[param.ChaincodeName].Contract.SubmitTransaction("initChaincode", data.DocType)
	if err != nil {
		// todo 重新初始化sdk并发起请求
		sdk_ := sdkpool.SdkPoll[param.ChaincodeName]
		sdk_.InitSdk()
		sdkpool.SdkPoll[param.ChaincodeName] = sdk_

		result, err = sdkpool.SdkPoll[param.ChaincodeName].Contract.SubmitTransaction("initChaincode", data.DocType)
		if err != nil {
			resp := model.Response{
				Code: 0,
				Msg:  err.Error(),
				Data: nil,
			}
			respByte, _ := json.Marshal(&resp)
			encResp, _ := utils.Encrypt(sdkpool.SdkPoll[param.ChaincodeName].Public, string(respByte))

			c.String(200, hex.EncodeToString(encResp))
			return
		}
	}
	log.Println("请求结果: ", string(result))

	resp := model.Response{
		Code: 1,
		Msg:  "success",
		Data: result,
	}
	respByte, _ := json.Marshal(&resp)
	encResp, _ := utils.Encrypt(sdkpool.SdkPoll[param.ChaincodeName].Public, string(respByte))

	c.String(200, hex.EncodeToString(encResp))
}

// 添加
func Add(c *gin.Context) {
	// 参数解析
	param := model.Request{}
	if err := c.ShouldBindJSON(&param); err != nil {
		resp := model.Response{
			Code: 0,
			Msg:  "参数解析错误: " + err.Error(),
			Data: nil,
		}

		c.JSON(200, resp)
		return
	}

	// 检查链码sdk是否在连接池中初始化
	if _, ok := sdkpool.SdkPoll[param.ChaincodeName]; !ok {
		resp := model.Response{
			Code: 0,
			Msg:  param.ChaincodeName + " sdk未初始化",
			Data: nil,
		}
		c.JSON(200, resp)
		return
	}

	// 解析16进制字符串为[]byte
	hexByte, err := hex.DecodeString(param.Data)
	if err != nil {
		resp := model.Response{
			Code: 0,
			Msg:  err.Error(),
			Data: nil,
		}
		respByte, _ := json.Marshal(&resp)
		encResp, _ := utils.Encrypt(sdkpool.SdkPoll[param.ChaincodeName].Public, string(respByte))

		c.String(200, hex.EncodeToString(encResp))
		return
	}

	// sm2解密请求数据
	dataDec, err := utils.Decrypt(sdkpool.SdkPoll[param.ChaincodeName].Private, hexByte)
	if err != nil {
		resp := model.Response{
			Code: 0,
			Msg:  err.Error(),
			Data: nil,
		}

		respByte, _ := json.Marshal(&resp)
		encResp, _ := utils.Encrypt(sdkpool.SdkPoll[param.ChaincodeName].Public, string(respByte))

		c.String(200, hex.EncodeToString(encResp))
		return
	}
	log.Println("解密结果[Data]: ", dataDec)

	data := new(model.AddData)
	if err := json.Unmarshal([]byte(dataDec), data); err != nil {
		resp := model.Response{
			Code: 0,
			Msg:  err.Error(),
			Data: nil,
		}
		respByte, _ := json.Marshal(&resp)
		encResp, _ := utils.Encrypt(sdkpool.SdkPoll[param.ChaincodeName].Public, string(respByte))

		c.String(200, hex.EncodeToString(encResp))
		return
	}

	keys, err := json.Marshal(data.Keys)
	if err != nil {
		resp := model.Response{
			Code: 0,
			Msg:  err.Error(),
			Data: nil,
		}
		respByte, _ := json.Marshal(&resp)
		encResp, _ := utils.Encrypt(sdkpool.SdkPoll[param.ChaincodeName].Public, string(respByte))

		c.String(200, hex.EncodeToString(encResp))
		return
	}
	log.Println("[keys]: ", string(keys))

	// sdk发起交易
	var result = []byte{}
	result, err = sdkpool.SdkPoll[param.ChaincodeName].Contract.SubmitTransaction("add", string(keys), data.Content)
	if err != nil {
		log.Println(err)
		// todo 重新初始化sdk并发起请求
		sdk_ := sdkpool.SdkPoll[param.ChaincodeName]
		sdk_.InitSdk()
		sdkpool.SdkPoll[param.ChaincodeName] = sdk_

		result, err = sdkpool.SdkPoll[param.ChaincodeName].Contract.SubmitTransaction("add", string(keys), data.Content)
		if err != nil {
			resp := model.Response{
				Code: 0,
				Msg:  err.Error(),
				Data: nil,
			}
			respByte, _ := json.Marshal(&resp)
			encResp, _ := utils.Encrypt(sdkpool.SdkPoll[param.ChaincodeName].Public, string(respByte))

			c.String(200, hex.EncodeToString(encResp))
			return
		}
	}
	log.Println("请求结果: ", string(result))

	resp := model.Response{
		Code: 1,
		Msg:  "success",
		Data: string(result), // 添加txid
	}

	respByte, _ := json.Marshal(&resp)
	encResp, _ := utils.Encrypt(sdkpool.SdkPoll[param.ChaincodeName].Public, string(respByte))

	c.String(200, hex.EncodeToString(encResp))
}

// 更新
func Update(c *gin.Context) {
	// 参数解析
	param := model.Request{}
	if err := c.ShouldBindJSON(&param); err != nil {
		resp := model.Response{
			Code: 0,
			Msg:  "参数解析错误: " + err.Error(),
			Data: nil,
		}
		c.JSON(200, resp)
		return
	}

	// 检查链码sdk是否在连接池中初始化
	if _, ok := sdkpool.SdkPoll[param.ChaincodeName]; !ok {
		resp := model.Response{
			Code: 0,
			Msg:  param.ChaincodeName + " sdk未初始化",
			Data: nil,
		}
		c.JSON(200, resp)
		return
	}

	// 解析16进制字符串为[]byte
	hexByte, err := hex.DecodeString(param.Data)
	if err != nil {
		resp := model.Response{
			Code: 0,
			Msg:  err.Error(),
			Data: nil,
		}
		respByte, _ := json.Marshal(&resp)
		encResp, _ := utils.Encrypt(sdkpool.SdkPoll[param.ChaincodeName].Public, string(respByte))

		c.String(200, hex.EncodeToString(encResp))
		return
	}

	// sm2解密请求数据
	dataDec, err := utils.Decrypt(sdkpool.SdkPoll[param.ChaincodeName].Private, hexByte)
	if err != nil {
		resp := model.Response{
			Code: 0,
			Msg:  err.Error(),
			Data: nil,
		}
		respByte, _ := json.Marshal(&resp)
		encResp, _ := utils.Encrypt(sdkpool.SdkPoll[param.ChaincodeName].Public, string(respByte))

		c.String(200, hex.EncodeToString(encResp))
		return
	}
	log.Println("解密结果[Data]: ", dataDec)

	data := new(model.AddData)
	if err := json.Unmarshal([]byte(dataDec), data); err != nil {
		resp := model.Response{
			Code: 0,
			Msg:  err.Error(),
			Data: nil,
		}
		respByte, _ := json.Marshal(&resp)
		encResp, _ := utils.Encrypt(sdkpool.SdkPoll[param.ChaincodeName].Public, string(respByte))

		c.String(200, hex.EncodeToString(encResp))
		return
	}

	// 解析请求数据
	keys, err := json.Marshal(data.Keys)
	if err != nil {
		resp := model.Response{
			Code: 0,
			Msg:  err.Error(),
			Data: nil,
		}
		respByte, _ := json.Marshal(&resp)
		encResp, _ := utils.Encrypt(sdkpool.SdkPoll[param.ChaincodeName].Public, string(respByte))

		c.String(200, hex.EncodeToString(encResp))
		return
	}

	// sdk发起交易
	var result = []byte{}
	result, err = sdkpool.SdkPoll[param.ChaincodeName].Contract.SubmitTransaction("update", string(keys), data.Content)
	if err != nil {
		log.Println(err)
		// todo 重新初始化sdk并发起请求
		sdk_ := sdkpool.SdkPoll[param.ChaincodeName]
		sdk_.InitSdk()
		sdkpool.SdkPoll[param.ChaincodeName] = sdk_

		result, err = sdkpool.SdkPoll[param.ChaincodeName].Contract.SubmitTransaction("update", string(keys), data.Content)
		if err != nil {
			resp := model.Response{
				Code: 0,
				Msg:  err.Error(),
				Data: nil,
			}
			respByte, _ := json.Marshal(&resp)
			encResp, _ := utils.Encrypt(sdkpool.SdkPoll[param.ChaincodeName].Public, string(respByte))

			c.String(200, hex.EncodeToString(encResp))
			return
		}
	}
	log.Println("请求结果: ", string(result))

	resp := model.Response{
		Code: 1,
		Msg:  "success",
		Data: string(result), // 添加txid
	}
	respByte, _ := json.Marshal(&resp)
	encResp, _ := utils.Encrypt(sdkpool.SdkPoll[param.ChaincodeName].Public, string(respByte))

	c.String(200, hex.EncodeToString(encResp))
}

// 删除
func Delete(c *gin.Context) {
	// 参数解析
	param := model.Request{}
	if err := c.ShouldBindJSON(&param); err != nil {
		resp := model.Response{
			Code: 0,
			Msg:  "参数解析错误: " + err.Error(),
			Data: nil,
		}
		c.JSON(200, resp)
		return
	}

	// 检查链码sdk是否在连接池中初始化
	if _, ok := sdkpool.SdkPoll[param.ChaincodeName]; !ok {
		resp := model.Response{
			Code: 0,
			Msg:  param.ChaincodeName + " sdk未初始化",
			Data: nil,
		}
		c.JSON(200, resp)
		return
	}

	// 解析16进制字符串为[]byte
	hexByte, err := hex.DecodeString(param.Data)
	if err != nil {
		resp := model.Response{
			Code: 0,
			Msg:  err.Error(),
			Data: nil,
		}
		respByte, _ := json.Marshal(&resp)
		encResp, _ := utils.Encrypt(sdkpool.SdkPoll[param.ChaincodeName].Public, string(respByte))

		c.String(200, hex.EncodeToString(encResp))
		return
	}

	// sm2解密请求数据
	decData, err := utils.Decrypt(sdkpool.SdkPoll[param.ChaincodeName].Private, hexByte)
	if err != nil {
		resp := model.Response{
			Code: 0,
			Msg:  err.Error(),
			Data: nil,
		}
		respByte, _ := json.Marshal(&resp)
		encResp, _ := utils.Encrypt(sdkpool.SdkPoll[param.ChaincodeName].Public, string(respByte))

		c.String(200, hex.EncodeToString(encResp))
		return
	}
	log.Println("解密结果[Data]: ", decData)

	data := new(model.KeysData)
	if err := json.Unmarshal([]byte(decData), data); err != nil {
		resp := model.Response{
			Code: 0,
			Msg:  err.Error(),
			Data: nil,
		}
		respByte, _ := json.Marshal(&resp)
		encResp, _ := utils.Encrypt(sdkpool.SdkPoll[param.ChaincodeName].Public, string(respByte))

		c.String(200, hex.EncodeToString(encResp))
		return
	}

	keys, err := json.Marshal(data.Keys)
	if err != nil {
		resp := model.Response{
			Code: 0,
			Msg:  err.Error(),
			Data: nil,
		}
		respByte, _ := json.Marshal(&resp)
		encResp, _ := utils.Encrypt(sdkpool.SdkPoll[param.ChaincodeName].Public, string(respByte))

		c.String(200, hex.EncodeToString(encResp))
		return
	}

	// sdk发起交易
	var result = []byte{}
	result, err = sdkpool.SdkPoll[param.ChaincodeName].Contract.SubmitTransaction("delete", string(keys))
	if err != nil {
		log.Println(err)
		// todo 重新初始化sdk并发起请求
		sdk_ := sdkpool.SdkPoll[param.ChaincodeName]
		sdk_.InitSdk()
		sdkpool.SdkPoll[param.ChaincodeName] = sdk_

		result, err = sdkpool.SdkPoll[param.ChaincodeName].Contract.SubmitTransaction("delete", string(keys))
		if err != nil {
			resp := model.Response{
				Code: 0,
				Msg:  err.Error(),
				Data: nil,
			}
			respByte, _ := json.Marshal(&resp)
			encResp, _ := utils.Encrypt(sdkpool.SdkPoll[param.ChaincodeName].Public, string(respByte))

			c.String(200, hex.EncodeToString(encResp))
			return
		}
	}
	log.Println("请求结果: ", string(result))

	resp := model.Response{
		Code: 1,
		Msg:  "success",
		Data: string(result), // 添加txid
	}
	respByte, _ := json.Marshal(&resp)
	encResp, _ := utils.Encrypt(sdkpool.SdkPoll[param.ChaincodeName].Public, string(respByte))

	c.String(200, hex.EncodeToString(encResp))
}

// 查询
func Query(c *gin.Context) {
	// 参数解析
	param := model.Request{}
	if err := c.ShouldBindJSON(&param); err != nil {
		resp := model.Response{
			Code: 0,
			Msg:  "参数解析错误: " + err.Error(),
			Data: nil,
		}
		c.JSON(200, resp)
		return
	}

	// 检查链码sdk是否在连接池中初始化
	if _, ok := sdkpool.SdkPoll[param.ChaincodeName]; !ok {
		resp := model.Response{
			Code: 0,
			Msg:  param.ChaincodeName + " sdk未初始化",
			Data: nil,
		}
		c.JSON(200, resp)
		return
	}

	// 解析16进制字符串为[]byte
	hexByte, err := hex.DecodeString(param.Data)
	if err != nil {
		resp := model.Response{
			Code: 0,
			Msg:  err.Error(),
			Data: nil,
		}
		respByte, _ := json.Marshal(&resp)
		encResp, _ := utils.Encrypt(sdkpool.SdkPoll[param.ChaincodeName].Public, string(respByte))

		c.String(200, hex.EncodeToString(encResp))
		return
	}

	// sm2解密请求数据
	decData, err := utils.Decrypt(sdkpool.SdkPoll[param.ChaincodeName].Private, hexByte)
	if err != nil {
		resp := model.Response{
			Code: 0,
			Msg:  err.Error(),
			Data: nil,
		}
		respByte, _ := json.Marshal(&resp)
		encResp, _ := utils.Encrypt(sdkpool.SdkPoll[param.ChaincodeName].Public, string(respByte))

		c.String(200, hex.EncodeToString(encResp))
		return
	}
	log.Println("解密结果[data]: ", decData)

	data := new(model.KeysData)
	if err := json.Unmarshal([]byte(decData), data); err != nil {
		resp := model.Response{
			Code: 0,
			Msg:  err.Error(),
			Data: nil,
		}
		respByte, _ := json.Marshal(&resp)
		encResp, _ := utils.Encrypt(sdkpool.SdkPoll[param.ChaincodeName].Public, string(respByte))

		c.String(200, hex.EncodeToString(encResp))
		return
	}

	keys, err := json.Marshal(data.Keys)
	if err != nil {
		resp := model.Response{
			Code: 0,
			Msg:  err.Error(),
			Data: nil,
		}
		respByte, _ := json.Marshal(&resp)
		encResp, _ := utils.Encrypt(sdkpool.SdkPoll[param.ChaincodeName].Public, string(respByte))

		c.String(200, hex.EncodeToString(encResp))
		return
	}

	// sdk发起交易
	var result = []byte{}
	result, err = sdkpool.SdkPoll[param.ChaincodeName].Contract.SubmitTransaction("query", string(keys))
	if err != nil {
		log.Println(err)
		// todo 重新初始化sdk并发起请求
		sdk_ := sdkpool.SdkPoll[param.ChaincodeName]
		sdk_.InitSdk()
		sdkpool.SdkPoll[param.ChaincodeName] = sdk_

		result, err = sdkpool.SdkPoll[param.ChaincodeName].Contract.SubmitTransaction("query", string(keys))
		if err != nil {
			resp := model.Response{
				Code: 0,
				Msg:  err.Error(),
				Data: nil,
			}
			respByte, _ := json.Marshal(&resp)
			encResp, _ := utils.Encrypt(sdkpool.SdkPoll[param.ChaincodeName].Public, string(respByte))

			c.String(200, hex.EncodeToString(encResp))
			return
		}
	}
	log.Println("请求结果: ", string(result))

	resp := model.Response{
		Code: 1,
		Msg:  "success",
		Data: string(result),
	}
	respByte, _ := json.Marshal(&resp)
	encResp, _ := utils.Encrypt(sdkpool.SdkPoll[param.ChaincodeName].Public, string(respByte))

	c.String(200, hex.EncodeToString(encResp))
}

// 查询所有
func QueryAll(c *gin.Context) {
	// 参数解析
	param := model.Request{}
	if err := c.ShouldBindJSON(&param); err != nil {
		resp := model.Response{
			Code: 0,
			Msg:  "参数解析错误: " + err.Error(),
			Data: nil,
		}
		c.JSON(200, resp)
		return
	}

	// 检查链码sdk是否在连接池中初始化
	if _, ok := sdkpool.SdkPoll[param.ChaincodeName]; !ok {
		resp := model.Response{
			Code: 0,
			Msg:  param.ChaincodeName + " sdk未初始化",
			Data: nil,
		}
		c.JSON(200, resp)
		return
	}

	// 解析16进制字符串为[]byte
	hexByte, err := hex.DecodeString(param.Data)
	if err != nil {
		resp := model.Response{
			Code: 0,
			Msg:  err.Error(),
			Data: nil,
		}
		respByte, _ := json.Marshal(&resp)
		encResp, _ := utils.Encrypt(sdkpool.SdkPoll[param.ChaincodeName].Public, string(respByte))

		c.String(200, hex.EncodeToString(encResp))
		return
	}

	// sm2解密请求数据
	decData, err := utils.Decrypt(sdkpool.SdkPoll[param.ChaincodeName].Private, hexByte)
	if err != nil {
		resp := model.Response{
			Code: 0,
			Msg:  err.Error(),
			Data: nil,
		}
		respByte, _ := json.Marshal(&resp)
		encResp, _ := utils.Encrypt(sdkpool.SdkPoll[param.ChaincodeName].Public, string(respByte))

		c.String(200, hex.EncodeToString(encResp))
		return
	}
	log.Println("解密结果[data]: ", decData)

	data := new(model.KeysData)
	if err := json.Unmarshal([]byte(decData), data); err != nil {
		resp := model.Response{
			Code: 0,
			Msg:  err.Error(),
			Data: nil,
		}
		respByte, _ := json.Marshal(&resp)
		encResp, _ := utils.Encrypt(sdkpool.SdkPoll[param.ChaincodeName].Public, string(respByte))

		c.String(200, hex.EncodeToString(encResp))
		return
	}

	keys, err := json.Marshal(data.Keys)
	if err != nil {
		resp := model.Response{
			Code: 0,
			Msg:  err.Error(),
			Data: nil,
		}
		respByte, _ := json.Marshal(&resp)
		encResp, _ := utils.Encrypt(sdkpool.SdkPoll[param.ChaincodeName].Public, string(respByte))

		c.String(200, hex.EncodeToString(encResp))
		return
	}

	// sdk发起交易
	var result = []byte{}
	result, err = sdkpool.SdkPoll[param.ChaincodeName].Contract.SubmitTransaction("queryAll", string(keys))
	if err != nil {
		log.Println(err)
		// todo 重新初始化sdk并发起请求
		sdk_ := sdkpool.SdkPoll[param.ChaincodeName]
		sdk_.InitSdk()
		sdkpool.SdkPoll[param.ChaincodeName] = sdk_

		result, err = sdkpool.SdkPoll[param.ChaincodeName].Contract.SubmitTransaction("queryAll", string(keys))
		if err != nil {
			resp := model.Response{
				Code: 0,
				Msg:  err.Error(),
				Data: nil,
			}
			respByte, _ := json.Marshal(&resp)
			encResp, _ := utils.Encrypt(sdkpool.SdkPoll[param.ChaincodeName].Public, string(respByte))

			c.String(200, hex.EncodeToString(encResp))
			return
		}
	}
	log.Println("请求结果: ", string(result))

	resp := model.Response{
		Code: 1,
		Msg:  "success",
		Data: string(result),
	}
	respByte, _ := json.Marshal(&resp)
	encResp, _ := utils.Encrypt(sdkpool.SdkPoll[param.ChaincodeName].Public, string(respByte))

	c.String(200, hex.EncodeToString(encResp))
}

// 分页查询
func QuerysByPagination(c *gin.Context) {
	// 参数解析
	param := model.Request{}
	if err := c.ShouldBindJSON(&param); err != nil {
		resp := model.Response{
			Code: 0,
			Msg:  "参数解析错误: " + err.Error(),
			Data: nil,
		}
		c.JSON(200, resp)
		return
	}

	// 检查链码sdk是否在连接池中初始化
	if _, ok := sdkpool.SdkPoll[param.ChaincodeName]; !ok {
		resp := model.Response{
			Code: 0,
			Msg:  param.ChaincodeName + " sdk未初始化",
			Data: nil,
		}
		c.JSON(200, resp)
		return
	}

	// 解析16进制字符串为[]byte
	hexByte, err := hex.DecodeString(param.Data)
	if err != nil {
		resp := model.Response{
			Code: 0,
			Msg:  err.Error(),
			Data: nil,
		}
		respByte, _ := json.Marshal(&resp)
		encResp, _ := utils.Encrypt(sdkpool.SdkPoll[param.ChaincodeName].Public, string(respByte))

		c.String(200, hex.EncodeToString(encResp))
		return
	}

	// sm2解密请求数据
	decData, err := utils.Decrypt(sdkpool.SdkPoll[param.ChaincodeName].Private, hexByte)
	if err != nil {
		resp := model.Response{
			Code: 0,
			Msg:  err.Error(),
			Data: nil,
		}
		respByte, _ := json.Marshal(&resp)
		encResp, _ := utils.Encrypt(sdkpool.SdkPoll[param.ChaincodeName].Public, string(respByte))

		c.String(200, hex.EncodeToString(encResp))
		return
	}
	log.Println("解密结果[data]: ", decData)

	data := new(model.PageData)
	if err := json.Unmarshal([]byte(decData), data); err != nil {
		resp := model.Response{
			Code: 0,
			Msg:  err.Error(),
			Data: nil,
		}
		respByte, _ := json.Marshal(&resp)
		encResp, _ := utils.Encrypt(sdkpool.SdkPoll[param.ChaincodeName].Public, string(respByte))

		c.String(200, hex.EncodeToString(encResp))
		return
	}

	keys, err := json.Marshal(data.Keys)
	if err != nil {
		resp := model.Response{
			Code: 0,
			Msg:  err.Error(),
			Data: nil,
		}
		respByte, _ := json.Marshal(&resp)
		encResp, _ := utils.Encrypt(sdkpool.SdkPoll[param.ChaincodeName].Public, string(respByte))

		c.String(200, hex.EncodeToString(encResp))
		return
	}

	// sdk发起交易
	var result = []byte{}
	result, err = sdkpool.SdkPoll[param.ChaincodeName].Contract.SubmitTransaction("querysByPagination", string(keys), data.Pagesize, data.Nextmark)
	if err != nil {
		log.Println(err)
		// todo 重新初始化sdk并发起请求
		sdk_ := sdkpool.SdkPoll[param.ChaincodeName]
		sdk_.InitSdk()
		sdkpool.SdkPoll[param.ChaincodeName] = sdk_

		result, err = sdkpool.SdkPoll[param.ChaincodeName].Contract.SubmitTransaction("querysByPagination", string(keys), data.Pagesize, data.Nextmark)
		if err != nil {
			resp := model.Response{
				Code: 0,
				Msg:  err.Error(),
				Data: nil,
			}
			respByte, _ := json.Marshal(&resp)
			encResp, _ := utils.Encrypt(sdkpool.SdkPoll[param.ChaincodeName].Public, string(respByte))

			c.String(200, hex.EncodeToString(encResp))
			return
		}
	}
	log.Println("请求结果: ", string(result))

	resp := model.Response{
		Code: 1,
		Msg:  "success",
		Data: string(result),
	}
	respByte, _ := json.Marshal(&resp)
	encResp, _ := utils.Encrypt(sdkpool.SdkPoll[param.ChaincodeName].Public, string(respByte))

	c.String(200, hex.EncodeToString(encResp))
}

// 查询log
func QueryLog(c *gin.Context) {
	// 参数解析
	param := model.Request{}
	if err := c.ShouldBindJSON(&param); err != nil {
		resp := model.Response{
			Code: 0,
			Msg:  "参数解析错误: " + err.Error(),
			Data: nil,
		}
		c.JSON(200, resp)
		return
	}

	// 检查链码sdk是否在连接池中初始化
	if _, ok := sdkpool.SdkPoll[param.ChaincodeName]; !ok {
		resp := model.Response{
			Code: 0,
			Msg:  param.ChaincodeName + " sdk未初始化",
			Data: nil,
		}
		c.JSON(200, resp)
		return
	}

	// 解析16进制字符串为[]byte
	hexByte, err := hex.DecodeString(param.Data)
	if err != nil {
		resp := model.Response{
			Code: 0,
			Msg:  err.Error(),
			Data: nil,
		}
		respByte, _ := json.Marshal(&resp)
		encResp, _ := utils.Encrypt(sdkpool.SdkPoll[param.ChaincodeName].Public, string(respByte))

		c.String(200, hex.EncodeToString(encResp))
		return
	}

	// sm2解密请求数据
	decData, err := utils.Decrypt(sdkpool.SdkPoll[param.ChaincodeName].Private, hexByte)
	if err != nil {
		resp := model.Response{
			Code: 0,
			Msg:  err.Error(),
			Data: nil,
		}
		respByte, _ := json.Marshal(&resp)
		encResp, _ := utils.Encrypt(sdkpool.SdkPoll[param.ChaincodeName].Public, string(respByte))

		c.String(200, hex.EncodeToString(encResp))
		return
	}
	log.Println("解密结果[data]: ", decData)

	data := new(model.KeysData)
	if err := json.Unmarshal([]byte(decData), data); err != nil {
		resp := model.Response{
			Code: 0,
			Msg:  err.Error(),
			Data: nil,
		}
		respByte, _ := json.Marshal(&resp)
		encResp, _ := utils.Encrypt(sdkpool.SdkPoll[param.ChaincodeName].Public, string(respByte))

		c.String(200, hex.EncodeToString(encResp))
		return
	}

	keys, err := json.Marshal(data.Keys)
	if err != nil {
		resp := model.Response{
			Code: 0,
			Msg:  err.Error(),
			Data: nil,
		}
		respByte, _ := json.Marshal(&resp)
		encResp, _ := utils.Encrypt(sdkpool.SdkPoll[param.ChaincodeName].Public, string(respByte))

		c.String(200, hex.EncodeToString(encResp))
		return
	}

	// sdk发起交易
	var result = []byte{}
	result, err = sdkpool.SdkPoll[param.ChaincodeName].Contract.SubmitTransaction("queryLog", string(keys))
	if err != nil {
		log.Println(err)
		// todo 重新初始化sdk并发起请求
		sdk_ := sdkpool.SdkPoll[param.ChaincodeName]
		sdk_.InitSdk()
		sdkpool.SdkPoll[param.ChaincodeName] = sdk_

		result, err = sdkpool.SdkPoll[param.ChaincodeName].Contract.SubmitTransaction("queryLog", string(keys))
		if err != nil {
			resp := model.Response{
				Code: 0,
				Msg:  err.Error(),
				Data: nil,
			}
			respByte, _ := json.Marshal(&resp)
			encResp, _ := utils.Encrypt(sdkpool.SdkPoll[param.ChaincodeName].Public, string(respByte))

			c.String(200, hex.EncodeToString(encResp))
			return
		}
	}
	log.Println("请求结果: ", string(result))

	resp := model.Response{
		Code: 1,
		Msg:  "success",
		Data: string(result),
	}
	respByte, _ := json.Marshal(&resp)
	encResp, _ := utils.Encrypt(sdkpool.SdkPoll[param.ChaincodeName].Public, string(respByte))

	c.String(200, hex.EncodeToString(encResp))
}

// 验证
func Check(c *gin.Context) {
	// 参数解析
	param := model.Request{}
	if err := c.ShouldBindJSON(&param); err != nil {
		resp := model.Response{
			Code: 0,
			Msg:  "参数解析错误: " + err.Error(),
			Data: nil,
		}
		c.JSON(200, resp)
		return
	}

	// 检查链码sdk是否在连接池中初始化
	if _, ok := sdkpool.SdkPoll[param.ChaincodeName]; !ok {
		resp := model.Response{
			Code: 0,
			Msg:  param.ChaincodeName + " sdk未初始化",
			Data: nil,
		}
		c.JSON(200, resp)
		return
	}

	// 解析16进制字符串为[]byte
	hexByte, err := hex.DecodeString(param.Data)
	if err != nil {
		resp := model.Response{
			Code: 0,
			Msg:  err.Error(),
			Data: nil,
		}
		respByte, _ := json.Marshal(&resp)
		encResp, _ := utils.Encrypt(sdkpool.SdkPoll[param.ChaincodeName].Public, string(respByte))

		c.String(200, hex.EncodeToString(encResp))
		return
	}

	// sm2解密请求数据
	decData, err := utils.Decrypt(sdkpool.SdkPoll[param.ChaincodeName].Private, hexByte)
	if err != nil {
		resp := model.Response{
			Code: 0,
			Msg:  err.Error(),
			Data: nil,
		}
		respByte, _ := json.Marshal(&resp)
		encResp, _ := utils.Encrypt(sdkpool.SdkPoll[param.ChaincodeName].Public, string(respByte))

		c.String(200, hex.EncodeToString(encResp))
		return
	}
	log.Println("解密结果[data]: ", decData)

	data := new(model.CheckData)
	if err := json.Unmarshal([]byte(decData), data); err != nil {
		resp := model.Response{
			Code: 0,
			Msg:  err.Error(),
			Data: nil,
		}
		respByte, _ := json.Marshal(&resp)
		encResp, _ := utils.Encrypt(sdkpool.SdkPoll[param.ChaincodeName].Public, string(respByte))

		c.String(200, hex.EncodeToString(encResp))
		return
	}

	keys, err := json.Marshal(data.Keys)
	if err != nil {
		resp := model.Response{
			Code: 0,
			Msg:  err.Error(),
			Data: nil,
		}
		respByte, _ := json.Marshal(&resp)
		encResp, _ := utils.Encrypt(sdkpool.SdkPoll[param.ChaincodeName].Public, string(respByte))

		c.String(200, hex.EncodeToString(encResp))
		return
	}
	log.Println(string(keys))
	compares, err := json.Marshal(data.Compares)
	if err != nil {
		resp := model.Response{
			Code: 0,
			Msg:  err.Error(),
			Data: nil,
		}
		respByte, _ := json.Marshal(&resp)
		encResp, _ := utils.Encrypt(sdkpool.SdkPoll[param.ChaincodeName].Public, string(respByte))

		c.String(200, hex.EncodeToString(encResp))
		return
	}
	log.Println(string(compares))
	// sdk发起交易
	var result = []byte{}
	result, err = sdkpool.SdkPoll[param.ChaincodeName].Contract.SubmitTransaction("check", string(keys), string(compares), data.Content, data.CheckType)
	if err != nil {
		log.Println(err)
		// todo 重新初始化sdk并发起请求
		sdk_ := sdkpool.SdkPoll[param.ChaincodeName]
		sdk_.InitSdk()
		sdkpool.SdkPoll[param.ChaincodeName] = sdk_

		result, err = sdkpool.SdkPoll[param.ChaincodeName].Contract.SubmitTransaction("check", string(keys), string(compares), data.Content, data.CheckType)
		if err != nil {
			resp := model.Response{
				Code: 0,
				Msg:  err.Error(),
				Data: nil,
			}
			respByte, _ := json.Marshal(&resp)
			encResp, _ := utils.Encrypt(sdkpool.SdkPoll[param.ChaincodeName].Public, string(respByte))

			c.String(200, hex.EncodeToString(encResp))
			return
		}
	}
	log.Println("请求结果: ", string(result))

	resp := model.Response{
		Code: 1,
		Msg:  "success",
		Data: string(result),
	}
	respByte, _ := json.Marshal(&resp)
	encResp, _ := utils.Encrypt(sdkpool.SdkPoll[param.ChaincodeName].Public, string(respByte))

	c.String(200, hex.EncodeToString(encResp))
}
