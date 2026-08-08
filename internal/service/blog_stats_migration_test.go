package service

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// A stats file written by an earlier version holds raw visitor addresses.
// Booting must hash them and leave no plaintext behind.
func TestBlogTrackerMigratesLegacyPlaintextAddresses(t *testing.T) {
	path := filepath.Join(t.TempDir(), "blog_stats.json")
	legacy := `{"totalViews":7,"uniqueIPs":["203.0.113.9","198.51.100.4"],"flags":{"PT":2}}`
	require.NoError(t, os.WriteFile(path, []byte(legacy), 0o644))

	t.Setenv("BLOG_STATS_PATH", path)
	t.Setenv("VISIT_SALT", "fixed-test-salt")

	tracker := NewBlogTracker()

	raw, err := os.ReadFile(path)
	require.NoError(t, err)

	assert.NotContains(t, string(raw), "203.0.113.9", "plaintext address survived the migration")
	assert.NotContains(t, string(raw), "198.51.100.4")
	assert.NotContains(t, string(raw), "uniqueIPs", "the legacy field must not be rewritten")

	var file blogStatsFile
	require.NoError(t, json.Unmarshal(raw, &file))
	assert.Len(t, file.Hashes, 2, "both visitors are still counted, just not identifiable")
	assert.Equal(t, 7, file.TotalViews)
	assert.Equal(t, 2, file.Flags["PT"])

	// The migrated visitors must still be recognised as returning, not new.
	assert.True(t, tracker.unique[hashWithSalt(tracker.salt, "203.0.113.9")])
}
