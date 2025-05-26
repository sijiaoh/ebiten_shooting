package core

import "reflect"

/*
* Entityは複数のComponentと子Entityから構成される、最も基本となるオブジェクト
*
* Entityは作成後一度だけEntityツリーに追加する必要がある
* それはEntityManagerにaddEntityするか、既存のEntityにaddChildrenすることで可能である
*
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
*   f.EntityBase.Awake()
* }
*
* func (f *Foo) Init() {
*   f.EntityBase.Init()
* }
*
* func (f *Foo) Update(delta float64) {
*   f.EntityBase.Update(delta)
* }
*
* func (f *Foo) Draw(dm *DrawerManager) {
*   f.EntityBase.Draw(dm)
* }
*
* func (f *Foo) OnDisposed() {
*   f.EntityBase.OnDisposed()
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

	Parent() Entity
	Children() []Entity
	AddChild(entity Entity)
	setParent(entity Entity)
	addChild(entity Entity)
	setChildren(entities []Entity)

	AddComponent(component Component)
	GetComponent(componentType reflect.Type) Component
	updateComponents(delta float64)
	drawComponents(drawManager *DrawerManager)
	disposeComponents()
}

type EntityBase struct {
	isActive  bool
	timeScale float64
	isInited  bool

	parent   Entity
	children []Entity

	componentManager componentManager
}

func NewEntityBase() *EntityBase {
	return &EntityBase{
		isActive:  true,
		timeScale: 1,

		parent:   nil,
		children: make([]Entity, 0),

		componentManager: *newComponentManager(),
	}
}

func (eb *EntityBase) Awake() {}

func (eb *EntityBase) Init() {}

func (eb *EntityBase) Update(delta float64) {}

func (eb *EntityBase) Draw(dm *DrawerManager) {}

func (eb *EntityBase) OnDisposed() {}

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

func (eb *EntityBase) Parent() Entity {
	return eb.parent
}

func (eb *EntityBase) Children() []Entity {
	return eb.children
}

func (eb *EntityBase) AddChild(entity Entity) {
	entity.setParent(eb)
	eb.addChild(entity)
	entity.Awake()
}

func (eb *EntityBase) setParent(entity Entity) {
	eb.parent = entity
}

func (eb *EntityBase) addChild(entity Entity) {
	eb.children = append(eb.children, entity)
}

func (eb *EntityBase) setChildren(entities []Entity) {
	eb.children = entities
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
