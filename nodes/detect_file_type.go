package nodes

import (
	"context"

	"github.com/h2non/filetype"
	"github.com/h2non/filetype/types"

	"christiangeorgelucas/filetype-tools/axiom"
	gen "christiangeorgelucas/filetype-tools/gen"
)

// DetectFileType inspects the magic bytes at the start of a byte sample and
// identifies its canonical extension, MIME type, and broad category. Returns
// matched=false (not an error) for an empty sample or bytes that do not
// match any of the 73 signatures this package recognizes.
func DetectFileType(ctx context.Context, ax axiom.Context, input *gen.FileSample) (*gen.FileTypeInfo, error) {
	buf := truncate(input.GetData())

	kind, _ := filetype.Match(buf)
	if kind == types.Unknown {
		return &gen.FileTypeInfo{Matched: false}, nil
	}

	return &gen.FileTypeInfo{
		Matched: true,
		Type:    descriptorOf(kind),
	}, nil
}
