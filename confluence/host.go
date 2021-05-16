package confluence

import (
	"bytes"
	b64 "encoding/base64"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

type Host struct {
	URL 	string	`yaml:"url"`
	ApiKey  string	`yaml:"api-key"`
	Email 	string	`yaml:"email"`
}

type SpaceArray struct {
	Results []Space `json:"results"`
}

type Space struct {
	Id   int    `json:"id"`
	Key  string `json:"key"`
	Name string `json:"name"`
}

func (i Host) GetSpaces()(* []Space, error) {
	_, responseBytes, _ := i.Get("/wiki/rest/api/space")
	spaceArray := SpaceArray{}
	err := json.Unmarshal(*responseBytes, &spaceArray)
	return &spaceArray.Results, err
}

func (i Host) Get(target string) (status string, data * []byte, err error) {
	request, err := i.newHttpRequest(target, http.MethodGet, nil)
	if err != nil {
		return "", nil, err
	}
	return makeRequest(request)
}

func (i Host) Post(target string, body []byte) (status string, data * []byte, err error) {
	request, err := i.newHttpRequest(target, http.MethodPost, bytes.NewBuffer(body))
	if err != nil {
		return "", nil, err
	}
	return makeRequest(request)
}

func makeRequest(request * http.Request) (status string, data * []byte, err error) {
	response, err := http.DefaultClient.Do(request)
	if err != nil && response != nil {
		return response.Status, nil, err
	} else if err != nil {
		return "", nil, err
	}
	responseByte, err := ioutil.ReadAll(response.Body)
	return response.Status, &responseByte, err
}

func (i Host) newHttpRequest(target string, method string, body io.Reader) (* http.Request, error) {
	request, err := http.NewRequest(
		method,
		i.URL + target,
		body,
		)

	if err == nil {
		i.addHeaders(request)
	}

	return request, err
}

func (i Host) addHeaders(r * http.Request) {
	encodedAuth := b64.StdEncoding.EncodeToString([]byte(i.Email + ":" + i.ApiKey))
	r.Header.Add("Authorization", "Basic " + encodedAuth)
	r.Header.Add("Content-Type", "application/json")
}

