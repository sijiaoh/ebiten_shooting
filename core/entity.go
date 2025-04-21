package core

import "reflect"

/*
* テンプレート
*
* package foo
*
* import (
* 	"github.com/hajimehoshi/ebiten/v2"
* 	"github.com/sijiaoh/ebiten_shooting/core"
* )
*
* type Foo struct {
* 	core.EntityBase
* }
*
* func NewFoo() *Foo {
* 	return &Foo{
* 		EntityBase: *core.NewEntityBase(),
* 	}
* }
*
* func (f *Foo) Awake() {
* }
*
* func (f *Foo) Init() {
* }
*
* func (f *Foo) Update(delta float64) {
* }
*
* func (f *Foo) Draw(dm *DrawerManager) {
* }
*
* func (f *Foo) OnDisposed() {
* }
 */
type Entity interface {
	// Lifecycle methods
	Awake()
	Init()
	Update(delta float64)
	Draw(drawManager *DrawerManager)
	OnDisposed()

	// Implemented by ComponentBase

	IsActive() bool
	Dispose()
	TimeScale() float64
	SetTimeScale(scale float64)

	getIsInited() bool
	setIsInited(isInited bool)

	AddComponent(component Component)
	GetComponent(componentType reflect.Type) Component
	updateComponents(delta float64)
	drawComponents(drawManager *DrawerManager)
	disposeComponents()
}

type EntityBase struct {
	isActive         bool
	timeScale        float64
	isInited         bool
	componentManager componentManager
}

func NewEntityBase() *EntityBase {
	return &EntityBase{
		isActive:         true,
		timeScale:        1,
		componentManager: *newComponentManager(),
	}
}

func (eb *EntityBase) IsActive() bool {
	return eb.isActive
}

func (eb *EntityBase) Dispose() {
	eb.isActive = false
}

func (eb *EntityBase) TimeScale() float64 {
	return eb.timeScale
}

func (eb *EntityBase) SetTimeScale(scale float64) {
	eb.timeScale = scale
}

func (eb *EntityBase) getIsInited() bool {
	return eb.isInited
}

func (eb *EntityBase) setIsInited(isInited bool) {
	eb.isInited = isInited
}

func (eb *EntityBase) AddComponent(component Component) {
	eb.componentManager.addComponent(component)
}

func (eb *EntityBase) GetComponent(componentType reflect.Type) Component {
	return eb.componentManager.getComponent(componentType)
}

func (eb *EntityBase) updateComponents(delta float64) {
	eb.componentManager.update(delta)
}

func (eb *EntityBase) drawComponents(dm *DrawerManager) {
	eb.componentManager.draw(dm)
}

func (eb *EntityBase) disposeComponents() {
	eb.componentManager.dispose()
}
