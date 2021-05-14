package settings

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Settings struct {
	Instances	[]Instance	`yaml:"confluence-instances"`
}

type Instance struct {
	URL 	string	`yaml:"url"`
	ApiKey  string	`yaml:"api-key"`
	Email 	string	`yaml:"email"`
}

func GetSettings() (* Settings, error) {
	var settings = Settings{}

	data, err := ioutil.ReadFile(".confluence-faker")
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(data, &settings)
	if err != nil {
		return nil, err
	}

	return &settings, nil
}

func SaveSettings(settings * Settings) error {
	data, err := yaml.Marshal(settings)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(".confluence-faker", data, 0)
	if err != nil {
		return err
	}

	return nil
}

func AddNewConfluenceInstance(instance * Instance) error {
	settings, err := GetSettings()
	if err != nil {
		 settings = &Settings{}
	}

	settings.Instances = append(settings.Instances, *instance)
	err = SaveSettings(settings)
	if err != nil {
		return err
	}
	return nil
}



