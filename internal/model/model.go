package model

type Request struct {
	ChaincodeName string `json:"chaincodeName" binding:"required"`
	Data          string `json:"data" binding:"required"`
}

// 链码初始化方法请求体
type InitData struct {
	DocType string `json:"docType" binding:"required"`
}

// 链码添加方法请求体
type AddData struct {
	Keys    []string `json:"keys" binding:"required"`
	Content string   `json:"content" binding:"required"`
}

// 链码添加方法请求体
type PageData struct {
	Keys     []string `json:"keys" binding:"required"`
	Pagesize string   `json:"pagesize" binding:"required"`
	Nextmark string   `json:"nextmark"`
}

// 链码删除方法请求体
type KeysData struct {
	Keys []string `json:"keys" binding:"required"`
}

// 链码验证方法请求体
type CheckData struct {
	Keys      []string `json:"keys" binding:"required"`
	Compares  []string `json:"compares" binding:"required"`
	Content   string   `json:"content" binding:"required"`
	CheckType string   `json:"checkType" binding:"required"`
}

// 响应体
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
