package CSV

import "testing"

func TestRead(t *testing.T) {
	t.Parallel()
	_, err := Read()
	if err != nil {
		t.Fail()
	}
}
