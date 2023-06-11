package main

import (
   "time"
	"errors"
   "github.com/luisfurquim/env"
   "github.com/luisfurquim/goose"
)

type OptionsT struct {
//	Verbose			   goose.Alert `env:"VERBOSE" default:"4"`
	Verbose			   int			`env:"VERBOSE" default:"2"`
}

var ErrDownloading error = errors.New("Error downloading page")

func main() {
   var err  error
   var options OptionsT

	goose.TraceOn()

	err = env.Read(&options)
	if err != nil {
		Goose.Init.Fatalf(1,"Error reading environment variables: %s", err)
	}

	goose.Geese{
		"viewer": &Goose,
//	}.Set(2)
	}.Set(options.Verbose)

	ctx, cancel := RunChrome("chrome", "10547", "about:blank")

	_ = ctx

	time.Sleep(45 * time.Second)

	cancel()
}

