package assets

import "embed"

//go:embed all:dist
var DistEmbed embed.FS
