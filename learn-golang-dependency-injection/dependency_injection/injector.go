//go:build wireinject
// +build wireinject

package dependency_injection

import (
	"database/sql"
	"io"
	"learn-golang-dependency-injection/controller"
	"learn-golang-dependency-injection/repository"
	"learn-golang-dependency-injection/service"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
)

// Injector Dependency Injection
func InitializeSimpleService() (*SimpleService, error) {
	wire.Build(NewSimpleRepository, NewSimpleService)
	return nil, nil
}

// Injector Dependency Injection with injector params
func InitializeController(db *sql.DB, validator *validator.Validate) controller.ProductController {
	wire.Build(repository.NewProductRepository, service.NewProductService, controller.NewProductController)
	return nil
}

// Multiple Binding
func InitializeDatabase() *DatabaseRepository {
	wire.Build(NewDatabaseRepository, NewDatabaseMongoDB, NewDatabaseMySQL)
	return nil
}

// Provider set
func InitializeFooBar() *FooBarService {
	wire.Build(fooSet, barSet, NewFooBarService)
	return nil
}

var fooSet = wire.NewSet(NewFooService, NewFooRepository)
var barSet = wire.NewSet(NewBarService, NewBarRepository)

// Wrong Binding Interface (because if struct implement func from interface, google wire need type interface itself, not the struct implemented in)
// func InitializeProductService() *ProductService {
// 	wire.Build(NewProductService, NewProductImpl)
// 	return nil
// }

// Binding Interface
func InitializeProductService() *ProductService {
	wire.Build(ProductSet, NewProductService)
	return nil
}

var ProductSet = wire.NewSet(
	NewProductImpl,
	// Pointer to struct implementation is a MUST
	wire.Bind(new(IProduct), new(*ProductImpl)),
)

// Struct Provider (rarely used)
func InitializeFooBarStruct() *NewFooBar {
	wire.Build(NewFoo, NewBar, wire.Struct(new(NewFooBar), "*"))
	return nil
}

// Binding Value (Struct)
func InitializeFooBarUsingValueStruct() *NewFooBar {
	wire.Build(wire.Value(fooValue), wire.Value(barValue), wire.Struct(new(NewFooBar), "*"))
	return nil
}

var fooValue = &Foo{}
var barValue = &Bar{}

// Binding Value (Interface)
func InitializeFooBarUsingValueInterface() io.Reader {
	wire.Build(wire.InterfaceValue(new(io.Reader), os.Stdin))
	return nil
}

// Struct Field Provider
func InitializeSmartphoneSamsung() *Smartphone {
	wire.Build(wire.FieldsOf(new(*Samsung), "Smartphone"), NewSamsung)
	return nil
}

// Cleanup Function
func InitializeConnection(name string) (*Connection, func()) {
	wire.Build(NewConnection, NewFile)
	return nil, nil
}
