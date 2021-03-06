package settings

import (
	"github.com/nicklpeterson/confluence-faker/confluence"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Settings struct {
	Hosts []confluence.Host `yaml:"confluence-hosts"`
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

func AddNewConfluenceHost(instance * confluence.Host) error {
	settings, err := GetSettings()
	if err != nil {
		 settings = &Settings{}
	}

	settings.Hosts = append(settings.Hosts, *instance)
	err = SaveSettings(settings)
	if err != nil {
		return err
	}
	return nil
}



