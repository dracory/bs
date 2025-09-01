package bs

import (
	"github.com/gouniverse/hb"
	"github.com/spf13/cast"
)

func Column(width int) *hb.Tag {
	return hb.Div().
		Class("col-sm-" + cast.ToString(width))
}
