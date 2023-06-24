package contract

import (
	"database/sql"

	"github.com/stretchr/testify/mock"
)

type Column struct {
	Field   string
	Type    string
	Null    string
	Key     string
	Default sql.NullString
	Extra   string
}

type TableStructureCheckerRepository interface {
	StructureGetter(tablename string) (columns []Column, failure error)
}

type TableStructureCheckerRepositoryMock struct {
	Mock mock.Mock
}

func (m *TableStructureCheckerRepositoryMock) StructureGetter(tablename string) (columns []Column, failure error) {
	args := m.Mock.Called(tablename)
	return args.Get(0).([]Column), args.Error(1)
}
