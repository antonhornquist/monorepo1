package main

import (
	"github.com/antonhornquist/monorepo1/httpservercommon"
	//"net/http/httputil"
	// "encoding/json"
	"io/ioutil"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
	"os"
	"path/filepath"
)

type photoEntityCollection []photoEntity

type photoEntity struct {
	Id           string    `json:"id"`
	Version      string    `json:"version"`
	Title        string    `json:"title,omitempty"`
	Content      string    `json:"content,omitempty"`
	Filename     string    `json:"filename,omitempty"`
	DateUploaded time.Time `json:"date_created,omitempty"`
	DateUpdated  time.Time `json:"date_updated,omitempty"`
	DateTaken    time.Time `json:"date_taken,omitempty"`
}

func main() {
	var (
		listenPort = flag.Int("port", 80, "Listen port.")
		eventstorePort = flag.Int("eventstore-port", 3000, "Event store port.") // TODO: URL
		materializedviewPort = flag.Int("materializedview-port", 4000, "Materialized view port.") // TODO: URL
	)

	flag.Parse()

	initializeTemplates()
	httpservercommon.Serve(*listenPort, newMux(*eventstorePort, *materializedviewPort))
}

func newMux(eventstorePort int, materializedviewPort int) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/photos/", photosHandler)
	mux.HandleFunc("/images/", imagesHandler)
	mux.HandleFunc("/thumbnails/", thumbnailsHandler)
	return mux
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	renderedTemplate, err := renderTemplate("index", nil)
	if err != nil {
		http.Error(w, fmt.Sprintf("An error has occurred: %v", err), http.StatusInternalServerError)
		return
	}
	w.Write([]byte(*renderedTemplate))
}

func photosHandler(w http.ResponseWriter, r *http.Request) {
	paths := splitPath(r.URL.Path)
	count := len(paths)

	// GET /photos - display a list of all photos
	if r.Method == http.MethodGet && count == 1 {
		displayAListOfAllPhotos(w)
	// GET /photos/new - return an HTML form for creating a new photo
	} else if r.Method == http.MethodPost && count == 2 && paths[1] == "new" {
		displayHTMLFormForCreatingNewPhoto(w)
	// POST /photos - create a new photo
	} else if r.Method == http.MethodPost && count == 1 {
		// TODO: upload photo file, append photo_added event to eventstore
		createNewPhoto(w, r)
	// GET /photos/:id - display a specific photo
	} else if count == 2 && r.Method == http.MethodGet {
		displayASpecificPhoto(w, paths[1])
	// GET /photos/:id/edit - return an HTML form for editing a photo
	} else if count == 3 && paths[2] == "edit" && r.Method == http.MethodGet {
		// TODO: update a specific photo, append photo_updated event to eventstore
		displayHTMLFormForEditingAPhoto(w, paths[1])
	// PATCH/PUT /photos/:id - update a specific photo
	} else if count == 2 && (r.Method == http.MethodPatch || r.Method == http.MethodPut) {
		// TODO: update a specific photo, append photo_updated event to eventstore
		updateASpecificPhoto(w, paths[1])
	// DELETE /photos/:id - delete a specific photo
	} else if count == 2 && r.Method == http.MethodDelete {
		// TODO: delete a specific photo, append photo_deleted event to eventstore
		deleteASpecificPhoto(w, paths[1])
	} else if r.Method == http.MethodOptions {
		setAllowedMethods(w)
		w.WriteHeader(http.StatusNoContent)
	} else {
		setAllowedMethods(w)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func thumbnailsHandler(w http.ResponseWriter, r *http.Request) {
	paths := splitPath(r.URL.Path)
	count := len(paths)

	// GET /thumbnails/:id
	if count == 2 && r.Method == http.MethodGet {
		id := paths[1]

		wd, err := os.Getwd()
		if err != nil {
			http.Error(w, fmt.Sprintf("An error has occurred: %v", err), http.StatusInternalServerError)
			return
		}

		sendFile(w, filepath.Join(wd, "thumbnails"), id)
	} else if r.Method == http.MethodOptions {
		setAllowedMethods(w)
		w.WriteHeader(http.StatusNoContent)
	} else {
		setAllowedMethods(w)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func imagesHandler(w http.ResponseWriter, r *http.Request) {
	paths := splitPath(r.URL.Path)
	count := len(paths)

	// GET /images/:id - display a specific photo
	if count == 2 && r.Method == http.MethodGet {
		id := paths[1]

		wd, err := os.Getwd()
		if err != nil {
			http.Error(w, fmt.Sprintf("An error has occurred: %v", err), http.StatusInternalServerError)
			return
		}

		sendFile(w, filepath.Join(wd, "images"), id)
	} else if r.Method == http.MethodOptions {
		setAllowedMethods(w)
		w.WriteHeader(http.StatusNoContent)
	} else {
		setAllowedMethods(w)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func setAllowedMethods(w http.ResponseWriter) {
	w.Header().Set("Allow", "GET, OPTIONS")
}

func sendFile(w http.ResponseWriter, root string, id string) {
	entity, ok := getMockMV()[id]
	if !ok {
		http.Error(w, "Bad request.", http.StatusBadRequest)
		return
	}

	path := filepath.Join(root, entity.Filename)
	fileBytes, err := ioutil.ReadFile(path)
	if err != nil {
		http.Error(w, "Bad request.", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/octet-stream")
	// TODO: Caching?
	w.Write(fileBytes)
}

func splitPath(path string) []string {
	paths := make([]string, 0)
	for _, p := range strings.Split(path, "/") {
		if p != "" {
			paths = append(paths, p)
		}
	}
	return paths
}

func displayAListOfAllPhotos(w http.ResponseWriter) {
	// TODO: query materialized view, sort by date taken

	body, err := renderTemplate("listOfPhotos", getMockMVAsEntityCollection())
	if err != nil {
		http.Error(w, fmt.Sprintf("An error has occurred: %v", err), http.StatusInternalServerError)
		return
	}
	w.Write([]byte(*body))
}

func displayHTMLFormForCreatingNewPhoto(w http.ResponseWriter) {
	log.Printf("POST /photos/new")
}

func displayASpecificPhoto(w http.ResponseWriter, id string) {
	log.Printf("GET /photos/:id")
	entity := getMockMV()[id]
	entity.Id = id
	body, err := renderTemplate("singlePhoto", entity)
	if err != nil {
		http.Error(w, fmt.Sprintf("An error has occurred: %v", err), http.StatusInternalServerError)
		return
	}
	w.Write([]byte(*body))
}

func displayHTMLFormForEditingAPhoto(w http.ResponseWriter, id string) {
	log.Printf("GET /photos/:id/edit")
	entity := getMockMV()[id]
	entity.Id = id
	body, err := renderTemplate("editPhoto", entity)
	if err != nil {
		http.Error(w, fmt.Sprintf("An error has occurred: %v", err), http.StatusInternalServerError)
		return
	}
	w.Write([]byte(*body))
}

func createNewPhoto(w http.ResponseWriter, r *http.Request) {
}

func updateASpecificPhoto(w http.ResponseWriter, id string) {
}

func deleteASpecificPhoto(w http.ResponseWriter, id string) {
}
