package klaviyo

import (
	"fmt"
)

type queryPage map[string]string

func (f queryPage) Set(key string, value interface{}) {
	if value != nil {
		f[fmt.Sprintf("page[%s]", key)] = fmt.Sprintf("%v", value)
	}
}

func (f queryPage) SetPageCursor(value string) {
	f.Set("cursor", value)
}

func (f queryPage) SetPageSize(value int) {
	f.Set("size", value)
}
