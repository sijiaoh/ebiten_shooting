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
	initEntities(em.entities)
	updateEntities(em.entities)
	removeDisposedEntities(true, em.entities)
}

func (em *entityManager) draw(dm *DrawerManager) {
	drawEntities(em.entities, dm)
}

func initEntities(entities []Entity) {
	for _, entity := range entities {
		if !entity.getIsInited() {
			entity.Init()
			entity.setIsInited(true)
		}

		initEntities(entity.Children())
	}
}

func updateEntities(entities []Entity) {
	for _, entity := range entities {
		if entity.IsActive() {
			delta := 1.0 / 60 * entity.TimeScale()
			entity.Update(delta)
			entity.updateComponents(delta)

			updateEntities(entity.Children())
		}
	}
}

func drawEntities(entities []Entity, dm *DrawerManager) {
	for _, entity := range entities {
		if entity.IsActive() {
			entity.Draw(dm)
			entity.drawComponents(dm)

			drawEntities(entity.Children(), dm)
		}
	}
}

func removeDisposedEntities(parentActive bool, entities []Entity) []Entity {
	var activeEntities []Entity
	for _, entity := range entities {
		realActive := parentActive && entity.IsActive()

		children := removeDisposedEntities(realActive, entity.Children())
		entity.setChildren(children)

		if realActive {
			activeEntities = append(activeEntities, entity)
		} else {
			entity.OnDisposed()
			entity.disposeComponents()

			entity.setParent(nil)
			entity.setChildren(nil)
		}
	}

	return activeEntities
}
