package service

import (
	"context"
	"io/fs"
	"strings"

	"github.com/adrg/frontmatter"
	"myapp/internal/db"
)

type blogFrontmatter struct {
	Title    string `yaml:"title" toml:"title"`
	Summary  string `yaml:"summary" toml:"summary"`
	Category string `yaml:"category" toml:"category"`
	Date     string `yaml:"date" toml:"date"`
}

func SeedBlogPosts(ctx context.Context, q *db.Queries, blogFS fs.FS) error {
	entries, err := fs.ReadDir(blogFS, "blog")
	if err != nil {
		return err
	}
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".md") {
			continue
		}
		data, err := fs.ReadFile(blogFS, "blog/"+e.Name())
		if err != nil {
			continue
		}
		var fm blogFrontmatter
		body, err := frontmatter.Parse(strings.NewReader(string(data)), &fm)
		if err != nil {
			continue
		}
		slug := strings.TrimSuffix(e.Name(), ".md")
		if err := q.UpsertBlogPost(ctx, db.UpsertBlogPostParams{
			Slug:     slug,
			Title:    fm.Title,
			Summary:  fm.Summary,
			Category: fm.Category,
			Date:     fm.Date,
			Body:     string(body),
		}); err != nil {
			return err
		}
	}
	return nil
}
