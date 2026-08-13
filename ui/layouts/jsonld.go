package layouts

import (
	"encoding/json"
	"strings"
	"time"
)

// JSON-LD helpers. Always built via encoding/json so free-text titles never break out of the script tag.

const personID = siteBaseURL + "/#person"

// BlogPostJSONLD returns BlogPosting + BreadcrumbList for a blog detail page.
func BlogPostJSONLD(slug, title, summary, category, date, updatedAt, imageURL string) string {
	pageURL := siteBaseURL + "/blog/" + slug
	published := parseDateISO(date)
	modified := parseDateISO(updatedAt)
	if modified == "" {
		modified = published
	}
	if published == "" {
		published = modified
	}
	if imageURL == "" {
		imageURL = siteBaseURL + "/assets/static/icon-512.png"
	}

	graph := []map[string]any{
		{
			"@type":               "BlogPosting",
			"@id":                 pageURL + "#article",
			"headline":            title,
			"description":         summary,
			"datePublished":       published,
			"dateModified":        modified,
			"inLanguage":          "en",
			"isAccessibleForFree": true,
			"mainEntityOfPage": map[string]any{
				"@type": "WebPage",
				"@id":   pageURL,
			},
			"url":   pageURL,
			"image": []string{imageURL},
			"author": map[string]any{
				"@id": personID,
			},
			"publisher": map[string]any{
				"@id": personID,
			},
		},
		breadcrumbList([]breadcrumbItem{
			{Name: "Home", URL: siteBaseURL + "/"},
			{Name: "Blog", URL: siteBaseURL + "/blog"},
			{Name: title, URL: pageURL},
		}),
	}
	if category != "" {
		graph[0]["articleSection"] = category
	}
	return scriptTag(map[string]any{
		"@context": "https://schema.org",
		"@graph":   graph,
	})
}

// ProjectJSONLD returns a CreativeWork page entity + BreadcrumbList for a project detail page.
// CreativeWork avoids SoftwareApplication's required Offer shape for portfolio pieces.
func ProjectJSONLD(slug, title, description, imageURL string) string {
	pageURL := siteBaseURL + "/projects/" + slug
	if imageURL == "" {
		imageURL = siteBaseURL + "/assets/static/icon-512.png"
	} else if strings.HasPrefix(imageURL, "/") {
		imageURL = siteBaseURL + imageURL
	}

	graph := []map[string]any{
		{
			"@type":       "CreativeWork",
			"@id":         pageURL + "#work",
			"name":        title,
			"description": description,
			"url":         pageURL,
			"image":       imageURL,
			"inLanguage":  "en",
			"author": map[string]any{
				"@id": personID,
			},
			"creator": map[string]any{
				"@id": personID,
			},
			"mainEntityOfPage": map[string]any{
				"@type": "WebPage",
				"@id":   pageURL,
			},
		},
		breadcrumbList([]breadcrumbItem{
			{Name: "Home", URL: siteBaseURL + "/"},
			{Name: "Projects", URL: siteBaseURL + "/projects"},
			{Name: title, URL: pageURL},
		}),
	}
	return scriptTag(map[string]any{
		"@context": "https://schema.org",
		"@graph":   graph,
	})
}

type breadcrumbItem struct {
	Name string
	URL  string
}

func breadcrumbList(items []breadcrumbItem) map[string]any {
	elements := make([]map[string]any, 0, len(items))
	for i, it := range items {
		elements = append(elements, map[string]any{
			"@type":    "ListItem",
			"position": i + 1,
			"name":     it.Name,
			"item":     it.URL,
		})
	}
	return map[string]any{
		"@type":           "BreadcrumbList",
		"itemListElement": elements,
	}
}

func scriptTag(payload any) string {
	b, err := json.Marshal(payload)
	if err != nil {
		return ""
	}
	return `<script type="application/ld+json">` + string(b) + `</script>`
}

// ParseBlogLastMod returns YYYY-MM-DD for sitemap lastmod, preferring updatedAt then date.
func ParseBlogLastMod(updatedAt, date string) string {
	if s := parseDateISO(updatedAt); s != "" {
		if len(s) >= 10 {
			return s[:10]
		}
		return s
	}
	if s := parseDateISO(date); s != "" {
		if len(s) >= 10 {
			return s[:10]
		}
		return s
	}
	return ""
}

func parseDateISO(raw string) string {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return ""
	}
	// Already date-only.
	if len(raw) == 10 && raw[4] == '-' && raw[7] == '-' {
		if _, err := time.Parse("2006-01-02", raw); err == nil {
			return raw
		}
	}
	layouts := []string{
		time.RFC3339,
		"2006-01-02 15:04:05",
		"2006-01-02T15:04:05Z07:00",
		"2006-01-02T15:04:05",
		"2006-01-02",
		"January 2, 2006",
		"Jan 2, 2006",
		"2 January 2006",
		"02 Jan 2006",
	}
	for _, layout := range layouts {
		if t, err := time.Parse(layout, raw); err == nil {
			return t.UTC().Format("2006-01-02")
		}
	}
	// SQLite datetime sometimes has no timezone; try local wall clock as UTC date.
	if t, err := time.ParseInLocation("2006-01-02 15:04:05", raw, time.UTC); err == nil {
		return t.Format("2006-01-02")
	}
	return ""
}
