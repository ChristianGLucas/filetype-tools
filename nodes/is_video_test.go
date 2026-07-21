package nodes_test

import (
	"context"
	"testing"

	gen "christiangeorgelucas/filetype-tools/gen"
	"christiangeorgelucas/filetype-tools/nodes"
)

func TestIsVideo(t *testing.T) {
	ctx := context.Background()
	ax := newTestContext(t)

	for _, c := range []struct {
		name string
		data []byte
		want bool
	}{
		{"mp4-is-video", mp4Sample, true},
		{"png-is-not-video", pngSample, false},
		{"garbage-is-not-video", garbageSample, false},
		{"empty-is-not-video", nil, false},
	} {
		t.Run(c.name, func(t *testing.T) {
			got, err := nodes.IsVideo(ctx, ax, &gen.FileSample{Data: c.data})
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got.GetMatches() != c.want {
				t.Errorf("matches = %v, want %v", got.GetMatches(), c.want)
			}
		})
	}
}
