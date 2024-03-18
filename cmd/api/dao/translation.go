package dao

import "golang.org/x/text/language"

type Translation struct {
	Text   string
	Source language.Tag
	Model  string
}
