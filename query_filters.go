package klaviyo

import (
	"fmt"
	"reflect"
	"time"
)

type QueryFilter []string

func (f *QueryFilter) negateFilter(filter string, negates ...bool) string {
	negate := false

	if len(negates) > 0 {
		negate = negates[0]
	}

	if negate {
		filter = fmt.Sprintf("not(%s)", filter)
	}

	return filter
}

// Creates a 'equals' filter
func (f *QueryFilter) CreateEqualsFilter(operator string, value interface{}, negate ...bool) {
	var (
		filterValue string
	)

	switch value := value.(type) {
	case string:
		filterValue = fmt.Sprintf(`"%s"`, value)
	case bool:
		filterValue = fmt.Sprintf("%t", value)
	case int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64,
		float32, float64, complex64, complex128:
		filterValue = fmt.Sprintf("%d", value)
	case time.Time:
		filterValue = value.Format("2006-01-02T15:04:05Z")
	default:
		// TODO: add arrays
		filterValue = fmt.Sprintf("%v", value)
	}

	filter := fmt.Sprintf("equals(%s,%s)", operator, filterValue)
	filter = f.negateFilter(filter, negate...)

	*f = append(*f, filter)
}

// Creates a 'greater than' filter for time.Time values
func (f *QueryFilter) CreateGreaterThanTimeFilter(operator string, value time.Time, negate ...bool) {
	filter := fmt.Sprintf("greater-than(%s,%s)", operator, value.Format("2006-01-02T15:04:05Z"))
	filter = f.negateFilter(filter, negate...)

	*f = append(*f, filter)
}

// Creates a 'greater than or equal' filter for time.Time values
func (f *QueryFilter) CreateGreaterThanOrEqualTimeFilter(operator string, value time.Time, negate ...bool) {
	filter := fmt.Sprintf("greater-or-equal(%s,%s)", operator, value.Format("2006-01-02T15:04:05Z"))
	filter = f.negateFilter(filter, negate...)

	*f = append(*f, filter)
}

// Creates a 'greater than' filter for uint64 values
func (f *QueryFilter) CreateGreaterThanNumberFilter(operator string, value uint64, negate ...bool) {
	filter := fmt.Sprintf("greater-than(%s,%d)", operator, value)
	filter = f.negateFilter(filter, negate...)

	*f = append(*f, filter)
}

// Creates a 'greater than or equal' filter for uint64 values
func (f *QueryFilter) CreateGreaterThanOrEqualNumberFilter(operator string, value uint64, negate ...bool) {
	filter := fmt.Sprintf("greater-or-equal(%s,%d)", operator, value)
	filter = f.negateFilter(filter, negate...)

	*f = append(*f, filter)
}

// Creates a 'less than' filter for time.Time values
func (f *QueryFilter) CreateLessThanTimeFilter(operator string, value time.Time, negate ...bool) {
	filter := fmt.Sprintf("less-than(%s,%s)", operator, value.Format("2006-01-02T15:04:05Z"))
	filter = f.negateFilter(filter, negate...)

	*f = append(*f, filter)
}

// Creates a 'less than or equal' filter for time.Time values
func (f *QueryFilter) CreateLessThanOrEqualsTimeFilter(operator string, value time.Time, negate ...bool) {
	filter := fmt.Sprintf("less-or-equal(%s,%s)", operator, value.Format("2006-01-02T15:04:05Z"))
	filter = f.negateFilter(filter, negate...)

	*f = append(*f, filter)
}

// Creates a 'less than' filter for uint64 values
func (f *QueryFilter) CreateLessThanNumberFilter(operator string, value uint64, negate ...bool) {
	filter := fmt.Sprintf("less-than(%s,%d)", operator, value)
	filter = f.negateFilter(filter, negate...)

	*f = append(*f, filter)
}

// Creates a 'less than or equal' filter for uint64 values
func (f *QueryFilter) CreateLessThanOrEqualsNumberFilter(operator string, value uint64, negate ...bool) {
	filter := fmt.Sprintf("less-or-equal(%s,%d)", operator, value)
	filter = f.negateFilter(filter, negate...)

	*f = append(*f, filter)
}

// Creates a 'contains' filter for string values
func (f *QueryFilter) CreateContainsFilter(operator string, value interface{}, negate ...bool) {
	filterValue := parseContainsFilterValues(value)

	if filterValue != "" {
		filter := fmt.Sprintf("contains(%s,%s)", operator, filterValue)
		filter = f.negateFilter(filter, negate...)
		*f = append(*f, filter)
	}
}

