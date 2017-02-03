package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/unixvoid/glogger"
	"gopkg.in/gcfg.v1"
)

type Config struct {
	Hloader struct {
		Loglevel   string
		S3endpoint string
		Authfile   string
	}
}

var (
	config = Config{}
)

func main() {
	// read in config
	readConf()

	// init logger
	initLogger()
}

func readConf() {
	err := gcfg.ReadFileInto(&config, "config.gcfg")
	if err != nil {
		fmt.Printf("Could not load config.gcfg, error: %s\n", err)
		return
	}
}

func initLogger() {
	// init logger
	if config.Hloader.Loglevel == "debug" {
		glogger.LogInit(os.Stdout, os.Stdout, os.Stdout, os.Stderr)
	} else if config.Hloader.Loglevel == "cluster" {
		glogger.LogInit(os.Stdout, os.Stdout, ioutil.Discard, os.Stderr)
	} else if config.Hloader.Loglevel == "info" {
		glogger.LogInit(os.Stdout, ioutil.Discard, ioutil.Discard, os.Stderr)
	} else {
		glogger.LogInit(ioutil.Discard, ioutil.Discard, ioutil.Discard, os.Stderr)
	}
}
