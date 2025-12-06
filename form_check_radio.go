package bs

import (
	"github.com/dracory/hb"
)

func FormCheckRadio() *hb.Tag {
	return hb.Input().Class("form-check-input").Type(hb.TYPE_RADIO)
}
