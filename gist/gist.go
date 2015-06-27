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
	msg := ansi.Color("Gist Document Infomation", "blue+b")
	fmt.Println(msg)
	
	fmt.Printf("  Gist User ID(default: %s)? ", conf.Gist.User)
	g.user = scan(conf.Gist.User)

	fmt.Printf("  Gist Document ID(default: %s)? ", conf.Gist.DocId)
	g.id = scan(conf.Gist.DocId)
	
	url := fmt.Sprintf(UrlTemplate, g.user, g.id, conf.Gist.FileName)

	fmt.Println("")
	fmt.Printf("info: Downloading Gist File '%s'\n",url)

	// Download Gist document to temp file
	res, _ := http.Get(url)
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		msg := ansi.Color("error: cannot get Gist file '%s'\nerror: status code %d\n", "red+b")
		fmt.Printf(msg, url, res.StatusCode)
		return nil, fmt.Errorf("Gist file download error %d", res.StatusCode)
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
		msg := ansi.Color("Error: Gist File cannot Download \nerror: %s", "red+b")
		fmt.Printf(msg, err)
		return nil, err
	}

	size, _ := f.Stat()
	fmt.Printf("Complete Downloading (%d Bytes)\n\n", size.Size())
	return f, nil
}

