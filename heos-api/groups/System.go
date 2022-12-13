package groups

var SystemGroup = "system"
var AccountCheck = "check_account"

type HeosAccountCheckResponse struct {
	Heos AccountCheckResponse `json:"heos"`
}

type AccountCheckResponse struct {
	Command string `json:"command"`
	Result  string `json:"result"`
	Message string `json:"message"`
}
