package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"os"
	"regexp"
	"time"
)

type Page struct {
	Title    string
	Body     []byte
	PageData []PageData
}

type PageData struct {
	Date         string
	Description  string
	RandomNumber int
}

func generateRandomNumber() int {
	return rand.Intn(41) - 20 // Generates a random number between -20 and 20 (inclusive).
}

func loadPage(title string) (*Page, error) {
	// filename := title + ".txt" // can be used to fetch files based on the input.
	log.Println("Set title as", title)
	filename := "view.html"
	body, err := os.ReadFile(filename)
	if err != nil {
		log.Println("Reading file failed", filename)
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		log.Println("Page not found", title)
		http.NotFound(w, r)
		return
	}

	// Create the PageData structure for the template.
	p.PageData = populatePageDataList()
	renderTemplate(w, "view", p)
}

func populatePageDataList() []PageData {
	var pageDataList []PageData

	for i := 0; i < 50; i++ {
		date := time.Now().AddDate(0, 0, -i) // Decrease date by i days.
		dateFormatted := date.Format("02/01/2006")
		description := fmt.Sprintf("Checking Stuff Day %v", date.YearDay())
		randomNumber := generateRandomNumber()

		pageData := PageData{
			Date:         dateFormatted,
			Description:  description,
			RandomNumber: randomNumber,
		}

		pageDataList = append(pageDataList, pageData)
	}

	return pageDataList
}

var templates = template.Must(template.ParseFiles("view.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		log.Println("renderTemplate failed ", tmpl)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

var validPath = regexp.MustCompile("^/(view)/([a-zA-Z0-9]+)$")

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			log.Println("FindStringSubmatch failed for path", r.URL.Path)
			resp := make(map[string]string)
			resp["message"] = "only Alphanumerics allowed after view in URL"
			jsonResp, err := json.Marshal(resp)
			if err != nil {
				log.Fatalf("Error happened in JSON marshal. Err: %s", err)
			}
			w.Write(jsonResp)
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

// Loading only files code start.
type justFilesFilesystem struct {
	fs http.FileSystem
}

func (fs justFilesFilesystem) Open(name string) (http.File, error) {
	f, err := fs.fs.Open(name)
	if err != nil {
		return nil, err
	}
	return neuteredReaddirFile{f}, nil
}

type neuteredReaddirFile struct {
	http.File
}

func (f neuteredReaddirFile) Readdir(count int) ([]os.FileInfo, error) {
	return nil, nil
}

// Loading only files code end.

func main() {
	log.Println("Application Startup")
	fs := justFilesFilesystem{http.Dir("resources/")}
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(fs))) // for loading static files
	http.HandleFunc("/view/", makeHandler(viewHandler))
	log.Fatal(http.ListenAndServe(":9000", nil))
}
