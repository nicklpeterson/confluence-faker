package generators

import (
	"github.com/bxcodec/faker/v3"
)

type fakePage struct {
	Id          string `faker:"uuid_digit" json:"id"`
	Title       string `faker:"sentence" json:"title"`
	ContentType string `json:"type"`
	Space       string `json:"space"`
	Status      string `faker:"oneof:current, historical, draft" json:"status"`
	Ancestors   ancestors `json:"ancestors"`// Blank if the page is top level, else this is the parent content
	Body        string `faker:"paragraph" json:"body"`
}

type ancestors []struct {
	Id string `json:"id"`
}

func newFakePage(space string) (fakePage, error) {
	page := fakePage{}
	err := faker.FakeData(&page)
	if err != nil {
		return page, err
	}
	page.ContentType = "page"
	page.Space = space
	return page, nil
}

func NewFakePageArray(space string, numPages int) ([]fakePage, error) {
	pageArray := make([]fakePage, numPages)
	for i := 0; i < numPages; i++ {
		if page, err := newFakePage(space); err != nil {
			return nil, err
		} else {
			pageArray[i] = page
		}

	}
	return pageArray, nil
}
