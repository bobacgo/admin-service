package user

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/bobacgo/admin-service/apps/repo/dto"
	"github.com/go-playground/validator/v10"
)

type UserService struct {
	repo      *UserRepo
	validator *validator.Validate
}

func NewUserService(r *UserRepo, v *validator.Validate) *UserService {
	return &UserService{repo: r, validator: v}
}

// Get /user/one 获取单个用户
func (s *UserService) GetOne(ctx context.Context, req *GetUserReq) (*User, error) {
	return s.repo.FindOne(ctx, req)
}

// Get /user/list 获取用户列表
func (s *UserService) GetList(ctx context.Context, req *UserListReq) (*dto.PageResp[User], error) {
	rows, total, err := s.repo.Find(ctx, req)
	if err != nil {
		return nil, err
	}
	for _, row := range rows {
		row.Password = ""
	}
	return dto.NewPageResp(total, rows), nil
}

// Post /user 创建用户
func (s *UserService) Post(ctx context.Context, req *User) (*User, error) {
	if err := s.validator.StructCtx(ctx, req); err != nil {
		return nil, err
	}

	now := time.Now().Unix()
	req.CreatedAt = now
	req.UpdatedAt = now
	req.RegisterAt = now
	// Operator field should be set by the caller or middleware

	if err := s.repo.Create(ctx, req); err != nil {
		return nil, err
	}

	req.Password = ""
	return req, nil
}

// Put /user 更新用户
func (s *UserService) Put(ctx context.Context, req *User) (interface{}, error) {
	if err := s.validator.StructCtx(ctx, req); err != nil {
		return nil, err
	}

	// 先查询现有用户数据，保留未修改的字段
	existUser, err := s.repo.FindOne(ctx, &GetUserReq{ID: uint(req.ID)})
	if err != nil {
		return nil, err
	}

	// 只更新前端发送的非空字段，保留原有数据
	if req.Account != "" {
		existUser.Account = req.Account
	}
	if req.Phone != "" {
		existUser.Phone = req.Phone
	}
	if req.Email != "" {
		existUser.Email = req.Email
	}
	if req.Status != 0 {
		existUser.Status = req.Status
	}
	if req.RoleIds != "" {
		existUser.RoleIds = req.RoleIds
	}
	if req.Operator != "" {
		existUser.Operator = req.Operator
	}

	now := time.Now().Unix()
	existUser.UpdatedAt = now

	if err := s.repo.Update(ctx, existUser); err != nil {
		return nil, err
	}

	return nil, nil
}

// Delete /user 删除用户
func (s *UserService) Delete(ctx context.Context, req *DeleteUserReq) (interface{}, error) {
	return nil, s.repo.Delete(ctx, req.IDs)
}

// Login 用户登录（特殊方法，不参与自动路由）
func (s *UserService) Login(ctx context.Context, req *LoginReq) (map[string]string, error) {
	row, err := s.repo.FindOne(ctx, &GetUserReq{
		Account: req.Account,
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("username or password error")
		}
		return nil, err
	}

	if req.Password != row.Password {
		return nil, errors.New("username or password error")
	}

	if row.Status != 1 {
		return nil, errors.New("user disabled")
	}

	row.LoginAt = time.Now().Unix()
	if err = s.repo.Update(ctx, row); err != nil {
		return nil, err
	}

	return map[string]string{"token": "xxxxx"}, nil
}

// Get /user/logout 用户登出
func (s *UserService) GetLogout(ctx context.Context, _ any) (any, error) {
	return nil, nil
}
