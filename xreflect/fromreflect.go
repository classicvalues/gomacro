/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2017 Massimiliano Ghilardi
 *
 *     This program is free software: you can redistribute it and/or modify
 *     it under the terms of the GNU Lesser General Public License as published
 *     by the Free Software Foundation, either version 3 of the License, or
 *     (at your option) any later version.
 *
 *     This program is distributed in the hope that it will be useful,
 *     but WITHOUT ANY WARRANTY; without even the implied warranty of
 *     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *     GNU Lesser General Public License for more details.
 *
 *     You should have received a copy of the GNU Lesser General Public License
 *     along with this program.  If not, see <https://www.gnu.org/licenses/lgpl>.
 *
 *
 * fromreflect.go
 *
 *  Created on May 07, 2017
 *      Author Massimiliano Ghilardi
 */

package xreflect

import (
	"go/token"
	"go/types"
	"reflect"
	"strings"
)

// TypeOf creates a Type corresponding to reflect.TypeOf() of given value.
// Note: conversions from Type to reflect.Type and back are not exact,
// because of the reasons listed in Type.ReflectType()
// Conversions from reflect.Type to Type and back are not exact for the same reasons.
func (v *Universe) TypeOf(rvalue interface{}) Type {
	return v.FromReflectType(reflect.TypeOf(rvalue))
}

// FromReflectType creates a Type corresponding to given reflect.Type
// Note: conversions from Type to reflect.Type and back are not exact,
// because of the reasons listed in Type.ReflectType()
// Conversions from reflect.Type to Type and back are not exact for the same reasons.
func (v *Universe) FromReflectType(rtype reflect.Type) Type {
	if rtype == nil {
		return nilT
	}
	if v.ThreadSafe {
		defer un(lock(v))
	}
	return v.fromReflectType(rtype)
}

func (v *Universe) fromReflectType(rtype reflect.Type) Type {
	if rtype == nil {
		return nilT
	}
	if v.BasicTypes == nil {
		v.init()
	}
	t := v.BasicTypes[rtype.Kind()]
	if unwrap(t) != nil && t.ReflectType() == rtype {
		return t
	}
	if t = v.ReflectTypes[rtype]; unwrap(t) != nil {
		// debugf("found rtype in cache: %v -> %v (%v)", rtype, t, t.ReflectType())
		// time.Sleep(100 * time.Millisecond)
		return t
	}
	name := rtype.Name()
	tryresolve := v.TryResolve
	if tryresolve != nil && len(name) != 0 {
		t = tryresolve(name, rtype.PkgPath())
		if unwrap(t) != nil {
			return t
		}
	}
	if v.RebuildDepth >= 0 {
		// decrement ONLY here and in fromReflectPtr() when calling fromReflectInterfacePtrStruct()
		v.RebuildDepth--
		defer func() {
			v.RebuildDepth++
		}()
	}
	// when converting a named type and v.Importer cannot locate it,
	// immediately register it in the cache because it may reference itself,
	// as for example type List struct { Elem int; Rest *List }
	// otherwise we may get an infinite recursion
	if len(name) != 0 {
		if !v.rebuild() {
			if t = v.namedTypeFromImport(rtype); unwrap(t) != nil {
				// debugf("found type in import: %v -> %v", t, t.ReflectType())
				return t
			}
		}
		t = v.namedOf(name, rtype.PkgPath())
		v.cache(rtype, t) // support self-refencing types
		// debugf("prepared named type %v", t)
	}

	var u Type
	switch k := rtype.Kind(); k {
	case reflect.Invalid:
		return nilT
	case reflect.Bool, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr,
		reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128, reflect.String,
		reflect.UnsafePointer:
		u = v.BasicTypes[k]
	case reflect.Array:
		u = v.fromReflectArray(rtype)
	case reflect.Chan:
		u = v.fromReflectChan(rtype)
	case reflect.Func:
		u = v.fromReflectFunc(rtype)
	case reflect.Interface:
		u = v.fromReflectInterface(rtype)
	case reflect.Map:
		u = v.fromReflectMap(rtype)
	case reflect.Ptr:
		u = v.fromReflectPtr(rtype)
	case reflect.Slice:
		u = v.fromReflectSlice(rtype)
	case reflect.Struct:
		u = v.fromReflectStruct(rtype)
	default:
		errorf(t, "unsupported reflect.Type %v", rtype)
	}
	if unwrap(t) == nil {
		t = u
		// cache before adding methods - otherwise we get an infinite recursion
		// if u is a pointer to named type with methods that reference the named type
		v.cache(rtype, t)
	} else {
		t.SetUnderlying(u)
		// t.ReflectType() is now u.ReflectType(). but we can do better... we know the exact rtype to set
		if !v.rebuild() {
			t.UnsafeForceReflectType(rtype)
		}
	}
	return v.addmethods(t, rtype)
}

