// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package main

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson6601e8cdDecodeTmpEasyjson(in *jlexer.Lexer, out *Settings) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "allowedHostPaths":
			if in.IsNull() {
				in.Skip()
				out.AllowedHostPaths = nil
			} else {
				in.Delim('[')
				if out.AllowedHostPaths == nil {
					if !in.IsDelim(']') {
						out.AllowedHostPaths = make([]HostPath, 0, 2)
					} else {
						out.AllowedHostPaths = []HostPath{}
					}
				} else {
					out.AllowedHostPaths = (out.AllowedHostPaths)[:0]
				}
				for !in.IsDelim(']') {
					var v1 HostPath
					(v1).UnmarshalEasyJSON(in)
					out.AllowedHostPaths = append(out.AllowedHostPaths, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson6601e8cdEncodeTmpEasyjson(out *jwriter.Writer, in Settings) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"allowedHostPaths\":"
		out.RawString(prefix[1:])
		if in.AllowedHostPaths == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.AllowedHostPaths {
				if v2 > 0 {
					out.RawByte(',')
				}
				(v3).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Settings) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6601e8cdEncodeTmpEasyjson(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Settings) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6601e8cdEncodeTmpEasyjson(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Settings) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6601e8cdDecodeTmpEasyjson(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Settings) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6601e8cdDecodeTmpEasyjson(l, v)
}
func easyjson6601e8cdDecodeTmpEasyjson1(in *jlexer.Lexer, out *HostPath) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "pathPrefix":
			out.PathPrefix = string(in.String())
		case "readOnly":
			out.ReadOnly = bool(in.Bool())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson6601e8cdEncodeTmpEasyjson1(out *jwriter.Writer, in HostPath) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"pathPrefix\":"
		out.RawString(prefix[1:])
		out.String(string(in.PathPrefix))
	}
	{
		const prefix string = ",\"readOnly\":"
		out.RawString(prefix)
		out.Bool(bool(in.ReadOnly))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v HostPath) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6601e8cdEncodeTmpEasyjson1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v HostPath) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6601e8cdEncodeTmpEasyjson1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *HostPath) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6601e8cdDecodeTmpEasyjson1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *HostPath) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6601e8cdDecodeTmpEasyjson1(l, v)
}
