package nlibgo

import (
	"io"
	"strings"
	"testing"
	"time"
)

func TestLogs(t *testing.T) {
	client := NewClient("http://localhost:9502", "nlib-go-test")
	if err := client.Debug("Debug from nlib-go"); err != nil {
		t.Fatal(err)
	}
	if err := client.Info("Info from nlib-go"); err != nil {
		t.Fatal(err)
	}
	if err := client.Warn("Warn from nlib-go"); err != nil {
		t.Fatal(err)
	}
	if err := client.Error("Error from nlib-go"); err != nil {
		t.Fatal(err)
	}
	if err := client.Fatal("Fatal from nlib-go"); err != nil {
		t.Fatal(err)
	}
}

func TestGetFile(t *testing.T) {
	client := NewClient("http://localhost:9502", "nlib-go-test")
	res, err := client.GetFile("abc.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer res.Close()
	buf, err := io.ReadAll(res)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(buf))
}

func TestPutFile(t *testing.T) {
	client := NewClient("http://localhost:9502", "nlib-go-test")
	err := client.PutFile("abc.txt", strings.NewReader("hi file!"))
	if err != nil {
		t.Fatal(err)
	}
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
