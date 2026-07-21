package nodes

import (
	"context"

	"github.com/h2non/filetype"

	"christiangeorgelucas/filetype-tools/axiom"
	gen "christiangeorgelucas/filetype-tools/gen"
)

// IsApplication checks whether a byte sample's magic bytes match any
// recognized general application/binary format (WebAssembly WASM, Android
// DEX/ODEX bytecode) not already covered by another category.
func IsApplication(ctx context.Context, ax axiom.Context, input *gen.FileSample) (*gen.TypeMatch, error) {
	buf := truncate(input.GetData())
	return &gen.TypeMatch{Matches: filetype.IsApplication(buf)}, nil
}
