package bs

import (
	"github.com/dracory/hb"
)

func FormCheckBox() *hb.Tag {
	return hb.Input().Class("form-check-input").Type(hb.TYPE_CHECKBOX)
}
