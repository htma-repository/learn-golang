package learn_golang_embed

import (
	_ "embed"
	"os"
	"testing"
)

//go:embed img.jpg
var image []byte

func TestEmbedByte(t *testing.T) {
	err := os.WriteFile("img_new.jpg", image, os.ModePerm)

	if err != nil {
		panic(err)
	}
}
