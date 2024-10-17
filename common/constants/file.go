package constants

import (
	"path/filepath"
	"runtime"
)

// ../ will take you one folder back so add it as per your folder location
var (
	_, b, _, _ = runtime.Caller(0)
	RootPath   = filepath.Join(filepath.Dir(b), "../..")
)
