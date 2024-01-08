package logger

import (
	"context"
	"reflect"
	"runtime"

	"github.com/jabardigitalservice/golog/logger"
	"github.com/rindibudiaramdhan/sw-go/constant"
)

type (
	Logger struct {
		Logger *logger.Logger
		Data   logger.LoggerData
	}

	Map map[string]interface{}

	ErrMap struct {
		Messsage   string `json:"message"`
		StackTrace string `json:"stackTrace"`
	}
)

var entries []ErrMap

func InitLogger(service string, appVersion string) *Logger {
	Logger := &Logger{
		Logger: logger.Init(),
		Data: logger.LoggerData{
			Service: service,
			Version: appVersion,
		},
	}

	return Logger
}

func (Logger *Logger) Log(ctx context.Context) *Logger {
	if ctx.Value(constant.CtxUserIDKey) != nil {
		Logger.Data.UserID = ctx.Value(constant.CtxUserIDKey).(string)
	}

	if ctx.Value(constant.CtxSessionIDKey) != nil {
		Logger.Data.SessionID = ctx.Value(constant.CtxSessionIDKey).(string)
	}

	if ctx.Value(constant.CtxClientIDKey) != nil {
		Logger.Data.ClientID = ctx.Value(constant.CtxClientIDKey).(string)
	}

	return Logger
}

func (Logger *Logger) WithFields(info Map) *Logger {
	Logger.Data.AdditionalInfo = info

	return Logger
}

func (Logger *Logger) WithError(funcName interface{}, err error) {
	collectError(&entries, funcName, err)
	Logger.Data.AdditionalInfo = map[string]interface{}{"errors": entries}
}

func (Logger *Logger) SetModule(module string) *Logger {
	Logger.Data.Module = module
	return Logger
}

func (Logger *Logger) SetMethod(method string) *Logger {
	Logger.Data.Method = method
	return Logger
}

func (Logger *Logger) SetCategory(category string) *Logger {
	switch category {
	case constant.LogCategoryApp:
		Logger.Data.Category = logger.LoggerApp
	case constant.LogCategoryRouter:
		Logger.Data.Category = logger.LoggerRouter
	case constant.LogCategoryUsecase:
		Logger.Data.Category = logger.LoggerUsecase
	case constant.LogCategoryExternal:
		Logger.Data.Category = logger.LoggerExternal
	default:
		Logger.Data.Category = logger.LoggerApp
	}

	return Logger
}

func (Logger *Logger) Success() {
	Logger.Logger.Info(&Logger.Data, "success")
}

func (Logger *Logger) Error(err error) {
	entries = []ErrMap{}
	Logger.Logger.Error(&Logger.Data, err)
}

func GetStackTrace(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func collectError(entries *[]ErrMap, funcName interface{}, err error) {
	*entries = append(*entries, ErrMap{
		Messsage:   err.Error(),
		StackTrace: GetStackTrace(funcName),
	})
}
