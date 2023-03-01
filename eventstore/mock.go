package main

import (
	id "github.com/antonhornquist/monorepo1/uniqueid"
)

func getMockStreams() map[string]*stream {
	photo1AggregateId := id.PseudoUniqueId()
	photo2AggregateId := id.PseudoUniqueId()
	photo3AggregateId := id.PseudoUniqueId()
	photo4AggregateId := id.PseudoUniqueId()

	return map[string]*stream{
		"stream_1_unique_id": {
			Id: "photos1",
			Records: []record{
				{
					Id:          id.PseudoUniqueId(),
					AggregateId: photo1AggregateId,
					Version:     "1",
					Data: photoEventData{
						EventType: "added",
						Title:     "my title",
						Content:   "my content",
						Filename:  "photo1.jpg"}},
				{
					Id:          id.PseudoUniqueId(),
					AggregateId: photo1AggregateId,
					Version:     "2",
					Data: photoEventData{
						EventType: "updated",
						Title:     "new title"}},
				{
					Id:          id.PseudoUniqueId(),
					AggregateId: photo2AggregateId,
					Version:     "1",
					Data: photoEventData{
						EventType: "added",
						Title:     "my title",
						Content:   "my content",
						Filename:  "photo2.jpg"}},
				{
					Id:          id.PseudoUniqueId(),
					AggregateId: photo1AggregateId,
					Version:     "3",
					Data: photoEventData{
						EventType: "deleted"}},
				{
					Id:          id.PseudoUniqueId(),
					AggregateId: photo2AggregateId,
					Version:     "2",
					Data: photoEventData{
						EventType: "updated",
						Content:   "new content"}},
				{
					Id:          id.PseudoUniqueId(),
					AggregateId: photo3AggregateId,
					Version:     "1",
					Data: photoEventData{
						EventType: "added",
						Title:     "another title",
						Content:   "new content"}},
				{
					Id:          id.PseudoUniqueId(),
					AggregateId: photo4AggregateId,
					Version:     "1",
					Data: photoEventData{
						EventType: "added",
						Title:     "yet another title",
						Content:   "bla bla",
						Filename:  "photo4.jpg"}}}},
		"stream_2_unique_id": {
			Id: "photos2",
			Records: []record{
				{
					Id:          id.PseudoUniqueId(),
					AggregateId: photo1AggregateId,
					Version:     "1",
					Data: photoEventData{
						EventType: "added",
						Title:     "my title",
						Content:   "my content",
						Filename:  "photo1.jpg"}},
				{
					Id:          id.PseudoUniqueId(),
					AggregateId: photo1AggregateId,
					Version:     "2",
					Data: photoEventData{
						EventType: "updated",
						Title:     "new title"}}}}}
}
