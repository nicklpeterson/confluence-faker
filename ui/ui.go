package ui

import (
	"errors"
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/nicklpeterson/confluence-faker/confluence"
	"github.com/nicklpeterson/confluence-faker/settings"
	"log"
	"regexp"
	"strings"
)

func GetConfluenceInstance(url string) * confluence.Instance {
	userSettings, err := settings.GetSettings()
	selectedInstance := &confluence.Instance{}
	selectedInstance.URL = ""
	if url != "" {
		for _, instance := range userSettings.Instances {
			if instance.URL == url {
				selectedInstance = &instance
				break
			}
		}
	}

	if selectedInstance.URL == "" && ( err != nil ||  len(userSettings.Instances) == 0) {
		log.Printf("%v", err)
		// Todo: Prompt user to enter confluence information
		selectedInstance = PromptUserForConfluenceInstance("Please add a confluence instance.")
		err := settings.AddNewConfluenceInstance(selectedInstance)
		if err != nil {
			fmt.Printf("Failed to save new settings: %v\n", err)
		}
	} else if selectedInstance.URL == "" {
		selectedInstance, err = SelectConfluenceInstance(userSettings)
		if err != nil {
			//log.Panicf("%v\n", err)
		}
	}
	return selectedInstance
}

func PromptUserForConfluenceInstance(prompt string) * confluence.Instance {
	log.Printf("%v\n", prompt)
	url := executePrompt(promptui.Prompt{
		Label: "confluence URL (*.atlassian.net)",
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
	return &confluence.Instance{
		URL: url,
		Email: email,
		ApiKey: apiKey,
	}
}

func SelectConfluenceInstance(userSettings * settings.Settings) (* confluence.Instance, error) {
	items := make([]string, len(userSettings.Instances))

	for index, instance := range userSettings.Instances {
		items[index] = instance.URL
	}

	index, _, err := SelectFromList(items,"Select a confluence instance for your fake data")

	if err != nil {
		return &confluence.Instance{}, errors.New("settings do not contain any confluence instances")
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
		//log.Panicf("Prompt Failed %v\n", err)
	}
	return index, result
}

func validateUrl(input string) error {
	//Todo: improve this.
	valid := strings.HasSuffix(input, ".atlassian.net")
	if !valid {
		return errors.New(`url must end with ".atlassian.net`)
	}
	valid = strings.HasPrefix(input, "https://")
	if ! valid {
		return errors.New(`url must begin with "https://""`)
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
