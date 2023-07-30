package config

import (
	"fmt"
	"testing"
)

func TestFromFile(t *testing.T) {
	var config Config
	var content = `
[server]
host = "0.0.0.0"
port = 8080

[database]
host = "127.0.0.1"
port = 5432
username = "xxxx"
password = "xxxx"
`
	config, err := FromString(content)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v", config)
}
