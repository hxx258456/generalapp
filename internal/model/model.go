package model

import "time"

// 请求体
type Reuqest struct {
	ChaincodeName string `json:"chaincodeName" binding:"required"`
	Data          string `json:"data" binding:"required"`
}

// 响应体
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// 链码添加方法参数序列化结构体
type Contract_Add_Params struct {
	Params  []string `json:"keys"`
	Content string   `json:"content"`
}

// 链码响应
type Contract_Reponse struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

// 链码分页想用
type Contract_PageResult struct {
	Code     int         `json:"code"`
	Count    int32       `json:"count"`    // 总数量
	Nextmark string      `json:"nextmark"` // 指向下一个的标记
	Data     interface{} `json:"data"`
}

// 链码log响应
type Contract_LogResult struct {
	Record    interface{} `json:"record"`
	TxId      string      `json:"txId"`
	Timestamp time.Time   `json:"timestamp"`
	IsDelete  bool        `json:"isDelete"`
}
