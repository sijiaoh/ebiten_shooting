package core

import "reflect"

type componentManager struct {
	components      map[reflect.Type]Component
	addedComponents []Component
}

func newComponentManager() *componentManager {
	return &componentManager{}
}

func (cm *componentManager) addComponent(component Component) {
	cm.addedComponents = append(cm.addedComponents, component)
	component.Awake()
}

func (cm *componentManager) update(delta float64) {
	cm.initComponents()

	for _, component := range cm.components {
		if component.IsActive() {
			component.Update(delta)
		}
	}

	cm.removeDisposedComponents()
}

func (cm *componentManager) draw(dm *DrawerManager) {
	for _, component := range cm.components {
		component.Draw(dm)
	}
}

func (cm *componentManager) dispose() {
	for _, component := range cm.components {
		component.Dispose()
	}
}

func (cm *componentManager) initComponents() {
	for _, component := range cm.addedComponents {
		component.Init()
		cm.components[reflect.TypeOf(component)] = component
	}
	cm.addedComponents = make([]Component, 0)
}

func (cm *componentManager) removeDisposedComponents() {
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
