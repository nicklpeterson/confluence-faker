package generators

import (
	"github.com/bxcodec/faker/v3"
)

type fakePage struct {
	Title       string 		`faker:"uuid_digit" json:"title"`
	ContentType string 		`faker:"-" json:"type"`
	Status      string 		`faker:"-" json:"status"`
	Body        body		`faker:"-" json:"body"`
	Space		space		`faker:"-" json:"space"`
}

type body struct {
	Value          string `faker:"paragraph" json:"value"`
	Representation string `faker:"-" json:"representation"`
}

type space struct {
	Key string `faker:"-" json:"key"`
}

func NewFakePage(space string) (fakePage, error) {
	page := fakePage{}
	err := faker.FakeData(&page)
	if err != nil {
		return page, err
	}
	page.ContentType = "page"
	page.Space.Key = space
	page.Body.Representation = "storage"
	page.Status = "current"
	return page, nil
}

func NewFakePageArray(space string, numPages int) ([]fakePage, error) {
	pageArray := make([]fakePage, numPages)
	for i := 0; i < numPages; i++ {
		if page, err := NewFakePage(space); err != nil {
			return nil, err
		} else {
			pageArray[i] = page
		}

	}
	return pageArray, nil
}
