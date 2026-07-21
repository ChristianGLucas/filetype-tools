package nodes

import (
	"context"

	"github.com/h2non/filetype"

	"christiangeorgelucas/filetype-tools/axiom"
	gen "christiangeorgelucas/filetype-tools/gen"
)

// MatchesType tests a byte sample's magic bytes against ONE specific
// expected type, given as an extension (e.g. "png") or a MIME type (e.g.
// "image/png"). If `extension` is set it takes precedence over `mime_type`;
// if neither is set, the result is always false (not an error).
func MatchesType(ctx context.Context, ax axiom.Context, input *gen.MatchRequest) (*gen.TypeMatch, error) {
	buf := truncate(input.GetData())
	ext := normalizeIdentifier(input.GetExtension())
	mime := normalizeIdentifier(input.GetMimeType())

	if ext != "" {
		return &gen.TypeMatch{Matches: filetype.Is(buf, ext)}, nil
	}
	if mime != "" {
		return &gen.TypeMatch{Matches: filetype.IsMIME(buf, mime)}, nil
	}
	return &gen.TypeMatch{Matches: false}, nil
}
