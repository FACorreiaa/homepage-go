-- name: ListBlogPosts :many
SELECT slug, title, summary, category, date, created_at, updated_at
FROM blog_posts
ORDER BY date DESC;

-- name: GetBlogPostBySlug :one
SELECT * FROM blog_posts WHERE slug = ? LIMIT 1;

-- name: UpsertBlogPost :exec
INSERT INTO blog_posts (slug, title, summary, category, date, body, updated_at)
VALUES (?, ?, ?, ?, ?, ?, datetime('now'))
ON CONFLICT(slug) DO UPDATE SET
    title      = excluded.title,
    summary    = excluded.summary,
    category   = excluded.category,
    date       = excluded.date,
    body       = excluded.body,
    updated_at = excluded.updated_at;

-- name: CreateBlogPost :exec
INSERT INTO blog_posts (slug, title, summary, category, date, body)
VALUES (?, ?, ?, ?, ?, ?);

-- name: UpdateBlogPost :exec
UPDATE blog_posts
SET title = ?, summary = ?, category = ?, date = ?, body = ?, updated_at = datetime('now')
WHERE slug = ?;
