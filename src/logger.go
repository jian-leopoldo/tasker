package main

import "os"

func writeLog(logFileName string, logMessage string) {
	messageToWrite := []byte(logMessage)
	err := os.WriteFile(logFileName, messageToWrite, 0644)
	raiseError(err)
}

func raiseError(e error) {
	if e != nil {
		panic(e)
	}
}
