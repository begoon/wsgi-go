package util

import (
	"embed"
	"fmt"
)

func StringFS(b *[]byte, prefix string, fs embed.FS) {
	if p, err := fs.ReadDir(prefix); err == nil {
		for _, f := range p {
			info, err := f.Info()
			if err != nil {
				panic(fmt.Sprintf("unable to read info of %v", f))
			}
			path := prefix + "/" + f.Name()
			if info.IsDir() {
				StringFS(b, path, fs)
			} else {
				*b = fmt.Appendf(*b, "%v (%d)\n", path, info.Size())
			}
		}
	} else {
		panic("embed files not found")
	}
}

func PrintFS(prefix string, fs embed.FS) {
	var b []byte
	StringFS(&b, "root", fs)
	fmt.Print(string(b))
}
