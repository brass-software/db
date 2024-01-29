package db

import "reflect"

type Type struct {
	IsString bool
	IsBool   bool
	IsInt    bool
	IsFloat  bool
	IsRef    bool
	IsStruct bool
	Fields   []Field
	IsList   bool
	IsMap    bool
	ElemType *Type
	Methods  map[string]MethodType
}

type MethodType struct {
	GoName      string
	InputTypes  []Field
	OutputTypes []Field
}

var TypeNames = map[string]*Type{
	// Base Types
	"string": {
		IsString: true,
	},
	"bool": {
		IsBool: true,
	},
	"int": {
		IsInt: true,
	},
	"float": {
		IsFloat: true,
	},

	// Builtin Types
}

func (t *Type) New() any {
	if t.IsString {
		return ""
	}
	if t.IsBool {
		return false
	}
	if t.IsInt {
		return 0
	}
	if t.IsFloat {
		return 0.0
	}
	if t.IsRef {
		elem := t.ElemType.New()
		return &elem
	}
	if t.IsList {
		elem := t.ElemType.New()
		t := reflect.SliceOf(reflect.TypeOf(elem))
		return reflect.MakeSlice(t, 0, 0).Interface()
	}
	if t.IsMap {
		elem := t.ElemType.New()
		t := reflect.MapOf(reflect.TypeOf(""), reflect.TypeOf(elem))
		return reflect.MakeSlice(t, 0, 0).Interface()
	}
	panic("invalid type")
}
