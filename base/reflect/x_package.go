// this file was generated by gomacro command: import _i "github.com/cosmos72/gomacro/base/reflect"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package reflect

import (
	r "reflect"
	"github.com/cosmos72/gomacro/imports"
)

// reflection: allow interpreted code to import "github.com/cosmos72/gomacro/base/reflect"
func init() {
	imports.Packages["github.com/cosmos72/gomacro/base/reflect"] = imports.Package{
	Binds: map[string]r.Value{
		"ConvertValue":	r.ValueOf(ConvertValue),
		"IsCategory":	r.ValueOf(IsCategory),
		"IsNillableKind":	r.ValueOf(IsNillableKind),
		"IsOptimizedKind":	r.ValueOf(IsOptimizedKind),
		"KindToCategory":	r.ValueOf(KindToCategory),
		"KindToType":	r.ValueOf(KindToType),
		"Nil":	r.ValueOf(&Nil).Elem(),
		"None":	r.ValueOf(&None).Elem(),
		"PackTypes":	r.ValueOf(PackTypes),
		"PackValues":	r.ValueOf(PackValues),
		"PackValuesAndTypes":	r.ValueOf(PackValuesAndTypes),
		"TypeOfBool":	r.ValueOf(&TypeOfBool).Elem(),
		"TypeOfComplex128":	r.ValueOf(&TypeOfComplex128).Elem(),
		"TypeOfComplex64":	r.ValueOf(&TypeOfComplex64).Elem(),
		"TypeOfFloat32":	r.ValueOf(&TypeOfFloat32).Elem(),
		"TypeOfFloat64":	r.ValueOf(&TypeOfFloat64).Elem(),
		"TypeOfInt":	r.ValueOf(&TypeOfInt).Elem(),
		"TypeOfInt16":	r.ValueOf(&TypeOfInt16).Elem(),
		"TypeOfInt32":	r.ValueOf(&TypeOfInt32).Elem(),
		"TypeOfInt64":	r.ValueOf(&TypeOfInt64).Elem(),
		"TypeOfInt8":	r.ValueOf(&TypeOfInt8).Elem(),
		"TypeOfString":	r.ValueOf(&TypeOfString).Elem(),
		"TypeOfUint":	r.ValueOf(&TypeOfUint).Elem(),
		"TypeOfUint16":	r.ValueOf(&TypeOfUint16).Elem(),
		"TypeOfUint32":	r.ValueOf(&TypeOfUint32).Elem(),
		"TypeOfUint64":	r.ValueOf(&TypeOfUint64).Elem(),
		"TypeOfUint8":	r.ValueOf(&TypeOfUint8).Elem(),
		"TypeOfUintptr":	r.ValueOf(&TypeOfUintptr).Elem(),
		"UnpackValues":	r.ValueOf(UnpackValues),
		"ValueInterface":	r.ValueOf(ValueInterface),
		"ValueType":	r.ValueOf(ValueType),
	}, 
	}
}
