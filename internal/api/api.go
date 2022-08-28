// 请求流程:解析参数=>解密参数=>请求链码=>加密请求结果=>返回加密结果
//
package api

import (
	"encoding/hex"
	"generalapp/internal/model"
	"generalapp/pkg/sdkpool"
	"generalapp/pkg/utils"
	"log"

	"github.com/gin-gonic/gin"
)

// 初始化链码
func InitChaincode(c *gin.Context) {
	// 参数解析
	param := model.InitChaincodeReuqest{}
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
	hex_byte, err := hex.DecodeString(param.DocType)
	if err != nil {
		resp := model.Response{
			Code: 0,
			Msg:  err.Error(),
			Data: nil,
		}
		c.JSON(200, resp)
		return
	}

	// sm2解密请求数据
	decParam, err := utils.Decrypt(sdkpool.SdkPoll[param.ChaincodeName].Private, hex_byte)
	if err != nil {
		resp := model.Response{
			Code: 0,
			Msg:  err.Error(),
			Data: nil,
		}
		c.JSON(200, resp)
		return
	}
	log.Println("解密结果: ", decParam)

	// sdk发起交易
	var result = []byte{}
	result, err = sdkpool.SdkPoll[param.ChaincodeName].Contract.SubmitTransaction("initChaincode", decParam)
	if err != nil {
		// todo 重新初始化sdk并发起请求
		sdk_ := sdkpool.SdkPoll[param.ChaincodeName]
		sdk_.InitSdk()
		sdkpool.SdkPoll[param.ChaincodeName] = sdk_

		result, err = sdkpool.SdkPoll[param.ChaincodeName].Contract.SubmitTransaction("initChaincode", decParam)
		if err != nil {
			resp := model.Response{
				Code: 0,
				Msg:  err.Error(),
				Data: nil,
			}
			c.JSON(200, resp)
			return
		}
	}
	log.Println("请求结果: ", string(result))

	resp := model.Response{
		Code: 1,
		Msg:  "success",
		Data: result,
	}
	c.JSON(200, resp)
}

// 添加
func Add(c *gin.Context) {
	// 参数解析
	param := model.AddReuqest{}
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
	keyByte, err := hex.DecodeString(param.Keys)
	if err != nil {
		resp := model.Response{
			Code: 0,
			Msg:  err.Error(),
			Data: nil,
		}
		c.JSON(200, resp)
		return
	}

	// sm2解密请求数据
	keyDec, err := utils.Decrypt(sdkpool.SdkPoll[param.ChaincodeName].Private, keyByte)
	if err != nil {
		resp := model.Response{
			Code: 0,
			Msg:  err.Error(),
			Data: nil,
		}
		c.JSON(200, resp)
		return
	}
	log.Println("解密结果[Keys]: ", keyDec)

	// 解析16进制字符串为[]byte
	contentByte, err := hex.DecodeString(param.Content)
	if err != nil {
		resp := model.Response{
			Code: 0,
			Msg:  err.Error(),
			Data: nil,
		}
		c.JSON(200, resp)
		return
	}

	// sm2解密请求数据
	contentDec, err := utils.Decrypt(sdkpool.SdkPoll[param.ChaincodeName].Private, contentByte)
	if err != nil {
		resp := model.Response{
			Code: 0,
			Msg:  err.Error(),
			Data: nil,
		}
		c.JSON(200, resp)
		return
	}
	log.Println("解密结果[Content]: ", contentDec)

	// 解析请求数据

	// sdk发起交易
	var result = []byte{}
	result, err = sdkpool.SdkPoll[param.ChaincodeName].Contract.SubmitTransaction("add", keyDec, contentDec)
	if err != nil {
		log.Println(err)
		// todo 重新初始化sdk并发起请求
		sdk_ := sdkpool.SdkPoll[param.ChaincodeName]
		sdk_.InitSdk()
		sdkpool.SdkPoll[param.ChaincodeName] = sdk_

		result, err = sdkpool.SdkPoll[param.ChaincodeName].Contract.SubmitTransaction("add", keyDec, contentDec)
		if err != nil {
			resp := model.Response{
				Code: 0,
				Msg:  err.Error(),
				Data: nil,
			}
			c.JSON(200, resp)
			return
		}
	}
	log.Println("请求结果: ", string(result))

	resp := model.Response{
		Code: 1,
		Msg:  "success",
		Data: result,
	}
	c.JSON(200, resp)
}

