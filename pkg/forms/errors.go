package forms

type errors map[string][]string

// Add method: adds an error message for our given field to our map
func (e errors) Add(field, message string) {
	e[field] = append(e[field], message)
}

// Get: will retrieve the first error message for a given field in the map
func (e errors) Get(field string) string {
	es := e[field]

	if len(es) == 0 {
		return ""
	}

	return es[0]
}
