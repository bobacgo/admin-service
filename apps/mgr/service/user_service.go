package service

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/bobacgo/admin-service/apps/common/dto"
	dto2 "github.com/bobacgo/admin-service/apps/mgr/dto"
	repo2 "github.com/bobacgo/admin-service/apps/mgr/repo"
	"github.com/bobacgo/admin-service/apps/mgr/repo/model"
	"github.com/bobacgo/admin-service/pkg/util"
	"github.com/go-playground/validator/v10"
)

type UserService struct {
	repo      *repo2.UserRepo
	validator *validator.Validate
}

func NewUserService(r *repo2.UserRepo, v *validator.Validate) *UserService {
	return &UserService{repo: r, validator: v}
}

// Get /user/one 获取单个用户
func (s *UserService) GetOne(ctx context.Context, req *dto2.GetUserReq) (*model.User, error) {
	return s.repo.FindOne(ctx, req)
}

// Get /user/list 获取用户列表
func (s *UserService) GetList(ctx context.Context, req *dto2.UserListReq) (*dto.PageResp[model.User], error) {
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
func (s *UserService) Post(ctx context.Context, req *model.User) (any, error) {
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

	return struct{}{}, nil
}

// Put /user 更新用户
func (s *UserService) Put(ctx context.Context, req *dto2.UpdateUserReq) (any, error) {
	if err := s.validator.StructCtx(ctx, req); err != nil {
		return nil, err
	}

	req.Operator = "system" // 示例值，实际应从上下文获取
	req.UpdatedAt = time.Now().Unix()
	if err := s.repo.Update(ctx, req); err != nil {
		return nil, err
	}

	return struct{}{}, nil
}

// Put /user/status 更新用户状态
func (s *UserService) PutStatus(ctx context.Context, req *dto2.UpdateUserStatusReq) (any, error) {
	if err := s.validator.StructCtx(ctx, req); err != nil {
		return nil, err
	}
	req.Operator = "system" // 示例值，实际应从上下文获取
	req.UpdatedAt = time.Now().Unix()
	if err := s.repo.UpdateStatus(ctx, req); err != nil {
		return nil, err
	}

	// TODO: 移除用户会话

	return struct{}{}, nil
}

// Put /user/role 更新用户角色
func (s *UserService) PutRole(ctx context.Context, req *dto2.UpdateUserRoleReq) (any, error) {
	if err := s.validator.StructCtx(ctx, req); err != nil {
		return nil, err
	}
	req.Operator = "system" // 示例值，实际应从上下文获取
	req.UpdatedAt = time.Now().Unix()
	if err := s.repo.UpdateRole(ctx, req); err != nil {
		return nil, err
	}

	// TODO: 移除用户会话

	return struct{}{}, nil
}

// Put /user/password 更新用户密码
func (s *UserService) PutPassword(ctx context.Context, req *dto2.UpdateUserPasswordReq) (any, error) {
	if err := s.validator.StructCtx(ctx, req); err != nil {
		return nil, err
	}
	req.Operator = "system" // 示例值，实际应从上下文获取
	req.UpdatedAt = time.Now().Unix()
	if err := s.repo.UpdatePassword(ctx, req); err != nil {
		return nil, err
	}

	// TODO: 移除用户会话

	return struct{}{}, nil
}

// Delete /user 删除用户
func (s *UserService) Delete(ctx context.Context, req *dto2.DeleteUserReq) (any, error) {

	// TODO: 移除用户会话

	return nil, s.repo.Delete(ctx, req.IDs)
}

// Post /Login 用户登录
func (s *UserService) PostLogin(ctx context.Context, req *dto2.LoginReq) (*dto2.LoginResp, error) {
	row, err := s.repo.FindOne(ctx, &dto2.GetUserReq{
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

	// 更新登录信息
	if err = s.repo.UpdateLoginInfo(ctx, &dto2.UpdateLoginInfoReq{
		Id:      row.ID,
		LoginAt: time.Now().Unix(),
		LoginIp: "127.0.0.1", // TODO: 示例值，实际应从上下文获取
	}); err != nil {
		return nil, err
	}

	// 生成并返回用户会话 token
	token, err := util.GenerateJWT(row.ID, row.Account, 12*time.Hour)
	if err != nil {
		return nil, err
	}
	// TODO: 记录登录日志

	return &dto2.LoginResp{Token: token}, nil
}

// Get /user/logout 用户登出
func (s *UserService) GetLogout(ctx context.Context, _ any) (any, error) {
	return struct{}{}, nil
}
