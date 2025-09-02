package logs

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger
var err error

const CUSTOM_LOG_FORMAT string = "TIME[${time}] PID[${pid}] REQUESTID[${locals:requestid}] RESSTATUS[${status}] - LATENCY[${latency}] METHOD[${method}] PATH[${path}] REFERER[${referer}] PROTOCOL[${protocol}] PORT[${port}] IP[${ip}] IPS[${ips}] HOST[${host}] UA[${ua}] REQHEADERS[${reqHeaders}] REQQUERYPARAMS[${queryParams}] \n URL[${url}]\n REQBODY[${body}] REQHEADER:[${header:}] REQHEADER:[${reqHeader:}] REQQUERY[${query:}] REQFORM[${form:}] REQCOOKIE[${cookie:}] \n RESBODY[${resBody}]\n BYTESSENT[${bytesSent}] BYTESRECEIVED[${bytesReceived}] ROUTE[${route}] ERROR[${error}]  RESPHEADER:[${respHeader:}]  LOCALS:[${locals:}]\n <------------------------------------------------------------------------------------> \n"

func init() {
	config := zap.NewProductionConfig()
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig.StacktraceKey = ""
	log, err = config.Build(zap.AddCallerSkip(1))
	if err != nil {
		log.Error(err.Error())
		return
	}
}

func Info(message string, fields ...zap.Field) {
	log.Info(message, fields...)
}
func Error(message interface{}, fields ...zap.Field) {
	switch v := message.(type) {
	case error:
		log.Error(v.Error(), fields...)
	case string:
		log.Error(v, fields...)
	}
}
