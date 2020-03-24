package main

import (
	"fmt"
	"testing"
)

func TestIsDir(t *testing.T) {
	testData := []struct {
		dir      string
		expected bool
	}{
		{"/root", true},
		{"/home/viz", true},
		{"/home/viz/.config", false},
		{"/home/viz/lib", true},
	}

	for _, d := range testData {
		t.Run(fmt.Sprintf("%s", d.dir), func(t *testing.T) {
			res := IsDir(d.dir)
			if res != d.expected {
				t.Errorf("Expected %t, got %t", d.expected, res)
			}
		})
	}
}

func TestIsManPage(t *testing.T) {
	testData := []struct {
		name, section string
		expected      bool
	}{
		{"st", "1", true},
		{"dmenu", "1", true},
		{"go", "1", false},
		{"st", "2", false},
	}

	for _, d := range testData {
		t.Run(fmt.Sprintf("%s.%s", d.name, d.section), func(t *testing.T) {
			res := IsManPage(d.name, d.section)
			if res != d.expected {
				t.Errorf("Expected %t, got %t", d.expected, res)
			}
		})
	}
}

func TestIsMan(t *testing.T) {
	testData := []struct {
		str      string
		expected bool
	}{
		{"st(1)", true},
		{"dmenu(asd", false},
		{"go", false},
		{"st(1p)", true},
	}

	for _, d := range testData {
		t.Run(fmt.Sprintf("%s", d.str), func(t *testing.T) {
			res := IsMan(d.str)
			if res != d.expected {
				t.Errorf("Expected %t, got %t", d.expected, res)
			}
		})
	}
}

func TestParseMan(t *testing.T) {
	testData := []struct {
		str      string
		expected bool
	}{
		{"st(1)", true},
		{"dmenu(asd", false},
		{"go", false},
		{"st(1p)", true},
	}

	for _, d := range testData {
		t.Run(fmt.Sprintf("%s", d.str), func(t *testing.T) {
			_, _, err := ParseMan(d.str)
			res := err == nil
			if res != d.expected {
				t.Errorf("Expected %t, got %t (error %s)", d.expected, res, err)
			}
		})
	}
}
