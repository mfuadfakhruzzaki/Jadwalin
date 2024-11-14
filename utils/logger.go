package utils

import (
	"log"
	"os"
)

var (
    InfoLog  *log.Logger
    ErrorLog *log.Logger
)

func InitLogger() {
    file, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
    if err != nil {
        log.Fatal(err)
    }

    InfoLog = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
    ErrorLog = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func LogInfo(message string) {
    InfoLog.Println(message)
}

func LogError(message string) {
    ErrorLog.Println(message)
}
