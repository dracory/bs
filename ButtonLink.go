package bs

import (
	"github.com/gouniverse/hb"
)

func ButtonLink() *hb.Tag {
	return hb.Hyperlink().Class("btn")
}
