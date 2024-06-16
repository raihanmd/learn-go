package golang_embed

import (
	"embed"
	"io/fs"
	"os"
	"testing"
)

//go:embed file/version.txt
var version string

//go:embed img/img.webp
var img []byte

//go:embed file/?.txt
var multiple embed.FS

func TestString(t *testing.T) {
	t.Log(version)
}

func TestByte(t *testing.T) {
	err := os.WriteFile("img/img_new.webp", img, fs.ModePerm)
	if err != nil {
		panic(err)
	}
}

func TestMultiple(t *testing.T) {
	entry, _ := multiple.ReadDir("file")
	for _, e := range entry {
		t.Log(e.Name())
		file, _ := multiple.ReadFile("file/" + e.Name())
		t.Log(string(file))
	}
}