// 更新
func Update(c *gin.Context) {
	// 参数解析
	param := model.AddReuqest{}
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
	keyByte, err := hex.DecodeString(param.Keys)
	if err != nil {
		resp := model.Response{
			Code: 0,
			Msg:  err.Error(),
			Data: nil,
		}
		c.JSON(200, resp)
		return
	}

	// sm2解密请求数据
	keyDec, err := utils.Decrypt(sdkpool.SdkPoll[param.ChaincodeName].Private, keyByte)
	if err != nil {
		resp := model.Response{
			Code: 0,
			Msg:  err.Error(),
			Data: nil,
		}
		c.JSON(200, resp)
		return
	}
	log.Println("解密结果[Keys]: ", keyDec)

	// 解析16进制字符串为[]byte
	contentByte, err := hex.DecodeString(param.Content)
	if err != nil {
		resp := model.Response{
			Code: 0,
			Msg:  err.Error(),
			Data: nil,
		}
		c.JSON(200, resp)
		return
	}

	// sm2解密请求数据
	contentDec, err := utils.Decrypt(sdkpool.SdkPoll[param.ChaincodeName].Private, contentByte)
	if err != nil {
		resp := model.Response{
			Code: 0,
			Msg:  err.Error(),
			Data: nil,
		}
		c.JSON(200, resp)
		return
	}
	log.Println("解密结果[Content]: ", contentDec)

	// 解析请求数据

	// sdk发起交易
	var result = []byte{}
	result, err = sdkpool.SdkPoll[param.ChaincodeName].Contract.SubmitTransaction("update", keyDec, contentDec)
	if err != nil {
		log.Println(err)
		// todo 重新初始化sdk并发起请求
		sdk_ := sdkpool.SdkPoll[param.ChaincodeName]
		sdk_.InitSdk()
		sdkpool.SdkPoll[param.ChaincodeName] = sdk_

		result, err = sdkpool.SdkPoll[param.ChaincodeName].Contract.SubmitTransaction("update", keyDec, contentDec)
		if err != nil {
			resp := model.Response{
				Code: 0,
				Msg:  err.Error(),
				Data: nil,
			}
			c.JSON(200, resp)
			return
		}
	}
	log.Println("请求结果: ", string(result))

	resp := model.Response{
		Code: 1,
		Msg:  "success",
		Data: result,
	}
	c.JSON(200, resp)
}

// 删除
func Delete(c *gin.Context) {
	// 参数解析
	param := model.Reuqest{}
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
	keyByte, err := hex.DecodeString(param.Keys)
	if err != nil {
		resp := model.Response{
			Code: 0,
			Msg:  err.Error(),
			Data: nil,
		}
		c.JSON(200, resp)
		return
	}

	// sm2解密请求数据
	keyDec, err := utils.Decrypt(sdkpool.SdkPoll[param.ChaincodeName].Private, keyByte)
	if err != nil {
		resp := model.Response{
			Code: 0,
			Msg:  err.Error(),
			Data: nil,
		}
		c.JSON(200, resp)
		return
	}
	log.Println("解密结果[Keys]: ", keyDec)

	// sdk发起交易
	var result = []byte{}
	result, err = sdkpool.SdkPoll[param.ChaincodeName].Contract.SubmitTransaction("delete", keyDec)
	if err != nil {
		log.Println(err)
		// todo 重新初始化sdk并发起请求
		sdk_ := sdkpool.SdkPoll[param.ChaincodeName]
		sdk_.InitSdk()
		sdkpool.SdkPoll[param.ChaincodeName] = sdk_

		result, err = sdkpool.SdkPoll[param.ChaincodeName].Contract.SubmitTransaction("delete", keyDec)
		if err != nil {
			resp := model.Response{
				Code: 0,
				Msg:  err.Error(),
				Data: nil,
			}
			c.JSON(200, resp)
			return
		}
	}
	log.Println("请求结果: ", string(result))

	resp := model.Response{
		Code: 1,
		Msg:  "success",
		Data: result,
	}
	c.JSON(200, resp)
}

