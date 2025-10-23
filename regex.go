package klaviyo

import "regexp"

const (
	pageCursorRegex = `^https:\/\/a\.klaviyo\.com\/api\/([a-z]+)\?page%5Bcursor%5D=([a-zA-Z0-9+\/=]+)$`
)

var (
	klaviyoPageCursorRegex = regexp.MustCompile(pageCursorRegex)
)