package nodes_test

import (
	"context"
	"testing"

	gen "christiangeorgelucas/filetype-tools/gen"
	"christiangeorgelucas/filetype-tools/nodes"
)

func TestIsAudio(t *testing.T) {
	ctx := context.Background()
	ax := newTestContext(t)

	for _, c := range []struct {
		name string
		data []byte
		want bool
	}{
		{"wav-is-audio", wavSample, true},
		{"mp4-is-not-audio", mp4Sample, false},
		{"garbage-is-not-audio", garbageSample, false},
		{"empty-is-not-audio", nil, false},
	} {
		t.Run(c.name, func(t *testing.T) {
			got, err := nodes.IsAudio(ctx, ax, &gen.FileSample{Data: c.data})
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got.GetMatches() != c.want {
				t.Errorf("matches = %v, want %v", got.GetMatches(), c.want)
			}
		})
	}
}
