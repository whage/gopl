package popcount

import "testing"

func TestPopcount(t *testing.T) {
	if PopCount(14) != 3 {
		t.Errorf("Expected popcount of 14 to equal %v, got %v instead", 3, PopCount(14))
	}

	if PopCount(255) != 8 {
		t.Errorf("Expected popcount of 255 to equal %v, got %v instead", 8, PopCount(255))
	}
}

func TestPopCountIterative(t *testing.T) {
	if PopCountIterative(14) != 3 {
		t.Errorf("Expectedd PopCountIterative(14) to equal %v, got %v instead", 3, PopCountIterative(14))
	}
	
	if PopCountIterative(255) != 8 {
		t.Errorf("Expectedd PopCountIterative(255) to equal %v, got %v instead", 8, PopCountIterative(255))
	}
}

func TestPopCountRec(t *testing.T) {
	if PopCountRec(14) != 3 {
		t.Errorf("Expectedd PopCountRec(14) to equal %v, got %v instead", 3, PopCountRec(14))
	}
	
	if PopCountRec(255) != 8 {
		t.Errorf("Expectedd PopCountRec(255) to equal %v, got %v instead", 8, PopCountRec(255))
	}
}
