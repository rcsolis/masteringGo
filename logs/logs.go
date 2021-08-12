package logs
// sample program to use logs

import (
	"fmt"
	"log"
	"log/syslog"
	"os"
	"path/filepath"
)

// start with capital letter to export it
func UsingSystemLogServer(message string){
	// Getting program name from command line arguments
	programName := filepath.Base(os.Args[0])
	// Sends log messages with priority of debug to the facility Mail using a given prefix tag
	sysLog, err := syslog.New(syslog.LOG_DEBUG|syslog.LOG_MAIL, programName)
	if err != nil{
		log.Fatal(err)
	}
	// sets the output destination for the logger
	log.SetOutput(sysLog)
	// send a message
	log.Println("GoLand says:"+message)
}
// Path to custom log
var LOGFILE = "/Users/integradesarrollo/GolandProjects/masteringGo/logs/mGo.log"

func CustomLog(message string){
	// Open a new file
	f, err := os.OpenFile(LOGFILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error:", err )
		log.Fatal("Exit program")
	}
	defer f.Close()
	iLog := log.New(f, os.Args[0] + " ", log.LstdFlags)
	iLog.SetFlags(log.LstdFlags | log.Lshortfile)
	iLog.Println(message)
}