package nodes

import (
	"context"

	"github.com/h2non/filetype"

	"christiangeorgelucas/filetype-tools/axiom"
	gen "christiangeorgelucas/filetype-tools/gen"
)

// IsVideo checks whether a byte sample's magic bytes match any recognized
// video format (MP4, MKV, WEBM, AVI, MOV/QuickTime, FLV, WMV, 3GP, and more).
func IsVideo(ctx context.Context, ax axiom.Context, input *gen.FileSample) (*gen.TypeMatch, error) {
	buf := truncate(input.GetData())
	return &gen.TypeMatch{Matches: filetype.IsVideo(buf)}, nil
}
