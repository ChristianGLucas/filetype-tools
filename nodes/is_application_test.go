package nodes_test

import (
	"context"
	"testing"

	gen "christiangeorgelucas/filetype-tools/gen"
	"christiangeorgelucas/filetype-tools/nodes"
)

func TestIsApplication(t *testing.T) {
	ctx := context.Background()
	ax := newTestContext(t)

	for _, c := range []struct {
		name string
		data []byte
		want bool
	}{
		{"wasm-is-application", wasmSample, true},
		{"png-is-not-application", pngSample, false},
		{"zip-is-not-application", zipSample, false}, // ZIP is categorized under Archive
		{"garbage-is-not-application", garbageSample, false},
		{"empty-is-not-application", nil, false},
	} {
		t.Run(c.name, func(t *testing.T) {
			got, err := nodes.IsApplication(ctx, ax, &gen.FileSample{Data: c.data})
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got.GetMatches() != c.want {
				t.Errorf("matches = %v, want %v", got.GetMatches(), c.want)
			}
		})
	}
}
