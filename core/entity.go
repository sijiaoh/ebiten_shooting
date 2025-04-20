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
	// ライフサイクルメソッド
	Awake()
	Init()
	Update(delta float64)
	Draw(drawManager *DrawerManager)
	OnDisposed()

	// 以下はEntityBaseに実装済み

	IsActive() bool
	Dispose()
	TimeScale() float64
	SetTimeScale(scale float64)

	AddComponent(component Component)
	UpdateComponents(delta float64)
	DrawComponents(drawManager *DrawerManager)
	DisposeComponents()
}

type EntityBase struct {
	isActive         bool
	timeScale        float64
	componentManager ComponentManager
}

func NewEntityBase() *EntityBase {
	return &EntityBase{
		isActive:         true,
		timeScale:        1,
		componentManager: *NewComponentManager(),
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

func (eb *EntityBase) AddComponent(component Component) {
	eb.componentManager.AddComponent(component)
}

func (eb *EntityBase) UpdateComponents(delta float64) {
	eb.componentManager.Update(delta)
}

func (eb *EntityBase) DrawComponents(dm *DrawerManager) {
	eb.componentManager.Draw(dm)
}

func (eb *EntityBase) DisposeComponents() {
	eb.componentManager.Dispose()
}
