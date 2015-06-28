swipe-go
========

Create Swipe.to slide from Gists markdown file.

* Uploading Markdown in [Gists](https://gist.github.com/ "Gists") to [Swipe.to](https://www.swipe.to/markdown/ "Swipe.to")

Example
=======

* This [gists markdown](https://gist.github.com/kaakaa/29ceacc3a8fa7b86f6bd "gists markdown") transforms to [Swipe document](https://www.swipe.to/0913ck "Swipe document")

Usage
=====

go get this repostiroy
----------------------

```
$ go get github.com/kaakaa/swipe-go
```

Write
-----

main.go
```
package main

import (
	"github.com/kaakaa/swipe-go"
)

func main() {
  swipe.SwipeUpload()
}
```

Run
---

```
$ go run main.go
Gist Document Infomation
  Gist User ID(default: kaakaa)?
  Gist Document ID(default: 29ceacc3a8fa7b86f6bd)?

info: Downloading Gist File 'https://gist.githubusercontent.com/kaakaa/29ceacc3a8fa7b86f6bd/raw/slide.md'
Complete Downloading (7153 Bytes)

Input Swipe.to Account Info
  Swipe Email(default: stooner.hoe@gmail.com)?
  Swipe Password?

Complete Uploading ===> https://www.swipe.to/edit/00000000000000000000000000000000

```

Configuration
-------------

* Follow is configuration file format for Gist / Swipe.to / slide format
  * place conf.json in current directory

conf.json
```
{
	"Gist" : {
		"User": "default",
		"DocId": "abcdefghijklmnopqrst",
    "FileName": "slide.md"
	},
	"Swipe" : {
		"Email": "foobar@exmpale.com",
		"Password": "pass",
		"Coloring": true
	}
}
```

Rules
=====

* Headers must be written in underline-ish style

```
OK

  H1
  ==
```

```
NG

  # H1
```

* Horizontal must be written in Asterisks style

```

OK

  * * *
```

```
NG

  ---
```
