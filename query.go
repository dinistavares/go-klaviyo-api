package klaviyo

import (
	"strings"
)

type QueryValues map[string]string

func (p QueryValues) encode() string {
	var query string

	count := 0

	for key, value := range p {
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

func (p QueryValues) setValues(fields map[string]string) {
	for key, value := range fields {
		p[key] = value
	}
}

func (p QueryValues) sort(value string) {
	p["sort"] = value
}

func (p QueryValues) filter(values QueryFilter) {
	if len(values) > 0 {
		p["filter"] = strings.Join(values, ",")
	}
}
