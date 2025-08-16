package logger

import (
	"log"
	"os"
)

type MLogger struct {
}

var (
	Logger *log.Logger
)

func (l *MLogger) Printf(format string, v ...interface{}) {
	// log.Printf("INFO: "+format, v...)
}

func (l *MLogger) Fatal(format string, v ...interface{}) {
	// log.Fatalf("ERROR: "+format, v...)
}

func (l *MLogger) Fatalf(format string, v ...interface{}) {
	// log.Fatalf("ERROR: "+format, v...)
}

func (l *MLogger) Println(v ...interface{}) {
	// log.Println(v...)
}

func init() {
	// file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	// if err != nil {
	// 	log.Fatal("Failed to open log file: ", err)
	// }

	// multiWriter := io.MultiWriter(os.Stdout, file)
	Logger = log.New(os.Stdout, "INFO: ", log.LstdFlags)
}
