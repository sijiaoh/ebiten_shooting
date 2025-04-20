package core

import "reflect"

type componentManager struct {
	components map[reflect.Type]Component
}

func newComponentManager() *componentManager {
	return &componentManager{
		components: make(map[reflect.Type]Component),
	}
}

func (cm *componentManager) addComponent(component Component) {
	cm.components[reflect.TypeOf(component)] = component
	component.Awake()
}

func (cm *componentManager) getComponent(componentType reflect.Type) Component {
	if component, ok := cm.components[componentType]; ok {
		return component
	}
	return nil
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
	for _, component := range cm.components {
		if !component.getIsInited() {
			component.Init()
			component.setIsInited(true)
		}
	}
}

func (cm *componentManager) removeDisposedComponents() {
	activeComponents := make(map[reflect.Type]Component)
	for _, component := range cm.components {
		if component.IsActive() {
			activeComponents[reflect.TypeOf(component)] = component
		} else {
			component.OnDisposed()
		}
	}
	cm.components = activeComponents
}
