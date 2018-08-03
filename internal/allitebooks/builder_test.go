package allitebooks

import "testing"

func TestBuilderNoExtra(t *testing.T) {
	borer := New().
		Build()
	if borer.GetStartPage() != 0 || borer.GetStartURL() != "" {
		t.Error("Builder did not produce instance with correct values")
	}
}

func TestBuilderExtraStartPage(t *testing.T) {
	borer := New().
		WithStartPage(30).
		Build()
	if borer.GetStartPage() != 30 || borer.GetStartURL() != "" {
		t.Error("Builder did not produce instance with correct values")
	}
}

func TestBuilderExtraStartURL(t *testing.T) {
	borer := New().
		WithStartURL("http://www.google.com").
		Build()
	if borer.GetStartPage() != 0 || borer.GetStartURL() != "http://www.google.com" {
		t.Error("Builder did not produce instance with correct values")
	}
}

func TestBuilderExtraStartPageAndStartURL(t *testing.T) {
	borer := New().
		WithStartPage(30).
		WithStartURL("http://www.google.com").
		Build()
	if borer.GetStartPage() != 30 || borer.GetStartURL() != "http://www.google.com" {
		t.Error("Builder did not produce instance with correct values")
	}
}
