package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/jamiekieranmartin/tryp"
)

const cliVersion = "0.0.1"

const helpMessage = `
tryp is a minimal wrapper for the Google Maps Platform Distance Matrix API.
	tryp v%s

Usage
	tryp -key='my super secret key'

Configurable via TOML
	tryp -key='my super secret key' -config "./my-config.toml"

Write to file
	tryp -key='my super secret key' -out "./result.json"

`

func main() {
	flag.Usage = func() {
		fmt.Printf(helpMessage, cliVersion)
		flag.PrintDefaults()
	}

	// cli arguments
	key := flag.String("key", "", "Google Maps Platform API key")
	path := flag.String("config", "./config.toml", "TOML configuration file")
	out := flag.String("out", "", "Output file")

	version := flag.Bool("version", false, "Print version string and exit")
	help := flag.Bool("help", false, "Print help message and exit")

	flag.Parse()

	// if asked for version, disregard everything else
	if *version {
		fmt.Printf("tryp v%s\n", cliVersion)
		return
	} else if *help {
		flag.Usage()
		return
	}

	// generate config
	config, err := tryp.ReadConfigFile(*path)
	if err != nil {
		fmt.Println(err)
		return
	}

	if *key == "" {
		key = &config.Key
	}

	// make new client
	client, err := tryp.NewClient(*key)
	if err != nil {
		fmt.Println(err)
		return
	}

	// get distance matrix
	res, err := client.Get(config.Request)
	if err != nil {
		fmt.Println(err)
		return
	}

	// translate to json
	data, err := json.Marshal(res)
	if err != nil {
		fmt.Println(err)
		return
	}

	if *out != "" {
		// write to file
		err = ioutil.WriteFile(*out, data, os.ModePerm)
		if err != nil {
			fmt.Printf("error: %s", err.Error())
			return
		}

		fmt.Printf("written to %s\n", *out)
		return
	}

	fmt.Println(string(data))
}
