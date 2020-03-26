package main

import "testing"

func TestOKBinary_tree(t *testing.T) {
	var tests = []struct {
		guess int
		list  []int
	}{
		{9, []int{1, 2, 9, 20, 31, 45, 63, 70, 100}},
		{70, []int{1, 2, 9, 20, 31, 45, 63, 70, 100}},
		{100, []int{1, 2, 9, 20, 31, 45, 63, 70, 100}},
		{31, []int{1, 2, 9, 20, 31, 45, 63, 70, 100}},
		{63, []int{1, 2, 9, 20, 31, 45, 63, 70, 100}},
	}
	for _, test := range tests {
		if !binarySearch(test.guess, test.list) {
			t.Errorf("binarySearch(%d, %v) = false", test.guess, test.list)
		}
	}
}

func TestFailBinary_tree(t *testing.T) {
	var tests = []struct {
		guess int
		list  []int
	}{
		{5, []int{1, 2, 9, 20, 31, 45, 63, 70, 100}},
		{0, []int{1, 2, 9, 20, 31, 45, 63, 70, 100}},
		{140, []int{1, 2, 9, 20, 31, 45, 63, 70, 100}},
		{34, []int{1, 2, 9, 20, 31, 45, 63, 70, 100}},
		{33, []int{1, 2, 9, 20, 31, 45, 63, 70, 100}},
	}
	for _, test := range tests {
		if binarySearch(test.guess, test.list) {
			t.Errorf("binarySearch(%d, %v) = true", test.guess, test.list)
		}
	}
}
