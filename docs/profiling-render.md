# Profiling rendering (CPU hot paths)

OpenNox still rasterizes the world in **software** into a 16-bit buffer; the GPU only presents the final image. On a fast GPU (e.g. RTX 4070 Ti Super), **CPU time dominates**.

## Micro-benchmarks (no game data)

From `src/`:

```bash
go test ./client/noxrender/ -bench=. -benchmem -count=3
```

Cases include particle glow (`BenchmarkDrawParticle*`) and render-state reset (`BenchmarkRenderDataReset`).

## In-game

1. Enable the FPS / perf overlay from Advanced Video options.
2. Use `opennox-hd.exe` at your target resolution; compare **nearest** vs **linear** filtering.
3. Optional **`video.vsync: true`** in `opennox.yml` (see vendored libs PR) to cap present rate; sim tick stays unchanged.

## Next steps (engine)

Tile draw, light grid, and drawable sorting remain largely in legacy C (`client__draw*.c`, `draw.go`). Profile with `NOX_DEBUG` env vars where available before porting additional paths to Go.