// Creates a 'contains-any' filter for string values
func (f *QueryFilter) CreateContainsAnyFilter(operator string, value interface{}, negate ...bool) {
	filterValue := parseContainsFilterValues(value)

	if filterValue != "" {
		filter := fmt.Sprintf("contains-any(%s,%s)", operator, filterValue)
		filter = f.negateFilter(filter, negate...)
		*f = append(*f, filter)
	}
}

// Creates a 'contains-all' filter for string values
func (f *QueryFilter) CreateContainsAllFilter(operator string, value interface{}, negate ...bool) {
	filterValue := parseContainsFilterValues(value)

	if filterValue != "" {
		filter := fmt.Sprintf("contains-all(%s,%s)", operator, filterValue)
		filter = f.negateFilter(filter, negate...)
		*f = append(*f, filter)
	}
}

// Creates a 'ends-with' filter for string values
func (f *QueryFilter) CreateEndsWithFilter(operator string, value interface{}, negate ...bool) {
	filterValue := parseContainsFilterValues(value)

	if filterValue != "" {
		filter := fmt.Sprintf("ends-with(%s,%s)", operator, filterValue)
		filter = f.negateFilter(filter, negate...)
		*f = append(*f, filter)
	}
}

// Creates a 'starts-with' filter for string values
func (f *QueryFilter) CreateStartsWithFilter(operator string, value interface{}, negate ...bool) {
	filterValue := parseContainsFilterValues(value)

	if filterValue != "" {
		filter := fmt.Sprintf("starts-with(%s,%s)", operator, filterValue)
		filter = f.negateFilter(filter, negate...)
		*f = append(*f, filter)
	}
}

// Creates a 'any' filter. Note that the value should be a slice (Accepts: []string, []int64, []uint64, []bool, []time.Time)
func (f *QueryFilter) CreateAnyFilter(operator string, value interface{}, negate ...bool) {
	var filterValue string

	firstSlice, isSlice := interfaceIsSlice(value)

	// Only accept slices
	if !isSlice {
		return
	}

	switch firstSlice.(type) {
	case string:
		stringSlice := value.([]string)

		for i, element := range stringSlice {
			if element != "" {
				if i > 0 {
					filterValue += fmt.Sprintf(`,"%s"`, element)
				} else {
					filterValue += fmt.Sprintf(`"%s"`, element)
				}
			}
		}
	case bool:
		boolSlice := value.([]bool)

		for i, element := range boolSlice {
			if i > 0 {
				filterValue += fmt.Sprintf(",%t", element)
			} else {
				filterValue += fmt.Sprintf("%t", element)
			}
		}
	case time.Time:
		timeSlice := value.([]time.Time)

		for i, element := range timeSlice {
			elementString := element.Format("2006-01-02T15:04:05Z")

			if i > 0 {
				filterValue += fmt.Sprintf(",%s", elementString)
			} else {
				filterValue += elementString
			}
		}
	case int:
		numberSlice := value.([]int)

		for i, element := range numberSlice {
			if i > 0 {
				filterValue += fmt.Sprintf(",%d", element)
			} else {
				filterValue += fmt.Sprintf("%d", element)
			}
		}
	case uint64:
		numberSlice := value.([]uint64)

		for i, element := range numberSlice {
			if i > 0 {
				filterValue += fmt.Sprintf(",%d", element)
			} else {
				filterValue += fmt.Sprintf("%d", element)
			}
		}
	default:
		return
	}

	if filterValue != "" {
		filter := fmt.Sprintf("any(%s,[%s])", operator, filterValue)
		filter = f.negateFilter(filter)

		*f = append(*f, filter)
	}
}

// Checks if an interface is a slice and returns the first slice value.
func interfaceIsSlice(i interface{}) (interface{}, bool) {
	v := reflect.ValueOf(i)

	if v.Kind() == reflect.Slice && v.Len() > 0 {
		first := v.Index(0).Interface()

		return first, true
	}

	return nil, false
}

func parseContainsFilterValues(value interface{}) string {
	var filterValue string

	firstSlice, isSlice := interfaceIsSlice(value)

	if isSlice {
		switch firstSlice.(type) {
		case string:
			valueSlice := value.([]string)

			for i, element := range valueSlice {
				if element != "" {
					if i > 0 {
						filterValue += fmt.Sprintf(`,"%s"`, element)
					} else {
						filterValue += fmt.Sprintf(`"%s"`, element)
					}
				}
			}
		default:
			// Ignore all other types
			return filterValue
		}

		if filterValue != "" {
			filterValue = fmt.Sprintf("[%s]", filterValue)
		}
	} else {
		filterValue = fmt.Sprintf(`"%s"`, value.(string))
	}

	return filterValue
}