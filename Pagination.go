package bs

import (
	"math"

	"github.com/dracory/hb"
	"github.com/spf13/cast"
)

type PaginationOptions struct {
	NumberItems       int
	CurrentPageNumber int
	PerPage           int
	PagesToShow       int
	URL               string
	FirstPageStartsAt int
}

func Pagination(options PaginationOptions) string {
	min, max, previousPageNumber, nextPageNumber := paginationMinMaxPrevNext(options.NumberItems, options.CurrentPageNumber, options.PerPage, options.PagesToShow, options.FirstPageStartsAt)

	if max < 2 {
		return ""
	}

	liStart := hb.LI().Class("page-item").Children([]hb.TagInterface{
		hb.Hyperlink().
			Class("page-link").
			Style("cursor:pointer;").
			Href(options.URL + cast.ToString(previousPageNumber)).
			HTML("&laquo;"),
	})

	liEnd := hb.LI().Class("page-item").Children([]hb.TagInterface{
		hb.Hyperlink().
			Class("page-link").
			Style("cursor:pointer;").
			Href(options.URL + cast.ToString(nextPageNumber)).
			HTML("&raquo;"),
	})

	pages := []hb.TagInterface{}

	if previousPageNumber >= options.FirstPageStartsAt {
		pages = append(pages, liStart)
	}

	for i := min; i < max; i++ {
		active := ""
		if i == options.CurrentPageNumber {
			active = " active"
		}

		li := hb.LI().Class("page-item" + active).Children([]hb.TagInterface{
			hb.Hyperlink().
				Class("page-link").
				Style("cursor:pointer;").
				Href(options.URL + cast.ToString(i)).
				HTML(cast.ToString(i + 1)),
		})

		pages = append(pages, li)
	}

	if nextPageNumber < max {
		pages = append(pages, liEnd)
	}

	nav := hb.Nav().Children([]hb.TagInterface{
		hb.UL().Class("pagination").Children(pages),
	})

	return nav.ToHTML()
}

func paginationMinMaxPrevNext(numberItems int, currentPageNumber int, perPage int, pagesToShow int, firstPageStartsAt int) (min int, max int, previousPageNumber int, nextPageNumber int) {
	previousPageNumber = currentPageNumber - 1
	nextPageNumber = currentPageNumber + 1

	numberOfPages := int(math.Ceil(float64(numberItems) / float64(perPage)))
	// adjacent := int(math.Floor(((float64(currentPageNumber) - 3) / 2)))

	min = currentPageNumber - pagesToShow/2
	if pagesToShow%2 == 0 {
		min++ // if even number
	}

	if min < firstPageStartsAt {
		min = firstPageStartsAt
	}

	max = min + pagesToShow - 1

	if max > numberOfPages {
		max = numberOfPages

		// too little pages, pad on the left
		min = max - pagesToShow
		if min < firstPageStartsAt {
			min = firstPageStartsAt
		}
	}

	return min, max, previousPageNumber, nextPageNumber
}
