package main

import (
	"slices"
	"testing"
)

func TestInitialModel(t *testing.T) {
	got := InitialModel()

	t.Run("testing initial model", func(t *testing.T) {
		want := []string{"Buy carrots", "Buy celery", "Buy kohlrabi"}
		if !slices.Equal(got.choices, want) {
			t.Errorf("got %v, but wanted %v", got.choices, want)
		}
	})
}
