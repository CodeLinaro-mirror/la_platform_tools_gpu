package template

import (
	"bytes"
	"fmt"
	"strings"
	"unicode"
	"reflect"
)

type stringList []string

// String returns the concatenation of all the string segments with no separator.
func (l stringList) String() string {
	return strings.Join([]string(l), "")
}


// stringify transforms the input parameters into a string list. Arrays and
// slices are flattened into a sequential list of strings.
func stringify(v ...interface{}) stringList {
	out := stringList{}
	for _, v := range v {
		switch v := v.(type) {
		case string:
			out = append(out, v)
		case []string:
			out = append(out, v...)
		case stringList:
			out = append(out, v...)
		default:
			switch reflect.TypeOf(v).Kind() {
			case reflect.Array, reflect.Slice:
				v := reflect.ValueOf(v)
				for i, c := 0, v.Len(); i < c; i++ {
					out = append(out, stringify(v.Index(i).Interface())...)
				}
			default:
				out = append(out, fmt.Sprintf("%v", v))
			}
		}
	}
	return out
}



// Join returns the concatenation of all the string segments with the specified separator.
func (Functions) JoinWith(sep string, v ...interface{}) string {
	l := stringify(v...)
	return strings.Join([]string(l), sep)
}

// Split slices each string segement into all substrings separated by sep. The returned stringList
// will not contain any occurances of sep.
func (Functions) SplitOn(sep string, v ...interface{}) stringList {
	l := stringify(v...)
	out := stringList{}
	for _, s := range l {
		for _, v := range strings.Split(s, sep) {
			if len(v) > 0 {
				out = append(out, v)
			}
		}
	}
	return out
}

// SplitUpperCase slices each string segment before and after each upper-case rune.
func (Functions) SplitUpperCase(v ...interface{}) stringList {
	l := stringify(v...)
	out := stringList{}
	for _, s := range l {
		str := ""
		for _, r := range s {
			if unicode.IsUpper(r) {
				if len(str) > 0 {
					out = append(out, str)
					str = ""
				}
				out = append(out, string(r))
			} else {
				str += string(r)
			}
		}
		if len(str) > 0 {
			out = append(out, str)
		}
	}
	return out
}

// SplitPascalCase slices each string segment at each transition from an letter rune to a upper-case
// letter rune.
func (Functions) SplitPascalCase(v ...interface{}) stringList {
	l := stringify(v...)
	out := stringList{}
	for _, str := range l {
		runes := bytes.Runes([]byte(str))
		str := ""
		p := 'x'
		for _, c := range runes {
			if unicode.IsLetter(p) && unicode.IsUpper(c) {
				if len(str) > 0 {
					out = append(out, str)
				}
				str = string(c)
			} else {
				str += string(c)
			}
			p = c
		}
		if len(str) > 0 {
			out = append(out, str)
		}
	}
	return out
}

// Title capitalizes each letter of each string segment.
func (Functions) Title(v ...interface{}) stringList {
	l := stringify(v...)
	out := make(stringList, len(l))
	for i, s := range l {
		first := true
		out[i] = strings.Map(func(r rune) rune {
			if first {
				first = false
				return unicode.ToTitle(r)
			} else {
				return r
			}
		}, s)
	}
	return out
}

// Untitle lower-cases each letter of each string segment.
func (Functions) Untitle(v ...interface{}) stringList {
	l := stringify(v...)
	out := make(stringList, len(l))
	for i, s := range l {
		first := true
		out[i] = strings.Map(func(r rune) rune {
			if first {
				first = false
				return unicode.ToLower(r)
			} else {
				return r
			}
		}, s)
	}
	return out
}

// Lower lower-cases all letters of each string segment.
func (Functions) Lower(v ...interface{}) stringList {
	l := stringify(v...)
	out := make(stringList, len(l))
	for i, s := range l {
		out[i] = strings.ToLower(s)
	}
	return out
}

// Upper upper-cases all letters of each string segment.
func (Functions) Upper(v ...interface{}) stringList {
	l := stringify(v...)
	out := make(stringList, len(l))
	for i, s := range l {
		out[i] = strings.ToUpper(s)
	}
	return out
}

// Contains returns true if any string segment contains substr.
func (Functions) Contains(substr string, v ...interface{}) bool {
	l := stringify(v...)
	for _, s := range l {
		if strings.Contains(s, substr) {
			return true
		}
	}
	return false
}

// Replace any occurance of old with new in tbe string segments.
func (Functions) Replace(old string, new string, v ...interface{}) stringList {
	l := stringify(v...)
	out := stringList{}
	for _, s := range l {
		s = strings.Replace(s, old, new, -1)
		if len(s) > 0 {
			out = append(out, s)
		}
	}
	return out
}

func (Functions) TrimLeft(cutset string, v ...interface{}) stringList {
	l := stringify(v...)
	out := stringList{}
	for _, s := range l {
		s = strings.TrimLeft(s, cutset)
		if len(s) > 0 {
			out = append(out, s)
		}
	}
	return out
}
