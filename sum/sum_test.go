package sum

import "testing"

func TestNewWithArgs(t *testing.T) {
	s := New("a","b")
	// insert values
	s.Insert("a")
	s.Insert("a")
	s.Insert("b")
	// ignore c
	s.Insert("c")
	s.Insert("c")

	r := s.Query()
	if len(r) != 2 {
		t.Errorf("Query result length should be 2, but was %d", len(r))
	}

	// verify 1st element
	if r[0].Value != "a" {
		t.Errorf("Result 1 value should be a, but was %s", r[0].Value)
	}

	if r[0].Count != 2 {
		t.Errorf("Result 1 count should be 2, but was %d", r[0].Count)
	}

	// verify 2nd element
	if r[1].Value != "b" {
		t.Errorf("Result 2 value should be b, but was %s", r[1].Value)
	}

	if r[1].Count != 1 {
		t.Errorf("Result 2 count should be 1, but was %d", r[1].Count)
	}
}

func TestNewWithoutArgs(t *testing.T) {
	s := New()
	// insert all values
	s.Insert("a")
	s.Insert("a")
	s.Insert("a")
	s.Insert("b")
	s.Insert("b")
	s.Insert("c")

	r := s.Query()

	// verify 1st element
	if r[0].Value != "a" {
		t.Errorf("Result 1 value should be a, but was %s", r[0].Value)
	}

	if r[0].Count != 3 {
		t.Errorf("Result 1 count should be 3, but was %d", r[0].Count)
	}

	// verify 2st element
	if r[1].Value != "b" {
		t.Errorf("Result 2 value should be b, but was %s", r[1].Value)
	}

	if r[1].Count != 2 {
		t.Errorf("Result 2 count should be 2, but was %d", r[1].Count)
	}

	// verify 3rd element
	if r[2].Value != "c" {
		t.Errorf("Result 3 value should be c, but was %s", r[2].Value)
	}

	if r[2].Count != 1 {
		t.Errorf("Result 3 count should be 1, but was %d", r[2].Count)
	}
}
