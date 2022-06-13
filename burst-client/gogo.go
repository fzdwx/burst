package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
)

func main() {

	buffer := bytes.NewBuffer([]byte("dnByaXgtY29t"))
	decoder := base64.NewDecoder(base64.StdEncoding, buffer)
	all, _ := ioutil.ReadAll(decoder)
	fmt.Println(string(all))
}