func (v *Universe) addmethods(t Type, rtype reflect.Type) Type {
	n := rtype.NumMethod()
	if n == 0 {
		return t
	}
	tm := t
	if !t.Named() && t.Kind() == reflect.Ptr {
		// methods on pointer-to-type. add them to the type itself
		tm = t.elem()
	}
	xt := unwrap(tm)
	if !xt.Named() {
		errorf(t, "cannot add methods to unnamed type %v", t)
	}
	if xt.kind == reflect.Interface {
		// debugf("NOT adding methods to interface %v", tm)
		return t
	}
	if xt.methodvalues != nil {
		// prevent another infinite recursion: Type.AddMethod() may reference the type itself in its methods
		// debugf("NOT adding again %d methods to %v", n, tm)
	} else {
		// debugf("adding %d methods to %v", n, tm)
		xt.methodvalues = make([]reflect.Value, 0, n)
		nilv := reflect.Value{}
		if v.rebuild() {
			v.RebuildDepth--

		}
		for i := 0; i < n; i++ {
			rmethod := rtype.Method(i)
			signature := v.fromReflectMethod(rmethod.Type)
			n1 := tm.NumExplicitMethod()
			tm.AddMethod(rmethod.Name, signature)
			n2 := tm.NumExplicitMethod()
			if n1 == n2 {
				// method was already present
				continue
			}
			for len(xt.methodvalues) < n2 {
				xt.methodvalues = append(xt.methodvalues, nilv)
			}
			xt.methodvalues[n1] = rmethod.Func
		}
	}
	return t
}

func (v *Universe) fromReflectField(rfield *reflect.StructField) StructField {
	t := v.fromReflectType(rfield.Type)
	name := rfield.Name
	anonymous := rfield.Anonymous

	if strings.HasPrefix(name, StrGensymEmbedded) {
		// this reflect.StructField emulates embedded field using our own convention.
		// eat our own dogfood and convert it back to an embedded field.
		rtype := rfield.Type
		typename := rtype.Name()
		if len(typename) == 0 {
			typename = name[len(StrGensymEmbedded):]
		}
		// rebuild the type's name and package
		t = v.named(t, typename, rtype.PkgPath())
		name = typename
		anonymous = true
	} else if strings.HasPrefix(name, StrGensymPrivate) {
		// this reflect.StructField emulates private (unexported) field using our own convention.
		// eat our own dogfood and convert it back to a private field.
		name = name[len(StrGensymPrivate):]
	}

	return StructField{
		Name:      name,
		Pkg:       v.loadPackage(rfield.PkgPath),
		Type:      t,
		Tag:       rfield.Tag,
		Offset:    rfield.Offset,
		Index:     rfield.Index,
		Anonymous: anonymous,
	}
}

func (v *Universe) fromReflectFields(rfields []reflect.StructField) []StructField {
	fields := make([]StructField, len(rfields))
	for i := range rfields {
		fields[i] = v.fromReflectField(&rfields[i])
	}
	return fields
}

