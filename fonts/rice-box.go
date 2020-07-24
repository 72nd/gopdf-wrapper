package utils

import (
	"time"

	"github.com/GeertJohan/go.rice/embedded"
)

func init() {

	// define files
	file2 := &embedded.EmbeddedFile{
		Filename:    "Lato-Heavy.ttf",
		FileModTime: time.Unix(1588446876, 0),

	}
	file3 := &embedded.EmbeddedFile{
		Filename:    "Lato-Regular.ttf",
		FileModTime: time.Unix(1588446876, 0),

	}

	// define dirs
	dir1 := &embedded.EmbeddedDir{
		Filename:   "",
		DirModTime: time.Unix(1588712779, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			file2, // "Lato-Heavy.ttf"
			file3, // "Lato-Regular.ttf"

		},
	}

	// link ChildDirs
	dir1.ChildDirs = []*embedded.EmbeddedDir{}

	// register embeddedBox
	embedded.RegisterEmbeddedBox(`utils`, &embedded.EmbeddedBox{
		Name: `utils`,
		Time: time.Unix(1588712779, 0),
		Dirs: map[string]*embedded.EmbeddedDir{
			"": dir1,
		},
		Files: map[string]*embedded.EmbeddedFile{
			"Lato-Heavy.ttf":   file2,
			"Lato-Regular.ttf": file3,
		},
	})
}