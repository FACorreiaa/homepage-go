package pages

import (
	"testing"

	"myapp/internal/service"

	"github.com/stretchr/testify/assert"
)

// sharePct feeds a width straight into markup, so the zero and overflow cases
// matter more than the ordinary one: a quiet day must not divide by zero and a
// bad aggregate must not push a bar past its row.
func TestSharePct(t *testing.T) {
	tests := []struct {
		name   string
		visits int
		max    int
		want   int
	}{
		{"full row", 40, 40, 100},
		{"quarter row", 10, 40, 25},
		{"empty list", 0, 0, 0},
		{"zero visits", 0, 40, 0},
		{"zero max is not a divide", 5, 0, 0},
		{"clamped above 100", 50, 40, 100},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, sharePct(tt.visits, tt.max))
		})
	}
}

func TestMaxVisits(t *testing.T) {
	assert.Equal(t, 0, maxCountryVisits(nil))
	assert.Equal(t, 0, maxPathVisits(nil))
	assert.Equal(t, 40, maxCountryVisits([]service.CountryCount{
		{Code: "DE", Visits: 10},
		{Code: "PT", Visits: 40},
	}))
	assert.Equal(t, 80, maxPathVisits([]service.PathCount{
		{Path: "/about", Visits: 20},
		{Path: "/projects", Visits: 80},
	}))
}
