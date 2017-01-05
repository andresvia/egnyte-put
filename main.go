package main

import (
	"errors"
	"github.com/urfave/cli"
	// "golang.org/x/oauth2"
	"os"
)

var from string
var to string

//var egnyte_chunk_size int64 = 104857600
var egnyte_chunk_size int64 = 100

func main() {
	app := cli.NewApp()
	app.Action = action
	app.Description = "put files into egnyte cloud"
	app.EnableBashCompletion = true
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "from",
			Usage:       "put from",
			Destination: &from,
		},
		cli.StringFlag{
			Name:        "to",
			Usage:       "put to",
			Destination: &to,
		},
	}
	app.Version = "v0.0.1"
	app.Run(os.Args)
}

func action(ctx *cli.Context) (action_err error) {
	var fd *os.File
	if fd, action_err = os.Open(from); action_err != nil {
		return
	} else {
		defer fd.Close()
	}

	var fi os.FileInfo
	if fi, action_err = fd.Stat(); action_err != nil {
		return
	}

	if fi.IsDir() {
		action_err = errors.New("from is a dir")
		return
	}

	chunks := (fi.Size() / egnyte_chunk_size) + 1
	fd.Close()

	for chunk := int64(0); chunk < chunks; chunk++ {
		upload_chunk(from, to, chunk*egnyte_chunk_size)
	}

	return
}

func upload_chunk(from, to string, start int64) (upload_chunk_err error) {

	var fd *os.File

	if fd, upload_chunk_err = os.Open(from); upload_chunk_err != nil {
		return
	} else {
		defer fd.Close()
	}

	if _, upload_chunk_err = fd.Seek(start, 0); upload_chunk_err != nil {
		return
	}

	upload_to(fd, to)
	return
}

func upload_to(fd *os.File, to string) {
}

func get_egnyte_token() {
	// curl https://<Egnyte Domain>.egnyte.com/puboauth/token?client_id=<API Key>&redirect_uri=<Callback URL>&scope=<SELECTED SCOPES>&state=<STRING>&response_type=code
}
