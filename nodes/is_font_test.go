package nodes_test

import (
	"context"
	"testing"

	gen "christiangeorgelucas/filetype-tools/gen"
	"christiangeorgelucas/filetype-tools/nodes"
)

func TestIsFont(t *testing.T) {
	ctx := context.Background()
	ax := newTestContext(t)

	for _, c := range []struct {
		name string
		data []byte
		want bool
	}{
		{"ttf-is-font", ttfSample, true},
		{"png-is-not-font", pngSample, false},
		{"garbage-is-not-font", garbageSample, false},
		{"empty-is-not-font", nil, false},
	} {
		t.Run(c.name, func(t *testing.T) {
			got, err := nodes.IsFont(ctx, ax, &gen.FileSample{Data: c.data})
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got.GetMatches() != c.want {
				t.Errorf("matches = %v, want %v", got.GetMatches(), c.want)
			}
		})
	}
}
