package nlibgo

type AddLogsRequest struct {
	Level   string      `json:"level"`
	Message string      `json:"message"`
	Details interface{} `json:"details"`
}

type WebSocketCallFunctionReq struct {
	FuncName string `json:"func_name" mapstructure:"func_name"`
	Params   string `json:"params" mapstructure:"params"`
}

type WebSocketCallFunctionRes struct {
	FuncName string `json:"func_name" mapstructure:"func_name"`
	Response string `json:"response" mapstructure:"response"`
}

type WebSocketMessage struct {
	MessageID     string      `json:"message_id"`
	PairMessageID string      `json:"pair_message_id"`
	Type          string      `json:"type"`
	SubType       string      `json:"sub_type"`
	Timestamp     int64       `json:"timestamp"`
	Payload       interface{} `json:"payload"`
}

const (
	WebSocketTypeDefault  = "default"
	WebSocketTypeRequest  = "request"
	WebSocketTypeResponse = "response"
)

const (
	WebSocketSubTypeDefault          = "default"
	WebSocketSubTypeStart            = "start"
	WebSocketSubTypeRegisterFunction = "register_function"
	WebSocketSubTypeCallFunction     = "call_function"
)
