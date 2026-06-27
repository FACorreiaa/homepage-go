package db_test

import (
	"context"
	"testing"

	"myapp/internal/db"

	"github.com/stretchr/testify/require"
)

func TestOpenAndSchema(t *testing.T) {
	conn, err := db.Open(":memory:")
	require.NoError(t, err)
	defer conn.Close()

	tables := []string{"client_users", "proposal_leads", "blog_posts"}
	for _, tbl := range tables {
		var name string
		err := conn.QueryRowContext(context.Background(),
			"SELECT name FROM sqlite_master WHERE type='table' AND name=?", tbl,
		).Scan(&name)
		require.NoError(t, err, "table %s missing", tbl)
	}
}

func TestSeedAdminUser(t *testing.T) {
	conn, err := db.Open(":memory:")
	require.NoError(t, err)
	defer conn.Close()

	q := db.New(conn)
	ctx := context.Background()

	require.NoError(t, db.SeedAdminUser(ctx, q))
	require.NoError(t, db.SeedAdminUser(ctx, q)) // idempotent

	user, err := q.GetUserByEmail(ctx, "fernandocorreia316@gmail.com")
	require.NoError(t, err)
	require.Equal(t, "Fernando Correia", user.Name)
}
