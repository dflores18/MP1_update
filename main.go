package main

import (
	"./buildTcp"
	"reflect"
)

type Config struct {
	maxDelay int
	minDelay int
	id       int
	ip       string
	port     string
}

/*var config2 Config
config2.maxDelay = 4
config2.minDelay = 1
config2.id = 1
config2.ip = "127.0.0.1"
config2.port = "1234"

var config1 Config
config1.maxDelay = 4
config1.minDelay = 1
config1.id = 2
config.ip = "127.0.0.1"
config.port = "8001"*/

func main() {
	//filled Config struct manually; couldn't implement readFile in time
	//we know it had to be read line by line and then filled out the struct
	c := Config{4, 1, 2, "127.0.0.1", ":8081"}
	v := reflect.ValueOf(c)
	//for loop to iterate through contents of Config struct
	for i := 0; i < v.NumField(); i++ {
		//imported buildTcp package to use/create clients for different processes in configuration.txt
		buildTcp.Send(5)
	}
}
