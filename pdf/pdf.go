package swipepdf

import(
	"io/ioutil"
	"strings"
	"github.com/kaakaa/gopdf"
)

func Output(id string) {
	// make url of slide as vertical form
	v_url := "https://www.swipe.to/vertical/" + id

	// make style sheet file
	fmt.Println("Making style sheet for making up PDF contents...")
	swipecss.Make(v_url)

	// make pdf file
	fmt.Println("Execute wkhtmltopdf for making PDF file...")
	result, _ := gopdf.Url2pdf(v_url, "-O", "Landscape", "--user-style-sheet", "style.css", "--page-width", "9in", "--page-height", "12in", "--zoom", "1.7")
	ioutil.WriteFile("./test.pdf", result, 0644)
}

