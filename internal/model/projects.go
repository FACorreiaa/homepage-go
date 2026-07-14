package model

type ProjectItem struct {
	Slug         string
	Title        string
	RoleTag      string
	Description  string
	Outcome      string
	Tags         []string
	Category     string
	DisplayGroup string
	GithubLink   string
	LiveLink     string
	AppStoreLink string
	HasLiveLink  bool
	Featured     bool
	Icon         string
	LogoAsset    string
	Status       string

	// Proof slots: templates render these only when populated.
	Metrics         []ProjectMetric
	Testimonial     Testimonial
	ScreenshotAsset string
	ScreenshotAlt   string
}

type ProjectMetric struct {
	Value string // e.g. "38ms"
	Label string // e.g. "p95 API latency"
}

type Testimonial struct {
	Quote  string
	Author string
	Role   string // e.g. "Founder, Fandemic"
}

type TrustItem struct {
	Icon  string
	Value string
	Label string
}

type DetailFeature struct {
	Title string
	Body  string
}

type DetailLink struct {
	Label string
	URL   string
}

type ProjectDetailData struct {
	Project     ProjectItem
	Tagline     string
	LongDesc    []string
	Features    []DetailFeature
	TechStack   []string
	SocialLinks []DetailLink
	BackendNote string
	BannerAsset string
}

type LuminaStep struct {
	Num   string
	Title string
	Body  string
}
