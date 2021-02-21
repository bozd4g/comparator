package configs

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path"
)

func New() Service {
	dirname, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	return service{
		configsPath: path.Join(dirname, "configs"),
	}
}

func (s service) GetAll() ([]Dto, error) {
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

func (s service) GetByName(name string) (*Dto, error) {
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
