package core

type entityManager struct {
	entities []Entity
}

func newEntityManager() *entityManager {
	return &entityManager{}
}

func (em *entityManager) addEntity(entity Entity) {
	em.entities = append(em.entities, entity)
	entity.Awake()
}

func (em *entityManager) update() {
	em.initEntities()

	for _, entity := range em.entities {
		if entity.IsActive() {
			delta := 1.0 / 60 * entity.TimeScale()
			entity.Update(delta)
			entity.updateComponents(delta)
		}
	}

	em.removeDisposedEntities()
}

func (em *entityManager) draw(dm *DrawerManager) {
	for _, entity := range em.entities {
		entity.Draw(dm)
		entity.drawComponents(dm)
	}
}

func (em *entityManager) initEntities() {
	for _, entity := range em.entities {
		if !entity.getIsInited() {
			entity.Init()
			entity.setIsInited(true)
		}
	}
}

func (em *entityManager) removeDisposedEntities() {
	var activeEntities []Entity
	for _, entity := range em.entities {
		if entity.IsActive() {
			activeEntities = append(activeEntities, entity)
		} else {
			entity.OnDisposed()
			entity.disposeComponents()
		}
	}
	em.entities = activeEntities
}
