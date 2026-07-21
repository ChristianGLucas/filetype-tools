package nodes_test

import (
	"context"
	"testing"

	gen "christiangeorgelucas/filetype-tools/gen"
	"christiangeorgelucas/filetype-tools/nodes"
)

func TestIsArchive(t *testing.T) {
	ctx := context.Background()
	ax := newTestContext(t)

	for _, c := range []struct {
		name string
		data []byte
		want bool
	}{
		{"zip-is-archive", zipSample, true},
		{"pdf-is-archive", pdfSample, true}, // h2non/filetype categorizes PDF under Archive, not Document
		{"png-is-not-archive", pngSample, false},
		{"garbage-is-not-archive", garbageSample, false},
		{"empty-is-not-archive", nil, false},
	} {
		t.Run(c.name, func(t *testing.T) {
			got, err := nodes.IsArchive(ctx, ax, &gen.FileSample{Data: c.data})
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got.GetMatches() != c.want {
				t.Errorf("matches = %v, want %v", got.GetMatches(), c.want)
			}
		})
	}
}
