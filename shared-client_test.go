package nlibgo

import (
	"os"
	"testing"
	"time"

	nlibshared "github.com/borerer/nlib-shared/go"
)

func init() {
	SetEndpoint(os.Getenv("NLIB_SERVER"))
	// SetEndpoint("https://nlib.home.iloahz.com")
	SetAppID("nlib-go")
	err := Connect()
	if err != nil {
		panic(err)
	}
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

func TestRegisterFunction(t *testing.T) {
	ch := make(chan bool)
	err := RegisterFunction("ping", func(in nlibshared.SimpleFunctionIn) nlibshared.SimpleFunctionOut {
		go func() {
			time.Sleep(time.Second)
			ch <- true
		}()
		return "pong"
	})
	if err != nil {
		t.Fatal(err)
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
