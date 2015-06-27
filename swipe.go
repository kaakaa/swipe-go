package swipe

import (
	"./conf"
	"./gist"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strings"
	"github.com/mgutz/ansi"
	"github.com/howeyc/gopass"
)

type Swipe struct {
	HomeUrl string
	LoginUrl string
	EditUrl string
	CreateUrl string
	UploadUrl string
}

type doc struct {
	Id	string
}

var (
	swipeInfo = Swipe {
		HomeUrl : "https://www.swipe.to/home",
		LoginUrl: "https://www.swipe.to/login",
		EditUrl: "https://www.swipe.to/edit",
		CreateUrl: "https://www.swipe.to/edit/create",
		UploadUrl: "https://www.swipe.to/edit/upload",
	}
)

func SwipeUpload() {
	conf, _ := conf.Parse("conf.json")

	g := new(gist.Gist)
	f, err := g.Download(conf)
	if err != nil {
		panic(err)
	}

	PostSwipeSlide(f, conf)
}

func scan(defaultValue string) string{
	var tmp string
	fmt.Scanln(&tmp)
	if strings.TrimSpace(tmp) == "" {
		return defaultValue
	}
	return tmp
}

func PostSwipeSlide(file *os.File, conf conf.Config) {
	msg := ansi.Color("Input Swipe.to Account Info", "blue+b")
	fmt.Println(msg)

	fmt.Printf("  Swipe Email(default: %s)? ", conf.Swipe.Email)
	email := scan(conf.Swipe.Email)

	fmt.Printf("  Swipe Password? ")
	tempPass := string(gopass.GetPasswd())
	
	pass := conf.Swipe.Password
	if strings.TrimSpace(tempPass) != "" {
		pass = tempPass
	}
	
	fmt.Println()

	// Login to Swipe
	client, _ := Login(email, pass)

	// Create Doc
	id, err := CreateDoc(client)
	if err != nil {
		return
	}

	// Upload Markdown
	b, contenttype, _ := CreateMultipartBody(file, id)
	if err = PostSlideFile(client, b, contenttype, id); err != nil {
		msg := ansi.Color("Error: Slide Upload Error\n", "red+b")
		fmt.Println(msg)
		panic(err)
	}

	// result
	fmt.Printf("Complete Uploading ===> %s/%s\n", swipeInfo.EditUrl, id)
}

func Login(email string, pass string) (client *http.Client, e error) {
	cookieJar, _ := cookiejar.New(nil)

	client = &http.Client {
		Jar: cookieJar,
	}

	// Login to swipe.to
	var str = []byte("email=" + url.QueryEscape(email) + "&" + "password=" + pass)
	req, _ := http.NewRequest("POST", swipeInfo.LoginUrl, bytes.NewBuffer(str))
	req.Header.Set("Referer", swipeInfo.HomeUrl)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client.Do(req)
	return client, nil
}

func CreateDoc(client *http.Client) (id string, err error){
	req, _ := http.NewRequest("POST", swipeInfo.CreateUrl, nil)
	req.Header.Set("Referer", swipeInfo.HomeUrl)
	res, err2 := client.Do(req)
	if err2 != nil {
		panic(err2)
	}
	defer res.Body.Close()

	text, _ := ioutil.ReadAll(res.Body)

	// get doc id
	d := &doc{}
	json.Unmarshal(text, &d)
	
	if strings.TrimSpace(d.Id) == "" {
		msg := ansi.Color("error: Creating Swipe document is failed\nerror: Login is failed maybe.\n", "red+b")
		fmt.Printf(msg)
		return "", fmt.Errorf("Creating Swipe document is failed.")
	}
	
	return d.Id, nil
}

func PostSlideFile(client *http.Client, b *bytes.Buffer, contenttype string, id string) (e error){
	req, _ := http.NewRequest("POST", swipeInfo.UploadUrl, b)
	req.Header.Set("Content-Type", contenttype)
	req.Header.Set("Referer", swipeInfo.EditUrl + "/" + id)

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
