package bs

import "github.com/dracory/hb"

func FormLabel(label string) *hb.Tag {
	return hb.Div().Class("form-label").HTML(label)
}
