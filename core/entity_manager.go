package core

type EntityManager struct {
	entities      []Entity
	addedEntities []Entity
}

func NewEntityManager() *EntityManager {
	return &EntityManager{}
}

func (em *EntityManager) AddEntity(entity Entity) {
	em.addedEntities = append(em.addedEntities, entity)
	entity.Awake()
}

func (em *EntityManager) Update() {
	em.initEntities()

	for _, entity := range em.entities {
		if entity.IsActive() {
			entity.Update(1.0 / 60 * entity.TimeScale())
		}
	}

	em.removeDisposedEntities()
}

func (em *EntityManager) Draw(dm *DrawerManager) {
	for _, entity := range em.entities {
		entity.Draw(dm)
	}
}

func (em *EntityManager) initEntities() {
	for _, entity := range em.addedEntities {
		entity.Init()
	}
	em.entities = append(em.entities, em.addedEntities...)
	em.addedEntities = make([]Entity, 0)
}

func (em *EntityManager) removeDisposedEntities() {
	var aliveEntities []Entity
	for _, entity := range em.entities {
		if entity.IsActive() {
			aliveEntities = append(aliveEntities, entity)
		} else {
			entity.OnDisposed()
		}
	}
	em.entities = aliveEntities
}
