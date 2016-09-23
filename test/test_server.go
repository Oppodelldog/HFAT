package test

import (
	"bytes"
	"net/http"
	"strconv"
)

const TEST_SERVER_RESPONSE_STATUS_CODE = 200

// Fires up a test http-server on localhost and the given port.
// Via the given channel it will return any request he receives.
func StartTestHttpServer(port int, responseString string, requestChannel chan *http.Request) {
	testServer := TestServer{
		responseString: responseString,
		requestChannel: requestChannel,
	}
	http.ListenAndServe(":"+strconv.Itoa(port), testServer)
}

type TestServer struct {
	requestChannel chan *http.Request
	responseString string
}

func (ts TestServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ts.requestChannel <- r
	w.WriteHeader(TEST_SERVER_RESPONSE_STATUS_CODE)
	w.Write(bytes.NewBufferString(ts.responseString).Bytes())

}
