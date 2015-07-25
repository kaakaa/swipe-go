package swipepdf

import (
	"bytes"
	"io/ioutil"
	"strconv"
	"github.com/PuerkitoBio/goquery"
)

func Make(url string){
	doc, _ := goquery.NewDocument(url)

	var buf bytes.Buffer
	buf.Write([]byte("h1 { page-break-before: always }\n"))
	buf.Write([]byte("h2 { page-break-before: always }\n"))
	doc.Find("#slides").Children().Each(func(i int, s *goquery.Selection){
		color, _ := s.Attr("data-color")
		bgcolor, _ := s.Attr("data-background")

		buf.Write([]byte("#slides div:nth-of-type(" + strconv.Itoa(i+1) + ") div { color: " + color + "; background-color: " + bgcolor + "; }\n"))
	})
	write("style.css", buf)
}

func write(name string, content bytes.Buffer){
	err := ioutil.WriteFile(name, content.Bytes(), 0644)
	if err != nil {
		panic(err)
	}
}
