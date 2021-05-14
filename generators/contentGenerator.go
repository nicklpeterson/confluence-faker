package generators

type Content struct {
	title       string `faker:"sentence" json:"title"`
	contentType string `faker:"customType" json:"type"`
	space       string `faker:"customSpace" json:""`
	status      string `faker:"" json:""`
	ancestors   string `faker:"" json:""`
	body        string `faker:"paragraph" json:"body"`
}

type Generator interface {
	generateContent() Content
}
