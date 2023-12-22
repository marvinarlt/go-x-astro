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
	Templates  []*Template
}

type Template struct {
	Path      string
	Name      string
	Extension string
	BaseName  string
	Content   []byte
}

type ParseHandlerFunc = func(t *Template, tmpl *template.Template) error

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

		astro.Templates = append(astro.Templates, &Template{
			Path:      path,
			Name:      strings.TrimSuffix(baseName, extension),
			Extension: strings.TrimPrefix(extension, "."),
			BaseName:  baseName,
			Content:   content,
		})

		return nil
	})
}

func (astro *Astro) ParseTemplates(handler ParseHandlerFunc) error {
	for _, t := range astro.Templates {
		tmpl, err := template.New(t.Name).Parse(string(t.Content))

		if err != nil {
			return err
		}

		if err := handler(t, tmpl); err != nil {
			return err
		}
	}

	return nil
}
