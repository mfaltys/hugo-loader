package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/unixvoid/glogger"
	"gopkg.in/gcfg.v1"
)

type Config struct {
	Hloader struct {
		Loglevel   string
		S3endpoint string
		Authfile   string
		Sourcedir  string
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

	// chdir and run hugo packager
	os.Chdir(config.Hloader.Sourcedir)
	cmd := exec.Command("hugo")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		glogger.Error.Println("error running 'hugo'")
	}
	// print output
	glogger.Debug.Printf("%s\n", out.String())
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
