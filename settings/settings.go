package settings

import (
	"github.com/nicklpeterson/confluence-faker/confluence"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Settings struct {
	Instances	[]confluence.Instance `yaml:"confluence-instances"`
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

func AddNewConfluenceInstance(instance * confluence.Instance) error {
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



