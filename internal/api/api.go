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

func InitChaincode(c *gin.Context) {
	// 参数解析
	param := model.Reuqest{}
	if err := c.ShouldBindJSON(&param); err != nil {
		resp := model.Response{
			Code: 0,
			Msg:  "参数解析错误",
			Data: nil,
		}
		c.JSON(200, resp)
		return
	}

	// 检查链码sdk是否在连接池中初始化
	if _, ok := sdkpool.SdkPoll[param.ChaincodeName]; !ok {
		resp := model.Response{
			Code: 0,
			Msg:  param.ChaincodeName + "sdk未初始化",
			Data: nil,
		}
		c.JSON(200, resp)
		return
	}

	// 16进制获取加密结果
	hex_byte, err := hex.DecodeString(param.Data)
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
	log.Println("解密结果: ", string(decParam))
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

	// 解析请求结果
	ccResponse := new(model.Contract_Reponse)
	if err := json.Unmarshal(result, ccResponse); err != nil {
		resp := model.Response{
			Code: 0,
			Msg:  err.Error(),
			Data: nil,
		}
		c.JSON(200, resp)
		return
	}

	log.Println("请求获取数据: ", ccResponse.Data)
	// 加密请求结果
	encResult, err := utils.Encrypt(sdkpool.SdkPoll[param.ChaincodeName].Public, string(ccResponse.Msg))
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
		Code: ccResponse.Code,
		Msg:  hex.EncodeToString(encResult),
		Data: ccResponse.Data,
	}
	c.JSON(200, resp)
}

func Add(c *gin.Context) {
	// 参数解析
	param := model.Reuqest{}
	if err := c.ShouldBindJSON(&param); err != nil {
		resp := model.Response{
			Code: 0,
			Msg:  "参数解析错误",
			Data: nil,
		}
		c.JSON(200, resp)
		return
	}

	// 检查链码sdk是否在连接池中初始化
	if _, ok := sdkpool.SdkPoll[param.ChaincodeName]; !ok {
		resp := model.Response{
			Code: 0,
			Msg:  param.ChaincodeName + "sdk未初始化",
			Data: nil,
		}
		c.JSON(200, resp)
		return
	}

	// sm2解密请求数据
	decParam, err := utils.Decrypt(sdkpool.SdkPoll[param.ChaincodeName].Private, []byte(param.Data))
	if err != nil {
		resp := model.Response{
			Code: 0,
			Msg:  err.Error(),
			Data: nil,
		}
		c.JSON(200, resp)
		return
	}

	// sdk发起交易
	var result = []byte{}
	result, err = sdkpool.SdkPoll[param.ChaincodeName].Contract.SubmitTransaction("add", decParam)
	if err != nil {
		// todo 重新初始化sdk并发起请求
		sdk_ := sdkpool.SdkPoll[param.ChaincodeName]
		sdk_.InitSdk()
		sdkpool.SdkPoll[param.ChaincodeName] = sdk_

		result, err = sdkpool.SdkPoll[param.ChaincodeName].Contract.SubmitTransaction("add", decParam)
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

	// 加密请求结果
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
		Code: 0,
		Msg:  "",
		Data: encResult,
	}
	c.JSON(200, resp)
}
