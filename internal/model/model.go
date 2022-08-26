package model

// 链码初始化方法请求体
type InitChaincodeReuqest struct {
	ChaincodeName string `json:"chaincodeName" binding:"required"`
	DocType       string `json:"docType" binding:"required"`
}

// 链码添加方法请求体
type AddReuqest struct {
	ChaincodeName string `json:"chaincodeName" binding:"required"`
	Keys          string `json:"keys" binding:"required"`
	Content       string `json:"content" binding:"required"`
}

// 链码添加方法请求体
type PageQueryReuqest struct {
	ChaincodeName string `json:"chaincodeName" binding:"required"`
	Keys          string `json:"keys" binding:"required"`
	Pagesize      string `json:"pagesize" binding:"required"`
	Nextmark      string `json:"nextmark"`
}

// 链码删除方法请求体
type Reuqest struct {
	ChaincodeName string `json:"chaincodeName" binding:"required"`
	Keys          string `json:"keys" binding:"required"`
}

// 响应体
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// 链码分页想用
type Contract_PageResult struct {
	Code     int         `json:"code"`
	Count    int32       `json:"count"`    // 总数量
	Nextmark string      `json:"nextmark"` // 指向下一个的标记
	Data     interface{} `json:"data"`
}
