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

var (
	defaultConf = Config {
		Gist: Gist{
			User: "default",
			DocId: "abcdefghijklmnopqrst",
			FileName: "slide.md",
		},
		Swipe: Swipe{
			Email: "foobar@example.com",
			Password: "12345678",
			Coloring: true,
		},
	}
)

func Parse(path string) (Config, error) {
	conf := defaultConf
	
	// Read conf file	
	c, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("error: %v", err)
		return conf, err
	}

	// apply configuration
	err = json.Unmarshal(c, &conf)
	if err != nil {
		fmt.Printf("error: %v", err)
		return conf, err
	}
	return conf, nil
}
