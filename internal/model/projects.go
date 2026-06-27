package model

type ProjectItem struct {
	Slug        string
	Title       string
	RoleTag     string
	Description string
	Tags        []string
	Category    string
	GithubLink  string
	LiveLink    string
	HasLiveLink bool
	Featured    bool
	Icon        string
	IconURL     string
	LogoAsset   string
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
