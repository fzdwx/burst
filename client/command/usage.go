package command

import "fmt"

func usage() {
	fmt.Println("Usage:")
	fmt.Println("  u: show usage")
	fmt.Println("  q: quit")
	fmt.Println("  ap: add proxy ")
	fmt.Println("      example: ap tcp::8888 tcp:192.168.1.1:9999 tcp:11.22.33.44:5555 ...")
}
