package fcc_test

import (
	"github.com/stretchr/testify/assert"
	"golang-techniques/fcc"
	"testing"
)

func TestFCC(t *testing.T) {
	t.Run("create domain collection", func(t *testing.T) {

		users := []fcc.User{
			{ID: "1", Name: "name1"},
			{ID: "2", Name: "name2"},
		}

		actual := fcc.NewUsers(users)

		expected := fcc.Users{
			{ID: "1", Name: "name1"},
			{ID: "2", Name: "name2"},
		}

		assert.Equal(t, expected, actual)
	})

	t.Run("immutable collection", func(t *testing.T) {

		users := []fcc.User{
			{ID: "1", Name: "name1"},
			{ID: "2", Name: "name2"},
		}

		actual := fcc.NewUsers(users)

		expected := fcc.Users{
			{ID: "1", Name: "name1"},
			{ID: "2", Name: "name2"},
		}

		// change original slice
		users = append(users[:1], users[2:]...)

		assert.Equal(t, expected, actual)
		assert.NotEqual(t, actual, users)
	})

	t.Run("filtering by condition", func(t *testing.T) {

		users := []fcc.User{
			{ID: "1", Name: "name1"},
			{ID: "2", Name: "name2"},
		}

		target := fcc.NewUsers(users)

		actual := target.Filter(func(user fcc.User) bool {
			return user.Name == "name1"
		})

		expected := fcc.Users{
			{ID: "1", Name: "name1"},
		}

		assert.Equal(t, expected, actual)
	})

	t.Run("immutable collection", func(t *testing.T) {

		users := []fcc.User{
			{ID: "1", Name: "name1"},
			{ID: "2", Name: "name2"},
		}

		original := []fcc.User{
			{ID: "1", Name: "name1"},
			{ID: "2", Name: "name2"},
		}

		target := fcc.NewUsers(users)

		actual := target.Filter(func(user fcc.User) bool {
			return user.Name == "name1"
		})

		expected := fcc.Users{
			{ID: "1", Name: "name1"},
		}

		assert.Equal(t, expected, actual)
		assert.Equal(t, original, users)
		assert.NotEqual(t, users, actual)
	})
}
