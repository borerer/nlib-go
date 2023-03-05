package nlibgo

import (
	"io"

	"github.com/borerer/nlib-go/logs"
	nlibshared "github.com/borerer/nlib-shared/go"
)

var (
	sharedClient *Client
)

func init() {
	sharedClient = &Client{}
}

func SetEndpoint(endpoint string) {
	sharedClient.Endpoint = endpoint
}

func SetAppID(appID string) {
	sharedClient.AppID = appID
}

func SetDebugMode(debugMode bool) {
	sharedClient.SetDebugMode(debugMode)
}

func SetLogsCallerSkip(skip int) {
	logs.SetCallerSkipForApp(skip)
}

func GetEndpoint() string {
	return sharedClient.Endpoint
}

func GetAppID() string {
	return sharedClient.AppID
}

func Connect() error {
	return sharedClient.Connect()
}

// KV

func GetKey(key string) (string, error) {
	return sharedClient.GetKey(key)
}

func GetJSON(key string, res interface{}) error {
	return sharedClient.GetJSON(key, res)
}

func SetKey(key string, value string) error {
	return sharedClient.SetKey(key, value)
}

func MustGetKey(key string) string {
	return sharedClient.MustGetKey(key)
}

func MustGetJSON(key string, res interface{}) {
	sharedClient.MustGetJSON(key, res)
}

func MustSetKey(key string, value string) {
	sharedClient.MustSetKey(key, value)
}

// Files

func GetFile(filename string) (io.ReadCloser, error) {
	return sharedClient.GetFile(filename)
}

func PutFile(filename string, reader io.Reader) error {
	return sharedClient.PutFile(filename, reader)
}

func DeleteFile(filename string) error {
	return sharedClient.DeleteFile(filename)
}

func MustGetFile(filename string) io.ReadCloser {
	return sharedClient.MustGetFile(filename)
}

func MustPutFile(filename string, reader io.Reader) {
	sharedClient.MustPutFile(filename, reader)
}

func MustDeleteFile(filename string) {
	sharedClient.MustDeleteFile(filename)
}

// Logs

func Debug(message string, args ...interface{}) error {
	return sharedClient.Debug(message, args...)
}

func Info(message string, args ...interface{}) error {
	return sharedClient.Info(message, args...)
}

func Warn(message string, args ...interface{}) error {
	return sharedClient.Warn(message, args...)
}

func Error(message string, args ...interface{}) error {
	return sharedClient.Error(message, args...)
}

func Fatal(message string, args ...interface{}) error {
	return sharedClient.Fatal(message, args...)
}

// Functions

func RegisterFunction(funcName string, f func(*nlibshared.Request) *nlibshared.Response) error {
	return sharedClient.RegisterFunction(funcName, f)
}
