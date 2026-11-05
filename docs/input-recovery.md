# Stuck keyboard / HUD after cutscenes (#754)

Some quest scripts disable the gameplay HUD and keyboard routing while a dialog or shop UI is active. If the NPC dies or the script is interrupted before the normal cleanup runs, the world can keep simulating while movement keys no longer reach the player.

## Console workaround

Open the in-game console (`~`) and run:

```
unstick
```

This clears known stale flags (chat modal, cutscene gameplay flag, shop camera mode) and resyncs HUD rendering so the keyboard capture window is shown again.

## Optional auto-recovery (off by default)

Add to `opennox.yml`:

```yaml
extensions:
  input:
    auto_recover_stuck: true
```

When enabled, the client detects the same stuck signature during normal gameplay and applies the same recovery steps automatically. Leave this **false** for vanilla-like behavior unless you hit the bug often in modded or speedrun content.
