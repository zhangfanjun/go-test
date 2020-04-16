package vo

import "time"

type ResponseVO struct {
	RequestId    string      `json:"requestId"`
	Code         int32       `json:"code"`
	UserMsg      string      `json:"userMsg"`
	DeveloperMsg string      `json:"developerMsg"`
	Timestamp    int64       `json:"timestamp"`
	Content      interface{} `json:"content"`
}

func (r *ResponseVO) BuildOk(content interface{}) {
	r.Code = 200
	r.Timestamp = time.Now().Unix() * 1000
	r.Content = content
}
func (r *ResponseVO) BuildFail() {
	r.Code = 400
	r.Timestamp = time.Now().Unix() * 1000
}
func (r *ResponseVO) BuildFailByMsg(msg string) {
	r.Code = 400
	r.Timestamp = time.Now().Unix() * 1000
	r.DeveloperMsg = msg
}
