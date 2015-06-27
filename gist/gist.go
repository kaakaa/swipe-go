package gist

import (
	"../conf"
	"fmt"
	"os"
	"strings"
	"net/http"
	"io/ioutil"
	"github.com/mgutz/ansi"
)

type Gist struct {
	user, id string
}

const (
	UrlTemplate = "https://gist.githubusercontent.com/%s/%s/raw/%s"
)

func scan(defaultValue string) string{
	var tmp string
	fmt.Scanln(&tmp)
	if strings.TrimSpace(tmp) == "" {
		return defaultValue
	}
	return tmp
}

func (g *Gist) GetGistCode(conf conf.Config) ([]byte, error) {
	// Make Gist document URL
	fmt.Printf("Gist User ID(default: %s)? ", conf.Gist.User)
	g.user = scan(conf.Gist.User)

	fmt.Printf("Gist Document ID(default: %s)? ", conf.Gist.DocId)
	g.id = scan(conf.Gist.DocId)
	
	url := fmt.Sprintf(UrlTemplate, g.user, g.id, conf.Gist.FileName)

	println("Downloading Gist File => " + url)

	// Download Gist document to temp file
	res, _ := http.Get(url)
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		msg := ansi.Color("Error: Gist File NotFound\n  => " + url, "red+b")
		return nil, fmt.Errorf(msg)
	}

	return ioutil.ReadAll(res.Body)
}

func (g *Gist) Download(conf conf.Config) (f *os.File, err error) {
	contents, err := g.GetGistCode(conf)
	if err != nil {
		return nil, err
	}

	// coloring
	contents = Color(contents)

	// Write Gists Markdown to temp file
	f, _ = ioutil.TempFile(os.TempDir(), conf.Gist.FileName)
	defer os.Remove(f.Name())
	if err = ioutil.WriteFile(f.Name(), contents, 0755); err != nil {
		msg := ansi.Color("Error: Gist File cannot Download\n", "red+b")
		return nil, fmt.Errorf(msg)
	}

	size, _ := f.Stat()
	fmt.Printf("Complete downloading (%d Bytes)\n\n", size.Size())
	return f, nil
}

