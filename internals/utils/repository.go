package utils

//go:generate mockery --name=Repository
type Repository interface {
	Create(ent interface{}) (err error)
	List(filters map[string]interface{}, order interface{}, out interface{}) (err error)
	Read(filters interface{}, ent interface{}) error
}
