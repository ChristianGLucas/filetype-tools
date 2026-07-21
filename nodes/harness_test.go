package nodes_test

import (
	"testing"

	"christiangeorgelucas/filetype-tools/axiom"
)

// testContext is a testing.T-backed axiom.Context for unit tests. Shared by
// every *_test.go file in this package — see barcode-tools for the pattern
// this mirrors.
type testContext struct {
	t          *testing.T
	secretsMap map[string]string
}

func newTestContext(t *testing.T) *testContext {
	return &testContext{t: t, secretsMap: map[string]string{}}
}

type testLogger struct{ t *testing.T }

func (l *testLogger) Debug(msg string, args ...any) { l.t.Logf("DEBUG  %s %v", msg, args) }
func (l *testLogger) Info(msg string, args ...any)  { l.t.Logf("INFO   %s %v", msg, args) }
func (l *testLogger) Warn(msg string, args ...any)  { l.t.Logf("WARN   %s %v", msg, args) }
func (l *testLogger) Error(msg string, args ...any) { l.t.Logf("ERROR  %s %v", msg, args) }

type testSecrets struct{ m map[string]string }

func (s testSecrets) Get(name string) (string, bool) { v, ok := s.m[name]; return v, ok }

type testFlowReflection struct{}

func (testFlowReflection) Nodes() []axiom.ReflectionNode     { return nil }
func (testFlowReflection) Edges() []axiom.ReflectionEdge     { return nil }
func (testFlowReflection) LoopEdges() []axiom.ReflectionEdge { return nil }
func (testFlowReflection) Position() axiom.FlowPosition      { return axiom.FlowPosition{} }
func (testFlowReflection) GraphID() string                   { return "" }

type testReflection struct{}

func (testReflection) Flow() axiom.FlowReflection { return testFlowReflection{} }

type testFlowMutation struct{}

func (testFlowMutation) AddNode(_, _ string, _ *axiom.CanvasPosition) uint32 { return 0 }
func (testFlowMutation) AddEdge(_, _ uint32, _ *axiom.EdgeCondition)         {}

type testMutation struct{}

func (testMutation) Flow() axiom.FlowMutation { return testFlowMutation{} }

func (c *testContext) Log() axiom.Logger            { return &testLogger{c.t} }
func (c *testContext) Secrets() axiom.Secrets       { return testSecrets{c.secretsMap} }
func (c *testContext) ExecutionID() string          { return "test-execution-id" }
func (c *testContext) FlowID() string               { return "test-flow-id" }
func (c *testContext) TenantID() string             { return "test-tenant-id" }
func (c *testContext) Reflection() axiom.Reflection { return testReflection{} }
func (c *testContext) Mutation() axiom.Mutation     { return testMutation{} }

var _ axiom.Context = (*testContext)(nil) // compile-time interface check

// ── Independent-oracle fixtures ─────────────────────────────────────────
//
// Every byte sequence below is a canonical magic-number constant taken
// directly from each format's own public specification (PNG, ISO-BMFF/MP4,
// RIFF/WAVE, sfnt/TTF, WebAssembly, the ZIP local-file-header, %PDF, and the
// OLE2/Compound-File-Binary header shared by legacy .doc/.xls/.ppt) — not
// derived from, or by reading, h2non/filetype's own source. They are the
// independent oracle: if the wrapped library's detection ever drifted from
// the real specs, these fixtures would catch it.

// pngSample is the complete 8-byte PNG file signature (ISO/IEC 15948 / the
// PNG spec's "File Header").
var pngSample = []byte{0x89, 'P', 'N', 'G', 0x0D, 0x0A, 0x1A, 0x0A}

// jpegSample is a minimal JPEG SOI + APP0 marker prefix (0xFFD8FFE0).
var jpegSample = []byte{0xFF, 0xD8, 0xFF, 0xE0, 0x00, 0x10, 'J', 'F', 'I', 'F', 0x00}

// mp4Sample is a minimal ISO-BMFF "ftyp" box declaring the "isom" brand.
var mp4Sample = []byte{0x00, 0x00, 0x00, 0x18, 'f', 't', 'y', 'p', 'i', 's', 'o', 'm'}

// wavSample is a minimal RIFF/WAVE header: "RIFF" + size + "WAVE".
var wavSample = []byte{'R', 'I', 'F', 'F', 0x24, 0x00, 0x00, 0x00, 'W', 'A', 'V', 'E'}

// ttfSample is the sfnt version tag for TrueType outlines (0x00010000).
var ttfSample = []byte{0x00, 0x01, 0x00, 0x00, 0x00}

// wasmSample is the WebAssembly binary magic ("\0asm") + version 1, per
// https://webassembly.github.io/spec/core/binary/modules.html#binary-magic.
var wasmSample = []byte{0x00, 'a', 's', 'm', 0x01, 0x00, 0x00, 0x00}

// zipSample is a ZIP local file header signature (PK\x03\x04).
var zipSample = []byte{'P', 'K', 0x03, 0x04, 0x14, 0x00, 0x00, 0x00}

// pdfSample is the "%PDF-" header every PDF file begins with.
var pdfSample = []byte("%PDF-1.4\n")

// oleDocSample is a 514-byte OLE2/Compound-File-Binary-Format buffer
// identifying specifically as legacy .doc: the CFB header (D0 CF 11 E0) at
// offset 0, plus the .doc discriminator bytes (0xEC, 0xA5) at offset
// 512-513, per the Apache OpenOffice bug-tracker reference h2non/filetype's
// own Doc() matcher cites (issue 111457). .doc/.xls/.ppt share an identical
// 4-byte CFB header and are only distinguishable past byte 513, so a sample
// shorter than that is real ambiguity in the wrapped library, not a bug —
// see is_document_test.go.
var oleDocSample = func() []byte {
	b := make([]byte, 514)
	b[0], b[1], b[2], b[3] = 0xD0, 0xCF, 0x11, 0xE0
	b[512], b[513] = 0xEC, 0xA5
	return b
}()

// oleAmbiguousSample is a short (8-byte) OLE2 header with no discriminator
// bytes — h2non/filetype v1.1.3 recognizes this as "some OLE2 compound-file
// document" (doc, xls, and ppt are indistinguishable below 514 bytes; which
// specific one it reports depends on Go's randomized map-iteration order,
// so tests must not assert a specific one of the three for this fixture).
var oleAmbiguousSample = []byte{0xD0, 0xCF, 0x11, 0xE0, 0xA1, 0xB1, 0x1A, 0xE1}

// garbageSample matches no known signature.
var garbageSample = []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08}
