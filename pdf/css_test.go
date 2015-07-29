package swipepdf

import(
	"testing"
)

func TestGetcss(t *testing.T) {
	url := "https://www.swipe.to/vertical/98747tzf1qqzn8x"
	content := getcss(url)

	expected := 957
	if len(content) != expected {
		t.Errorf("wrong size(expected: %d, actual: %d)", expected, len(content))
	}
}
