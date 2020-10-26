// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Filter struct {
	Key      string      `json:"key"`
	Operator string      `json:"operator"`
	Value    interface{} `json:"value"`
}

type NewNode struct {
	Type       string                 `json:"type"`
	Attributes map[string]interface{} `json:"attributes"`
}

type Node struct {
	ID         string                 `json:"id"`
	Type       string                 `json:"type"`
	Attributes map[string]interface{} `json:"attributes"`
}

type QueryNodes struct {
	Type    string    `json:"type"`
	Filters []*Filter `json:"filters"`
	Limit   int       `json:"limit"`
}
