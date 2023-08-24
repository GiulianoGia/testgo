package service

import (
	"gotest/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

var groceryService GroceryService

func TestSearchGroceriesFromUser(t *testing.T) {
	type args struct {
		query string
	}
	tests := []struct {
		name string
		args args
		want []types.Grocery
	}{
		{name: "test1", args: args{"Te"}, want: []types.Grocery{{ID: 1, Name: "Test", Quantity: 12, Done: false}}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, groceryService.SearchGroceriesFromUser(test.args.query))
		})
	}
}

func TestGetAllGroceriesFromUser(t *testing.T) {
	type args struct {
		userId string
	}
	tests := []struct {
		name string
		args args
		want []types.Grocery
	}{
		{name: "test1", args: args{userId: "00000000-0000-0000-0000-000000000000"}, want: []types.Grocery{{ID: 3, Name: "Apple", Quantity: 2, Done: true}}},
	}
}
