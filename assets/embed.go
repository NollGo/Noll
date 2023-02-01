package assets

import "embed"

//go:embed *
// Dir is assets dir embed fs
var Dir embed.FS
