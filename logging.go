package main

// the log package - for free-form output
// log/slog package - for structured logging

import (
	"log"
	"bytes"
	"fmt"
	"log/slog"
	"os"
)

func main(){
	// simply invoking functions like 
	// Println from the log packages uses the standard logger, which is already pre-configured 
	// for reasonable logging output to os.Stderr. Additional methods like
	// Fatal* and Panic* will exit the program after alogging

	log.Println("Standard logger")


	// Loggers can be configured with flags to set their output format
	// By default, the standard logger has the log.Ldate and log.Ltime flags set.
	// and these are collected in the log.LstdFlag. We change its flag to emit time in microsecond accuracy
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Println("Log with microsecond accuracy")

	// Adding flags to emit the filename and line number
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Log with file name and line number")


	// It is useful to create a custom logger and pass it around. When creating a custom logger it
	// we can set a prefix to distinguish its oupt from  other loggers.
	mylog := log.New(os.Stdout, "my:", log.LstdFlags)
	mylog.Println("This is my log")

	// setting prefix on existing loggers
	mylog.SetPrefix("ohmy:")
	mylog.Println("from mylog with new prefix")

	// logger can have custom log output targets
	// any io.Writer works

	var buf bytes.Buffer
	buflog := log.New(&buf, "buf:", log.LstdFlags)

	buflog.Println("This is a log entry in the buffer")


	fmt.Println("from buflog:", buf.String())


	// The slog package provides structured log output. For example
	// logging in JSON format is straightforward.
	jsonHandler := slog.NewJSONHandler(os.Stderr, nil)
	myslog := slog.New(jsonHandler)
	myslog.Info("hi there")

	// in addition to the message, slog output can contain an arbitary number of key=value pairs.
	myslog.Info("hello again", "key", "val", "age", 25)
}

