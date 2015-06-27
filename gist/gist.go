package gist

import (
	"../conf"
	"fmt"
	"os"
	"net/http"
	"io/ioutil"
	"github.com/mgutz/ansi"
)

type Gist struct {
	user, id string
}

const (
	UrlTemplate = "https://gist.githubusercontent.com/%s/%s/raw/%s"
	SlideFileName = "slide.md"
)

func (g *Gist) GetGistCode() ([]byte, error) {
	fmt.Printf("Gist User ID: ")
	fmt.Scanln(&g.user)

	fmt.Printf("Gist ID: ")
	fmt.Scanln(&g.id)
	url := fmt.Sprintf(UrlTemplate, g.user, g.id, SlideFileName)

	println("Downloading Gist File => " + url)

	res, _ := http.Get(url)
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		msg := ansi.Color("Error: Gist File NotFound\n  => " + url, "red+b")
		return nil, fmt.Errorf(msg)
	}

	return ioutil.ReadAll(res.Body)
}

func (g *Gist) Download(conf conf.Config) (f *os.File, err error) {
	contents, err := g.GetGistCode()
	if err != nil {
		return nil, err
	}

	// coloring
	contents = Color(contents)

	// Write Gists Markdown to temp file
	f, _ = ioutil.TempFile(os.TempDir(), SlideFileName)
	defer os.Remove(f.Name())
	if err = ioutil.WriteFile(f.Name(), contents, 0755); err != nil {
		msg := ansi.Color("Error: Gist File cannot Download\n", "red+b")
		return nil, fmt.Errorf(msg)
	}

	size, _ := f.Stat()
	fmt.Printf("Complete downloading (%d Bytes)\n\n", size.Size())
	return f, nil
}

