package nodes

import (
	"context"

	"github.com/h2non/filetype"

	"christiangeorgelucas/filetype-tools/axiom"
	gen "christiangeorgelucas/filetype-tools/gen"
)

// IsArchive checks whether a byte sample's magic bytes match any recognized
// archive, container, or executable format (ZIP, TAR, GZIP, RAR, 7Z, XZ,
// ZSTD, PDF, RTF, SQLite, ELF, Mach-O, EXE, DEB, RPM, ISO 9660, and more).
func IsArchive(ctx context.Context, ax axiom.Context, input *gen.FileSample) (*gen.TypeMatch, error) {
	buf := truncate(input.GetData())
	return &gen.TypeMatch{Matches: filetype.IsArchive(buf)}, nil
}
