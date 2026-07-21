// Package nodes implements christiangeorgelucas/filetype-tools — thin,
// stateless wrappers around github.com/h2non/filetype (MIT, pure Go, zero
// dependencies). This file holds logic shared by more than one node:
// input-size bounding, and deriving a broad category for a matched type
// (the upstream library exposes seven per-category matcher maps but no
// single "category of this Type" lookup, so this glue composes them).
package nodes

import (
	"sort"
	"strings"

	"github.com/h2non/filetype/matchers"
	"github.com/h2non/filetype/types"

	gen "christiangeorgelucas/filetype-tools/gen"
)

// maxInputBytes bounds how much of a byte sample this package inspects.
// h2non/filetype's own matchers never need more than this to disambiguate
// any of the 73 signatures it recognizes — the deepest is the ISO 9660
// boot-record check at offset 32773, and MS-OOXML (docx/xlsx/pptx vs. plain
// zip) disambiguation scans up to roughly 13KB into the ZIP local file
// headers. Capping here bounds the base64/JSON decode and matcher-scan cost
// per invocation regardless of how much a caller sends — detection only
// ever needs the header, so callers should send at most this much, and
// sending more is truncated rather than rejected.
const maxInputBytes = 64 * 1024

// truncate bounds buf to maxInputBytes. Detection never needs more than the
// header for any format this package recognizes, so truncation does not
// change the result for a real file's leading bytes — it only protects the
// deployed instance's CPU/memory from an oversized payload.
func truncate(buf []byte) []byte {
	if len(buf) > maxInputBytes {
		return buf[:maxInputBytes]
	}
	return buf
}

// categoryOrder fixes the iteration order used to resolve a Type's category
// and to build the deterministic catalog in ListSupportedTypes. Every
// registered Type belongs to exactly one of these seven maps (the upstream
// library partitions its matchers this way), so order does not change any
// result — it just makes output order reproducible.
var categoryOrder = []struct {
	name string
	set  matchers.Map
}{
	{"archive", matchers.Archive},
	{"document", matchers.Document},
	{"font", matchers.Font},
	{"audio", matchers.Audio},
	{"video", matchers.Video},
	{"image", matchers.Image},
	{"application", matchers.Application},
}

// categoryOf returns the broad category ("image", "archive", etc.) that a
// recognized Type belongs to, or "" if it is not a member of any of the
// seven category matcher maps (only possible for types.Unknown, or a
// custom type registered outside this package).
func categoryOf(kind types.Type) string {
	for _, c := range categoryOrder {
		if _, ok := c.set[kind]; ok {
			return c.name
		}
	}
	return ""
}

// lookupByMIME finds a registered Type by its exact (lowercase) MIME
// string. types.Types is a sync.Map keyed by extension, not MIME, so this
// mirrors the linear scan filetype.IsMIMESupported does internally — there
// is no faster public lookup. Returns types.Unknown if no type has this
// MIME value.
func lookupByMIME(mime string) types.Type {
	found := types.Unknown
	types.Types.Range(func(_, v interface{}) bool {
		kind := v.(types.Type)
		if kind.MIME.Value == mime {
			found = kind
			return false
		}
		return true
	})
	return found
}

// allTypes returns every registered Type except the "unknown" sentinel,
// sorted by extension for a deterministic, reproducible catalog order.
func allTypes() []types.Type {
	var all []types.Type
	types.Types.Range(func(_, v interface{}) bool {
		kind := v.(types.Type)
		if kind.Extension != "" && kind.Extension != types.Unknown.Extension {
			all = append(all, kind)
		}
		return true
	})
	sort.Slice(all, func(i, j int) bool { return all[i].Extension < all[j].Extension })
	return all
}

// normalizeIdentifier trims whitespace and lowercases an extension or MIME
// string, matching the lowercase form every type is registered under.
func normalizeIdentifier(s string) string {
	return strings.ToLower(strings.TrimSpace(s))
}

// descriptorOf converts a matched library Type into the package's public
// TypeDescriptor, filling in the category via categoryOf.
func descriptorOf(kind types.Type) *gen.TypeDescriptor {
	return &gen.TypeDescriptor{
		Extension: kind.Extension,
		MimeType:  kind.MIME.Value,
		Category:  categoryOf(kind),
	}
}
