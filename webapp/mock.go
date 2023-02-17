package main

import (
	"sort"
	"time"
)

type photoMV map[string]photoEntity

func getMockMVAsEntityCollection() photoEntityCollection {
	collection := photoMVToEntityCollection(getMockMV())
	sort.Slice(collection, func(i, j int) bool {
		return collection[i].DateTaken.After(collection[j].DateTaken)
	})
	return collection
}

func photoMVToEntityCollection(mv photoMV) photoEntityCollection {
	collection := make(photoEntityCollection, 0)
	for id, content := range mv {
		collection = append(
			collection,
			photoEntity{
				Id:          id,
				Version:     content.Version,
				Title:       content.Title,
				Content:     content.Content,
				Filename:    content.Filename,
				DateUploaded: content.DateUploaded,
				DateUpdated: content.DateUpdated,
				DateTaken:   content.DateTaken,
			})
	}
	return collection
}

func timeParse(str string) time.Time {
	const dateFormat = "2006-01-02 15:04:05"

	t, err := time.Parse(dateFormat, str)
	if err != nil {
		panic(err)
	}
	return t
}

func getMockMV() photoMV {
	mv := photoMV{
		"6o3vy": {
			Version:     "2",
			Title:       "my title",
			Content:     "new content",
			Filename:    "DSCF8948-2.jpg",
			DateUploaded: timeParse("2023-01-26 20:40:25"),
			DateTaken:   timeParse("2021-12-10 10:15:00"),
		},
		"8fskr": {
			Version:     "2",
			Title:       "my title",
			Content:     "new content",
			Filename:    "DSCF.jpg",
			DateUploaded: timeParse("2023-01-26 20:40:25"),
			DateTaken:   timeParse("2021-12-10 10:15:00"),
		},
		"dwsy6": {
			Version:     "1",
			Title:       "yet another title",
			Content:     "bla bla",
			Filename:    "DSCF8972-2.jpg",
			DateUploaded: timeParse("2023-01-26 20:40:22"),
			DateTaken:   timeParse("2018-03-19 22:42:00"),
		},
		"v2seq": {
			Version:     "1",
			Title:       "another title",
			Content:     "new content",
			Filename:    "DSCF8978-2.jpg",
			DateUploaded: timeParse("2023-01-22 10:40:25"),
			DateTaken:   timeParse("2020-02-26 14:39:00"),
		},
	}

	return mv
}

