package nodes_test

import (
	"context"
	"testing"

	gen "christiangeorgelucas/filetype-tools/gen"
	"christiangeorgelucas/filetype-tools/nodes"
)

// wantCategoryCounts is the number of distinct types the pinned h2non/filetype
// v1.1.3 registers per category — hand-counted directly from the pinned
// module's own matchers/*.go source (len() of each exported category Map,
// read from the actual v1.1.3 module cache, not the library's unreleased
// GitHub HEAD — HEAD has since added avif/exr/parquet/odt/ods/odp that
// v1.1.3 does not have), independent of this package's own categoryOf glue.
// Total = 73.
var wantCategoryCounts = map[string]int{
	"image":       13,
	"video":       10,
	"audio":       9,
	"font":        4,
	"archive":     28,
	"document":    6,
	"application": 3,
}

func TestListSupportedTypes_AllCategories(t *testing.T) {
	ctx := context.Background()
	ax := newTestContext(t)

	got, err := nodes.ListSupportedTypes(ctx, ax, &gen.SupportedTypesQuery{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	wantTotal := 0
	for _, n := range wantCategoryCounts {
		wantTotal += n
	}
	if int(got.GetCount()) != wantTotal {
		t.Fatalf("count = %d, want %d", got.GetCount(), wantTotal)
	}
	if len(got.GetTypes()) != wantTotal {
		t.Fatalf("len(types) = %d, want %d", len(got.GetTypes()), wantTotal)
	}

	counted := map[string]int{}
	seenExt := map[string]bool{}
	for _, d := range got.GetTypes() {
		if d.GetExtension() == "" || d.GetMimeType() == "" {
			t.Errorf("entry with blank extension/mime: %+v", d)
		}
		if d.GetCategory() == "" {
			t.Errorf("entry with no category: %+v", d)
		}
		if seenExt[d.GetExtension()] {
			t.Errorf("duplicate extension in catalog: %q", d.GetExtension())
		}
		seenExt[d.GetExtension()] = true
		counted[d.GetCategory()]++
	}
	for cat, want := range wantCategoryCounts {
		if counted[cat] != want {
			t.Errorf("category %q count = %d, want %d", cat, counted[cat], want)
		}
	}

	// Deterministic order: sorted by extension.
	for i := 1; i < len(got.GetTypes()); i++ {
		if got.GetTypes()[i-1].GetExtension() >= got.GetTypes()[i].GetExtension() {
			t.Fatalf("types not sorted by extension at index %d: %q >= %q",
				i, got.GetTypes()[i-1].GetExtension(), got.GetTypes()[i].GetExtension())
		}
	}
}

func TestListSupportedTypes_CategoryFilter(t *testing.T) {
	ctx := context.Background()
	ax := newTestContext(t)

	got, err := nodes.ListSupportedTypes(ctx, ax, &gen.SupportedTypesQuery{Category: "IMAGE"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if int(got.GetCount()) != wantCategoryCounts["image"] {
		t.Fatalf("count = %d, want %d", got.GetCount(), wantCategoryCounts["image"])
	}
	for _, d := range got.GetTypes() {
		if d.GetCategory() != "image" {
			t.Errorf("entry %+v leaked into image filter", d)
		}
	}
}

func TestListSupportedTypes_UnknownCategoryFilter(t *testing.T) {
	ctx := context.Background()
	ax := newTestContext(t)

	got, err := nodes.ListSupportedTypes(ctx, ax, &gen.SupportedTypesQuery{Category: "not-a-real-category"})
	if err != nil {
		t.Fatalf("unexpected error (must not crash): %v", err)
	}
	if got.GetCount() != 0 || len(got.GetTypes()) != 0 {
		t.Errorf("got %d entries for a nonexistent category filter, want 0", got.GetCount())
	}
}