// named creates a new named Type based on t, having the given name and pkgpath
func (v *Universe) named(t Type, name string, pkgpath string) Type {
	if t.Name() != name || t.PkgPath() != pkgpath {
		t2 := v.namedOf(name, pkgpath)
		t2.SetUnderlying(v.maketype(t.gunderlying(), t.ReflectType()))
		t = t2
	}
	return t
}

// fromReflectArray converts a reflect.Type with Kind reflect.Array into a Type
func (v *Universe) fromReflectArray(rtype reflect.Type) Type {
	count := rtype.Len()
	elem := v.fromReflectType(rtype.Elem())
	if v.rebuild() {
		rtype = reflect.ArrayOf(count, elem.ReflectType())
	}
	return v.maketype(types.NewArray(elem.GoType(), int64(count)), rtype)
}

// fromReflectChan converts a reflect.Type with Kind reflect.Chan into a Type
func (v *Universe) fromReflectChan(rtype reflect.Type) Type {
	dir := rtype.ChanDir()
	elem := v.fromReflectType(rtype.Elem())
	if v.rebuild() {
		rtype = reflect.ChanOf(dir, elem.ReflectType())
	}
	gdir := dirToGdir(dir)
	return v.maketype(types.NewChan(gdir, elem.GoType()), rtype)
}

// fromReflectFunc converts a reflect.Type with Kind reflect.Func into a function Type
func (v *Universe) fromReflectFunc(rtype reflect.Type) Type {
	nin, nout := rtype.NumIn(), rtype.NumOut()
	in := make([]Type, nin)
	out := make([]Type, nout)
	for i := 0; i < nin; i++ {
		in[i] = v.fromReflectType(rtype.In(i))
	}
	for i := 0; i < nout; i++ {
		out[i] = v.fromReflectType(rtype.Out(i))
	}
	gin := toGoTuple(in)
	gout := toGoTuple(out)
	variadic := rtype.IsVariadic()

	if v.rebuild() {
		rin := toReflectTypes(in)
		rout := toReflectTypes(out)
		rtype = reflect.FuncOf(rin, rout, variadic)
	}
	return v.maketype(
		types.NewSignature(nil, gin, gout, variadic),
		rtype,
	)
}

// fromReflectMethod converts a reflect.Type with Kind reflect.Func into a method Type,
// i.e. into a function with receiver
func (v *Universe) fromReflectMethod(rtype reflect.Type) Type {
	nin, nout := rtype.NumIn(), rtype.NumOut()
	if nin == 0 {
		errorf(nilT, "fromReflectMethod: function type has zero arguments, cannot use first one as receiver: <%v>", rtype)
	}
	in := make([]Type, nin)
	out := make([]Type, nout)
	for i := 0; i < nin; i++ {
		in[i] = v.fromReflectType(rtype.In(i))
	}
	for i := 0; i < nout; i++ {
		out[i] = v.fromReflectType(rtype.Out(i))
	}
	grecv := toGoParam(in[0])
	gin := toGoTuple(in[1:])
	gout := toGoTuple(out)
	variadic := rtype.IsVariadic()

	if v.RebuildDepth >= 1 {
		rin := toReflectTypes(in)
		rout := toReflectTypes(out)
		rtype = reflect.FuncOf(rin, rout, variadic)
	}
	return v.maketype(
		types.NewSignature(grecv, gin, gout, variadic),
		rtype,
	)
}

// fromReflectMethod converts a reflect.Type with Kind reflect.Func into a method Type,
// manually adding the given type as receiver
func (v *Universe) fromReflectInterfaceMethod(rtype, rmethod reflect.Type) Type {
	return v.fromReflectMethod(addreceiver(rtype, rmethod))
}

