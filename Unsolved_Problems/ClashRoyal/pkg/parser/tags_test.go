package parser_test

import (
	"testing"

	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/parser"
)

func TestToRequestTag(t *testing.T) {
	tag := parser.ToRequestTag("abc")

	if tag != "%25bc" {
		t.Errorf("expected %%25bc got %s", tag)
	}
}
