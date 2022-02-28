package cartridge

import (
	"embed"
	"image"
	"math/rand"

	"github.com/TheMightyGit/marv/marvlib"
)

//go:embed "resources/*"
var Resources embed.FS

const (
	SpriteStart = iota
	SpriteEnd   = 127
)

const (
	MapBankWorking = iota
)

const (
	GfxBankFont = iota
)

var (
	SpriteSize = image.Point{320, 200}
)

func Start() {

	area := marvlib.API.MapBanksGet(MapBankWorking).AllocArea(image.Point{256, 256})
	p := image.Point{}
	for p.Y = 0; p.Y < 256; p.Y++ {
		for p.X = 0; p.X < 256; p.X++ {
			area.Set(p, uint8(rand.Intn(8)), uint8(rand.Intn(8)), uint8(rand.Intn(15)), uint8(16))
		}
	}

	for i := SpriteStart; i <= SpriteEnd; i++ {
		marvlib.API.SpritesGet(i).Show(GfxBankFont, area)
		marvlib.API.SpritesGet(i).ChangePos(image.Rectangle{
			Min: image.Point{
				X: rand.Intn(320),
				Y: rand.Intn(200),
			},
			Max: SpriteSize,
		})
		//		marv.Sprites[i].ChangeViewport(image.Point{
		//			X: 320,
		//			Y: 200,
		//		})
	}
}

var (
	cnt = 0.0
)

func Update() {
	for i := SpriteStart; i <= SpriteEnd; i++ {
		marvlib.API.SpritesGet(i).ChangePos(image.Rectangle{
			Min: image.Point{
				X: int((cnt * float64(i))) % 320,
				// X: rand.Intn(320),
				// Y: rand.Intn(200),
			},
			Max: SpriteSize,
		})
	}
	cnt += 0.01
}
