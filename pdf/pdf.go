package swipepdf

import(
	"io/ioutil"
	"strings"
	"github.com/kaakaa/gopdf"
)

func Output(url string) {
	// make url of slide as vertical form
	e := strings.Split(url, "/")
	v_url := "https://www.swipe.to/vertical/" + e[len(e)-1]

	// make style sheet file
	MakeCss(v_url)

	// make pdf file
	result, _ := gopdf.Url2pdf(v_url, "-O", "Landscape", "--user-style-sheet", "style.css", "--page-width", "9in", "--page-height", "12in", "--zoom", "1.7")
	ioutil.WriteFile("./test.pdf", result, 0644)
}