// fromReflectInterface converts a reflect.Type with Kind reflect.Interface into a Type
func (v *Universe) fromReflectInterface(rtype reflect.Type) Type {
	if rtype == v.TypeOfInterface.ReflectType() {
		return v.TypeOfInterface
	}
	n := rtype.NumMethod()
	gmethods := make([]*types.Func, n)
	for i := 0; i < n; i++ {
		rmethod := rtype.Method(i)
		method := v.fromReflectInterfaceMethod(rtype, rmethod.Type)
		pkg := v.loadPackage(rmethod.PkgPath)
		gmethods[i] = types.NewFunc(token.NoPos, (*types.Package)(pkg), rmethod.Name, method.GoType().(*types.Signature))
	}
	// no way to extract embedded interfaces from reflect.Type
	if v.rebuild() {
		rfields := make([]reflect.StructField, 1+n)
		rfields[0] = approxInterfaceHeader()
		for i := 0; i < n; i++ {
			rmethod := rtype.Method(i)
			rmethodtype := rmethod.Type
			if v.RebuildDepth >= 1 {
				// needed? method := v.FromReflectType(rmethod.Type) above
				// should already rebuild rmethod.Type.ReflectType()
				rmethodtype = v.fromReflectInterfaceMethod(rtype, rmethod.Type).ReflectType()
			}
			rfields[i+1] = approxInterfaceMethod(rmethod.Name, rmethodtype)
		}
		// interfaces may have lots of methods, thus a lot of fields in the proxy struct.
		// Then use a pointer to the proxy struct: InterfaceOf() does that, and we must behave identically
		rtype = reflect.PtrTo(reflect.StructOf(rfields))
	}
	return v.maketype(types.NewInterface(gmethods, nil).Complete(), rtype)
}

// isReflectInterfaceStruct returns true if rtype is a reflect.Type with Kind reflect.Struct,
// that contains our own conventions to emulate an interface
func isReflectInterfaceStruct(rtype reflect.Type) bool {
	if rtype.Kind() == reflect.Struct {
		if n := rtype.NumField(); n != 0 {
			rfield := rtype.Field(0)
			return rfield.Name == StrGensymInterface && rfield.Type == reflectTypeOfInterfaceHeader
		}
	}
	return false
}

// fromReflectInterfacePtrStruct converts a reflect.Type with Kind reflect.Ptr,
// that contains our own conventions to emulate an interface, into a Type
func (v *Universe) fromReflectInterfacePtrStruct(rtype reflect.Type) Type {
	if rtype.Kind() != reflect.Ptr || rtype.Elem().Kind() != reflect.Struct {
		errorf(nilT, "internal error: fromReflectInterfacePtrStruct expects pointer-to-struct reflect.Type, found: %v", rtype)
	}
	rebuild := v.rebuild()
	rtype = rtype.Elem()
	n := rtype.NumField()
	// skip rtype.Field(0), it is just approxInterfaceSelf()
	var gmethods []*types.Func
	var gembeddeds []*types.Named
	var rebuildfields []reflect.StructField
	if rebuild {
		rebuildfields = make([]reflect.StructField, n)
		rebuildfields[0] = approxInterfaceHeader()
	}
	for i := 1; i < n; i++ {
		rfield := rtype.Field(i)
		name := rfield.Name
		if name == StrGensymEmbedded {
			ts := v.fromReflectInterfaceEmbeddeds(rtype, rfield.Type)
			for _, t := range ts {
				gembeddeds = append(gembeddeds, t.GoType().(*types.Named))
			}
			if rebuild {
				rebuildfields[i] = approxInterfaceEmbeddeds(ts)
			}
		} else {
			if strings.HasPrefix(name, StrGensymPrivate) {
				name = name[len(StrGensymPrivate):]
			}
			t := v.fromReflectFunc(rfield.Type)
			if t.Kind() != reflect.Func {
				errorf(t, "FromReflectType: reflect.Type <%v> is an emulated interface containing the method <%v>.\n\tExtracting the latter returned a non-function: %v", t)
			}
			gtype := t.GoType().Underlying()
			pkg := v.loadPackage(rfield.PkgPath)
			gmethods = append(gmethods, types.NewFunc(token.NoPos, (*types.Package)(pkg), name, gtype.(*types.Signature)))
			if rebuild {
				rebuildfields[i] = approxInterfaceMethod(name, t.ReflectType())
			}
		}
	}
	if rebuild {
		rtype = reflect.PtrTo(reflect.StructOf(rebuildfields))
	}
	return v.maketype(types.NewInterface(gmethods, gembeddeds).Complete(), rtype)
}

