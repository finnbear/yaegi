package stdlib

// Generated by 'goexports encoding/xml'. Do not edit!

import (
	"encoding/xml"
	"reflect"
)

func init() {
	Value["encoding/xml"] = map[string]reflect.Value{
		"CopyToken":       reflect.ValueOf(xml.CopyToken),
		"Escape":          reflect.ValueOf(xml.Escape),
		"EscapeText":      reflect.ValueOf(xml.EscapeText),
		"HTMLAutoClose":   reflect.ValueOf(xml.HTMLAutoClose),
		"HTMLEntity":      reflect.ValueOf(xml.HTMLEntity),
		"Header":          reflect.ValueOf(xml.Header),
		"Marshal":         reflect.ValueOf(xml.Marshal),
		"MarshalIndent":   reflect.ValueOf(xml.MarshalIndent),
		"NewDecoder":      reflect.ValueOf(xml.NewDecoder),
		"NewEncoder":      reflect.ValueOf(xml.NewEncoder),
		"NewTokenDecoder": reflect.ValueOf(xml.NewTokenDecoder),
		"Unmarshal":       reflect.ValueOf(xml.Unmarshal),
	}
	Type["encoding/xml"] = map[string]reflect.Type{
		"Attr":                 reflect.TypeOf((*xml.Attr)(nil)).Elem(),
		"CharData":             reflect.TypeOf((*xml.CharData)(nil)).Elem(),
		"Comment":              reflect.TypeOf((*xml.Comment)(nil)).Elem(),
		"Decoder":              reflect.TypeOf((*xml.Decoder)(nil)).Elem(),
		"Directive":            reflect.TypeOf((*xml.Directive)(nil)).Elem(),
		"Encoder":              reflect.TypeOf((*xml.Encoder)(nil)).Elem(),
		"EndElement":           reflect.TypeOf((*xml.EndElement)(nil)).Elem(),
		"Marshaler":            reflect.TypeOf((*xml.Marshaler)(nil)).Elem(),
		"MarshalerAttr":        reflect.TypeOf((*xml.MarshalerAttr)(nil)).Elem(),
		"Name":                 reflect.TypeOf((*xml.Name)(nil)).Elem(),
		"ProcInst":             reflect.TypeOf((*xml.ProcInst)(nil)).Elem(),
		"StartElement":         reflect.TypeOf((*xml.StartElement)(nil)).Elem(),
		"SyntaxError":          reflect.TypeOf((*xml.SyntaxError)(nil)).Elem(),
		"TagPathError":         reflect.TypeOf((*xml.TagPathError)(nil)).Elem(),
		"Token":                reflect.TypeOf((*xml.Token)(nil)).Elem(),
		"TokenReader":          reflect.TypeOf((*xml.TokenReader)(nil)).Elem(),
		"UnmarshalError":       reflect.TypeOf((*xml.UnmarshalError)(nil)).Elem(),
		"Unmarshaler":          reflect.TypeOf((*xml.Unmarshaler)(nil)).Elem(),
		"UnmarshalerAttr":      reflect.TypeOf((*xml.UnmarshalerAttr)(nil)).Elem(),
		"UnsupportedTypeError": reflect.TypeOf((*xml.UnsupportedTypeError)(nil)).Elem(),
	}
}