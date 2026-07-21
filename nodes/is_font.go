package nodes

import (
	"context"

	"github.com/h2non/filetype"

	"christiangeorgelucas/filetype-tools/axiom"
	gen "christiangeorgelucas/filetype-tools/gen"
)

// IsFont checks whether a byte sample's magic bytes match any recognized
// font format (TTF, OTF, WOFF, WOFF2, EOT).
func IsFont(ctx context.Context, ax axiom.Context, input *gen.FileSample) (*gen.TypeMatch, error) {
	buf := truncate(input.GetData())
	return &gen.TypeMatch{Matches: filetype.IsFont(buf)}, nil
}
