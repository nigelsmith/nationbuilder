package nationbuilder

import "strings"

type errorCollection []error

func (e *errorCollection) Error() string {
	s := make([]string)

	for _, err := range e {
		s = append(s, err.Error())
	}

	return strings.Join(s, "\n")
}
