package model

type LogStruct struct {
	Datetime     string      `json:"datetime"`
	Timestamp    int32       `json:"timestamp"`
	UniqueRemark string      `json:"unique_remark"`
	CnRemark     string      `json:"cn_remark"`
	Project      string      `json:"project"`
	UserId       int32       `json:"user_id"`
	Path         string      `json:"path"`
	Module       string      `json:"module"`
	Host         string      `json:"host"`
	Url          string      `json:"url"`
	Level        string      `json:"level"`
	Context      interface{} `json:"context"`
	Backtrace    interface{} `json:"backtrace"`
	PostData     interface{} `json:"postData"`
	GetData      interface{} `json:"getData"`
}
