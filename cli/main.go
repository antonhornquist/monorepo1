package main

import (
	// "io/ioutil"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
)

type handler struct {
	photos   *photoMV
	position string
}

type photoMV map[string]photoEntity

type photoEntity struct {
	Version  string `json:"version"`
	Title    string `json:"title,omitempty"`
	Content  string `json:"content,omitempty"`
	Filename string `json:"filename,omitempty"`
}

type photoEventData struct {
	EventType string `json:"type"`
	Title     string `json:"title,omitempty"`
	Content   string `json:"content,omitempty"`
	Filename  string `json:"filename,omitempty"`
}

type record struct {
	Id          string         `json:"id"`
	AggregateId string         `json:"aggregate_id,omitempty"`
	Version     string         `json:"version,omitempty"`
	Data        photoEventData `json:"data"`
}

type photosStream struct {
	Id      string   `json:"id"`
	Records []record `json:"records"`
}

func main() {
	var (
		// listenPort = flag.Int("port", 4000, "Listen port.")
		eventstorePort = flag.Int("port", 3000, "Listen port.")
	)

	flag.Parse()

	requestURL := fmt.Sprintf("http://localhost:%d/stream_1_unique_id", *eventstorePort)
	res, err := http.Get(requestURL)
	if err != nil {
		log.Fatal(fmt.Sprintf("HTTP request error: %v\n", err))
	}

	log.Printf("We're ok!")

	var stream photosStream
	err = json.NewDecoder(res.Body).Decode(&stream)
	if err != nil {
		log.Fatal("Could not parse response: %s\n", err)
	}

	log.Printf("--- Stream ---")
	logAsJsonIfPossible(stream)

	log.Printf("--- View ---")
	photos, position := reconstructMV(stream.Records)
	logAsJsonIfPossible(position)
	logAsJsonIfPossible(photos)

	//httpservercommon.Serve(*listenPort, newMux())
}

func reconstructMV(records []record) (photoMV, int) {
	getUniqueAggregateIds := func(records []record) map[string]bool {
		aggregateIds := make(map[string]bool) // map is here used as a set datastructure

		for _, rec := range records {
			aggregateIds[rec.AggregateId] = true
		}

		return aggregateIds
	}

	getRecordsForAggregateId := func(records []record, aggregateId string) []record {
		result := make([]record, 0)
		for _, rec := range records {
			if rec.AggregateId == aggregateId {
				result = append(result, rec)
			}
		}
		return result
	}

	photos := make(photoMV)

	uniqueAggregateIds := getUniqueAggregateIds(records)

	for aggregateId := range uniqueAggregateIds {
		for _, rec := range getRecordsForAggregateId(records, aggregateId) {
			logRecordIgnored := func() {
				log.Printf("Record ignored: %v", rec)
			}

			logRecordInconsistent := func() {
				log.Printf("Record ignored (inconsistent): %v", rec)
			}

			photoExists := func(string) bool {
				_, exists := photos[aggregateId]
				return exists
			}

			// logAsJsonIfPossible(rec)

			photoEventData := rec.Data
			switch eventType := photoEventData.EventType; eventType {
			case "added":
				if photoExists(aggregateId) {
					logRecordInconsistent()
				} else {
					photos[aggregateId] = photoEntity{
						Version:  rec.Version,
						Title:    photoEventData.Title,
						Content:  photoEventData.Content,
						Filename: photoEventData.Filename}
				}
			case "deleted":
				if photoExists(aggregateId) {
					delete(photos, aggregateId)
				} else {
					logRecordInconsistent()
				}
			case "updated":
				if photoExists(aggregateId) {
					photo := photos[aggregateId]
					if photoEventData.Title != "" {
						photo.Title = photoEventData.Title
					}
					if photoEventData.Content != "" {
						photo.Content = photoEventData.Content
					}
					if photoEventData.Filename != "" {
						photo.Filename = photoEventData.Filename
					}
					photo.Version = rec.Version
					photos[aggregateId] = photo
				} else {
					logRecordInconsistent()
				}
			default:
				logRecordIgnored()
			}
		}
	}
	return photos, len(records)
}

func logAsJsonIfPossible(p interface{}) {
	bytes, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		log.Printf("Error converting to JSON: %v", p)
	}
	log.Printf(string(bytes))
}
