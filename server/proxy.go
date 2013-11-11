package server

import (
	"fmt"
	"github.com/julianduniec/jsonproxy/configuration"
	"io/ioutil"
	"net/http"
	"strings"
)

type Proxy struct {
	//The query-param name where
	//the jsonp-callback-function is fetched
	callbackQueryStringParameterName string

	//The services to use
	services []configuration.Service
}

/*
	Main serve-method of the proxy.
	Fetches data from the origin-url and responds
	with the raw content.

	If there is a JsonP-parameter present, it wraps all content in a jsonp method.
*/
func (f *Proxy) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	originUrl := f.getOrigin(req)
	fmt.Println(originUrl)
	resp, body, err := fetchHttpContent(originUrl)

	if err != nil {
		fmt.Fprintf(w, "Error %s", err)
		return
	}

	//Copy headers from
	copyHeaders(resp, w)

	//Fetch the jsonp-parameter from the request
	callbackFunctionName := req.FormValue(f.callbackQueryStringParameterName)

	//Serve as jsonp or regular proxy
	if len(callbackFunctionName) != 0 {
		fmt.Fprintf(w, "%s(%s)", callbackFunctionName, body)
	} else {
		fmt.Fprintf(w, "%s", body)
	}
}

/*
	Fetch remote url from either path mapped to service,
	or the form-value "url"
*/
func (f *Proxy) getOrigin(req *http.Request) string {


	path := req.URL.Path
	for _, service := range f.services {
		fmt.Println(path, service.BasePath, strings.HasPrefix(path, service.BasePath))
		if strings.HasPrefix(path, service.BasePath) {
			append := strings.Replace(path, service.BasePath, "", 1)
			return service.RemotePath + append
		}
	}
	return req.FormValue("url")
}

/********** INTERNAL METHODS **********/

/*
	Helper method to fecth the body and response from an url
*/
func fetchHttpContent(url string) (*http.Response, []byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	return resp, body, err

}

/*
	Helper method to copy headers from a http.Response onto a http.ResponseWriter
*/
func copyHeaders(resp *http.Response, w http.ResponseWriter) {
	for key, values := range resp.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}
}
