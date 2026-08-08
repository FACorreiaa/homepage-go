package service

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestMain points the salt at a temp directory for the whole package. Without
// it, any tracker built without DATABASE_PATH writes its salt next to the
// working directory — which means a test run drops a file into the repo.
func TestMain(m *testing.M) {
	dir, err := os.MkdirTemp("", "visit-salt-test")
	if err != nil {
		panic(err)
	}
	if err := os.Setenv("VISIT_SALT_PATH", filepath.Join(dir, "visit-salt")); err != nil {
		panic(err)
	}

	code := m.Run()

	_ = os.RemoveAll(dir)
	os.Exit(code)
}

func TestVisitSaltPersistsAcrossRestarts(t *testing.T) {
	dir := t.TempDir()
	t.Setenv("VISIT_SALT_PATH", filepath.Join(dir, "visit-salt"))
	t.Setenv("VISIT_SALT", "")

	first := resolveVisitSalt()
	second := resolveVisitSalt()

	assert.NotEmpty(t, first)
	assert.Equal(t, first, second, "a restart must reuse the stored salt, or unique counts inflate")
}

func TestVisitSaltFileIsOwnerOnly(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "visit-salt")
	t.Setenv("VISIT_SALT_PATH", path)
	t.Setenv("VISIT_SALT", "")

	resolveVisitSalt()

	info, err := os.Stat(path)
	require.NoError(t, err)
	assert.Equal(t, os.FileMode(0o600), info.Mode().Perm(), "the salt must not be world-readable")
}

func TestExplicitVisitSaltWins(t *testing.T) {
	t.Setenv("VISIT_SALT_PATH", filepath.Join(t.TempDir(), "visit-salt"))
	t.Setenv("VISIT_SALT", "from-the-environment")

	assert.Equal(t, []byte("from-the-environment"), resolveVisitSalt())
}

func TestVisitSaltSurvivesUnwritableLocation(t *testing.T) {
	t.Setenv("VISIT_SALT_PATH", "/proc/definitely-not-writable/visit-salt")
	t.Setenv("VISIT_SALT", "")

	assert.NotEmpty(t, resolveVisitSalt(), "an unwritable volume must not stop the process")
}

func TestVisitSaltPathFollowsDatabaseLocation(t *testing.T) {
	t.Setenv("VISIT_SALT_PATH", "")
	t.Setenv("DATABASE_PATH", "/var/lib/facorreia-site/studio.sqlite")

	assert.Equal(t, "/var/lib/facorreia-site/visit-salt", resolveVisitSaltPath())
}
