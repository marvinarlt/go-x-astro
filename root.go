package root

import "embed"

//go:embed all:client/dist
var DistFileSystem embed.FS
