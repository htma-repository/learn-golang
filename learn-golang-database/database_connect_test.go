package learn_golang_database

import (
	"testing"
)

func TestConnectDB(t *testing.T) {
	db := ConnectionDB()

	defer db.Close()

}
