package swipepdf

import(
	"testing"
	"os"
)

var https_proxy = os.Getenv("https_proxy");

func TestGetargs(t *testing.T) {
	os.Setenv("https_proxy", "")
	args := getargs()

	if len(args) != 10 {
		t.Errorf("Elements of args is wrong number\n  => %s", args)
	}

	os.Setenv("https_proxy", https_proxy)
}


func TestGetargs_withproxy(t *testing.T) {
	os.Setenv("https_proxy", "http://test.proxy.com:8080")
	args := getargs()

	if len(args) != 12 {
		t.Errorf("Elements of args is wrong number\n  => %s", args)
	}

	os.Setenv("https_proxy", https_proxy)
}

func TestGetslide(t *testing.T) {
	url := "https://www.swipe.to/vertical/98747tzf1qqzn8x"
	content := getslide(url)

	expected := 1312393
	if len(content) != expected {
		t.Errorf("wrong size(expected: %d, actual: %d)", expected, len(content))
	}
}
