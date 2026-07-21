package nodes_test

import (
	"context"
	"testing"

	gen "christiangeorgelucas/filetype-tools/gen"
	"christiangeorgelucas/filetype-tools/nodes"
)

func TestDetectFileType(t *testing.T) {
	cases := []struct {
		name     string
		data     []byte
		wantExt  string
		wantMime string
		wantCat  string
	}{
		{"png", pngSample, "png", "image/png", "image"},
		{"jpeg", jpegSample, "jpg", "image/jpeg", "image"},
		{"mp4", mp4Sample, "mp4", "video/mp4", "video"},
		{"wav", wavSample, "wav", "audio/x-wav", "audio"},
		{"ttf", ttfSample, "ttf", "application/font-sfnt", "font"},
		{"wasm", wasmSample, "wasm", "application/wasm", "application"},
		{"zip", zipSample, "zip", "application/zip", "archive"},
		{"pdf", pdfSample, "pdf", "application/pdf", "archive"},
		{"ole-doc", oleDocSample, "doc", "application/msword", "document"},
	}

	ctx := context.Background()
	ax := newTestContext(t)
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, err := nodes.DetectFileType(ctx, ax, &gen.FileSample{Data: c.data})
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !got.GetMatched() {
				t.Fatalf("matched = false, want true")
			}
			ty := got.GetType()
			if ty.GetExtension() != c.wantExt {
				t.Errorf("extension = %q, want %q", ty.GetExtension(), c.wantExt)
			}
			if ty.GetMimeType() != c.wantMime {
				t.Errorf("mime_type = %q, want %q", ty.GetMimeType(), c.wantMime)
			}
			if ty.GetCategory() != c.wantCat {
				t.Errorf("category = %q, want %q", ty.GetCategory(), c.wantCat)
			}
		})
	}
}

func TestDetectFileType_NoMatch(t *testing.T) {
	ctx := context.Background()
	ax := newTestContext(t)

	for _, c := range []struct {
		name string
		data []byte
	}{
		{"empty", nil},
		{"garbage", garbageSample},
	} {
		t.Run(c.name, func(t *testing.T) {
			got, err := nodes.DetectFileType(ctx, ax, &gen.FileSample{Data: c.data})
			if err != nil {
				t.Fatalf("unexpected error (no-match must not crash): %v", err)
			}
			if got.GetMatched() {
				t.Fatalf("matched = true for %s input, want false", c.name)
			}
			if got.GetType() != nil {
				t.Errorf("type = %v, want nil/unset when matched=false", got.GetType())
			}
		})
	}
}

// TestDetectFileType_TruncatesOversizedInput proves the 64KiB input cap
// (nodes/helper.go's maxInputBytes) does not break detection for a real
// file whose magic bytes sit at the very start — the cap must protect the
// deployed instance from an oversized payload without changing correctness.
func TestDetectFileType_TruncatesOversizedInput(t *testing.T) {
	ctx := context.Background()
	ax := newTestContext(t)

	buf := make([]byte, 300*1024) // 300 KiB, well over the 64 KiB cap
	copy(buf, pngSample)

	got, err := nodes.DetectFileType(ctx, ax, &gen.FileSample{Data: buf})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !got.GetMatched() || got.GetType().GetExtension() != "png" {
		t.Fatalf("got = %+v, want a matched png despite oversized input", got)
	}
}
