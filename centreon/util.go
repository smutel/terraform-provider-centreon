package centreon

func expandToStringSlice(v []interface{}) []string {
	s := make([]string, len(v))
	for i, val := range v {
		if strVal, ok := val.(string); ok {
			s[i] = strVal
		}
	}

	return s
}

func diffSlices(oldSlice []string, newSlice []string) []string {
	var diff []string

	for _, x := range oldSlice {
		found := false
		for _, y := range newSlice {
			if x == y {
				found = true
			}
		}

		if found == false {
			diff = append(diff, x)
		}
	}

	return diff
}
