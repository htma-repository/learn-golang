package learn_golang_embed

import (
	_ "embed"
	"fmt"
	"testing"
)

//go:embed embed.txt
var embedString string

func TestEmbedString(t *testing.T) {
	fmt.Println(embedString)
}