// 查询
func Query(c *gin.Context) {
	// 参数解析
	param := model.Reuqest{}
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
	keyByte, err := hex.DecodeString(param.Keys)
	if err != nil {
		resp := model.Response{
			Code: 0,
			Msg:  err.Error(),
			Data: nil,
		}
		c.JSON(200, resp)
		return
	}

	// sm2解密请求数据
	keyDec, err := utils.Decrypt(sdkpool.SdkPoll[param.ChaincodeName].Private, keyByte)
	if err != nil {
		resp := model.Response{
			Code: 0,
			Msg:  err.Error(),
			Data: nil,
		}
		c.JSON(200, resp)
		return
	}
	log.Println("解密结果[Keys]: ", keyDec)

	// sdk发起交易
	var result = []byte{}
	result, err = sdkpool.SdkPoll[param.ChaincodeName].Contract.SubmitTransaction("query", keyDec)
	if err != nil {
		log.Println(err)
		// todo 重新初始化sdk并发起请求
		sdk_ := sdkpool.SdkPoll[param.ChaincodeName]
		sdk_.InitSdk()
		sdkpool.SdkPoll[param.ChaincodeName] = sdk_

		result, err = sdkpool.SdkPoll[param.ChaincodeName].Contract.SubmitTransaction("query", keyDec)
		if err != nil {
			resp := model.Response{
				Code: 0,
				Msg:  err.Error(),
				Data: nil,
			}
			c.JSON(200, resp)
			return
		}
	}
	log.Println("请求结果: ", string(result))

	if result == nil {
		resp := model.Response{
			Code: 1,
			Msg:  "success",
			Data: result,
		}
		c.JSON(200, resp)
		return
	}

	encResult, err := utils.Encrypt(sdkpool.SdkPoll[param.ChaincodeName].Public, string(result))
	if err != nil {
		resp := model.Response{
			Code: 0,
			Msg:  err.Error(),
			Data: nil,
		}
		c.JSON(200, resp)
		return
	}

	resp := model.Response{
		Code: 1,
		Msg:  "success",
		Data: hex.EncodeToString(encResult),
	}
	c.JSON(200, resp)
}

// 查询所有
func QueryAll(c *gin.Context) {
	// 参数解析
	param := model.Reuqest{}
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
	keyByte, err := hex.DecodeString(param.Keys)
	if err != nil {
		resp := model.Response{
			Code: 0,
			Msg:  err.Error(),
			Data: nil,
		}
		c.JSON(200, resp)
		return
	}

	// sm2解密请求数据
	keyDec, err := utils.Decrypt(sdkpool.SdkPoll[param.ChaincodeName].Private, keyByte)
	if err != nil {
		resp := model.Response{
			Code: 0,
			Msg:  err.Error(),
			Data: nil,
		}
		c.JSON(200, resp)
		return
	}
	log.Println("解密结果[Keys]: ", keyDec)

	// sdk发起交易
	var result = []byte{}
	result, err = sdkpool.SdkPoll[param.ChaincodeName].Contract.SubmitTransaction("queryAll", keyDec)
	if err != nil {
		log.Println(err)
		// todo 重新初始化sdk并发起请求
		sdk_ := sdkpool.SdkPoll[param.ChaincodeName]
		sdk_.InitSdk()
		sdkpool.SdkPoll[param.ChaincodeName] = sdk_

		result, err = sdkpool.SdkPoll[param.ChaincodeName].Contract.SubmitTransaction("queryAll", keyDec)
		if err != nil {
			resp := model.Response{
				Code: 0,
				Msg:  err.Error(),
				Data: nil,
			}
			c.JSON(200, resp)
			return
		}
	}
	log.Println("请求结果: ", string(result))

	if result == nil {
		resp := model.Response{
			Code: 1,
			Msg:  "success",
			Data: result,
		}
		c.JSON(200, resp)
	}

	encResult, err := utils.Encrypt(sdkpool.SdkPoll[param.ChaincodeName].Public, string(result))
	if err != nil {
		resp := model.Response{
			Code: 0,
			Msg:  err.Error(),
			Data: nil,
		}
		c.JSON(200, resp)
		return
	}

	resp := model.Response{
		Code: 1,
		Msg:  "success",
		Data: hex.EncodeToString(encResult),
	}
	c.JSON(200, resp)
}

