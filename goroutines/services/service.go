// services/service.go

package services

import (
	"goroutines/utils"
)

type Service struct{}

func (s *Service) GetDbResult() []utils.DbResult {
	return []utils.DbResult{
		{Id: "1", Name: "John"},
		{Id: "2", Name: "Doe"},
	}
}
