package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/urfave/cli"
)

const usage = `
  _          _ _   _                         _ _
 | |___ __ _(_) |_| |_ ___ _ _ ___ _ __  ___| | |___ _ _
 |  _\ V  V / |  _|  _/ -_) '_|___| '_ \/ _ \ | / -_) '_|
  \__|\_/\_/|_|\__|\__\___|_|     | .__/\___/_|_\___|_|
                                  |_|

polls a list of urls, reporting an aggregate of success count by
application/version
`

func main() {
	app := cli.NewApp()

	app.Name = "twitter-poller"
	app.Usage = usage
	app.Version = "1.0.0"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "infile,i",
			Usage: "file containing newline delimited list of urls to poll",
			Value: "servers.txt",
		},
		cli.StringFlag{
			Name:  "outfile,o",
			Usage: "file that machine-parseable results will get written to",
			Value: "result.json",
		},
		cli.IntFlag{
			Name:  "pollers,p",
			Usage: "number of pollers to launch",
			Value: 2,
		},
	}
	app.Action = func(c *cli.Context) error {
		file, err := os.Open(c.String("infile")) // open file for reading
		if err != nil {
			return cliErr(err)
		}

		formatter := NewFormatter()    // formats output for human reading
		aggregator := NewAggregator()  // aggregates statuses
		client := &http.Client{}       // default http client
		scanner := NewURLScanner(file) // url scanner for input file
		urls := make(chan Poller)      // channel for scanner to pass urls
		statuses := make(chan Status)  // channel for pollers to pass statuses

		go scanner.Scan(urls)             // launch the scanner
		go aggregator.Aggregate(statuses) // launch the aggregator

		// launch some pollers
		var wg sync.WaitGroup
		for i := 0; i < c.Int("pollers"); i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				Poll(client, urls, statuses)
			}()
		}

		wg.Wait()       // wait until pollers have finished
		close(statuses) // close the statuses channel

		// writes json to outfile
		byt, err := json.Marshal(aggregator.Data)
		if err != nil {
			log.Printf("%v\n", err)
		}
		err = ioutil.WriteFile(c.String("outfile"), byt, 0644)
		if err != nil {
			log.Printf("%v\n", err)
		}

		// writes human readable output to stdout
		formatter.Format(os.Stdout, aggregator.Data)

		return nil
	}

	app.Run(os.Args)
}

func cliErr(err error) *cli.ExitError {
	return cli.NewExitError(fmt.Sprintf("main: %s", err), 1)
}
