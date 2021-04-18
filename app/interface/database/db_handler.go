package database

type DbHandler interface {
	Create(value interface{}, arg ...interface{}) (Result, error)
	Read(value interface{}, arg ...interface{}) (Result, error)
	Update(value interface{}, arg ...interface{}) (Result, error)
	Delete(value interface{}, arg ...interface{}) (Result, error)
}
