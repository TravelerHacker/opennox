# Vendored `opennox/libs`

This fork vendors [opennox/libs](https://github.com/opennox/libs) so rendering, SDL, and OpenGL changes can ship in **TravelerHacker/opennox** without a separate libs repository.

`src/go.mod` contains:

```go
replace github.com/opennox/libs => ../lib/opennox-libs
```

## Patches in this copy

- **Go 1.25** toolchain alignment (`go.mod`).
- **Streaming present path**: `TexSubImage2D` after first frame (`seat/opengl/surface.go`, `seat.VideoStreamUpload`).
- **Optional vsync**: `seat.VideoVsync` wired from OpenNox config (`video.vsync`, default `false`).

To upstream later, extract commits from this directory into github.com/opennox/libs or a TravelerHacker/opennox-libs fork and drop the `replace` directive.
