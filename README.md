# Website Content Checker in Go Language


The application takes a website URL as an input and provides general information about the contents of the requested page such as:
- HTML Version
- Page Title
- Headings count by level
- Amount of internal and external links
- Amount of inaccessible links
- If a page contains a login form

The application has programmed with [Go language](https://golang.org/) and it is useful for Search Engine Optimization (SEO). 

## Installation

1. Please [download and install Go](https://golang.org/doc/install)
2. Verify that you've installed Go. 
	- Open Command Prompt (cmd) and in the Command Prompt window, type the following command:
	
			$ go version
	
3. Download all files and folders. The project contains templates, css and images. So make sure you download them all.

4. Change directory to where you download/ extracted the files.

5. Open Command Promt (cmd) and change directory to the location of the files.

6. Run the the index.go application by typing the following command:
	
		$ go run index.go
	
7. You are all setup.

Please note the project requires **goquery**. You may install it by using

    $ go get github.com/PuerkitoBio/goquery


**Please note** the program might take some time to fully load the information depends on the requested URL. The program fetches all links and check the accessability of the links. So it might take some time depends on the URL.


## Imported packages

```
import (
	"fmt"
	"log"
	"net/http"
	"html/template"
	"io/ioutil"
	"strings"
	"net/url"
	"time"
	"golang.org/x/net/html"
    	"github.com/PuerkitoBio/goquery"
	)
```


## Support

There are a number of ways you can support the project:

* Use it, star it, build something with it, spread the word!
  - If you do build something open-source or otherwise publicly-visible!
* Raise issues to improve the project.
  - Please search existing issues before opening a new one - it may have already been adressed.
* Pull requests: please discuss new code in an issue first, unless the fix is really trivial.
  - Make sure new code is tested.
