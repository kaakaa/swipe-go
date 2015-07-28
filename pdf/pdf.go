package swipepdf

import(
	"io/ioutil"
	"os"
	"strings"
	"github.com/kaakaa/gopdf"
)

const(
	vurl = "https://www.swipe.to/vertical/"
)

// Output swipe presenation in pdf format
func Output(url string) {
	// make url of slide as vertical form
	e := strings.Split(url, "/")
	slideurl := vurl + e[len(e)-1]

	// make style sheet file
	MakeCss(slideurl)

	// make pdf file
	result := getslide(slideurl)
	write("./test.pdf", result)
}

// get slide contents
func getslide(url string) (r []byte){
	args := getargs()
	result, _ := gopdf.Url2pdf(url, args...)

	return result
}

// get command line arguments for wkhtmltopdf
func getargs() (args []string) {
	args = []string{}

	args = append(args, "--orientation", "Landscape")
	args = append(args, "--user-style-sheet", "style.css")
	args = append(args, "--page-width", "9in")
	args = append(args, "--page-height", "12in")
	args = append(args, "--zoom", "1.7")

	proxy := os.Getenv("https_proxy")
	if proxy != "" {
		args = append(args, "--proxy", proxy)
	}

	return args
}

// write content to file named name
func write(name string, content []byte) {
	err := ioutil.WriteFile(name, content, 0644)
	if err != nil {
		panic(err)
	}
}
