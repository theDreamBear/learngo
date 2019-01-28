package normalTest

import (
	"testing"
)

func Test_add(t *testing.T) {
	var c = []struct{ a, b, c int }{
		{0, 1, 1},
		{1, 2, 3},
		{1000, 1, 1001},
	}

	for _, cc := range c {
		if actual := add(cc.b, cc.a); cc.c != actual {
			t.Errorf("add(%d, %d): got %d, should be %d", cc.a, cc.b, actual, cc.c)
		}
	}

}
