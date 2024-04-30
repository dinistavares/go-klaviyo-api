package klaviyo

import (
	"fmt"
	"strings"
)

type queryAdditionalFields map[string]string

func (f queryAdditionalFields) Set(key string, values []string) {
	if len(values) > 0 {
		f[fmt.Sprintf("additional-fields[%s]", key)] = strings.Join(values, ",")
	}
}

func (f queryAdditionalFields) SetProfileAdditionalFields(values []string) {
	f.Set("profile", values)
}