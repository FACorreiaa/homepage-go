package assets

import (
	"image"
	_ "image/png"
	"testing"
)

func TestOGImageIsShareCardSize(t *testing.T) {
	f, err := Assets.Open("static/og.png")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if closeErr := f.Close(); closeErr != nil {
			t.Errorf("close: %v", closeErr)
		}
	})
	cfg, format, err := image.DecodeConfig(f)
	if err != nil {
		t.Fatal(err)
	}
	if format != "png" {
		t.Fatalf("format=%s", format)
	}
	if cfg.Width != 1200 || cfg.Height != 630 {
		t.Fatalf("og.png is %dx%d, want 1200x630", cfg.Width, cfg.Height)
	}
}
