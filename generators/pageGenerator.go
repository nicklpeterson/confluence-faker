package generators

import (
	"fmt"
	"github.com/bxcodec/faker/v3"
)

type Page struct {
	id        string `faker:"uuid_digit" json:"id"`
	title     string `faker:"sentence" json:"title"`
	type      string `faker:""`
	space     string ``
	status    string ``
	ancestors string ``
	body

}

func generateFakePage(space string) Page {
	page := Page{}
	err := faker.FakeData(&page)

	if err != nil {
		// Improve Error logging
		fmt.Println(err)
	}

	return page
}
