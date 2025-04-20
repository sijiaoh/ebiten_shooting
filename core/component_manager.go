package core

import "reflect"

type ComponentManager struct {
	components      map[reflect.Type]Component
	addedComponents []Component
}

func NewComponentManager() *ComponentManager {
	return &ComponentManager{}
}

func (cm *ComponentManager) AddComponent(component Component) {
	cm.addedComponents = append(cm.addedComponents, component)
	component.Awake()
}

func (cm *ComponentManager) Update(delta float64) {
	cm.initComponents()

	for _, component := range cm.components {
		if component.IsActive() {
			component.Update(delta)
		}
	}

	cm.removeDisposedComponents()
}

func (cm *ComponentManager) Draw(dm *DrawerManager) {
	for _, component := range cm.components {
		component.Draw(dm)
	}
}

func (cm *ComponentManager) Dispose() {
	for _, component := range cm.components {
		component.Dispose()
	}
}

func (cm *ComponentManager) initComponents() {
	for _, component := range cm.addedComponents {
		component.Init()
		cm.components[reflect.TypeOf(component)] = component
	}
	cm.addedComponents = make([]Component, 0)
}

func (cm *ComponentManager) removeDisposedComponents() {
	var activeComponents map[reflect.Type]Component
	for _, component := range cm.components {
		if component.IsActive() {
			activeComponents[reflect.TypeOf(component)] = component
		} else {
			component.OnDisposed()
		}
	}
	cm.components = activeComponents
}
