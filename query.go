package klaviyo

import (
	"reflect"
	"strings"
)

type QueryValues map[string]string

func (v *QueryValues) getQueryValues() QueryValues {
	return *v
}

func (v QueryValues) encode() string {
	var query string

	count := 0

	for key, value := range v {
		if count > 0 {
			query += "&"
		} else {
			query = "?"
		}

		query += key + "=" + value

		count++
	}

	return query
}

func (v QueryValues) setValues(fields map[string]string) {
	for key, value := range fields {
		v[key] = value
	}
}

func (v QueryValues) sort(value string) {
	v["sort"] = value
}

func (v QueryValues) include(values []string) {
	v["include"] = strings.Join(values, ",")
}

func (v QueryValues) filter(values QueryFilter) {
	if len(values) > 0 {
		v["filter"] = strings.Join(values, ",")
	}
}

func isPointerWithQueryValues(i interface{}) (interface{}, bool) {
	if i == nil {
		return nil, false
	}

	val := reflect.ValueOf(i)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() == reflect.Struct {
		field := val.FieldByName("QueryValues")

		if field.IsValid() {
			return field.Interface(), true
		}
	}

	return nil, false
}
