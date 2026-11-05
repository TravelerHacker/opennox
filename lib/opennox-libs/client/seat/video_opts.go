package seat

// VideoVsync enables GL swap interval (vsync) when true. Default is off for parity with legacy pacing.
var VideoVsync bool

// VideoStreamUpload uploads frames with TexSubImage2D after the first full texture allocation.
var VideoStreamUpload = true
