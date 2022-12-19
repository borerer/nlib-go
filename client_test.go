package nlibgo

import (
	"io"
	"strings"
	"testing"
	"time"
)

func getClient() *Client {
	return NewClient("https://nlib.home.iloahz.com", "nlib-go")
	// return NewClient(os.Getenv("NLIB_SERVER"), "nlib-go")
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
	client.RegisterFunction("ping", func(in map[string]interface{}) interface{} {
		return "pong"
	})
	time.Sleep(time.Second * 100)
}

func TestRegisterFunction_WithParams(t *testing.T) {
	client := getClient()
	client.RegisterFunction("hi", func(in map[string]interface{}) interface{} {
		return "hello " + in["name"].(string)
	})
	time.Sleep(time.Second * 100)
}
