# Hybrid 2× (`walleky/opennox-hd`) evaluation

**Status:** not integrated in this fork yet.

[walleky/opennox-hd](https://github.com/walleky/opennox-hd) ships an experimental **Hybrid 2×** client with pre-scaled assets and a Windows installer. It targets the same upstream OpenNox engine and requires an owned GOG copy.

## Why not merged here

- **Separate release pipeline** (installer, asset blobs) not present in TravelerHacker/opennox.
- **License/review**: hybrid assets and redistribution terms must be verified against EA/Westwood modding guidelines before bundling.
- **Maintenance**: would split testing between vanilla GOG data and hybrid asset packs.

## Recommended path for playtesting

1. Use this fork’s **`opennox-hd.exe`** with GOG data and `opennox-hd` resolution settings.
2. Optionally install **walleky/opennox-hd** side-by-side and compare visuals; cherry-pick engine fixes (e.g. m2xim PRs #751–753) already ported here.

Integration PRs welcome once asset licensing and CI story are defined.
