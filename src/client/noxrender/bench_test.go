package noxrender

import (
	"image"
	"log/slog"
	"testing"
)

func BenchmarkDrawParticleWhite(b *testing.B) {
	benchmarkDrawParticle(b, particleOpt{rad: 10, blur: 0, intens: 0xff, color: RGB{0xff, 0xff, 0xff}})
}

func BenchmarkDrawParticleBlur32(b *testing.B) {
	benchmarkDrawParticle(b, particleOpt{rad: 10, blur: 32, intens: 0x80, color: RGB{140, 220, 80}})
}

func benchmarkDrawParticle(b *testing.B, opt particleOpt) {
	img := genParticle(opt)
	off, sz, _ := img.Meta()
	csz := image.Point{X: (sz.X + off.X) * 2, Y: (sz.Y + off.Y) * 2}
	pos := image.Pt(sz.X/2, sz.Y/2)
	pix := newBlack16(csz.X, csz.Y)
	d := newRenderData(csz.X, csz.Y)
	r := NewRender(slog.Default(), nil)
	r.SetPixBuffer(pix)
	r.SetData(d)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r.DrawImage16(img, pos)
	}
}

func BenchmarkRenderDataReset(b *testing.B) {
	d := newRenderData(1024, 768)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d.Reset()
		d.SetClipRect(image.Rect(0, 0, 1024, 768))
	}
}
