package main

import "github.com/fzdwx/burst/cmd"

func main() {
	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
