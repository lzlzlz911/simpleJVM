package main

import "fmt"

func main() {
	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Println(cmd.versionFlag)
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		printCmd(cmd)
	}
}
