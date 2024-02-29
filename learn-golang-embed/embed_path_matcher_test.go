package learn_golang_embed

import (
	"embed"
	"fmt"
	"testing"
)

//go:embed files/*.txt
var path embed.FS

func TestEmbedPathMatcher(t *testing.T) {
	dirEntries, _ := path.ReadDir("files")

	for _, entry := range dirEntries {
		if !entry.IsDir() {
			fmt.Println(entry.Name())
			file, _ := path.ReadFile("files/" + entry.Name())
			fmt.Println(string(file))
		}
	}
}
