package ui

import (
	"errors"
	"github.com/manifoldco/promptui"
	"github.com/nicklpeterson/confluence-faker/settings"
	"log"
	"regexp"
	"strings"
)

func PromptUserForConfluenceInstance(prompt string) * settings.Instance {
	log.Printf("%v\n", prompt)
	url := executePrompt(promptui.Prompt{
		Label: "Confluence URL (*.atlassian.net)",
		Validate: validateUrl,
	})
	email := executePrompt(promptui.Prompt{
		Label: "Email",
		Validate: validateEmail,
	})
	apiKey := executePrompt(promptui.Prompt{
		Label: "Api Key",
		Validate: nil,
		HideEntered: true,
	})
	return &settings.Instance{
		URL: url,
		Email: email,
		ApiKey: apiKey,
	}
}

func SelectConfluenceInstance(userSettings * settings.Settings) (* settings.Instance, error) {
	items := make([]string, len(userSettings.Instances))

	for index, instance := range userSettings.Instances {
		items[index] = instance.URL
	}

	index, _, err := SelectFromList(items,"Select a Confluence instance for your fake data")

	if err != nil {
		return &settings.Instance{}, errors.New("settings do not contain any confluence instances")
	}

	return &userSettings.Instances[index], nil
}

func SelectFromList(items []string, prompt string) (int, string, error) {
	if len(items) == 0 {
		return 0, "", errors.New("items must be a non empty list of string")
	}

	index, result := executeSelect(promptui.Select{
		Label: prompt,
		Items: items,
	})

	return index, result, nil

}

func executePrompt(prompt promptui.Prompt) string {
	result, err := prompt.Run()
	if err != nil {
		log.Panicf("Prompt Failed %v\n", err)
	}
	return result
}

func executeSelect(prompt promptui.Select) (int, string) {
	index, result, err := prompt.Run()
	if err != nil {
		log.Panicf("Prompt Failed %v\n", err)
	}
	return index, result
}

func validateUrl(input string) error {
	//Todo: improve this.
	valid := strings.HasSuffix(input, ".atlassian.net")
	if !valid {
		return errors.New("url must end with .atlassian.net")
	}
	return nil
}

func validateEmail(input string) error {
	var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if len(input) < 3 && len(input) > 254 || !emailRegex.MatchString(input){
		return errors.New("invlaid email")
	}
	return nil
}
