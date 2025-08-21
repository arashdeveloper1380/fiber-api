package pkg

import (
	"fmt"
	"os"
)

type Logger interface {
	Log(message string)
}

type ConsoleLogger struct{}

func (c *ConsoleLogger) Log(message string) {
	fmt.Println("[Console]", message)
}

type FileLogger struct {
	file *os.File
}

func (fl *FileLogger) Log(message string) {
	fmt.Fprintln(fl.file, "[File]", message)
}

func Factory(loggerType string) Logger {
	if loggerType == "console" {
		return &ConsoleLogger{}
	} else if loggerType == "file" {
		file, _ := os.Create("app.log")
		return &FileLogger{file: file}
	}

	return nil
}

func FactoryWithT(t Logger) Logger {
	if t == new(ConsoleLogger) {
		return &ConsoleLogger{}
	} else if t == new(FileLogger) {
		return &FileLogger{}
	}

	return nil
}

func selectLogger() {
	logger := Factory("console")
	logger.Log("this is sample log")
}

func selectLoggerWithT() {
	logger := FactoryWithT(&ConsoleLogger{})
	logger.Log("test with T")
}
