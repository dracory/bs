package bs

import (
	"github.com/dracory/hb"
)

func Modal() *hb.Tag {
	return hb.Div().Class("modal")
}

func ModalDialog() *hb.Tag {
	return hb.Div().Class("modal-dialog")
}

func ModalContent() *hb.Tag {
	return hb.Div().Class("modal-content")
}

func ModalHeader() *hb.Tag {
	return hb.Div().Class("modal-header")
}

func ModalBody() *hb.Tag {
	return hb.Div().Class("modal-body")
}

func ModalFooter() *hb.Tag {
	return hb.Div().Class("modal-footer")
}
