package dao

import "golang.org/x/text/language"

type detection struct {
	Language   language.Tag
	Confidence float64
	IsReliable bool
}
