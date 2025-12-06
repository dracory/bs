package bs

import (
	"strings"
	"testing"
)

func TestPaginationTwoPages(t *testing.T) {
	// Test case for the bug: 2 pages starting from 1
	options := PaginationOptions{
		NumberItems:       10,
		CurrentPageNumber: 1,
		PerPage:           5,
		PagesToShow:       5,
		URL:               "/page/",
		FirstPageStartsAt: 1,
	}

	html := Pagination(options)

	// Should contain both page 1 and page 2
	if !strings.Contains(html, ">1<") {
		t.Error("Expected page 1 to be displayed")
	}
	if !strings.Contains(html, ">2<") {
		t.Error("Expected page 2 to be displayed")
	}

	// Page 1 should be active
	if !strings.Contains(html, "active") {
		t.Error("Expected page 1 to be marked as active")
	}
}

func TestPaginationSinglePage(t *testing.T) {
	// Test with only 1 page total
	options := PaginationOptions{
		NumberItems:       5,
		CurrentPageNumber: 1,
		PerPage:           5,
		PagesToShow:       5,
		URL:               "/page/",
		FirstPageStartsAt: 1,
	}

	html := Pagination(options)

	// Should return empty string for less than 2 pages
	if html != "" {
		t.Error("Expected empty string for single page pagination")
	}
}

func TestPaginationMultiplePages(t *testing.T) {
	// Test with 5 pages, viewing page 3
	options := PaginationOptions{
		NumberItems:       50,
		CurrentPageNumber: 3,
		PerPage:           10,
		PagesToShow:       5,
		URL:               "/page/",
		FirstPageStartsAt: 1,
	}

	html := Pagination(options)

	// Should contain page 3 and surrounding pages
	if !strings.Contains(html, ">1<") {
		t.Error("Expected page 1 to be displayed")
	}
	if !strings.Contains(html, ">3<") {
		t.Error("Expected page 3 to be displayed")
	}
	if !strings.Contains(html, ">5<") {
		t.Error("Expected page 5 to be displayed")
	}

	// Page 3 should be active
	if !strings.Contains(html, `<li class="page-item active"><a class="page-link" href="/page/3" style="cursor:pointer;">3</a></li>`) {
		t.Error("Expected page 3 to be marked as active")
	}

	// Should have previous and next links
	if !strings.Contains(html, "&laquo;") {
		t.Error("Expected previous link")
	}
	if !strings.Contains(html, "&raquo;") {
		t.Error("Expected next link")
	}
}

func TestPaginationFirstPage(t *testing.T) {
	// Test on first page with multiple pages
	options := PaginationOptions{
		NumberItems:       100,
		CurrentPageNumber: 1,
		PerPage:           10,
		PagesToShow:       5,
		URL:               "/page/",
		FirstPageStartsAt: 1,
	}

	html := Pagination(options)

	// Should not have previous link when on first page
	if strings.Count(html, "&laquo;") > 0 {
		t.Error("Expected no previous link on first page")
	}

	// Should have next link
	if !strings.Contains(html, "&raquo;") {
		t.Error("Expected next link")
	}

	// Should display correct pages
	if !strings.Contains(html, ">1<") {
		t.Error("Expected page 1 to be displayed")
	}
}

func TestPaginationLastPage(t *testing.T) {
	// Test on last page with multiple pages
	options := PaginationOptions{
		NumberItems:       100,
		CurrentPageNumber: 10,
		PerPage:           10,
		PagesToShow:       5,
		URL:               "/page/",
		FirstPageStartsAt: 1,
	}

	html := Pagination(options)

	// Should have previous link
	if !strings.Contains(html, "&laquo;") {
		t.Error("Expected previous link")
	}

	// Should not have next link when on last page
	if strings.Count(html, "&raquo;") > 0 {
		t.Error("Expected no next link on last page")
	}

	// Should display page 10
	if !strings.Contains(html, ">10<") {
		t.Error("Expected page 10 to be displayed")
	}
}

func TestPaginationFirstPageStartsFromZero(t *testing.T) {
	// Test with FirstPageStartsAt = 0 (zero-indexed pages)
	options := PaginationOptions{
		NumberItems:       20,
		CurrentPageNumber: 0,
		PerPage:           5,
		PagesToShow:       5,
		URL:               "/page/",
		FirstPageStartsAt: 0,
	}

	html := Pagination(options)

	// Should contain page 0 and page 1
	if !strings.Contains(html, ">0<") {
		t.Error("Expected page 0 to be displayed")
	}
	if !strings.Contains(html, ">1<") {
		t.Error("Expected page 1 to be displayed")
	}
}

func TestPaginationCorrectPageNumbers(t *testing.T) {
	// Test that page numbers are correct (not off by one)
	options := PaginationOptions{
		NumberItems:       30,
		CurrentPageNumber: 2,
		PerPage:           10,
		PagesToShow:       3,
		URL:               "/page/",
		FirstPageStartsAt: 1,
	}

	html := Pagination(options)

	// Should contain pages 1, 2, 3
	if !strings.Contains(html, ">1<") {
		t.Error("Expected page 1 to be displayed")
	}
	if !strings.Contains(html, ">2<") {
		t.Error("Expected page 2 to be displayed")
	}
	if !strings.Contains(html, ">3<") {
		t.Error("Expected page 3 to be displayed")
	}

	// Should not contain page 4
	if strings.Contains(html, ">4<") {
		t.Error("Did not expect page 4 to be displayed")
	}
}

