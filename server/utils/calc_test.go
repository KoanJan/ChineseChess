package utils

import "testing"

func Test_SquaredEucDist(t *testing.T) {

	var (
		x1, y1 int32 = 1, 2
		x2, y2 int32 = 3, 4
		sed    int32 = 8
	)
	t.Logf("(x1, y1) is (%d, %d)\n", x1, y1)
	t.Logf("(x2, y2) is (%d, %d)\n", x2, y2)
	result := SquaredEucDist(x1, y1, x2, y2)
	if result != sed {
		t.Errorf("error answer: %d\n", result)
	}

}
