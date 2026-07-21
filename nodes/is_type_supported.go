package nodes

import (
	"context"

	"github.com/h2non/filetype"
	"github.com/h2non/filetype/types"

	"christiangeorgelucas/filetype-tools/axiom"
	gen "christiangeorgelucas/filetype-tools/gen"
)

// IsTypeSupported looks up whether an extension or MIME type string is one
// this package recognizes — a pure catalog lookup by name, no file bytes
// involved. If `extension` is set it takes precedence over `mime_type`; if
// neither is set, the result is matched=false.
func IsTypeSupported(ctx context.Context, ax axiom.Context, input *gen.TypeIdentifier) (*gen.FileTypeInfo, error) {
	ext := normalizeIdentifier(input.GetExtension())
	mime := normalizeIdentifier(input.GetMimeType())

	if ext != "" {
		if !filetype.IsSupported(ext) {
			return &gen.FileTypeInfo{Matched: false}, nil
		}
		return &gen.FileTypeInfo{Matched: true, Type: descriptorOf(types.Get(ext))}, nil
	}
	if mime != "" {
		if !filetype.IsMIMESupported(mime) {
			return &gen.FileTypeInfo{Matched: false}, nil
		}
		return &gen.FileTypeInfo{Matched: true, Type: descriptorOf(lookupByMIME(mime))}, nil
	}
	return &gen.FileTypeInfo{Matched: false}, nil
}
