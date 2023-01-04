package nlibgo

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/borerer/nlib-go/har"
	"github.com/borerer/nlib-go/logs"
	"go.uber.org/zap"
)

func init() {
	// SetEndpoint(os.Getenv("NLIB_SERVER"))
	SetEndpoint(os.Getenv("NLIB_SERVER_DEV"))
	SetAppID("nlib-go")
	SetDebugMode(true)
	err := Connect()
	if err != nil {
		panic(err)
	}
}

func get(url string) (int, string, map[string][]string) {
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	buf, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	code := res.StatusCode
	content := string(buf)
	headers := res.Header
	return code, content, headers
}

func TestGetKey(t *testing.T) {
	value, err := GetKey("some_key")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(value)
}

func TestSetKey(t *testing.T) {
	err := SetKey("some_key", "some_value")
	if err != nil {
		t.Fatal(err)
	}
}

// func TestGetFile(t *testing.T) {
// 	client := getClient()
// 	res, err := client.GetFile("abc.txt")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	defer res.Close()
// 	buf, err := io.ReadAll(res)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	t.Log(string(buf))
// }

// func TestPutFile(t *testing.T) {
// 	client := getClient()
// 	err := client.PutFile("abc.txt", strings.NewReader("hi file!"))
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// }

// func TestRegisterSimpleFunction(t *testing.T) {
// 	ch := make(chan bool)
// 	err := RegisterFunction("ping", func(in nlibshared.SimpleFunctionIn) nlibshared.SimpleFunctionOut {
// 		go func() {
// 			time.Sleep(time.Millisecond)
// 			ch <- true
// 		}()
// 		return "pong"
// 	})
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	res := get(fmt.Sprintf("%s/api/app/%s/ping", GetEndpoint(), GetAppID()))
// 	if res != "pong" {
// 		t.Fatal("expect res to be pong")
// 	}
// 	<-ch
// }

// func TestRegisterSimpleFunctionWithParams(t *testing.T) {
// 	ch := make(chan bool)
// 	err := RegisterFunction("add", func(in nlibshared.SimpleFunctionIn) nlibshared.SimpleFunctionOut {
// 		sa := in["a"].(string)
// 		sb := in["b"].(string)
// 		a, _ := strconv.Atoi(sa)
// 		b, _ := strconv.Atoi(sb)
// 		go func() {
// 			time.Sleep(time.Millisecond)
// 			ch <- true
// 		}()
// 		return a + b
// 	})
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	res := get(fmt.Sprintf("%s/api/app/%s/add?a=3&b=4", GetEndpoint(), GetAppID()))
// 	if res != "7" {
// 		t.Fatal("expect res to be 7, but got:", res)
// 	}
// 	<-ch
// }

func TestRegisterFunction(t *testing.T) {
	ch := make(chan bool)
	err := RegisterFunction("summary", func(in *FunctionIn) (*FunctionOut, error) {
		logs.Info("", zap.Any("in", in))
		go func() {
			time.Sleep(time.Millisecond)
			ch <- true
		}()
		return har.Text(in.Method + " " + in.URL), nil
	})
	if err != nil {
		t.Fatal(err)
	}
	code, content, headers := get(fmt.Sprintf("%s/api/app/%s/summary", GetEndpoint(), GetAppID()))
	if code != 200 {
		t.Fatal("expect code to be 200, but got:", code)
	}
	if content != "GET /api/app/nlib-go/summary" {
		t.Fatal("expect content to be GET /api/app/nlib-go/summary, but got:", content)
	}
	contentType := headers["Content-Type"]
	if contentType[0] != "text/plain" {
		t.Fatal("expect content type to be text/plain, but got:", contentType)
	}
	<-ch
}

func TestLogLevels(t *testing.T) {
	if err := Debug("debug from nlib-go"); err != nil {
		t.Fatal(err)
	}
	if err := Info("info from nlib-go"); err != nil {
		t.Fatal(err)
	}
	if err := Warn("warn from nlib-go"); err != nil {
		t.Fatal(err)
	}
	if err := Error("error from nlib-go"); err != nil {
		t.Fatal(err)
	}
	if err := Fatal("fatal from nlib-go"); err != nil {
		t.Fatal(err)
	}
}

func TestLogWithDetails(t *testing.T) {
	if err := Info("info from nlib-go", "who", "me", "happy", true, "birth", 1992, "hobby", []string{"homelab", "badminton"}); err != nil {
		t.Fatal(err)
	}
}
