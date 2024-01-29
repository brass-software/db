package db

import "github.com/mikerybka/brass"

type Schema struct {
	Types    map[string]brass.Type
	RootType string
}
