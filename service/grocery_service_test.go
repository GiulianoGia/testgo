package service

import (
	"gotest/db/mocks"
	"gotest/types"
	"testing"

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
		{name: "test1", want: []types.Grocery{{ID: 3, Name: "Apple", Quantity: 2, Done: true}}, prepareMocks: func(db *mocks.DataStore) {
			db.On("GetAllGroceries").Return([]types.Grocery{{ID: 3, Name: "Apple", Quantity: 2, Done: true}}, nil)
		}, wantErr: nil},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, []types.Grocery{{ID: 3, Name: "Apple", Quantity: 2, Done: true}})
		})
	}
}
