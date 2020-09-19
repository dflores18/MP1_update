package configFile

/*
import (
	"bufio"
	"os"
	"strings"
)
*/

type Config struct {
	MaxDelay int
	MinDelay int
	Id       int
	Ip       string
	Port     string
}

func readData() Config {
	var config Config
	config.MaxDelay = 4
	config.MinDelay = 1
	config.Id = 1
	config.Ip = "127.0.0.1"
	config.Port = "1234"

	var config1 Config
	config1.MaxDelay = 4
	config1.MinDelay = 1
	config1.Id = 2
	config.Ip = "127.0.0.1"
	config.Port = "8001"

	return config
}
