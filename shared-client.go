package nlibgo

import "io"

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

// Files

func GetFile(filename string) (io.ReadCloser, error) {
	return sharedClient.GetFile(filename)
}

func SaveFile(filename string, reader io.Reader) error {
	return sharedClient.SaveFile(filename, reader)
}

func DeleteFile(filename string) error {
	return sharedClient.DeleteFile(filename)
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

func RegisterFunction(funcName string, f interface{}) error {
	return sharedClient.RegisterFunction(funcName, f)
}
