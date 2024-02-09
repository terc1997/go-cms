package db

import (
	"fmt"
	"testing"
)

func TestAuthor(t *testing.T) {

	m := NewDBConfig()
	email := "terc@abc.com"

	t.Run("Create author", func(t *testing.T) {
		name := "tarsis"
		id, err := m.CreateAuthor(name, email)
		if err != nil {
			panic("Failed to create author")
		}
		fmt.Printf("User ID: %v\n", id)
	})

	t.Run("Get author", func(t *testing.T) {
		result, err := m.GetAuthor(email)
		if err != nil {
			fmt.Printf("User not found %s\n", err)
		}
		for i := 0; i < len(result); i++ {
			fmt.Printf("username: %v, email: %v\n", result[i].Name, result[i].Email)
		}
	})

	t.Run("Update author", func(t *testing.T) {
		name := "tigre"
		email := "terc@abc.com"

		err := m.UpdateAuthor(name, email)
		if err != nil {
			panic("Error updating author")
		}

		fmt.Println("Succesfully updated")
	})

	t.Run("Deleting author", func(t *testing.T) {
		email := "terc@abc.com"
		err := m.DeleteAuthor(email)

		if err != nil {
			panic("Failed to delete user")
		}

		fmt.Printf("Succesfully deleted ")
	})

}
