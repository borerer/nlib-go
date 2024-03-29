package nlibgo

import (
	"fmt"
	"runtime"

	"github.com/borerer/nlib-go/logs"
	"github.com/borerer/nlib-go/network"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func (c *Client) Debug(message string, args ...interface{}) error {
	return c.log(zap.DebugLevel, message, args...)
}

func (c *Client) Info(message string, args ...interface{}) error {
	return c.log(zap.InfoLevel, message, args...)
}

func (c *Client) Warn(message string, args ...interface{}) error {
	return c.log(zap.WarnLevel, message, args...)
}

func (c *Client) Error(message string, args ...interface{}) error {
	return c.log(zap.ErrorLevel, message, args...)
}

func (c *Client) Fatal(message string, args ...interface{}) error {
	return c.log(zap.FatalLevel, message, args...)
}

func arrayToMap(args ...interface{}) map[string]interface{} {
	res := map[string]interface{}{}
	for i := 0; i+1 < len(args); i += 2 {
		s, ok := args[i].(string)
		if !ok {
			continue
		}
		res[s] = args[i+1]
	}
	return res
}

func (c *Client) logToStdout(level zapcore.Level, message string, details map[string]interface{}) {
	var fields []zapcore.Field
	for k, v := range details {
		fields = append(fields, zap.Any(k, v))
	}
	logs.GetZapLoggerForApp().Log(level, message, fields...)
}

func getStackTrace(skip int) string {
	pc := make([]uintptr, 100)
	n := runtime.Callers(skip+2, pc)
	res := ""
	frames := runtime.CallersFrames(pc[:n])
	for {
		frame, more := frames.Next()
		res = res + fmt.Sprintf("%s:%d\n", frame.File, frame.Line)
		if !more {
			break
		}
	}
	return res
}

func (c *Client) log(level zapcore.Level, message string, args ...interface{}) error {
	details := arrayToMap(args...)
	c.logToStdout(level, message, details)
	details["app_id"] = c.AppID
	details["stacktrace"] = getStackTrace(c.LogsSkip + 1)
	body := map[string]interface{}{
		"level":   level.String(),
		"message": message,
		"details": details,
	}
	req, err := network.NewHTTPRequestBuilder("POST", fmt.Sprintf("%s/api/app/logs/log", c.Endpoint)).Body(body).Build()
	if err != nil {
		return err
	}
	res, err := network.HttpClient.Do(req)
	if err != nil {
		return err
	}
	if !network.StatusOK(res.StatusCode) {
		return fmt.Errorf("http error, path: %s, status code %d", req.URL.Path, res.StatusCode)
	}
	return nil
}
