package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"github.com/vooon/esp-ota-server/server"
	"os"
)

var opts struct {
	Bind    string `short:"s" long:"bind" env:"EOBIND" default:":8092"`
	BaseUrl string `short:"u" long:"base-url" env:"EOBASEURL" default:"http://localhost:8092"`
	DataDir string `short:"d" long:"data-dir" env:"EODATADIR" required:"true"`
}

func main() {
	_, err := flags.Parse(&opts)
	if err != nil {
		if flerr, ok := err.(*flags.Error); ok && flerr.Type == flags.ErrHelp {
			os.Exit(0)
		} else {
			fmt.Fatal(err)
		}
	}

	config := server.Config{
		Bind:    opts.Bind,
		BaseUrl: opts.BaseUrl,
		DataDir: opts.DataDir,
	}

	server.Serve(config)
}