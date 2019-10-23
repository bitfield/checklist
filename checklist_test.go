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
	
	item, err := checklist.NextItem()
	if err != nil {
		t.Fatal(err)
	}
	if item.text != "Line 1" {
		t.Errorf("Expecting 'Line 1', got %q", item.text)
	}

	item, err = checklist.NextItem()
	if err != nil {
		t.Fatal(err)
	}
	if item.text != "Line 2" {
		t.Errorf("Expecting 'Line 2', got %q", item.text)
	}

	item, err = checklist.NextItem()
	if err != nil {
		t.Fatal(err)
	}
	if item.text != "Line 3" {
		t.Errorf("Expecting 'Line 3', got %q", item.text)
	}

	item, err = checklist.NextItem()
	if err == nil {
		t.Errorf("Want error, got nil.")
	}

}
