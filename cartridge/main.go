package cartridge

import (
	"image"
	"math/rand"

	"github.com/TheMightyGit/marv/marvtypes"
)

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

var API marvtypes.MarvAPI

func Start(api marvtypes.MarvAPI) {
	API = api

	area := API.MapBanksGet(MapBankWorking).AllocArea(image.Point{256, 256})
	p := image.Point{}
	for p.Y = 0; p.Y < 256; p.Y++ {
		for p.X = 0; p.X < 256; p.X++ {
			area.Set(p, uint8(rand.Intn(8)), uint8(rand.Intn(8)), uint8(rand.Intn(15)), uint8(16))
		}
	}

	for i := SpriteStart; i <= SpriteEnd; i++ {
		API.SpritesGet(i).Show(GfxBankFont, area)
		API.SpritesGet(i).ChangePos(image.Rectangle{
			Min: image.Point{
				X: rand.Intn(320),
				Y: rand.Intn(200),
			},
			Max: SpriteSize,
		})
	}
}

var (
	cnt = 0.0
)

func Update() {
	for i := SpriteStart; i <= SpriteEnd; i++ {
		API.SpritesGet(i).ChangePos(image.Rectangle{
			Min: image.Point{
				X: int((cnt * float64(i))) % 320,
				// X: rand.Intn(320),
				// Y: rand.Intn(200),
			},
			Max: SpriteSize,
		})
	}
	cnt += 0.01
	// fmt.Println("CART UPDATE:", api)
	// fmt.Printf("About to Bar: %#v\n", fn)
	//fn.Bar()
	//fmt.Println("UPDATE!")
}
