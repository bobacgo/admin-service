package orm

type Model interface {
	TableName() string
	Mapping(bool) map[string]any
}
