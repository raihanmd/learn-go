//go:build wireinject
// +build wireinject

package simple

import (
	"io"
	"os"

	"github.com/google/wire"
)

func InitializeService(isError bool) (*SimpleService, error) {
	wire.Build(NewSimpleService, NewSimpleRepository)
	return nil, nil
}

func InitializeDatabaseRepository() *DatabaseRepository {
	wire.Build(
		NewDatabasePostgreSQL,
		NewDatabaseMongoDB,
		NewDatabaseRepository,
	)
	return nil
}

var fooSet = wire.NewSet(NewFooRepository, NewFooService)
var barSet = wire.NewSet(NewBarRepository, NewBarService)

func InitializeFoobarService() *FooBarService {
	wire.Build(fooSet, barSet, NewFooBarService)
	return nil
}

var FooBarSet = wire.NewSet(NewFoo, NewBar)

func InitializeFooBar() *FooBar {
	wire.Build(
		FooBarSet,
		wire.Struct(new(FooBar), "*"),
	)
	return nil
}

var HelloSet = wire.NewSet(
	NewSayHelloImpl,
	wire.Bind(new(SayHello), new(*SayHelloImpl)),
)

func InitializeHelloService() *HelloService {
	wire.Build(NewHelloService, HelloSet)
	return nil
}

var FooBarValueSet = wire.NewSet(
	wire.Value(&Foo{}),
	wire.Value(&Bar{}),
)

func InitializeFooBarUsingValue() *FooBar {
	wire.Build(FooBarValueSet, wire.Struct(new(FooBar), "*"))
	return nil
}

func InitializeRader() io.Reader {
	wire.Build(wire.InterfaceValue(new(io.Reader), os.Stdin))
	return nil
}

func InitializedConfiguration() *Configuration {
	wire.Build(
		NewApplication,
		wire.FieldsOf(new(*Application), "Configuration"),
	)
	return nil
}

func InitializedConnection(name string) (*Connection, func()) {
	wire.Build(NewConnection, NewFile)
	return nil, nil
}
