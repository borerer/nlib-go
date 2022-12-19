package nlibgo

import (
	"testing"
)

func TestLogLevels(t *testing.T) {
	client := getClient()
	if err := client.Debug("debug from nlib-go"); err != nil {
		t.Fatal(err)
	}
	if err := client.Info("info from nlib-go"); err != nil {
		t.Fatal(err)
	}
	if err := client.Warn("warn from nlib-go"); err != nil {
		t.Fatal(err)
	}
	if err := client.Error("error from nlib-go"); err != nil {
		t.Fatal(err)
	}
	if err := client.Fatal("fatal from nlib-go"); err != nil {
		t.Fatal(err)
	}
}

func TestLogWithDetails(t *testing.T) {
	client := getClient()
	if err := client.Info("info from nlib-go", "url", "https://www.baidu.com", "user", "some_fancy_user_name"); err != nil {
		t.Fatal(err)
	}
}
