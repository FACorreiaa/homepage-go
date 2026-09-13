package model

import "testing"

func TestProjectItem_IOSLinkLabel(t *testing.T) {
	tests := []struct {
		name string
		url  string
		want string
	}{
		{name: "testflight join url", url: "https://testflight.apple.com/join/s1qKaww4", want: "TestFlight"},
		{name: "app store url", url: "https://apps.apple.com/pt/app/norviq/id6765849578?l=en-GB", want: "App Store"},
		{name: "empty url", url: "", want: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := ProjectItem{AppStoreLink: tt.url}
			if got := p.IOSLinkLabel(); got != tt.want {
				t.Fatalf("IOSLinkLabel() = %q, want %q", got, tt.want)
			}
		})
	}
}
