package configFile

/*
import (
	"bufio"
	"os"
	"strings"
)
*/

type Config struct {
	maxDelay int
	minDelay int
	id       int
	ip       string
	port     string
}

func readData() Config{
	var config Config
	config.maxDelay = 4
	config.minDelay = 1
	config.id = 1
	config.ip = "127.0.0.1"
	config.port = "1234"

	var config1 Config
	config1.maxDelay = 4
	config1.minDelay = 1
	config1.id = 2
	config.ip = "127.0.0.1"
	config.port = "8001"

	return config
}

