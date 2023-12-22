package astro

import (
	"embed"
	"io/fs"
	"log"
	"path/filepath"
	"strings"
	"text/template"
)

type Astro struct {
	FileSystem *embed.FS
	Files      []*File
}

type File struct {
	Path      string
	Name      string
	Extension string
	BaseName  string
	Pattern   string
	Content   []byte
}

type ParseHandlerFunc = func(file *File, tmpl *template.Template) error

func New(fs *embed.FS) *Astro {
	return &Astro{
		FileSystem: fs,
	}
}

func (astro *Astro) LoadTemplates(directory string) error {
	return fs.WalkDir(astro.FileSystem, directory, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Printf("failed to walk directory %s: %v\n", path, err)
			return nil
		}

		if d.IsDir() {
			return nil
		}

		content, err := fs.ReadFile(astro.FileSystem, path)

		if err != nil {
			return err
		}

		baseName := filepath.Base(path)
		extension := filepath.Ext(path)
		pattern := strings.TrimPrefix(path, "dist")
		pattern = strings.TrimSuffix(pattern, "/index.html")

		if len(pattern) == 0 {
			pattern = "/"
		}

		astro.Files = append(astro.Files, &File{
			Path:      path,
			Name:      strings.TrimSuffix(baseName, extension),
			Extension: strings.TrimPrefix(extension, "."),
			BaseName:  baseName,
			Pattern:   pattern,
			Content:   content,
		})

		return nil
	})
}

func (astro *Astro) ParseFiles(handler ParseHandlerFunc) error {
	for _, file := range astro.Files {
		tmpl, err := template.New(file.Name).Parse(string(file.Content))

		if err != nil {
			return err
		}

		if err := handler(file, tmpl); err != nil {
			return err
		}
	}

	return nil
}
