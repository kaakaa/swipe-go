package conf

import (
	"io/ioutil"
	"fmt"
	"encoding/json"
)

type Config struct {
	Gist Gist
	Swipe Swipe
}

type Gist struct {
	User string
	DocId string
	FileName string
}

type Swipe struct {
	Email string
	Password string
	Coloring bool
}

func Parse(path string) (Config, error) {
	var conf Config
	
	c, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("error: %v", err)
		return conf, err
	}

	err = json.Unmarshal(c, &conf)
	if err != nil {
		fmt.Printf("error: %v", err)
		return conf, err
	}

	return conf, nil
}
