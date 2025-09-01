package bs

import (
	"github.com/dracory/hb"
)

func FormText(text string) *hb.Tag {
	return hb.Div().Class("form-text").HTML(text)
}
