package main

import (
	"github.com/julianduniec/jsonproxy/configuration"
	"github.com/julianduniec/jsonproxy/server"
)

func main() {
	c := configuration.Load("./configuration.yaml")
	server.Start(c.Server.Port, c.JsonP.CallbackQueryStringParameterName, c.Services)
}
