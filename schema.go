package db

type Schema struct {
	Types    map[string]Type
	RootType string
}
