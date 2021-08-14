package main

import "encoding/json"

type TODO struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type TodoOptions struct {
	Title   string
	Content string
}

func (t *TODO) ToJSON() string {
	data, err := json.Marshal(t)
	if err != nil {
		panic(err)
	}
	return string(data)
}
