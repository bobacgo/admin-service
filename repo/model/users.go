package model

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

func (*User) TableName() string {
	return "users"
}

func (m *User) Mapping() map[string]func(*User) any {
	return map[string]func(*User) any{
		"account":     func(u *User) any { return u.Account },
		"password":    func(u *User) any { return u.Password },
		"phone":       func(u *User) any { return u.Phone },
		"email":       func(u *User) any { return u.Email },
		"status":      func(u *User) any { return u.Status },
		"register_at": func(u *User) any { return u.RegisterAt },
		"register_ip": func(u *User) any { return u.RegisterIp },
		"login_at":    func(u *User) any { return u.LoginAt },
		"login_ip":    func(u *User) any { return u.LoginIp },
		"created_at":  func(u *User) any { return u.CreatedAt },
		"updated_at":  func(u *User) any { return u.UpdatedAt },
	}
}

func (m *User) MappingSelect() map[string]func(*User) any {
	return map[string]func(*User) any{
		"id":          func(u *User) any { return &u.ID },
		"account":     func(u *User) any { return &u.Account },
		"password":    func(u *User) any { return &u.Password },
		"phone":       func(u *User) any { return &u.Phone },
		"email":       func(u *User) any { return &u.Email },
		"status":      func(u *User) any { return &u.Status },
		"register_at": func(u *User) any { return &u.RegisterAt },
		"register_ip": func(u *User) any { return &u.RegisterIp },
		"login_at":    func(u *User) any { return &u.LoginAt },
		"login_ip":    func(u *User) any { return &u.LoginIp },
		"created_at":  func(u *User) any { return &u.CreatedAt },
		"updated_at":  func(u *User) any { return &u.UpdatedAt },
	}
}