func TestPaginationURLCorrectness(t *testing.T) {
	// Test that URLs are constructed correctly
	options := PaginationOptions{
		NumberItems:       20,
		CurrentPageNumber: 1,
		PerPage:           5,
		PagesToShow:       2,
		URL:               "/users?page=",
		FirstPageStartsAt: 1,
	}

	html := Pagination(options)

	// Should contain correct URLs
	if !strings.Contains(html, "/users?page=1") {
		t.Error("Expected page 1 URL to be correct")
	}
	if !strings.Contains(html, "/users?page=2") {
		t.Error("Expected page 2 URL to be correct")
	}
}

func TestPaginationMinMaxCalculation(t *testing.T) {
	testCases := []struct {
		name              string
		numberItems       int
		currentPageNumber int
		perPage           int
		pagesToShow       int
		firstPageStartsAt int
		expectedMin       int
		expectedMax       int
		expectedPrevious  int
		expectedNext      int
	}{
		{
			name:              "Two pages, page 1",
			numberItems:       10,
			currentPageNumber: 1,
			perPage:           5,
			pagesToShow:       5,
			firstPageStartsAt: 1,
			expectedMin:       1,
			expectedMax:       2,
			expectedPrevious:  0,
			expectedNext:      2,
		},
		{
			name:              "Five pages, page 3",
			numberItems:       50,
			currentPageNumber: 3,
			perPage:           10,
			pagesToShow:       5,
			firstPageStartsAt: 1,
			expectedMin:       1,
			expectedMax:       5,
			expectedPrevious:  2,
			expectedNext:      4,
		},
		{
			name:              "Ten pages, page 1",
			numberItems:       100,
			currentPageNumber: 1,
			perPage:           10,
			pagesToShow:       5,
			firstPageStartsAt: 1,
			expectedMin:       1,
			expectedMax:       6,
			expectedPrevious:  0,
			expectedNext:      2,
		},
		{
			name:              "Ten pages, page 10",
			numberItems:       100,
			currentPageNumber: 10,
			perPage:           10,
			pagesToShow:       5,
			firstPageStartsAt: 1,
			expectedMin:       5,
			expectedMax:       10,
			expectedPrevious:  9,
			expectedNext:      11,
		},
		{
			name:              "Four pages, zero-indexed",
			numberItems:       40,
			currentPageNumber: 0,
			perPage:           10,
			pagesToShow:       4,
			firstPageStartsAt: 0,
			expectedMin:       0,
			expectedMax:       4,
			expectedPrevious:  -1,
			expectedNext:      1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			min, max, prev, next := paginationMinMaxPrevNext(
				tc.numberItems,
				tc.currentPageNumber,
				tc.perPage,
				tc.pagesToShow,
				tc.firstPageStartsAt,
			)

			if min != tc.expectedMin {
				t.Errorf("Expected min %d, got %d", tc.expectedMin, min)
			}
			if max != tc.expectedMax {
				t.Errorf("Expected max %d, got %d", tc.expectedMax, max)
			}
			if prev != tc.expectedPrevious {
				t.Errorf("Expected previous %d, got %d", tc.expectedPrevious, prev)
			}
			if next != tc.expectedNext {
				t.Errorf("Expected next %d, got %d", tc.expectedNext, next)
			}
		})
	}
}

func TestPaginationPageCountBoundaries(t *testing.T) {
	testCases := []struct {
		name              string
		numberItems       int
		currentPageNumber int
		perPage           int
		expectedPageCount int
	}{
		{
			name:              "1 item per page = 5 pages",
			numberItems:       5,
			currentPageNumber: 1,
			perPage:           1,
			expectedPageCount: 5,
		},
		{
			name:              "10 items, 3 per page = 4 pages",
			numberItems:       10,
			currentPageNumber: 1,
			perPage:           3,
			expectedPageCount: 4,
		},
		{
			name:              "100 items, 10 per page = 10 pages",
			numberItems:       100,
			currentPageNumber: 1,
			perPage:           10,
			expectedPageCount: 10,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			options := PaginationOptions{
				NumberItems:       tc.numberItems,
				CurrentPageNumber: tc.currentPageNumber,
				PerPage:           tc.perPage,
				PagesToShow:       tc.expectedPageCount,
				URL:               "/page/",
				FirstPageStartsAt: 1,
			}

			html := Pagination(options)

			// For expectedPageCount = 1, should return empty string
			if tc.expectedPageCount < 2 && html != "" {
				t.Error("Expected empty string for less than 2 pages")
			}

			// For expectedPageCount >= 2, should contain pages
			if tc.expectedPageCount >= 2 {
				if !strings.Contains(html, ">1<") {
					t.Error("Expected page 1 to be displayed")
				}
			}
		})
	}
}
