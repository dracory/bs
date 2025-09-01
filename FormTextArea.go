package bs

import (
	"github.com/dracory/hb"
)

func FormTextArea() *hb.Tag {
	return hb.TextArea().Class("form-control")
}
