.PHONY: all

all:
	GOOS=windows GOARCH=386 go build -o swipe-go.exe swipe.go
	zip swipe-go-win32.zip swipe-go.exe 
	GOOS=windows GOARCH=amd64 go build -o swipe-go.exe swipe.go
	zip swipe-go-win64.zip swipe-go.exe
	GOOS=linux GOARCH=386 go build -o swipe-go.exe swipe.go
	zip swipe-go-linux32.zip swipe-go.exe
	GOOS=windows GOARCH=amd64 go build -o swipe-go.exe swipe.go
	zip swipe-go-linux64.zip swipe-go.exe
	GOOS=darwin GOARCH=386 go build -o swipe-go.exe swipe.go
	zip swipe-go-mac32.zip swipe-go.exe
	GOOS=darwin GOARCH=amd64 go build -o swipe-go.exe swipe.go
	zip swipe-go-mac64.zip swipe-go.exe

clean:
	rm *.exe
	rm *.zip
