package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func Generate() *cli.App {

	app := cli.NewApp()
	app.Name = "Command line application"
	app.Usage = "Return IPs and Servers from internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "google.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Return public IP from host",
			Flags:  flags,
			Action: getIp,
		}, {
			Name:   "server",
			Usage:  "Return servers from internet from host",
			Flags:  flags,
			Action: getServer,
		},
	}

	return app
}

func getIp(c *cli.Context) {
	host := c.String("host")

	ips, err := net.LookupIP(host)

	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func getServer(c *cli.Context) {
	host := c.String("host")

	servers, err := net.LookupNS(host)

	if err != nil {
		log.Fatal(err)
	}

	for _, server := range servers {
		fmt.Println(server.Host)
	}
}
