package core

import (
	"sort"

	"github.com/hajimehoshi/ebiten/v2"
)

type DrawerManager struct {
	drawers []*Drawer
}

func newDrawerManager() *DrawerManager {
	return &DrawerManager{}
}

func (dm *DrawerManager) AddDrawer(drawer *Drawer) {
	dm.drawers = append(dm.drawers, drawer)
}

func (dm *DrawerManager) clearDrawers() {
	dm.drawers = dm.drawers[:0]
}

func (dm *DrawerManager) draw(screen *ebiten.Image) {
	sort.SliceStable(dm.drawers, func(i, j int) bool {
		if dm.drawers[i].Layer == dm.drawers[j].Layer {
			return dm.drawers[i].Order < dm.drawers[j].Order
		}
		return dm.drawers[i].Layer < dm.drawers[j].Layer
	})

	for _, drawer := range dm.drawers {
		drawer.Draw(screen)
	}
}
