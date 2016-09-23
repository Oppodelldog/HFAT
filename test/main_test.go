package test

import (
	"HFAT/server"
	"bytes"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

const TEST_SERVER_1_PORT = 10000
const TEST_SERVER_2_PORT = 10001
const TEST_SERVER_1_RESPONSE_BODY = "TEST SERVER 1 RESPONSE"
const TEST_SERVER_2_RESPONSE_BODY = "TEST SERVER 2 RESPONSE"
const LOCALHOST = "localhost"
const HFAT_SERVER_PORT = 8000
const TEST_PATH_AND_QUERY = "/whateverPath?TEST=98&sha=27892572805702572"

func TestTimeConsuming(t *testing.T) {

	requestSnifferChannel1 := make(chan *http.Request)
	requestSnifferChannel2 := make(chan *http.Request)
	go StartTestHttpServer(TEST_SERVER_1_PORT, TEST_SERVER_1_RESPONSE_BODY, requestSnifferChannel1)
	go StartTestHttpServer(TEST_SERVER_2_PORT, TEST_SERVER_2_RESPONSE_BODY, requestSnifferChannel2)

	go server.StartHFATServer(HFAT_SERVER_PORT, []server.ForwardingTarget{
		{Server: LOCALHOST, Port: TEST_SERVER_1_PORT},
		{Server: LOCALHOST, Port: TEST_SERVER_2_PORT, Primary:true},
	})

	testRequest, err := http.NewRequest("GET", "http://" + LOCALHOST + ":" + strconv.Itoa(HFAT_SERVER_PORT) + TEST_PATH_AND_QUERY, bytes.NewBufferString(""))
	if err != nil {
		t.Error(err)
	}
	responseChannel := make(chan *http.Response)
	go sendTestRequest(testRequest, responseChannel)

	sniffedRequest1 := <-requestSnifferChannel1
	sniffedRequest2 := <-requestSnifferChannel2
	response := <-responseChannel

	if !reflect.DeepEqual(testRequest.URL.Path, sniffedRequest1.URL.Path) {
		t.Error("url path values sniffed at server1 are not equal to test request path value")
	}

	if !reflect.DeepEqual(testRequest.URL.RawQuery, sniffedRequest1.URL.RawQuery) {
		t.Error("raw query sniffed at server1 is not equal to test request raw query")
	}

	if !reflect.DeepEqual(testRequest.URL.Path, sniffedRequest2.URL.Path) {
		t.Error("url path values sniffed at server2 are not equal to test request path value")
	}

	if !reflect.DeepEqual(testRequest.URL.RawQuery, sniffedRequest2.URL.RawQuery) {
		t.Error("raw query sniffed at server2 is not equal to test request raw query")
	}

	if response.StatusCode != TEST_SERVER_RESPONSE_STATUS_CODE {
		t.Error("Did expect status code %v, but not: %v\n", TEST_SERVER_RESPONSE_STATUS_CODE, response.StatusCode)
	}

	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Error(err)
	}
	if !strings.ContainsAny(string(responseBody), TEST_SERVER_2_RESPONSE_BODY) {
		t.Errorf("Did expect that: %v \nwould contain: %v\n", string(responseBody), TEST_SERVER_2_RESPONSE_BODY)
	}
}

func sendTestRequest(request *http.Request, responseChannel chan *http.Response) {
	client := &http.Client{}
	response, err1 := client.Do(request)
	if err1 != nil {
		panic(err1)
	}

	responseChannel <- response
}
