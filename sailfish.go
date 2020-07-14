package sailfish

import (
	"fmt"
	"reflect"
	"strings"
)

type Gql interface {
	Parse(args, res interface{}) string
}

type gql struct {
	Operation string
	Schema    string
}

func NewQuery(operation, schema string) Gql {
	return &gql{
		Operation: operation,
		Schema:    schema,
	}
}

func (q *gql) ParseArgs(a interface{}) string {
	return iterateArgs(reflect.ValueOf(a).Elem())
}

func iterateArgs(s reflect.Value) string {
	var r string
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		tag := s.Type().Field(i).Tag.Get("gql")
		switch f.Kind() {
		case reflect.Bool:
			r += fmt.Sprintf(` %s:%t `, tag, f.Bool())
		case reflect.Int:
			r += fmt.Sprintf(` %s:%d `, tag, f.Int())
			break
		case reflect.Float32:
			r += fmt.Sprintf(` %s:%f `, tag, f.Float())
			break
		case reflect.Float64:
			r += fmt.Sprintf(` %s:%f `, tag, f.Float())
			break
		case reflect.String:
			r += fmt.Sprintf(` %s:"%s"`, tag, f.String())
			break
		case reflect.Slice:
			d := reflect.ValueOf(f.Interface())
			r += fmt.Sprintf(` %s %s `, tag, iterateSlice(d))
			break
		case reflect.Struct:
			r += fmt.Sprintf(` %s { %s } `, tag, iterateArgs(f))
			break
		default:
			break
		}
	}

	return r
}

func iterateSlice(d reflect.Value) string {
	var o, r string
	for i := 0; i < d.Len(); i++ {
		switch d.Index(i).Kind() {
		case reflect.String:
			o += fmt.Sprintf(` "%s" `, d.Index(i).String())
		case reflect.Int:
			o += fmt.Sprintf(` %d `, d.Index(i).Int())
		case reflect.Struct:
			o += fmt.Sprintf(` { %s } `, iterateArgs(d.Index(i)))
		}
	}
	if d.Len() > 0 {
		r += fmt.Sprintf(` [ %s ] `, o)
	}

	return r
}

func (q *gql) ParseResolver(r interface{}) string {
	var res string
	switch reflect.TypeOf(r).Kind() {
	case reflect.Slice:
		for _, t := range reflect.ValueOf(r).Interface().([]string) {
			res += fmt.Sprintf(" %s ", t)
		}
	}

	return res
}

func (q *gql) Parse(args interface{}, res interface{}) string {
	return fmt.Sprintf(`%s %s(%s){%s}`, strings.ToLower(q.Operation), q.Schema, q.ParseArgs(args), q.ParseResolver(res))
}
