package klaviyo

import (
	"fmt"
	"strings"
)

type queryFields map[string]string

func (f queryFields) set(key string, values []string) {
	if len(values) > 0 {
		f[fmt.Sprintf("fields[%s]", key)] = strings.Join(values, ",")
	}
}

func (f queryFields) setListFields(values []string) {
	f.set("list", values)
}

func (f queryFields) setProfileFields(values []string) {
	f.set("profile", values)
}

func (f queryFields) setSegmentFields(values []string) {
	f.set("segment", values)
}

func (f queryFields) setEventFields(values []string) {
	f.set("event", values)
}

func (f queryFields) setMetricFields(values []string) {
	f.set("metric", values)
}

func (f queryFields) setTagFields(values []string) {
	f.set("tag", values)
}