func (v *Universe) fromReflectInterfaceEmbeddeds(rinterf, rtype reflect.Type) []Type {
	if rtype.Kind() != reflect.Array || rtype.Len() != 0 || rtype.Elem().Kind() != reflect.Struct {
		return nil
	}
	rtype = rtype.Elem()
	n := rtype.NumField()
	ts := make([]Type, n)
	for i := 0; i < n; i++ {
		f := rtype.Field(i)
		t := v.fromReflectInterface(f.Type)
		if t.Kind() != reflect.Interface {
			errorf(t, `FromReflectType: reflect.Type <%v> is an emulated interface containing the embedded interface <%v>.
	Extracting the latter returned a non-interface: %v`, rinterf, f.Type, t)
		}
		ts[i] = t
	}
	return ts
}

// fromReflectMap converts a reflect.Type with Kind reflect.map into a Type
func (v *Universe) fromReflectMap(rtype reflect.Type) Type {
	key := v.fromReflectType(rtype.Key())
	elem := v.fromReflectType(rtype.Elem())
	if v.rebuild() {
		rtype = reflect.MapOf(key.ReflectType(), elem.ReflectType())
	}
	return v.maketype(types.NewMap(key.GoType(), elem.GoType()), rtype)
}

// fromReflectPtr converts a reflect.Type with Kind reflect.Ptr into a Type
func (v *Universe) fromReflectPtr(rtype reflect.Type) Type {
	relem := rtype.Elem()
	var gtype types.Type
	rebuild := v.rebuild()
	if isReflectInterfaceStruct(relem) {
		if rebuild {
			v.RebuildDepth--
			defer func() {
				v.RebuildDepth++
			}()
		}
		t := v.fromReflectInterfacePtrStruct(rtype)
		if rebuild {
			relem = t.ReflectType().Elem()
		}
		gtype = t.GoType()
	} else {
		elem := v.fromReflectType(relem)
		gtype = types.NewPointer(elem.GoType())
	}
	if rebuild {
		rtype = reflect.PtrTo(relem)
	}
	return v.maketype(gtype, rtype)
}

// fromReflectPtr converts a reflect.Type with Kind reflect.Slice into a Type
func (v *Universe) fromReflectSlice(rtype reflect.Type) Type {
	elem := v.fromReflectType(rtype.Elem())
	if v.rebuild() {
		rtype = reflect.SliceOf(elem.ReflectType())
	}
	return v.maketype(types.NewSlice(elem.GoType()), rtype)
}

// fromReflectStruct converts a reflect.Type with Kind reflect.Struct into a Type
func (v *Universe) fromReflectStruct(rtype reflect.Type) Type {
	n := rtype.NumField()
	fields := make([]StructField, n)
	for i := 0; i < n; i++ {
		rfield := rtype.Field(i)
		fields[i] = v.fromReflectField(&rfield)
	}
	vars := toGoFields(fields)
	tags := toTags(fields)

	// use reflect.StructOf to recreate reflect.Type only if requested, because it's not 100% accurate:
	// reflect.StructOf does not support unexported or anonymous fields,
	// and go/reflect cannot create named types, interfaces and self-referencing types
	if v.rebuild() {
		rfields := toReflectFields(fields, true)
		rtype = reflect.StructOf(rfields)
	}
	return v.maketype(types.NewStruct(vars, tags), rtype)
}
