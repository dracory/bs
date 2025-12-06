package bs

import (
	"github.com/dracory/hb"
)

func NavLink() *hb.Tag {
	return hb.Hyperlink().Class("nav-link")
}
