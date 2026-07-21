package nodes

import (
	"context"

	"github.com/h2non/filetype"

	"christiangeorgelucas/filetype-tools/axiom"
	gen "christiangeorgelucas/filetype-tools/gen"
)

// IsDocument checks whether a byte sample's magic bytes match any recognized
// office document format: legacy binary DOC/XLS/PPT (OLE2 compound file), or
// OOXML DOCX/XLSX/PPTX. A legacy DOC/XLS/PPT sample shorter than 514 bytes is
// still recognized as a document (all three share an identical short OLE2
// header), just not disambiguated to a specific one of the three.
func IsDocument(ctx context.Context, ax axiom.Context, input *gen.FileSample) (*gen.TypeMatch, error) {
	buf := truncate(input.GetData())
	return &gen.TypeMatch{Matches: filetype.IsDocument(buf)}, nil
}
