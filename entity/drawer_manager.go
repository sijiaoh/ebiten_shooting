package entity

import (
	"sort"

	"github.com/hajimehoshi/ebiten/v2"
)

type DrawerManager struct {
	drawers []Drawer
}

func NewDrawerManager() DrawerManager {
	return DrawerManager{}
}

func (dm *DrawerManager) AddDrawer(drawer Drawer) {
	dm.drawers = append(dm.drawers, drawer)
}

func (dm *DrawerManager) Draw(screen *ebiten.Image) {
	sort.SliceStable(dm.drawers, func(i, j int) bool {
		if dm.drawers[i].Layer == dm.drawers[j].Layer {
			return dm.drawers[i].Order < dm.drawers[j].Order
		}
		return dm.drawers[i].Layer < dm.drawers[j].Layer
	})

	for _, drawer := range dm.drawers {
		drawer.Draw(screen)
	}

	dm.drawers = dm.drawers[:0]
}
