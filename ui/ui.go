package ui

import (
	"errors"
	"github.com/briandowns/spinner"
	"github.com/manifoldco/promptui"
	"github.com/nicklpeterson/confluence-faker/confluence"
	"github.com/nicklpeterson/confluence-faker/settings"
	"log"
	"os"
	"regexp"
	"strings"
	"time"
)


func NewSpinner(loadingText string) *spinner.Spinner {
	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	if loadingText != "" {
		s.Suffix = "  :" + loadingText
	}
	s.Color("fgHiGreen")
	return s
}

func GetConfluenceHost(url string) * confluence.Host {
	userSettings, err := settings.GetSettings()
	selectedHost := &confluence.Host{}
	selectedHost.URL = ""
	if url != "" {
		for _, instance := range userSettings.Hosts {
			if instance.URL == url {
				selectedHost = &instance
				break
			}
		}
	}

	if selectedHost.URL == "" && ( err != nil ||  len(userSettings.Hosts) == 0) {
		log.Printf("%v", err)
		// Todo: Prompt user to enter confluence information
		selectedHost = PromptUserForConfluenceInstance("Please add a confluence instance.")
		err := settings.AddNewConfluenceHost(selectedHost)
		if err != nil {
			log.Printf("Failed to save new settings: %v\n", err)
		}
	} else if selectedHost.URL == "" {
		selectedHost, err = SelectConfluenceHost(userSettings)
		if err != nil {
			log.Printf("An Error occured, unable to continue: \n%v\n", err)
			os.Exit(-1)
		}
	}
	return selectedHost
}

func PromptUserForConfluenceInstance(prompt string) * confluence.Host {
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
	return &confluence.Host{
		URL: url,
		Email: email,
		ApiKey: apiKey,
	}
}

func SelectConfluenceHost(userSettings * settings.Settings) (* confluence.Host, error) {
	items := make([]string, len(userSettings.Hosts))

	for index, host := range userSettings.Hosts {
		items[index] = host.URL
	}

	index, _, err := SelectFromList(items,"Select a confluence instance for your fake data")

	if err != nil {
		return &confluence.Host{}, errors.New("settings do not contain any confluence instances")
	}

	return &userSettings.Hosts[index], nil
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
	//TODO: accept urls without https://
	valid := strings.HasSuffix(input, ".atlassian.net")
	if !valid {
		return errors.New(`url must end with ".atlassian.net`)
	}
	valid = strings.HasPrefix(input, "https://")
	if !valid {
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