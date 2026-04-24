package util

import (
	"log"
	"os"
)

var logger *log.Logger

func init() {
	logger = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.LUTC|log.Lmicroseconds)
}

func generalLog(category string, v string) {
	logger.Printf("[%v] %v", category, v)
}

func InfoLog(v string) {
	generalLog("INFO", v)
}

func WarnLog(v string) {
	generalLog("WARN", v)
}

func ErrorLog(v string) {
	generalLog("ERROR", v)
}

func FatalLog(v string) {
	logger.Fatalf("[Fatal] %v", v)
}
