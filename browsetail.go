package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {

	filePathPtr := flag.String("file", "", "File for showing log from")
	//portPtr := flag.Int("port", 4040, "-port number")

	flag.Parse()
	if *filePathPtr == "" {
		log.Fatal("-file parameter is missing")
		return
	}

	if _, err := os.Stat(*filePathPtr); os.IsNotExist(err) {
		fmt.Println("File does not exist :"+*filePathPtr)
		return
	}

	//fmt.Println(*filePathPtr)
	//fmt.Println(*portPtr)

	cmd := exec.Command("tail", "-f", *filePathPtr)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}


/*
	go build browsetail.go && ./browsetail -file=/var/log/mysql/mysql.log

	browsetail -file=/var/log/mysql/mysql.log -port=2343
	- check if port is free
	- if port not provided 4040 will be default port
	- -file parameter is mandatory, and check file is exist
	- -username=asdfsdf
	- -password=asdfsdf  if username and password provided both must be provided

*/