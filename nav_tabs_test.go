package bs

import (
	"testing"
)

func TestNavTabs(t *testing.T) {
	nav := NavTabs().ToHTML()
	if nav != "<ul class=\"nav nav-tabs\"></ul>" {
		t.Log("Expected: ", `<ul class="nav-tabs"></ul>`, " but found:", nav)
		t.Fail()
	}
}
