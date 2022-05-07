package dist

import _ "embed"

//go:embed index.html
var Index string

//go:embed favicon.ico

var Icon []byte
