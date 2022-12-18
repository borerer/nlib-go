package nlibgo

type RegisterFunctionOptions struct {
	FuncName string
	Endpoint string
}

type NLIBFunc func(map[string]interface{}) map[string]interface{}
