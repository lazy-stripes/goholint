package states

// State of a PPU at a given step.
type State int

// Mere collection of constants because I was procrastinating when I did that for package cpu.

// PPU states, ordered so the enum corresponds to its STAT mode number.
const (
	HBlank        = iota // Mode 0
	VBlank               // Mode 1
	OAMSearch            // Mode 2, initial value for PPU
	PixelTransfer        // Mode 3
)

// Fetcher states. Using bitfield to allow combinations. ([TUGBT] 00:37:52)
const (
	ReadTileID = 1 << iota
	ReadTileData0
	ReadTileData1
	PushToFIFO
	ReadSpriteID
	ReadSpriteFlags
	ReadSpriteData0
	ReadSpriteData1
	MixInFIFO

	// Useful combinations.
	FetchingSprite = ReadSpriteID | ReadSpriteFlags | ReadSpriteData0 |
		ReadSpriteData1 | MixInFIFO
)

// OAM states.
const (
	ReadSpriteY = iota
	ReadSpriteX
)
