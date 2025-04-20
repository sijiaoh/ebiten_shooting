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
* func NewFoo() *Foo {
* 	return &Foo{
* 		ComponentBase: *core.NewComponentBase(),
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
type Component interface {
	// ライフサイクルメソッド
	Awake()
	Init()
	Update(delta float64)
	Draw(drawManager *DrawerManager)
	OnDisposed()

	// Implemented by ComponentBase
	IsActive() bool
	Dispose()
}

type ComponentBase struct {
	Entity   Entity
	isActive bool
}

func NewComponentBase(entity Entity) *ComponentBase {
	return &ComponentBase{Entity: entity}
}

func (cb *ComponentBase) IsActive() bool {
	return cb.isActive
}

func (cb *ComponentBase) Dispose() {
	cb.isActive = false
}
