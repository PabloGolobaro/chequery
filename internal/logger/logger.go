package logger

import (
	"go.uber.org/zap"
	"sync"
)

var Global *zap.SugaredLogger
var once sync.Once

func Get() *zap.SugaredLogger {
	once.Do(func() {
		newProduction, err := zap.NewDevelopment()
		if err != nil {
			panic(err)
		}
		Global = newProduction.Sugar()
	})

	return Global
}
