package confluence

import (
	b64 "encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Instance struct {
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

func (i Instance) GetSpaces()(* []Space, error) {
	responseBytes, _ := i.Get("/wiki/rest/api/space")
	spaceArray := SpaceArray{}
	err := json.Unmarshal(*responseBytes, &spaceArray)
	return &spaceArray.Results, err
}

func (i Instance) Get(target string) (* []byte, error) {
	request, err := http.NewRequest(
		http.MethodGet,
		i.URL + target,
		nil,
		)
	if err != nil {
		return nil, err
	}

	i.addHeaders(request)
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	responseByte, err := ioutil.ReadAll(response.Body)
	return &responseByte, err
}

func (i Instance) addHeaders(r * http.Request) {
	encodedAuth := b64.StdEncoding.EncodeToString([]byte(i.Email + ":" + i.ApiKey))
	r.Header.Add("Authorization", "Basic " + encodedAuth)
	r.Header.Add("Content-Type", "application/json")
}

