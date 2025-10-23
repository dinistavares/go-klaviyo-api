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

// Set page cursor. Value must be a string, it can be either the cursor value or the next \
//    cursor endpoint returned from the previous request.
func (f queryPage) setPageCursor(value string) {
	match := klaviyoPageCursorRegex.FindStringSubmatch(value)

	// Only use cursor if the whole URL is provided.
	if len(match) == 3 {
		value = match[2]
	}

	f.set("cursor", value)
}

func (f queryPage) setPageSize(value int) {
	f.set("size", value)
}
