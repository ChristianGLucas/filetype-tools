package nodes

import (
	"context"

	"christiangeorgelucas/filetype-tools/axiom"
	gen "christiangeorgelucas/filetype-tools/gen"
)

// ListSupportedTypes enumerates every file type this package can
// recognize (extension, MIME type, category), optionally filtered to one
// category ("image", "video", "audio", "font", "archive", "document",
// "application"). An unrecognized or empty category returns the full
// catalog. Results are sorted by extension for a stable, reproducible order.
func ListSupportedTypes(ctx context.Context, ax axiom.Context, input *gen.SupportedTypesQuery) (*gen.SupportedTypesList, error) {
	filter := normalizeIdentifier(input.GetCategory())

	all := allTypes()
	out := make([]*gen.TypeDescriptor, 0, len(all))
	for _, kind := range all {
		d := descriptorOf(kind)
		if filter != "" && d.Category != filter {
			continue
		}
		out = append(out, d)
	}

	return &gen.SupportedTypesList{Types: out, Count: int32(len(out))}, nil
}
