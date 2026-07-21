package nodes_test

import (
	"context"
	"testing"

	gen "christiangeorgelucas/filetype-tools/gen"
	"christiangeorgelucas/filetype-tools/nodes"
)

func TestMatchesType(t *testing.T) {
	ctx := context.Background()
	ax := newTestContext(t)

	cases := []struct {
		name string
		req  *gen.MatchRequest
		want bool
	}{
		{"png-matches-own-extension", &gen.MatchRequest{Data: pngSample, Extension: "png"}, true},
		{"png-matches-own-mime", &gen.MatchRequest{Data: pngSample, MimeType: "image/png"}, true},
		{"png-does-not-match-jpg", &gen.MatchRequest{Data: pngSample, Extension: "jpg"}, false},
		{"png-does-not-match-other-mime", &gen.MatchRequest{Data: pngSample, MimeType: "image/jpeg"}, false},
		{"case-insensitive-extension", &gen.MatchRequest{Data: pngSample, Extension: "PNG"}, true},
		{"case-insensitive-mime", &gen.MatchRequest{Data: pngSample, MimeType: "IMAGE/PNG"}, true},
		{"extension-takes-precedence-over-mime", &gen.MatchRequest{Data: pngSample, Extension: "png", MimeType: "image/jpeg"}, true},
		{"neither-set-is-false-not-error", &gen.MatchRequest{Data: pngSample}, false},
		{"empty-data-is-false", &gen.MatchRequest{Extension: "png"}, false},
		{"garbage-data-is-false", &gen.MatchRequest{Data: garbageSample, Extension: "png"}, false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, err := nodes.MatchesType(ctx, ax, c.req)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got.GetMatches() != c.want {
				t.Errorf("matches = %v, want %v", got.GetMatches(), c.want)
			}
		})
	}
}
