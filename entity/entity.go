package entity

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

	IsActive() bool
}

type EntityBase struct{}

func NewEntityBase() EntityBase {
	return EntityBase{}
}

func (eb *EntityBase) IsActive() bool {
	return true
}
