package dependency_injection

type Database struct {
	Name string
}

// if Provider func used same data, used alias type to fix (Multiple Binding)
type DatabaseMySQL Database
type DatabaseMongoDB Database

type DatabaseRepository struct {
	MongoDB *DatabaseMongoDB
	MySQL   *DatabaseMySQL
}

func NewDatabaseMongoDB() *DatabaseMongoDB {
	database := &Database{
		Name: "MongoDB",
	}
	return (*DatabaseMongoDB)(database)
}

func NewDatabaseMySQL() *DatabaseMySQL {
	database := &Database{
		Name: "MySQL",
	}
	return (*DatabaseMySQL)(database)
}

func NewDatabaseRepository(MongoDB *DatabaseMongoDB, MySQL *DatabaseMySQL) *DatabaseRepository {
	return &DatabaseRepository{
		MongoDB,
		MySQL,
	}
}
