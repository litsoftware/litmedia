package core

import "time"

type ObjectOptions struct {
}

type ObjectInfo struct {
	Name        string
	Type        string
	Size        int64
	Hash        string
	ContentType string
	Metadata    map[string]string
}

type optionType string
type optionValue struct {
	Value interface{}
	Type  optionType
}
type Option func(map[string]optionValue) error
type CopyObjectResult struct {
	LastModified time.Time `json:"last_modified"`
	ETag         string    `json:"e_tag"`
}
