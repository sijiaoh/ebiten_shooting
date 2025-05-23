package core

import (
	"sort"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
)

type DrawerManager struct {
	pool    sync.Pool
	drawers []*drawer
}

func newDrawerManager() *DrawerManager {
	return &DrawerManager{
		pool: sync.Pool{
			New: func() interface{} {
				return &drawer{}
			},
		},
	}
}

func (dm *DrawerManager) NewDrawer() *drawer {
	drawer := dm.pool.New().(*drawer)
	drawer.reset()
	return drawer
}

func (dm *DrawerManager) AddDrawer(drawer *drawer) {
	dm.drawers = append(dm.drawers, drawer)
}

func (dm *DrawerManager) clearDrawers() {
	for _, drawer := range dm.drawers {
		dm.pool.Put(drawer)
	}
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
