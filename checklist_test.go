package checklist

import "testing"

func TestLoad(t *testing.T) {
	checklist, err := Load("testdata/checklist.txt")
	if err != nil {
		t.Fatal(err)
	}
	if checklist.Len() != 3 {
		t.Errorf("Len(): want 3, got %d", checklist.Len())
	}
}
