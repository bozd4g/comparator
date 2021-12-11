package configs

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"gopkg.in/yaml.v2"
)

func New() Service {
	dirname, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	return Service{
		configsPath: path.Join(dirname, "config"),
	}
}

func (s Service) GetAll() ([]Dto, error) {
	files, err := ioutil.ReadDir(s.configsPath)
	if err != nil {
		return nil, err
	}

	var models []Dto
	for _, file := range files {
		content, err := ioutil.ReadFile(path.Join(s.configsPath, file.Name()))
		if err != nil {
			return nil, err
		}

		var model Dto
		err = yaml.Unmarshal(content, &model)
		if err != nil {
			return nil, err
		}

		models = append(models, model)
	}
	return models, nil
}

func (s Service) GetByName(name string) (*Dto, error) {
	models, err := s.GetAll()
	if err != nil {
		return nil, err
	}

	for _, v := range models {
		if v.Name == name {
			return &v, nil
		}
	}

	return nil, errors.New(fmt.Sprintf("There is no config model for %s", name))
}
