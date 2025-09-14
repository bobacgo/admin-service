package model

const UsersTable = "users"

type User struct {
	Model
	Account    string `json:"account"`
	Password   string `json:"password"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
	Status     int8   `json:"status"` // 状态 1:正常 2:禁用
	RegisterAt int64  `json:"register_at"`
	RegisterIp string `json:"register_ip"`
	LoginAt    int64  `json:"login_at"`
	LoginIp    string `json:"login_ip"`
}

func (User) TableName() string {
	return UsersTable
}

func (m User) Mapping(ptr bool) map[string]any {
	return map[string]any{
		"account":     ptrFunc(ptr, &m.Account, m.Account),
		"password":    ptrFunc(ptr, &m.Password, m.Password),
		"phone":       ptrFunc(ptr, &m.Phone, m.Phone),
		"email":       ptrFunc(ptr, &m.Email, m.Email),
		"status":      ptrFunc(ptr, &m.Status, m.Status),
		"register_at": ptrFunc(ptr, &m.RegisterAt, m.RegisterAt),
		"register_ip": ptrFunc(ptr, &m.RegisterIp, m.RegisterIp),
		"login_at":    ptrFunc(ptr, &m.LoginAt, m.LoginAt),
		"login_ip":    ptrFunc(ptr, &m.LoginIp, m.LoginIp),
		"created_at":  ptrFunc(ptr, &m.CreatedAt, m.CreatedAt),
		"updated_at":  ptrFunc(ptr, &m.UpdatedAt, m.UpdatedAt),
	}
}
