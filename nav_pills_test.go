package bs

import (
	"testing"
)

func TestNavPills(t *testing.T) {
	nav := NavPills().ToHTML()
	if nav != "<ul class=\"nav nav-pills\"></ul>" {
		t.Log("Expected: ", `<ul class="nav-pills"></ul>`, " but found:", nav)
		t.Fail()
	}
}
