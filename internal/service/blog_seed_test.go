package service

import (
	"io/fs"
	"strings"
	"testing"

	"github.com/adrg/frontmatter"
	"github.com/stretchr/testify/require"

	"myapp/content"
)

func TestKhepriAndLociPostParses(t *testing.T) {
	data, err := fs.ReadFile(content.Blog, "blog/khepri-and-loci.md")
	require.NoError(t, err)
	var fm blogFrontmatter
	body, err := frontmatter.Parse(strings.NewReader(string(data)), &fm)
	require.NoError(t, err)
	require.Equal(t, "Khepri and Loci are live", fm.Title)
	require.NotEmpty(t, fm.Summary)
	require.Equal(t, "2026-09-07", fm.Date)
	require.Contains(t, string(body), "kheprios.com")
	require.Contains(t, string(body), "lociai.fyi")
}