// 分页查询
func QuerysByPagination(c *gin.Context) {
	// 参数解析
	param := model.PageQueryReuqest{}
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
	keyByte, err := hex.DecodeString(param.Keys)
	if err != nil {
		resp := model.Response{
			Code: 0,
			Msg:  err.Error(),
			Data: nil,
		}
		c.JSON(200, resp)
		return
	}

	// sm2解密请求数据
	keyDec, err := utils.Decrypt(sdkpool.SdkPoll[param.ChaincodeName].Private, keyByte)
	if err != nil {
		resp := model.Response{
			Code: 0,
			Msg:  err.Error(),
			Data: nil,
		}
		c.JSON(200, resp)
		return
	}
	log.Println("解密结果[Keys]: ", keyDec)

	// 解析16进制字符串为[]byte
	psByte, err := hex.DecodeString(param.Pagesize)
	if err != nil {
		resp := model.Response{
			Code: 0,
			Msg:  err.Error(),
			Data: nil,
		}
		c.JSON(200, resp)
		return
	}

	// sm2解密请求数据
	psDec, err := utils.Decrypt(sdkpool.SdkPoll[param.ChaincodeName].Private, psByte)
	if err != nil {
		resp := model.Response{
			Code: 0,
			Msg:  err.Error(),
			Data: nil,
		}
		c.JSON(200, resp)
		return
	}
	log.Println("解密结果[pageSize]: ", psDec)

	var nextmarkDec string
	if param.Nextmark != "" {
		// 解析16进制字符串为[]byte
		nextmarkByte, err := hex.DecodeString(param.Nextmark)
		if err != nil {
			resp := model.Response{
				Code: 0,
				Msg:  err.Error(),
				Data: nil,
			}
			c.JSON(200, resp)
			return
		}

		// sm2解密请求数据
		nextmarkDec, err = utils.Decrypt(sdkpool.SdkPoll[param.ChaincodeName].Private, nextmarkByte)
		if err != nil {
			resp := model.Response{
				Code: 0,
				Msg:  err.Error(),
				Data: nil,
			}
			c.JSON(200, resp)
			return
		}
		log.Println("解密结果[nexkMark]: ", nextmarkDec)
	}

	// sdk发起交易
	var result = []byte{}
	result, err = sdkpool.SdkPoll[param.ChaincodeName].Contract.SubmitTransaction("querysByPagination", keyDec, psDec, nextmarkDec)
	if err != nil {
		log.Println(err)
		// todo 重新初始化sdk并发起请求
		sdk_ := sdkpool.SdkPoll[param.ChaincodeName]
		sdk_.InitSdk()
		sdkpool.SdkPoll[param.ChaincodeName] = sdk_

		result, err = sdkpool.SdkPoll[param.ChaincodeName].Contract.SubmitTransaction("querysByPagination", keyDec, psDec, nextmarkDec)
		if err != nil {
			resp := model.Response{
				Code: 0,
				Msg:  err.Error(),
				Data: nil,
			}
			c.JSON(200, resp)
			return
		}
	}
	log.Println("请求结果: ", string(result))

	if result == nil {
		resp := model.Response{
			Code: 1,
			Msg:  "success",
			Data: result,
		}
		c.JSON(200, resp)
	}

	encResult, err := utils.Encrypt(sdkpool.SdkPoll[param.ChaincodeName].Public, string(result))
	if err != nil {
		resp := model.Response{
			Code: 0,
			Msg:  err.Error(),
			Data: nil,
		}
		c.JSON(200, resp)
		return
	}

	resp := model.Response{
		Code: 1,
		Msg:  "success",
		Data: hex.EncodeToString(encResult),
	}
	c.JSON(200, resp)
}

// 查询log
func QueryLog(c *gin.Context) {
	// 参数解析
	param := model.Reuqest{}
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
	keyByte, err := hex.DecodeString(param.Keys)
	if err != nil {
		resp := model.Response{
			Code: 0,
			Msg:  err.Error(),
			Data: nil,
		}
		c.JSON(200, resp)
		return
	}

	// sm2解密请求数据
	keyDec, err := utils.Decrypt(sdkpool.SdkPoll[param.ChaincodeName].Private, keyByte)
	if err != nil {
		resp := model.Response{
			Code: 0,
			Msg:  err.Error(),
			Data: nil,
		}
		c.JSON(200, resp)
		return
	}
	log.Println("解密结果[Keys]: ", keyDec)

	// sdk发起交易
	var result = []byte{}
	result, err = sdkpool.SdkPoll[param.ChaincodeName].Contract.SubmitTransaction("queryLog", keyDec)
	if err != nil {
		log.Println(err)
		// todo 重新初始化sdk并发起请求
		sdk_ := sdkpool.SdkPoll[param.ChaincodeName]
		sdk_.InitSdk()
		sdkpool.SdkPoll[param.ChaincodeName] = sdk_

		result, err = sdkpool.SdkPoll[param.ChaincodeName].Contract.SubmitTransaction("queryLog", keyDec)
		if err != nil {
			resp := model.Response{
				Code: 0,
				Msg:  err.Error(),
				Data: nil,
			}
			c.JSON(200, resp)
			return
		}
	}
	log.Println("请求结果: ", string(result))

	if result == nil {
		resp := model.Response{
			Code: 1,
			Msg:  "success",
			Data: result,
		}
		c.JSON(200, resp)
	}

	encResult, err := utils.Encrypt(sdkpool.SdkPoll[param.ChaincodeName].Public, string(result))
	if err != nil {
		resp := model.Response{
			Code: 0,
			Msg:  err.Error(),
			Data: nil,
		}
		c.JSON(200, resp)
		return
	}

	resp := model.Response{
		Code: 1,
		Msg:  "success",
		Data: hex.EncodeToString(encResult),
	}
	c.JSON(200, resp)
}
