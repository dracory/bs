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

// Pagination generates a pagination component given the options.
//
// The options are:
// - NumberItems: The total number of items that are being paginated.
// - CurrentPageNumber: The current page number that is being displayed.
// - PerPage: The number of items per page.
// - PagesToShow: The number of pages to show in the pagination component.
// - URL: The base URL of the pagination component.
// - FirstPageStartsAt: The page number at which the first page starts.
//
// The pagination component will generate a list of page numbers from min to max,
// where min is the minimum page number that should be displayed, and max is the maximum
// page number that should be displayed. The page numbers will be hyperlinked to the
// URL with the page number appended to the end. The page numbers will be displayed
// in a <nav> element with a <ul> element containing the page numbers. If the
// previous page number is greater than or equal to the first page starts at, then a
// previous page link will be displayed. If the next page number is less than the maximum
// page number, then a next page link will be displayed. The previous and next page links
// will be displayed in <li> elements with the class "page-item". The page numbers
// will be displayed in <li> elements with the class "page-item" and the class "active"
// if the page number matches the current page number.
//
// The function will return an empty string if the maximum page number is less than 2.
//
// Example:
//
//	options := PaginationOptions{
//		NumberItems: 10,
//		CurrentPageNumber: 2,
//		PerPage: 5,
//		PagesToShow: 5,
//		URL: "/users/profile/",
//		FirstPageStartsAt: 1,
//	}
//
// pagination := Pagination(options)
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

	for i := min; i <= max; i++ {
		active := ""
		if i == options.CurrentPageNumber {
			active = " active"
		}

		li := hb.LI().Class("page-item" + active).Children([]hb.TagInterface{
			hb.Hyperlink().
				Class("page-link").
				Style("cursor:pointer;").
				Href(options.URL + cast.ToString(i)).
				HTML(cast.ToString(i)),
		})

		pages = append(pages, li)
	}

	if nextPageNumber <= max {
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

	max = min + pagesToShow

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
