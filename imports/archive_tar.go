// this file was generated by gomacro command: import _b "archive/tar"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	tar "archive/tar"
)

// reflection: allow interpreted code to import "archive/tar"
func init() {
	Packages["archive/tar"] = Package{
	Name: "tar",
	Binds: map[string]Value{
		"ErrFieldTooLong":	ValueOf(&tar.ErrFieldTooLong).Elem(),
		"ErrHeader":	ValueOf(&tar.ErrHeader).Elem(),
		"ErrWriteAfterClose":	ValueOf(&tar.ErrWriteAfterClose).Elem(),
		"ErrWriteTooLong":	ValueOf(&tar.ErrWriteTooLong).Elem(),
		"FileInfoHeader":	ValueOf(tar.FileInfoHeader),
		"FormatGNU":	ValueOf(tar.FormatGNU),
		"FormatPAX":	ValueOf(tar.FormatPAX),
		"FormatUSTAR":	ValueOf(tar.FormatUSTAR),
		"FormatUnknown":	ValueOf(tar.FormatUnknown),
		"NewReader":	ValueOf(tar.NewReader),
		"NewWriter":	ValueOf(tar.NewWriter),
		"TypeBlock":	ValueOf(tar.TypeBlock),
		"TypeChar":	ValueOf(tar.TypeChar),
		"TypeCont":	ValueOf(tar.TypeCont),
		"TypeDir":	ValueOf(tar.TypeDir),
		"TypeFifo":	ValueOf(tar.TypeFifo),
		"TypeGNULongLink":	ValueOf(tar.TypeGNULongLink),
		"TypeGNULongName":	ValueOf(tar.TypeGNULongName),
		"TypeGNUSparse":	ValueOf(tar.TypeGNUSparse),
		"TypeLink":	ValueOf(tar.TypeLink),
		"TypeReg":	ValueOf(tar.TypeReg),
		"TypeRegA":	ValueOf(tar.TypeRegA),
		"TypeSymlink":	ValueOf(tar.TypeSymlink),
		"TypeXGlobalHeader":	ValueOf(tar.TypeXGlobalHeader),
		"TypeXHeader":	ValueOf(tar.TypeXHeader),
	}, Types: map[string]Type{
		"Format":	TypeOf((*tar.Format)(nil)).Elem(),
		"Header":	TypeOf((*tar.Header)(nil)).Elem(),
		"Reader":	TypeOf((*tar.Reader)(nil)).Elem(),
		"Writer":	TypeOf((*tar.Writer)(nil)).Elem(),
	}, Untypeds: map[string]string{
		"TypeBlock":	"rune:52",
		"TypeChar":	"rune:51",
		"TypeCont":	"rune:55",
		"TypeDir":	"rune:53",
		"TypeFifo":	"rune:54",
		"TypeGNULongLink":	"rune:75",
		"TypeGNULongName":	"rune:76",
		"TypeGNUSparse":	"rune:83",
		"TypeLink":	"rune:49",
		"TypeReg":	"rune:48",
		"TypeRegA":	"rune:0",
		"TypeSymlink":	"rune:50",
		"TypeXGlobalHeader":	"rune:103",
		"TypeXHeader":	"rune:120",
	}, 
	}
}
