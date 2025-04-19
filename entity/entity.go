package entity

/*
* テンプレート
*
* package foo
*
* import (
* 	"github.com/hajimehoshi/ebiten/v2"
* 	"github.com/sijiaoh/ebiten_shooting/scenes/world/entities"
* )
*
* type Foo struct {
* 	entity.EntityBase
* }
*
* func NewFoo() Foo {
* 	return Foo{
* 		EntityBase: entity.NewEntityBase(),
* 	}
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
	// ライフサイクルメソッド
	Init()
	Update(delta float64)
	Draw(drawManager *DrawerManager)
	OnDisposed()

	// 以下はEntityBaseに実装済み

	IsActive() bool
	Dispose()
	TimeScale() float64
}

type EntityBase struct {
	isActive  bool
	timeScale float64
}

func NewEntityBase() EntityBase {
	return EntityBase{
		isActive:  true,
		timeScale: 1,
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
