package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	
	"github.com/reiver/go-telnet"
)

func telnetMain(dest string) {
	var caller telnet.Caller = telnet.StandardCaller
	
	if( !strings.Contains(dest,":") ) {
		dest = dest+":23"
	}
	log.Println("Trying to connect to "+dest+" ...")
	telnet.DialToAndCall(dest, caller)
}

func main() {
	if( len(os.Args)>1 ) {
		telnetMain(os.Args[1])
	} else {
		fmt.Println("Usage: telnet hostname")
	}
}