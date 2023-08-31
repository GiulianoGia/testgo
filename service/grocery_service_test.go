package service

import (
	"gotest/db/mocks"
	"gotest/types"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

type prepareMocks func(db *mocks.DataStore)

func TestGetAllGroceries(t *testing.T) {
	tests := []struct {
		name         string
		want         []types.Grocery
		prepareMocks prepareMocks
		wantErr      error
	}{
		{
			name: "test1",
			want: []types.Grocery{{ID: 3, Name: "Apple", Quantity: 2, Done: true}},
			prepareMocks: func(db *mocks.DataStore) {
				db.On("GetAllGroceries").Return([]types.Grocery{{ID: 3, Name: "Apple", Quantity: 2, Done: true}}, nil)
			}, wantErr: nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			db := &mocks.DataStore{}
			test.prepareMocks(db)
			service := NewServiceStruct(db)
			assert.Equal(t, test.want, service.GetAllGroceries())
		})
	}
}

func TestGetAllGroceriesFromUser(t *testing.T) {
	type args struct {
		userId string
	}

	tests := []struct {
		name         string
		args         args
		want         []types.Grocery
		err          error
		prepareMocks prepareMocks
		wantErr      error
	}{
		{
			name: "Test1",
			args: args{userId: "00000000-0000-0000-0000-000000000000"},
			want: []types.Grocery{{ID: 3, Name: "Fish", Quantity: 3, Done: false}}, err: nil,
			prepareMocks: func(db *mocks.DataStore) {
				db.On("GetAllGroceriesFromUser", "00000000-0000-0000-0000-000000000000").Return([]types.UserGrocery{{UserID: uuid.MustParse("00000000-0000-0000-0000-000000000000"), GroceryID: 3}}, nil)
				db.On("FindGroceryWithId", 3).Return(types.Grocery{ID: 3, Name: "Fish", Quantity: 3, Done: false}, nil)
			},
			wantErr: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			db := &mocks.DataStore{}
			test.prepareMocks(db)
			service := NewServiceStruct(db)
			groceries, err := service.GetAllGroceriesFromUser(test.args.userId)
			assert.Equal(t, test.want, groceries)
			assert.Equal(t, test.wantErr, err)

			db.AssertExpectations(t)
		})
	}
}

func TestGetGroceryByName(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name         string
		args         args
		want         []types.Grocery
		err          error
		prepareMocks prepareMocks
		wantErr      error
	}{
		{
			name: "Test1",
			args: args{name: "Apple"},
			want: []types.Grocery{{ID: 3, Name: "Apple", Quantity: 12, Done: false}},
			prepareMocks: func(db *mocks.DataStore) {
				db.On("GetGroceriesByName", "Apple").Return([]types.Grocery{{ID: 3, Name: "Apple", Quantity: 12, Done: false}}, nil)
			},
		},
		{
			name: "Test2",
			args: args{name: "Appl"},
			want: []types.Grocery{{ID: 3, Name: "Apple", Quantity: 12, Done: false}},
			prepareMocks: func(db *mocks.DataStore) {
				db.On("GetGroceriesByName", "Appl").Return([]types.Grocery{{ID: 3, Name: "Apple", Quantity: 12, Done: false}}, nil)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			db := &mocks.DataStore{}
			test.prepareMocks(db)
			service := NewServiceStruct(db)
			groceries := service.GetGroceryByName(test.args.name)
			assert.Equal(t, test.want, groceries)
		})
	}
}
