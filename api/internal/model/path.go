package model

type BrowseResult struct {
	Paths   []string     `json:"paths"`
	Entries []*PathEntry `json:"entries"`
}

type PathEntry struct {
	Path        string `json:"path"`
	IsDirectory bool   `json:"is_directory"`
}
