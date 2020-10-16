package main

import (
	"fmt"
	"os"
	"time"

	"github.com/NahuelDT/port-scanner.git/port"

	"github.com/urfave/cli"
)

func main() {
	fmt.Println("Hello WildLife!")

	app := cli.NewApp()
	app.Name = "Port Scanner"
	app.Usage = "Scans ports (in range given) of HostName given"

	myFlags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "google.com",
		},
		cli.IntFlag{
			Name:  "start",
			Value: 0,
		},
		cli.IntFlag{
			Name:  "end",
			Value: 100,
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "scan",
			Usage: "Scans hostame ports in range given",
			Flags: myFlags,
			Action: func(c *cli.Context) {
				host := c.String("host")
				start := c.Int("start")
				end := c.Int("end")
				rng := port.Range{Start: start, End: end}

				port.GetOpenPorts(host, rng)
				return
			},
		},
	}
	start := time.Now()
	app.Run(os.Args)
	elapsed := time.Since(start)

	fmt.Println("Scan duration:", elapsed)
}
