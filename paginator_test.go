package paginator_test

import (
	"testing"

	"github.com/kaDDDar/paginator"
)

func TestPaginator(t *testing.T) {
	tables := []struct {
		CurrentPage int
		TotalPages  int
		Boundaries  int
		Around      int

		SkipChar string
		Result   string
	}{
		{5, 10, 2, 0, "...", "1 2 ... 5 ... 9 10"},
		{5, 10, 1, 1, "...", "1 ... 4 5 6 ... 10"},
		{2, 10, 1, 1, "...", "1 2 3 ... 10"},
		{2, 10, 1, 2, "...", "1 2 3 4 ... 10"},
		{10, 10, 1, 1, "...", "1 ... 9 10"},
		{10, 10, 1, 2, "...", "1 ... 8 9 10"},
		{1, 10, 0, 0, "...", "1 ..."},
		{1, 10, 0, 2, "...", "1 2 3 ..."},
		{5, 10, 0, 2, "...", "... 3 4 5 6 7 ..."},
		{10, 10, 0, 0, "...", "... 10"},
		{10, 10, 0, 2, "...", "... 8 9 10"},
		{9, 10, 0, 2, "...", "... 7 8 9 10"},
		{2, 10, 0, 2, "...", "1 2 3 4 ..."},
		{2, 3, 0, 0, "...", "... 2 ..."},
	}

	for _, table := range tables {
		paginator, _ := paginator.Create(table.CurrentPage, table.TotalPages, table.Boundaries, table.Around, table.SkipChar)
		result := paginator.Generate()
		if result != table.Result {
			t.Errorf("Got: %s, Need: %s", result, table.Result)
		}
	}
}
