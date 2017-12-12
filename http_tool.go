package main

import (
	"gopkg.in/urfave/cli.v1"
	"net/http"
	"os"
)

func main() {
	dataFlag := cli.StringFlag{
		Usage: "-d \"foo=bar&bar=baz\"",
		Name:  "d",
	}

	app := cli.NewApp()
	app.Name = "Http_tool"
	app.Usage = "Send HTTP requests like you own it."
	app.Commands = []cli.Command{
		{
			Name:  http.MethodGet,
			Usage: "Send a GET request to <url>",
			Flags: []cli.Flag{
				dataFlag,
			},
			Action: func(context *cli.Context) {
				Get(context.Args().Get(0), context.String("d"))
			},
		},
		{
			Name:  http.MethodPost,
			Usage: "Send a POST request to <url>",
			Flags: []cli.Flag{
				dataFlag,
			},
			Action: func(context *cli.Context) {
				Post(context.Args().Get(0), "application/x-www-form-urlencoded", context.String("d"))
			},
		},
	}

	app.Run(os.Args)
}

func Get(url string, data string) {
	get := NewGetRequest(url, NewUrlEncodedPayload(data))
	get.Send()

	get.Response.BodyContentToStdout()
}

func Post(url string, contentType string, data string) {
	post := NewPostRequest(url, contentType, NewUrlEncodedPayload(data))
	post.Send()

	post.Response.BodyContentToStdout()
}
