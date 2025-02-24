package actors

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
* 	"github.com/sijiaoh/ebiten_shooting/actors"
* )
*
* type Foo struct {
* 	actors.ActorBase
* }
*
* func NewFoo() Foo {
* 	return Foo{
* 		ActorBase: actors.NewActorBase(),
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
* func (f *Foo) OnDead() {
* }
 */
type Actor interface {
	// ライフサイクルメソッド
	Init()
	Update()
	Draw(screen *ebiten.Image)
	OnDead()

	// 以下はActorBaseに実装済み

	IsInited() bool
	EndInit()
	IsAlive() bool
}

type ActorBase struct {
	isInited bool
}

func NewActorBase() ActorBase {
	return ActorBase{
		isInited: false,
	}
}

func (ub *ActorBase) IsInited() bool {
	return ub.isInited
}

func (ub *ActorBase) EndInit() {
	ub.isInited = true
}

func (ub *ActorBase) IsAlive() bool {
	return true
}
