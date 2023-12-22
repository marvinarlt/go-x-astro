package root

import "embed"

//go:embed all:dist
var DistFileSystem embed.FS
