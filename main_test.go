package main

import "testing"

func TestAnalyzeRange(t *testing.T) {
	t.Run("start and end are the same number", func(t *testing.T) {
		start, end := 5, 5
		assertAnalyzeRangeReturns(t, start, end, 0, 1, 1)
	})

	t.Run("start is greater than end", func(t *testing.T) {
		start, end := 3, 2
		assertAnalyzeRangeReturns(t, start, end, 0, 0, 0)
	})

	t.Run("generally proper ranges", func(t *testing.T) {
		tests := []struct {
			name string
			start, end int
			wantEvens, wantOdds, wantTotal int
		}{
			{"0 through 10", 0, 10, 6, 5, 11},
			{"5 through 37", 5, 37, 16, 17, 33},
			{"55 through 234", 55, 234, 90, 90, 180},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				assertAnalyzeRangeReturns(
					t,
					test.start,
					test.end,
					test.wantEvens,
					test.wantOdds,
					test.wantTotal,
				)
			})
		}
	})
}

func assertAnalyzeRangeReturns(t *testing.T, start, end, wantEvens, wantOdds, wantTotal int) {
	t.Helper()
	gotEvens, gotOdds, gotTotal := analyzeRange(start, end)

	if wantEvens != gotEvens {
		t.Errorf("expected evens to equal %d, got %d", wantEvens, gotEvens)
	}

	if wantOdds != gotOdds {
		t.Errorf("expected odds to equal %d, got %d", wantOdds, gotOdds)
	}

	if wantTotal != gotTotal {
		t.Errorf("expected total to equal %d, got %d", wantTotal, gotTotal)
	}
}
