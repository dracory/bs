package bs

import (
	"testing"
)

func TestFormTextArea(t *testing.T) {
	nav := FormTextArea().ToHTML()
	if nav != "<textarea class=\"form-control\"></textarea>" {
		t.Log("Expected: ", `<textarea class="form-control"></textarea>`, " but found:", nav)
		t.Fail()
	}
}
