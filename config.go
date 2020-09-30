package main

import (
	"os"

	"github.com/jessevdk/go-flags"
)

type config struct {
	Domain string `short:"d" long:"domain" description:"domain to crawl" required:"true"`

	NumWorkers  uint `short:"n" long:"workers" description:"number of concurrent workers" default:"1"`
	DialTimeout uint `short:"t" long:"timeout" description:"Maximum dial timeout duration in seconds" default:"15"`

	LogLevel string `long:"loglevel" description:"log level {trace, debug, info, error, critical}" default:"info"`
}

func loadConfig() (*config, error) {
	conf := config{}
	_, err := flags.Parse(&conf)
	if err != nil {
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			os.Exit(0)
		}
		return nil, err
	}

	return &conf, nil
}
