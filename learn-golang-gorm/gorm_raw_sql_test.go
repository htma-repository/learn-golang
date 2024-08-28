package learn_golang_gorm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecRawGorm(t *testing.T) {
	err := Db.Exec("INSERT INTO users(first_name, last_name, email, password) VALUES(?, ?, ?, ?)", "Tama", "Tri", "tama@example.com", "password").Error
	assert.NoError(t, err)

	err = Db.Exec("INSERT INTO users(first_name, last_name, email, password) VALUES(?, ?, ?, ?)", "Tri", "Rahmanto", "tri@example.com", "password").Error
	assert.NoError(t, err)

	err = Db.Exec("INSERT INTO users(first_name, last_name, email, password) VALUES(?, ?, ?, ?)", "Rahmanto", "Tama", "rahmanto@example.com", "password").Error
	assert.NoError(t, err)
}

func TestQueryRawGorm(t *testing.T) {
	var newUser User
	err := Db.Raw("SELECT id, first_name, last_name, email, password, created_at, updated_at FROM users WHERE id = ?", 1).Scan(&newUser).Error

	assert.NoError(t, err)
	assert.Equal(t, "Tama", newUser.Name.FirstName)

	//

	var newUsers []User
	err = Db.Raw("SELECT id, first_name, last_name, email, password, created_at, updated_at FROM users").Scan(&newUsers).Error

	assert.NoError(t, err)
	assert.Equal(t, "Tama", newUsers[0].Name.FirstName)

	///

	var newUserRow User
	row := Db.Raw("SELECT id, first_name, last_name, email, password, created_at, updated_at FROM users WHERE id = ?", 1).Row()

	err = row.Scan(&newUserRow.ID, &newUserRow.Name.FirstName, &newUserRow.Name.LastName, &newUserRow.Email, &newUserRow.Password, &newUserRow.CreatedAt, &newUserRow.UpdatedAt)
	assert.NoError(t, err)

	assert.Equal(t, "Tama", newUserRow.Name.FirstName)

	//

	var newUserRows []User
	rows, err := Db.Raw("SELECT id, first_name, last_name, email, password, created_at, updated_at FROM users").Rows()
	assert.NoError(t, err)
	defer rows.Close()

	for rows.Next() {
		var newUserRowItem User
		err = rows.Scan(&newUserRowItem.ID, &newUserRowItem.Name.FirstName, &newUserRowItem.Name.LastName, &newUserRowItem.Email, &newUserRowItem.Password, &newUserRowItem.CreatedAt, &newUserRowItem.UpdatedAt)
		assert.NoError(t, err)

		newUserRows = append(newUserRows, newUserRowItem)
	}

	assert.Equal(t, "Tama", newUserRows[0].Name.FirstName)

	//

	var newUserRowsScan []User
	rows, err = Db.Raw("SELECT id, first_name, last_name, email, password, created_at, updated_at FROM users").Rows()
	assert.NoError(t, err)
	defer rows.Close()

	for rows.Next() {
		err := Db.ScanRows(rows, &newUserRowsScan)
		assert.NoError(t, err)
	}

	assert.Equal(t, "Tama", newUserRowsScan[0].Name.FirstName)
}
