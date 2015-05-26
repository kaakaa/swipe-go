swipe-go
========

Create Swipe.to slide from Gists markdown file.

* Uploading Markdown in [Gists](https://gist.github.com/ "Gists") to [Swipe.to](https://www.swipe.to/markdown/ "Swipe.to")

Example
=======

* This [gists markdown](https://gist.github.com/kaakaa/29ceacc3a8fa7b86f6bd "gists markdown") transforms to [New deck - Swipe](https://www.swipe.to/8563ch "New deck - Swipe")
  * but including many bugs...

Usage
=====

* go get this

```
$ go get github.com/kaakaa/swipe-go
```

* write main.go

```
package main

import (
	"github.com/kaakaa/swipe-go"
)

func main() {
  swipe.SwipeUpload()
}
```

* go run main.go

```
$ go run main.go
Gist User ID: kaakaa
Gist ID: 29ceacc3a8fa7b86f6bd
Downloading Gist File => https://gist.githubusercontent.com/kaakaa/29ceacc3a8fa7b86f6bd/raw/slide.md
truetruetrueComplete downloading (4780 Bytes)

Uploading Markdown to www.swipe.to
Swipe Email: hoge@example.com
Password:
Complete Uploading => https://www.swipe.to/edit/00000000000000000000000000000000
```

Rules
=====

* Markdown file in Gists must be named `slide.md`

* headers must be written in underline-ish style

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
