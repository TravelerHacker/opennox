# Experimental video options (off by default)

Set in `opennox.yml`:

```yaml
video:
  widescreen_fov: false      # widen visible area on ultrawide; MP/visual parity risk
  uncapped_present: false    # skip client frame-limit sleep; sim tick unchanged
```

With **#4** (vendored libs) merged you can also use:

```yaml
video:
  vsync: false
```

See [profiling-render.md](profiling-render.md) for benchmarks.

## Keyboard stuck after cutscene (#754)

See [input-recovery.md](input-recovery.md) for the `unstick` console command and optional auto-recover config (PR #7).
