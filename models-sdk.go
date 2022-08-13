package nlibgo

type RegisterFunctionOptions struct {
	FuncName string
	Endpoint string
}

type NLIBFunc func(string) string
