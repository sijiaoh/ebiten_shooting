package entities

import (
	"github.com/hajimehoshi/ebiten/v2"
)

/*
* テンプレート
*
* package foo
*
* import (
* 	"github.com/hajimehoshi/ebiten/v2"
* 	"github.com/sijiaoh/ebiten_shooting/entities"
* )
*
* type Foo struct {
* 	entities.EntityBase
* }
*
* func NewFoo() Foo {
* 	return Foo{
* 		EntityBase: entities.NewEntityBase(),
* 	}
* }
*
* func (f *Foo) Init() {
* }
*
* func (f *Foo) Update() {
* }
*
* func (f *Foo) Draw(screen *ebiten.Image) {
* }
*
* func (f *Foo) OnDisposed() {
* }
 */
type Entity interface {
	// ライフサイクルメソッド
	Init()
	Update()
	Draw(screen *ebiten.Image)
	OnDisposed()

	// 以下はEntityBaseに実装済み

	IsInited() bool
	EndInit()
	IsActive() bool
}

type EntityBase struct {
	isInited bool
}

func NewEntityBase() EntityBase {
	return EntityBase{
		isInited: false,
	}
}

func (eb *EntityBase) IsInited() bool {
	return eb.isInited
}

func (eb *EntityBase) EndInit() {
	eb.isInited = true
}

func (eb *EntityBase) IsActive() bool {
	return true
}
