package nodes_test

import (
	"context"
	"testing"

	gen "christiangeorgelucas/filetype-tools/gen"
	"christiangeorgelucas/filetype-tools/nodes"
)

func TestIsDocument(t *testing.T) {
	ctx := context.Background()
	ax := newTestContext(t)

	for _, c := range []struct {
		name string
		data []byte
		want bool
	}{
		{"ole-doc-is-document", oleDocSample, true},
		{"ole-ambiguous-is-still-document", oleAmbiguousSample, true},
		{"pdf-is-not-document", pdfSample, false}, // PDF is categorized under Archive, not Document
		{"png-is-not-document", pngSample, false},
		{"garbage-is-not-document", garbageSample, false},
		{"empty-is-not-document", nil, false},
	} {
		t.Run(c.name, func(t *testing.T) {
			got, err := nodes.IsDocument(ctx, ax, &gen.FileSample{Data: c.data})
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got.GetMatches() != c.want {
				t.Errorf("matches = %v, want %v", got.GetMatches(), c.want)
			}
		})
	}
}
