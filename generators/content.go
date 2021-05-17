package generators

import (
	"encoding/json"
	"github.com/bxcodec/faker/v3"
	"github.com/nicklpeterson/confluence-faker/confluence"
	"github.com/nicklpeterson/confluence-faker/logging"
	"sync"
)

type content struct {
	Title       string 		`faker:"uuid_digit" json:"title"`
	ContentType string 		`faker:"-" json:"type"`
	Status      string 		`faker:"-" json:"status"`
	Body        Body		`json:"body"`
	Space		Space		`faker:"-" json:"space"`
}

type Body struct {
	Storage struct {
		Value          string `faker:"paragraph" json:"value"`
		Representation string `faker:"-" json:"representation"`
	} `json:"storage"`
}

type Space struct {
	Key string `faker:"-" json:"key"`
}

func NewFakePage(space string) (content, error) {
	return newFakeContent(space, "page")
}

func NewFakeBlog(space string) (content, error) {
	return newFakeContent(space, "blogpost")
}

func newFakeContent(space string, contentType string) (content, error) {
	content := content{}
	err := faker.FakeData(&content)
	if err != nil {
		return content, err
	}
	content.Space.Key = space
	content.Body.Storage.Representation = "storage"
	content.Status = "current"
	content.ContentType = contentType
	return content, nil
}

func ContentWorker(id int, wg *sync.WaitGroup,
	logger *logging.Logger,
	host *confluence.Host,
	space string,
	generator func(s string) (content, error)) {
	defer wg.Done()
	blog, err := generator(space)
	if err != nil {
		logger.Info("Worker %d failed: to generate data", id)
	} else {
		body, err := json.Marshal(blog)
		if err == nil {
			status, _, err := host.Post("/wiki/rest/api/content", body)
			logger.Debug("Worker %d: http response: %v\n", id, status)
			if err != nil {
				logger.Info("Worker %d: Unable to create content, skipping", id)
			}
		} else {
			logger.Info("Worker %d: Unable to create content, skipping", id)
		}
	}
}