package orm

type model struct{}

func (m *model) TableName() string {
	return ""
}

func (m *model) Mapping(ptr bool) map[string]any {
	return nil
}

func Ptr(ptr bool, p, v any) any {
	if ptr {
		return p
	}
	return v
}

type TestModel struct {
	ID int
	A  string
	B  int
	C  string
}

func (m *TestModel) TableName() string {
	return "test_model"
}

func (m *TestModel) Mapping(ptr bool) map[string]any {
	return map[string]any{
		"id": Ptr(ptr, &m.ID, m.ID),
		"a":  Ptr(ptr, &m.A, m.A),
		"b":  Ptr(ptr, &m.B, m.B),
		"c":  Ptr(ptr, &m.C, m.C),
	}
}
