//go:build !server

package opennox

import (
	noxflags "github.com/opennox/opennox/v1/common/flags"
	"github.com/opennox/opennox/v1/common/memmap"
	"github.com/opennox/opennox/v1/legacy"
)

var autoRecoverStuckInput bool

func init() {
	configBoolPtr("extensions.input.auto_recover_stuck", "", false, &autoRecoverStuckInput)
}

// clientInputLooksStuck reports whether client-side GUI/input state can block
// gameplay keys while the world is still simulating (see issue #754).
func clientInputLooksStuck() bool {
	if gameGetPlayState() != 3 {
		return false
	}
	if legacy.Sub_47A260() != 0 {
		return false
	}
	if noxClient != nil && noxClient.GUI.Captured() != nil {
		return false
	}
	if legacy.Get_dword_5d4594_1064868() != 0 {
		return true
	}
	if noxflags.HasGamePlay(noxflags.GameplayFlag4) {
		return true
	}
	if legacy.Sub_478030() != 0 {
		return true
	}
	if !nox_client_renderGUI_80828 {
		if memmap.Uint32(0x587000, 153436) == 1 {
			return true
		}
		if memmap.Uint32(0x5D4594, 1098612) == 1 {
			return true
		}
	}
	return false
}

// clientRecoverStuckInput clears known stale modal/cutscene flags and resyncs
// the keyboard capture window via sub_437100. Returns human-readable actions.
func clientRecoverStuckInput() []string {
	var out []string
	if legacy.Sub_47A260() == 1 {
		guiCloseNPCDialog()
		out = append(out, "closed NPC dialog")
	}
	if legacy.Sub_478030() == 1 {
		legacy.Sub_479280()
		out = append(out, "ended shop/camera GUI mode")
	}
	if legacy.Get_dword_5d4594_1064868() != 0 {
		*memmap.PtrUint32(0x5D4594, 1064868) = 0
		out = append(out, "cleared chat input modal")
	}
	if noxflags.HasGamePlay(noxflags.GameplayFlag4) {
		noxflags.UnsetGamePlay(noxflags.GameplayFlag4)
		out = append(out, "cleared gameplay cutscene flag")
	}
	if !nox_client_renderGUI_80828 {
		if memmap.Uint32(0x587000, 153436) == 1 || memmap.Uint32(0x5D4594, 1098612) == 1 {
			enableGUIDrawing(true)
			legacy.Sub_437100()
			out = append(out, "restored gameplay HUD and keyboard focus")
		}
	}
	return out
}

func clientMaybeAutoRecoverStuckInput() {
	if !autoRecoverStuckInput || !clientInputLooksStuck() {
		return
	}
	clientRecoverStuckInput()
}
