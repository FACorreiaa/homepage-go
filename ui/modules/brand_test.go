package modules

import (
	"context"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func renderBrandMark(t *testing.T, invert bool, size string) string {
	t.Helper()
	var sb strings.Builder
	require.NoError(t, BrandMark(invert, size).Render(context.Background(), &sb))
	return sb.String()
}

func TestBrandMarkRendersFC2SGrid(t *testing.T) {
	body := renderBrandMark(t, false, "")
	assert.Contains(t, body, `class="brand-mark"`)
	for _, letter := range []string{"F", "C", "2", "S"} {
		assert.Contains(t, body, ">"+letter+"</span>")
	}
	assert.NotContains(t, body, "brand-mark--invert")
	assert.NotContains(t, body, "brand-mark--sm")
	assert.NotContains(t, body, "brand-mark--lg")
}

func TestBrandMarkSizeAndInvertModifiers(t *testing.T) {
	inverted := renderBrandMark(t, true, "lg")
	assert.Contains(t, inverted, "brand-mark--invert")
	assert.Contains(t, inverted, "brand-mark--lg")

	small := renderBrandMark(t, false, "sm")
	assert.Contains(t, small, "brand-mark--sm")
	assert.NotContains(t, small, "brand-mark--invert")
}
