package dao

import "golang.org/x/text/language"

type Options struct {
	Source language.Tag
	Format Format
	Model  string
}
