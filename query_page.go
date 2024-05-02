package klaviyo

import (
	"fmt"
)

type queryPage map[string]string

func (f queryPage) set(key string, value interface{}) {
	if value != nil {
		f[fmt.Sprintf("page[%s]", key)] = fmt.Sprintf("%v", value)
	}
}

func (f queryPage) setPageCursor(value string) {
	f.set("cursor", value)
}

func (f queryPage) setPageSize(value int) {
	f.set("size", value)
}
