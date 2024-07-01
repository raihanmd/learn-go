package simple

import "errors"

type SimpleRepository struct {
	Error bool
}

func NewSimpleRepository(isErr bool) *SimpleRepository {
	return &SimpleRepository{
		Error: isErr,
	}
}

type SimpleService struct {
	SimpleRepository *SimpleRepository
}

func NewSimpleService(simpleRepository *SimpleRepository) (*SimpleService, error) {
	if simpleRepository.Error {
		return nil, errors.New("failed create service")
	} else {
		return &SimpleService{
			SimpleRepository: simpleRepository,
		}, nil
	}
}

type Database struct {
	Name string
}

type DatabasePostgreSQL Database
type DatabaseMongoDB Database

func NewDatabasePostgreSQL() *DatabasePostgreSQL {
	return &DatabasePostgreSQL{"PostgreSQL"}
}

func NewDatabaseMongoDB() *DatabaseMongoDB {
	return &DatabaseMongoDB{"MongoDB"}
}

type DatabaseRepository struct {
	*DatabasePostgreSQL
	*DatabaseMongoDB
}

func NewDatabaseRepository(postgre *DatabasePostgreSQL, mongoDB *DatabaseMongoDB) *DatabaseRepository {
	return &DatabaseRepository{postgre, mongoDB}
}
