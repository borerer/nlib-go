package nlibgo

import (
	"io"
	"os"
	"strings"
	"testing"
	"time"
)

func getClient() *Client {
	return NewClient(os.Getenv("NLIB_SERVER"), "nlib-go")
}

func TestLogs(t *testing.T) {
	client := getClient()
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
	client := getClient()
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
	client := getClient()
	err := client.PutFile("abc.txt", strings.NewReader("hi file!"))
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetKey(t *testing.T) {
	client := getClient()
	value, err := client.GetKey("some_key")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(value)
}

func TestSetKey(t *testing.T) {
	client := getClient()
	err := client.SetKey("some_key", "some_value")
	if err != nil {
		t.Fatal(err)
	}
}

func TestRegisterFunction(t *testing.T) {
	client := getClient()
	client.RegisterFunction("ping", func(in map[string]interface{}) map[string]interface{} {
		return map[string]interface{}{"message": "pong"}
	})
	time.Sleep(time.Second * 100)
}

func TestRegisterFunction_WithParams(t *testing.T) {
	client := getClient()
	client.RegisterFunction("hi", func(in map[string]interface{}) map[string]interface{} {
		return map[string]interface{}{"message": "hello " + in["name"].(string)}
	})
	time.Sleep(time.Second * 100)
}
