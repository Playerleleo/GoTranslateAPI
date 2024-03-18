package dao

import "golang.org/x/text/language"

type Language struct {
	// Name is the human-readable name of the language.
	Name string

	// Tag is a standard code for the language.
	Tag language.Tag
}
