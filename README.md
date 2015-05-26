swipe-go
========

Markdown slide maker

* Uploading Markdown in [Gists](https://gist.github.com/ "Gists") to [Swipe.to](https://www.swipe.to/markdown/ "Swipe.to")

Example
=======

* [swipe-go](https://gist.github.com/kaakaa/29ceacc3a8fa7b86f6bd "swipe-go") transforms to [New deck - Swipe](https://www.swipe.to/8563ch "New deck - Swipe")
  * including many bugs...

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


