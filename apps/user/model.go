package user

import (
	"github.com/bobacgo/admin-service/apps/repo/model"
	"github.com/bobacgo/orm"
)

const UsersTable = "users"

type User struct {
	model.Model
	Account    string `json:"account"`     // 账号
	Password   string `json:"password"`    // 密码
	Phone      string `json:"phone"`       // 手机号
	Email      string `json:"email"`       // 邮箱
	Status     int8   `json:"status"`      // 状态 1:正常 2:禁用
	RegisterAt int64  `json:"register_at"` // 注册时间
	RegisterIp string `json:"register_ip"` // 注册IP
	LoginAt    int64  `json:"login_at"`    // 登录时间
	LoginIp    string `json:"login_ip"`    // 登录IP
	RoleCodes  string `json:"role_codes"`  // 角色编码，多个用逗号隔开
}

const (
	Account    string = "account"
	Password   string = "password"
	Phone      string = "phone"
	Email      string = "email"
	RegisterAt string = "register_at"
	RegisterIp string = "register_ip"
	LoginAt    string = "login_at"
	LoginIp    string = "login_ip"
	RoleCodes  string = "role_codes"
)

func (m *User) Mapping() []*orm.Mapping {
	return []*orm.Mapping{
		{Column: model.Id, Result: &m.ID, Value: m.ID},
		{Column: Account, Result: &m.Account, Value: m.Account},
		{Column: Password, Result: &m.Password, Value: m.Password},
		{Column: Phone, Result: &m.Phone, Value: m.Phone},
		{Column: Email, Result: &m.Email, Value: m.Email},
		{Column: model.Status, Result: &m.Status, Value: m.Status},
		{Column: RegisterAt, Result: &m.RegisterAt, Value: m.RegisterAt},
		{Column: RegisterIp, Result: &m.RegisterIp, Value: m.RegisterIp},
		{Column: LoginAt, Result: &m.LoginAt, Value: m.LoginAt},
		{Column: LoginIp, Result: &m.LoginIp, Value: m.LoginIp},
		{Column: RoleCodes, Result: &m.RoleCodes, Value: m.RoleCodes},
		{Column: model.Operator, Result: &m.Operator, Value: m.Operator},
		{Column: model.CreatedAt, Result: &m.CreatedAt, Value: m.CreatedAt},
		{Column: model.UpdatedAt, Result: &m.UpdatedAt, Value: m.UpdatedAt},
	}
}
