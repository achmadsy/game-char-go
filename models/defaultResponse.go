package models

type OutStruct struct {
	IsSuccess bool        `json:"isSuccess"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
}

// SetResponse : set template of the return
func SetResponse(s bool, m string, d interface{}) OutStruct {
	out := OutStruct{
		IsSuccess: s,
		Message:   m,
		Data:      d,
	}

	return out
}
