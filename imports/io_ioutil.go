// this file was generated by gomacro command: import "io/ioutil"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	pkg "io/ioutil"
	. "reflect"
)

func Package_io_ioutil() (map[string]Value, map[string]Type) {
	return map[string]Value{
			"Discard":   ValueOf(&pkg.Discard).Elem(),
			"NopCloser": ValueOf(pkg.NopCloser),
			"ReadAll":   ValueOf(pkg.ReadAll),
			"ReadDir":   ValueOf(pkg.ReadDir),
			"ReadFile":  ValueOf(pkg.ReadFile),
			"TempDir":   ValueOf(pkg.TempDir),
			"TempFile":  ValueOf(pkg.TempFile),
			"WriteFile": ValueOf(pkg.WriteFile),
		}, map[string]Type{}
}

func init() {
	binds, types := Package_io_ioutil()
	Binds["io/ioutil"] = binds
	Types["io/ioutil"] = types
}