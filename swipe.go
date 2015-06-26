package swipe

import (
	"bytes"
	"fmt"
	"os"
	"encoding/json"
	"mime/multipart"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"io/ioutil"
	"github.com/mgutz/ansi"
	"github.com/howeyc/gopass"
)

const (
	gisturl = "https://gist.githubusercontent.com/%s/%s/raw/%s"
	SlideFileName = "slide.md"
)

func GetGistCode() (f *os.File, err error) {
	// Get Gists Markdown
	var user string
	var id string

	fmt.Printf("Gist User ID: ")
	fmt.Scanln(&user)

	fmt.Printf("Gist ID: ")
	fmt.Scanln(&id)

	gist := fmt.Sprintf(gisturl, user, id, SlideFileName)

	println("Downloading Gist File => " + gist)

	res, _ := http.Get(gist)
	defer res.Body.Close()

	if res.StatusCode != 200 {
		msg := ansi.Color("Error: Gist File NotFound\n  => " + gist, "red+b")
		return nil, fmt.Errorf(msg)
	}

    // Write Gists Markdown to temp file
	f, _ = ioutil.TempFile(os.TempDir(), SlideFileName)
	defer os.Remove(f.Name())

	contents, _ := ioutil.ReadAll(res.Body)
	
	// coloring
	contents = Color(contents)
	
	if err = ioutil.WriteFile(f.Name(), contents, 0755); err != nil {
		msg := ansi.Color("Error: Gist File cannot Download\n  => " + gist, "red+b")
		return nil, fmt.Errorf(msg)
	}

	size, _ := f.Stat()
	fmt.Printf("Complete downloading (%d Bytes)\n\n", size.Size())
	return f, nil
}

func SwipeUpload() {
	f, err := GetGistCode()
	if err != nil {
		fmt.Println(err)
	}

	PostSwipeSlide(f)
}

func PostSwipeSlide(file *os.File) {
	println("Uploading Markdown to www.swipe.to")

	var email string
	fmt.Printf("Swipe Email: ")
	fmt.Scanln(&email)

	fmt.Printf("Password: ")
	pass := string(gopass.GetPasswd())

	// Login to Swipe
	client, err := Login(email, pass)

	// Create Doc
	id, _ := CreateDoc(client)

	fmt.Println("id: ", id)
	// Upload Markdown
	b, contenttype, _ := CreateMultipartBody(file, id)
	if err = PostSlideFile(client, b, contenttype, id); err != nil {
		msg := ansi.Color("Error: Slide Upload Error\n", "red+b")
		fmt.Println(msg)
		return
	}

	// result
	println("Complete Uploading => https://www.swipe.to/edit/" + id)
}

func Login(email string, pass string) (client *http.Client, err error) {
	cookieJar, _ := cookiejar.New(nil)

	client = &http.Client {
		Jar: cookieJar,
	}

	// Login to swipe.to
	var str = []byte("email=" + url.QueryEscape(email) + "&" + "password=" + pass)
	req, _ := http.NewRequest("POST", "https://www.swipe.to/login", bytes.NewBuffer(str))
	req.Header.Set("Referer", "https://www.swipe.to/home")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client.Do(req)
	return client, nil
}

type doc struct {
	Id	string
}

func CreateDoc(client *http.Client) (id string, err error){
	req, _ := http.NewRequest("POST", "https://www.swipe.to/edit/create", nil)
	req.Header.Set("Referer", "https://www.swipe.to/home")
	res, err2 := client.Do(req) 
	if err2 != nil {
		panic(err2)
	}
	defer res.Body.Close()

	text, _ := ioutil.ReadAll(res.Body)

	d := &doc{}
	// get doc id
	json.Unmarshal(text, &d)
	return d.Id, nil
}

func PostSlideFile(client *http.Client, b *bytes.Buffer, contenttype string, id string) (e error){
	req, _ := http.NewRequest("POST", "https://www.swipe.to/edit/upload", b)
	req.Header.Set("Content-Type", contenttype)
	req.Header.Set("Referer", "https://www.swipe.to/edit/" + id)

	res, err := client.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		e = fmt.Errorf("bad status: %s", res.Status)
		return e
	}
	return nil
}

func SetFormFile(writer *multipart.Writer, file *os.File) {
	// upload markdown file name must have the ext '.md'
	w, _ := writer.CreateFormFile("files[]", "./dummy.md")

	filecontents, _ := ioutil.ReadAll(file)
	w.Write(filecontents)
}

func SetFormFieldValue(writer *multipart.Writer, field string, value string) {
	w, _ := writer.CreateFormField(field)
	w.Write([]byte(value))
}

func CreateMultipartBody(file *os.File, id string)(b *bytes.Buffer, contenttype string, err error){
	b = new(bytes.Buffer)
	w := multipart.NewWriter(b)

	SetFormFile(w, file)

	SetFormFieldValue(w, "pid", id)
	SetFormFieldValue(w, "pos", "0")
	SetFormFieldValue(w, "action", "add")

	w.Close()
	return b, w.FormDataContentType(), nil
}

