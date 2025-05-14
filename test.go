package main

import "testing"

func TestAdd(t *testing.T) {
	if add(2, 3) != 5 {
		t.Errorf("Add(2, 3) = %d; want 5", add(2, 3))
	}
	if add(-1, 1) != 0 {
		t.Errorf("Add(-1, 1) = %d; want 0", add(-1, 1))
	}
}
