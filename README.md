# filetype-tools

Composable [Axiom](https://axiomide.com) nodes for file-type / MIME detection from
magic bytes. Wraps [h2non/filetype](https://github.com/h2non/filetype) (MIT,
pure Go, zero dependencies) — a library that identifies a file's type by
inspecting the leading bytes of its content, the same technique the Unix
`file` command and `libmagic` use, without linking `libmagic` or any other
native dependency.

This package **identifies** a file's type; it never **processes** its
content. That makes it a natural upstream router for content-processing
packages (image, PDF, office-document, etc.): detect the type here first,
then hand the same bytes to the package built for that format.

Built for the Axiom marketplace, handle `christiangeorgelucas`.

## Nodes

| Node | Input | Output | What it does |
|---|---|---|---|
| `DetectFileType` | `FileSample` | `FileTypeInfo` | Identify a file's canonical extension, MIME type, and category from its magic bytes. |
| `MatchesType` | `MatchRequest` | `TypeMatch` | Test whether a byte sample matches one specific expected extension or MIME type. |
| `IsTypeSupported` | `TypeIdentifier` | `FileTypeInfo` | Look up whether an extension/MIME string is one this package recognizes (no bytes involved). |
| `ListSupportedTypes` | `SupportedTypesQuery` | `SupportedTypesList` | List the full (or category-filtered) catalog of recognizable types. |
| `IsImage` | `FileSample` | `TypeMatch` | Category predicate: JPEG, JPEG2000, PNG, GIF, WEBP, TIFF, BMP, HEIF, PSD, ICO, CR2, DWG. |
| `IsVideo` | `FileSample` | `TypeMatch` | Category predicate: MP4, M4V, MKV, WEBM, MOV, AVI, WMV, MPEG, FLV, 3GP. |
| `IsAudio` | `FileSample` | `TypeMatch` | Category predicate: MP3, WAV, FLAC, OGG, MIDI, AAC, AIFF, AMR, M4A. |
| `IsFont` | `FileSample` | `TypeMatch` | Category predicate: TTF, OTF, WOFF, WOFF2, EOT. |
| `IsArchive` | `FileSample` | `TypeMatch` | Category predicate: ZIP, TAR, GZIP, RAR, 7Z, XZ, ZSTD, PDF, RTF, SQLite, ELF, Mach-O, EXE, DEB, RPM, ISO 9660, and more. |
| `IsDocument` | `FileSample` | `TypeMatch` | Category predicate: legacy binary DOC/XLS/PPT, or OOXML DOCX/XLSX/PPTX. |
| `IsApplication` | `FileSample` | `TypeMatch` | Category predicate: WebAssembly WASM, Android DEX/ODEX bytecode. |

73 recognized types across 7 categories, pinned to `h2non/filetype` v1.1.3.

## Design

Every detection node takes a `FileSample{ bytes data }` — only the leading
header of a file is needed to identify it, so callers should send at most a
file's first few hundred bytes (the deepest signature check, ISO 9660, needs
32773 bytes; MS-OOXML docx/xlsx/pptx disambiguation can need up to ~13KB). To
keep the deployed instance's cost bounded regardless of what a caller sends,
every node truncates input to the first 64 KiB before inspecting it — this
never changes the result for a real file, since detection never looks past
the header.

## License

MIT. See [LICENSE](./LICENSE). h2non/filetype is itself MIT-licensed and has
zero runtime dependencies (verified against its `go.mod`).
