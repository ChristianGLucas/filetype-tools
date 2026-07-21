package nodes_test

import (
	"context"
	"testing"

	gen "christiangeorgelucas/filetype-tools/gen"
	"christiangeorgelucas/filetype-tools/nodes"
)

func TestIsImage(t *testing.T) {
	ctx := context.Background()
	ax := newTestContext(t)

	for _, c := range []struct {
		name string
		data []byte
		want bool
	}{
		{"png-is-image", pngSample, true},
		{"jpeg-is-image", jpegSample, true},
		{"wav-is-not-image", wavSample, false},
		{"garbage-is-not-image", garbageSample, false},
		{"empty-is-not-image", nil, false},
	} {
		t.Run(c.name, func(t *testing.T) {
			got, err := nodes.IsImage(ctx, ax, &gen.FileSample{Data: c.data})
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got.GetMatches() != c.want {
				t.Errorf("matches = %v, want %v", got.GetMatches(), c.want)
			}
		})
	}
}
