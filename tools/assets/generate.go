// +build ignore

package main

import (
	"log"

	"github.com/shurcooL/vfsgen"

	"github.com/var512/ddda-save-corrupter/internal/assets"
)

func main() {
	err := vfsgen.Generate(assets.Assets, vfsgen.Options{
		Filename:     "./internal/assets/fs_prod.go",
		PackageName:  "assets",
		BuildTags:    "!dev",
		VariableName: "Assets",
	})

	if err != nil {
		log.Fatal(err)
	}
}
