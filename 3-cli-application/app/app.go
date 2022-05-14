package app

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"net"
)

func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "app via cli"
	app.Usage = "search IP and DNS on Internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "devbook.com.br", // valor default caso nao insira nenhum valor
		},
	}

	app.Commands = []cli.Command{ // slice para receber varios comandos
		{
			Name:   "ip",
			Usage:  "Search IP and DNS on Internet",
			Flags:  flags,
			Action: searchIps,
		},
		{
			Name:   "server",
			Usage:  "Search Servers on Internet",
			Flags:  flags,
			Action: searchServers,
		},
	}

	return app
}

func searchIps(c *cli.Context) {
	host := c.String("host") // pega o valor da tag host mais acima e salva na variavel

	// func do pacote NET
	ips, erro := net.LookupIP(host) // retorna slice de IP e um erro
	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func searchServers(c *cli.Context) {
	host := c.String("host")

	// retornar servers
	servers, erro := net.LookupNS(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, server := range servers {
		fmt.Println(server)
		fmt.Println(server.Host)
	}
}
