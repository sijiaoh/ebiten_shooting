package core

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
* 	core.ComponentBase
* }
*
* func NewFoo(entity core.Entity) *Foo {
* 	return &Foo{
* 		ComponentBase: *core.NewComponentBase(entity),
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
* func (f *Foo) Draw(dm *core.DrawerManager) {
* }
*
* func (f *Foo) OnDisposed() {
* }
 */
type Component interface {
	// Lifecycle methods
	Awake()
	Init()
	Update(delta float64)
	Draw(drawManager *DrawerManager)
	OnDisposed()

	// Implemented by ComponentBase
	IsActive() bool
	Dispose()

	getIsInited() bool
	setIsInited(isInited bool)
}

type ComponentBase struct {
	Entity   Entity
	isActive bool
	isInited bool
}

func NewComponentBase(entity Entity) *ComponentBase {
	return &ComponentBase{
		Entity:   entity,
		isActive: true,
	}
}

func (cb *ComponentBase) IsActive() bool {
	return cb.isActive
}

func (cb *ComponentBase) Dispose() {
	cb.isActive = false
}

func (cb *ComponentBase) getIsInited() bool {
	return cb.isInited
}

func (cb *ComponentBase) setIsInited(isInited bool) {
	cb.isInited = isInited
}
