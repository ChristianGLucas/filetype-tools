package nodes

import (
	"context"

	"github.com/h2non/filetype"

	"christiangeorgelucas/filetype-tools/axiom"
	gen "christiangeorgelucas/filetype-tools/gen"
)

// IsImage checks whether a byte sample's magic bytes match any recognized
// image format (JPEG, JPEG2000, PNG, GIF, WEBP, TIFF, BMP, HEIF, PSD, ICO,
// Canon CR2 raw, DWG, and more).
func IsImage(ctx context.Context, ax axiom.Context, input *gen.FileSample) (*gen.TypeMatch, error) {
	buf := truncate(input.GetData())
	return &gen.TypeMatch{Matches: filetype.IsImage(buf)}, nil
}
