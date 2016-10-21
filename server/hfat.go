package server

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

// StartHFATServer starts HFAT server
func StartHFATServer(port int, forwardingTargets []ForwardingTarget) {
	handler := HFatServer{targets: forwardingTargets}
	http.ListenAndServe(":"+strconv.Itoa(port), handler)
}

// HFatServer represents the HFAT server
type HFatServer struct {
	targets []ForwardingTarget
}

// http.Handler interface
func (hfs HFatServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	var preferredResponse *http.Response
	for _, target := range hfs.targets {
		targetRequest := hfs.prepareTargetRequest(r, target)
		response, err := hfs.forwardRequest(targetRequest)
		if err != nil {
			fmt.Errorf("error forwarding http reqest: %v\n", err)
		}
		if target.Primary {
			preferredResponse = response
		}
	}

	if preferredResponse != nil {
		defer preferredResponse.Body.Close()

		responseText, err := ioutil.ReadAll(preferredResponse.Body)
		if err != nil {
			fmt.Errorf("error parsing http response text: %v\n", err)
		}
		w.WriteHeader(preferredResponse.StatusCode)
		w.Write(responseText)
	}
}

func (hfs HFatServer) prepareTargetRequest(r *http.Request, target ForwardingTarget) *http.Request {
	targetRequest := &http.Request{}
	*targetRequest = *r
	targetRequest.RequestURI = ""
	targetRequest.RemoteAddr = ""
	targetRequest.URL.Host = fmt.Sprintf("%v:%v", target.Server, target.Port)
	targetRequest.URL.Scheme = "http"

	return targetRequest
}
func (hfs HFatServer) forwardRequest(r *http.Request) (*http.Response, error) {
	client := http.Client{}
	return client.Do(r)
}
