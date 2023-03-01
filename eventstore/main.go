package main

import (
	//"net/http/httputil"
	// "io/ioutil"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/antonhornquist/monorepo1/httpservercommon"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

type handler struct {
	stream *stream
}

type photoEventData struct {
	EventType string `json:"type"`
	Title     string `json:"title,omitempty"`
	Content   string `json:"content,omitempty"`
	Filename  string `json:"filename,omitempty"`
}

type record struct {
	Id          string      `json:"id"`
	AggregateId string      `json:"aggregate_id,omitempty"`
	Version     string      `json:"version,omitempty"`
	Data        interface{} `json:"data"`
}

type stream struct {
	mutex   sync.RWMutex
	Id      string   `json:"id"`
	Records []record `json:"records"`
}

func main() {
	var (
		listenPort = flag.Int("port", 3000, "Listen port.")
	)

	flag.Parse()

	httpservercommon.Serve(*listenPort, newMux())
}

func newMux() *http.ServeMux {
	mux := http.NewServeMux()

	rand.Seed(time.Now().UnixNano())

	streams := getMockStreams()

	mux.Handle("/stream_1_unique_id", handler{stream: streams["stream_1_unique_id"]})
	mux.Handle("/stream_2_unique_id", handler{stream: streams["stream_2_unique_id"]})
	mux.HandleFunc("/", indexHandler)

	// logAsJsonIfPossible(streams["stream_2_unique_id"])

	return mux
}

func logAsJsonIfPossible(p interface{}) {
	bytes, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		log.Printf("Error converting to JSON: %v", p)
	}
	log.Printf(string(bytes))
}

/*
func dumpRequest(r *http.Request) {
	dump, err := httputil.DumpRequest(r, true)

	if err != nil {
		log.Printf("cannot dump request")
	} else {
		log.Printf("%q", dump)
	}
}
*/

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// dumpRequest(r)

	switch r.Method {
	case http.MethodGet:
		h.stream.mutex.RLock()
		defer h.stream.mutex.RUnlock()

		bytes, err := json.MarshalIndent(h.stream, "", "  ")

		if err != nil {
			http.Error(w, fmt.Sprintf("An error has occurred: %v", err), http.StatusBadRequest)
			return
		}

		w.Write(bytes)
	case http.MethodPost:
		h.stream.mutex.Lock()
		defer h.stream.mutex.Unlock()

		/*
			bytes, err := ioutil.ReadAll(r.Body)
			if err != nil {
				http.Error(w, fmt.Sprintf("An error has occurred: %v", err), http.StatusInternalServerError)
				return
			}
			log.Printf(string(bytes))
		*/

		var record record
		err := json.NewDecoder(r.Body).Decode(&record)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		h.stream.Records = append(h.stream.Records, record)

		logAsJsonIfPossible(h.stream)

		w.WriteHeader(http.StatusNoContent)
	case http.MethodOptions:
		w.Header().Set("Allow", "GET, POST, OPTIONS")
		w.WriteHeader(http.StatusNoContent)
	default:
		w.Header().Set("Allow", "GET, POST, OPTIONS")
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Woo!"))
}
