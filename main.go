package main

import (
	"github.com/urfave/cli"
	"os"
	"time"
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/stiltskincode/http-tools/cmd"
)

var log = logrus.New()

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func main(){
	availableMethods := []string{"GET", "HEAD", "PUT"}
	app := cli.NewApp()
	app.Name = "web api test"
	app.Version = "00.00.1"
	app.Compiled = time.Now()
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Lukasz Glowacki",
			Email: "lukasz.glowacki@acaisoft.com",
		},
	}
	app.Usage = "tools for testing web api"
	app.EnableBashCompletion = true

	app.Commands = []cli.Command{
		cli.Command{
			Name:        "ping",
			Category:    "multi-threading",
			Description: "test endpoint request",

			Flags:[]cli.Flag {
				cli.StringFlag{
					Name: "method, m",
					Value: "GET",
					Usage: "http request method",
				},
				cli.UintFlag{
					Name: "threads, t",
					Value: 1,
					Usage: "number of threads",
				},
				cli.UintFlag{
					Name: "requests, r",
					Value: 1,
					Usage: "number of requests",
				},cli.BoolFlag{
					Name: "postfix, p",
					Usage: "number of requests",
				},
			},
			Action:func(c *cli.Context) error {
				if c.NArg() == 0 {
					fmt.Println("Please provide address url.")
					return nil
				}

				method := c.String("method")


				if !contains(availableMethods, method) {
					fmt.Println("Available metods are GET HEAD.")
					return nil
				}

				threads := c.Int("threads")
				requests := c.Int("requests")
				postfix := c.Bool("postfix")
				url := c.Args().Get(0)

				cmd.HttpEndpointBenchmark(method, url, threads, requests, postfix)

				return nil
			},
			BashComplete:func(c *cli.Context) {
				// This will complete if no args are passed
				if c.NArg() > 0 {
					return
				}
				for _, t := range availableMethods {
					fmt.Println(t)
				}
			},
		},
	}

	app.Run(os.Args)
}

