package nodes_test

import (
	"context"
	"testing"

	gen "christiangeorgelucas/filetype-tools/gen"
	"christiangeorgelucas/filetype-tools/nodes"
)

func TestIsTypeSupported(t *testing.T) {
	ctx := context.Background()
	ax := newTestContext(t)

	t.Run("known-extension", func(t *testing.T) {
		got, err := nodes.IsTypeSupported(ctx, ax, &gen.TypeIdentifier{Extension: "PNG"})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !got.GetMatched() {
			t.Fatal("matched = false, want true for known extension")
		}
		if got.GetType().GetMimeType() != "image/png" || got.GetType().GetCategory() != "image" {
			t.Errorf("type = %+v, want mime_type=image/png category=image", got.GetType())
		}
	})

	t.Run("known-mime", func(t *testing.T) {
		got, err := nodes.IsTypeSupported(ctx, ax, &gen.TypeIdentifier{MimeType: "application/pdf"})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !got.GetMatched() || got.GetType().GetExtension() != "pdf" {
			t.Errorf("got = %+v, want matched pdf", got)
		}
	})

	t.Run("unknown-extension", func(t *testing.T) {
		got, err := nodes.IsTypeSupported(ctx, ax, &gen.TypeIdentifier{Extension: "notareal extension"})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if got.GetMatched() {
			t.Errorf("matched = true, want false for a nonexistent extension")
		}
	})

	t.Run("unknown-mime", func(t *testing.T) {
		got, err := nodes.IsTypeSupported(ctx, ax, &gen.TypeIdentifier{MimeType: "application/x-not-a-real-mime"})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if got.GetMatched() {
			t.Errorf("matched = true, want false for a nonexistent mime")
		}
	})

	t.Run("neither-set", func(t *testing.T) {
		got, err := nodes.IsTypeSupported(ctx, ax, &gen.TypeIdentifier{})
		if err != nil {
			t.Fatalf("unexpected error (must not crash): %v", err)
		}
		if got.GetMatched() {
			t.Errorf("matched = true, want false when neither field is set")
		}
	})
}
