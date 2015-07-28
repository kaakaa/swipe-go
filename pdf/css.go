package swipepdf

import (
	"bytes"
	"strconv"
	"github.com/PuerkitoBio/goquery"
)

// make css file for coloring & pagination
func MakeCss(url string){
	content := getcss(url)

	write("style.css", content)
}

func getcss(url string) (b []byte) {
	doc, _ := goquery.NewDocument(url)

	var buf bytes.Buffer
	buf.Write([]byte("h1 { page-break-before: always }\n"))
	buf.Write([]byte("h2 { page-break-before: always }\n"))
	doc.Find("#slides").Children().Each(func(i int, s *goquery.Selection){
		color, _ := s.Attr("data-color")
		bgcolor, _ := s.Attr("data-background")

		buf.Write([]byte("#slides div:nth-of-type(" + strconv.Itoa(i+1) + ") div { color: " + color + "; background-color: " + bgcolor + "; }\n"))
	})

	return buf.Bytes()
}
