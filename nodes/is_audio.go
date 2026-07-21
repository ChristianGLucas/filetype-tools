package nodes

import (
	"context"

	"github.com/h2non/filetype"

	"christiangeorgelucas/filetype-tools/axiom"
	gen "christiangeorgelucas/filetype-tools/gen"
)

// IsAudio checks whether a byte sample's magic bytes match any recognized
// audio format (MP3, WAV, FLAC, OGG, MIDI, AAC, AIFF, AMR, and more).
func IsAudio(ctx context.Context, ax axiom.Context, input *gen.FileSample) (*gen.TypeMatch, error) {
	buf := truncate(input.GetData())
	return &gen.TypeMatch{Matches: filetype.IsAudio(buf)}, nil
}
