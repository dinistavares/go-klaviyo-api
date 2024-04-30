package klaviyo

import (
	"fmt"
	"strings"
)

type queryFields map[string]string

func (f queryFields) Set(key string, values []string) {
	if len(values) > 0 {
		f[fmt.Sprintf("fields[%s]", key)] = strings.Join(values, ",")
	}
}

func (f queryFields) SetProfileFields(values []string) {
	f.Set("profile", values)
}

func (f queryFields) SetSegmentsFields(values []string) {
	f.Set("segments", values)
}