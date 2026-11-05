//go:build !server

package opennox

import (
	"image"

	"github.com/spf13/viper"

	"github.com/opennox/opennox/v1/client/noxrender"
)

const (
	configVideoWidescreenFOV  = "video.widescreen_fov"
	configVideoUncappedPresent = "video.uncapped_present"
)

func init() {
	viper.SetDefault(configVideoWidescreenFOV, false)
	viper.SetDefault(configVideoUncappedPresent, false)
	registerOnConfigRead(func() {
		setEnableFrameLimit(!viper.GetBool(configVideoUncappedPresent))
	})
}

func videoApplyWidescreenFOV(vp *noxrender.Viewport, bsz image.Point) {
	if !viper.GetBool(configVideoWidescreenFOV) {
		return
	}
	if bsz.X <= 0 || bsz.Y <= 0 {
		return
	}
	aspect := float64(bsz.X) / float64(bsz.Y)
	const ref = 4.0 / 3.0
	if aspect <= ref+0.01 {
		return
	}
	// Widen the visible world area on ultrawide displays. Off by default: can desync
	// perception in multiplayer and differs from vanilla FOV.
	extra := int(float64(bsz.X) * (1 - ref/aspect) / 2)
	vp.Screen.Min.X -= extra
	if vp.Screen.Min.X < 0 {
		vp.Screen.Min.X = 0
	}
	vp.Screen.Max.X = bsz.X - vp.Screen.Min.X - 1
	if vp.Screen.Max.X >= bsz.X {
		vp.Screen.Max.X = bsz.X - 1
	}
	vp.Size.X = vp.Screen.Dx() + 1
}
