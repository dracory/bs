package bs

import (
	"github.com/dracory/hb"
)

func CardImageTop(src string) *hb.Tag {
	return hb.Image(src).Class("card-image-top")
}
