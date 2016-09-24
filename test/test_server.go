package test

import (
	"bytes"
	"net/http"
	"strconv"
)

// Fires up a test http-server on localhost and the given port.
// Via the given channel it will return any request he receives.
func StartTestHttpServer(port int, responseString string, statusCode int, requestChannel chan *http.Request) {
	testServer := TestServer{
		responseString: responseString,
		requestChannel: requestChannel,
		statusCode:     statusCode,
	}
	http.ListenAndServe(":"+strconv.Itoa(port), testServer)
}

type TestServer struct {
	requestChannel chan *http.Request
	responseString string
	statusCode     int
}

func (ts TestServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ts.requestChannel <- r
	w.WriteHeader(ts.statusCode)
	w.Write(bytes.NewBufferString(ts.responseString).Bytes())

}
