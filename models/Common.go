package models

//公共返回json 结构
type ReturnJson struct {
	Code 	int `json:"code"`
	Message string `json:"message"`
	Data	interface{} `json:"data"`
}

