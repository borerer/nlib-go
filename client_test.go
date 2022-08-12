package nlibgo

import (
	"testing"
	"time"
)

func TestAddLogs(t *testing.T) {
	client := NewClient("localhost:9502", "nlib-go-test")
	client.LogMessage("hello from nlib-go")
}

func TestRegisterFunction(t *testing.T) {
	client := NewClient("localhost:9502", "nlib-go-test")
	client.RegisterFunction(func(in string) string {
		return "pong"
	}, RegisterFunctionOptions{
		FuncName: "ping",
	})
	time.Sleep(time.Second * 100)
}

func TestRegisterFunction_WithParams(t *testing.T) {
	client := NewClient("localhost:9502", "nlib-go-test")
	client.RegisterFunction(func(in string) string {
		return "hello " + in
	}, RegisterFunctionOptions{
		FuncName: "hi",
	})
	time.Sleep(time.Second * 100)
}
