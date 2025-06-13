package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.SugaredLogger

func Init() error {
	file, err := os.Create("logs.log")
	if err != nil {
		return err
	}

	encoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	core := zapcore.NewCore(encoder, zapcore.AddSync(file), zapcore.InfoLevel)
	Log = zap.New(core).Sugar()
	return nil
}

func StartLog() {
	err := Init()
	if err != nil {
		panic("error initializing worker: " + err.Error())
	}
	defer Log.Sync()

}
