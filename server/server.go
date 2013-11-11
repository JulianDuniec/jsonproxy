package server

import (
	"github.com/julianduniec/jsonproxy/configuration"
	"net/http"
)

/********** EXPORTED METHODS **********/
func Start(portNumber, callbackQueryStringParameterName string, services []configuration.Service) {
	port := ":" + portNumber
	http.ListenAndServe(
		port,
		&Proxy{
			callbackQueryStringParameterName,
			services})
}

/********** INTERNAL METHODS **********/
