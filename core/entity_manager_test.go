package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockEntity struct {
	EntityBase

	awakeCallCount  int
	initCallCount   int
	updateCallCount int
}

func NewMockEntity() *MockEntity {
	return &MockEntity{EntityBase: *NewEntityBase()}
}

func (me *MockEntity) Awake() {
	me.awakeCallCount++
}

func (me *MockEntity) Init() {
	me.initCallCount++
}

func (me *MockEntity) Update(delta float64) {
	me.updateCallCount++
}

func TestEntityManager_update_CallsEntityLifecycleMethods(t *testing.T) {
	em := newEntityManager()
	mockEntity := NewMockEntity()
	em.addEntity(mockEntity)

	em.update()

	assert.Equal(t, 1, mockEntity.awakeCallCount, "Awake should be called once")
	assert.Equal(t, 1, mockEntity.initCallCount, "Init should be called once")
	assert.Equal(t, 1, mockEntity.updateCallCount, "Update should be called once")
}
