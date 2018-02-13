package main

import (
	"bufio"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"

	"github.com/gorilla/mux"
)

func main() {
	serveWeb()
}

var themeName = getThemeName()
var staticPages = populateStaticPages()

func serveWeb() {

	gorillaRoute := mux.NewRouter()

	// for email sending, added by ubs121
	//https://github.com/ReK2Fernandez/stealthycybersecurity/blob/master/main.go
	//gorillaRoute.HandleFunc("/sendEmail", sendEmail)

	gorillaRoute.HandleFunc("/", serveContent)
	gorillaRoute.HandleFunc("/{page_alias}", serveContent)

	http.HandleFunc("/img/", serveResource)
	http.HandleFunc("/css/", serveResource)
	http.HandleFunc("/js/", serveResource)
	http.HandleFunc("/fonts/", serveResource)

	http.Handle("/", gorillaRoute)
	http.ListenAndServe(":8080", nil)
}

func serveContent(w http.ResponseWriter, r *http.Request) {
	urlParams := mux.Vars(r)
	page_alias := urlParams["page_alias"]
	if page_alias == "" {
		page_alias = "home"
	}

	staticPage := staticPages.Lookup(page_alias + ".html")
	if staticPage == nil {
		staticPage = staticPages.Lookup("404.html")
		w.WriteHeader(404)
	}
	staticPage.Execute(w, nil)
}

func getThemeName() string {
	return "bs4"
}

func populateStaticPages() *template.Template {
	result := template.New("templates")
	templatePaths := new([]string)

	basePath := "pages"
	templateFolder, _ := os.Open(basePath)
	defer templateFolder.Close()
	templatePathsRaw, _ := templateFolder.Readdir(-1)

	for _, pathInfo := range templatePathsRaw {
		log.Println(pathInfo.Name())
		*templatePaths = append(*templatePaths, basePath+"/"+pathInfo.Name())
	}

	basePath = "themes/" + themeName
	templateFolder, _ = os.Open(basePath)
	defer templateFolder.Close()
	templatePathsRaw, _ = templateFolder.Readdir(-1)

	for _, pathInfo := range templatePathsRaw {
		log.Println(pathInfo.Name())
		*templatePaths = append(*templatePaths, basePath+"/"+pathInfo.Name())
	}

	result.ParseFiles(*templatePaths...)
	return result
}

func serveResource(w http.ResponseWriter, req *http.Request) {
	path := "public/" + themeName + req.URL.Path
	var contentType string

	if strings.HasSuffix(path, ".css") {
		contentType = "text/css; charset=utf-8"
	} else if strings.HasSuffix(path, ".png") {
		contentType = "image/png; charset=utf-8"
	} else if strings.HasSuffix(path, ".jpg") {
		contentType = "image/jpg; charset=utf-8"
	} else if strings.HasSuffix(path, ".js") {
		contentType = "application/javascript; charset=utf-8"
	} else if strings.HasSuffix(path, ".ttf") {
		contentType = "application/x-font-ttf; charset=utf-8"
	} else if strings.HasSuffix(path, ".svg") {
		contentType = "image/svg+xml; charset=utf-8"
	} else if strings.HasSuffix(path, ".woff") {
		contentType = "application/font-woff; charset=utf-8"
	} else if strings.HasSuffix(path, ".woff2") {
		contentType = "application/font-woff2; charset=utf-8"
	} else if strings.HasSuffix(path, ".eot") {
		contentType = "application/vnd.ms-fontobject; charset=utf-8"
	} else {
		contentType = "text/plain; charset=utf-8"
	}
	log.Println(path)

	f, err := os.Open(path)
	if err == nil {
		defer f.Close()
		w.Header().Add("Content-Type", contentType)
		br := bufio.NewReader(f)
		br.WriteTo(w)
	} else {
		w.WriteHeader(404)
	}
}
