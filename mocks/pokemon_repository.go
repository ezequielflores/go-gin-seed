// Code generated by MockGen. DO NOT EDIT.
// Source: ./repository.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	pokemon "github.com/redbeestudios/go-seed/internal/application/model/pokemon"
)

// MockPokemonRepository is a mock of PokemonRepository interface.
type MockPokemonRepository struct {
	ctrl     *gomock.Controller
	recorder *MockPokemonRepositoryMockRecorder
}

// MockPokemonRepositoryMockRecorder is the mock recorder for MockPokemonRepository.
type MockPokemonRepositoryMockRecorder struct {
	mock *MockPokemonRepository
}

// NewMockPokemonRepository creates a new mock instance.
func NewMockPokemonRepository(ctrl *gomock.Controller) *MockPokemonRepository {
	mock := &MockPokemonRepository{ctrl: ctrl}
	mock.recorder = &MockPokemonRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPokemonRepository) EXPECT() *MockPokemonRepositoryMockRecorder {
	return m.recorder
}

// GetByName mocks base method.
func (m *MockPokemonRepository) GetByName(ctx context.Context, name string) (*pokemon.Pokemon, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByName", ctx, name)
	ret0, _ := ret[0].(*pokemon.Pokemon)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByName indicates an expected call of GetByName.
func (mr *MockPokemonRepositoryMockRecorder) GetByName(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByName", reflect.TypeOf((*MockPokemonRepository)(nil).GetByName), ctx, name)
}

// SavePokemon mocks base method.
func (m *MockPokemonRepository) SavePokemon(ctx context.Context, pokemon *pokemon.Pokemon) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SavePokemon", ctx, pokemon)
	ret0, _ := ret[0].(error)
	return ret0
}

// SavePokemon indicates an expected call of SavePokemon.
func (mr *MockPokemonRepositoryMockRecorder) SavePokemon(ctx, pokemon interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SavePokemon", reflect.TypeOf((*MockPokemonRepository)(nil).SavePokemon), ctx, pokemon)
}
