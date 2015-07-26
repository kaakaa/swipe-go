.PHONY: all

all:
	GOOS=windows GOARCH=386 go build -o swipe-go.exe
	zip swipe-go-win32.zip swipe-go.exe

clean:
	rm *.exe
	rm *.zip
