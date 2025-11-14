package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"time"

	"github.com/bobacgo/admin-service/apps/repo/dto"
	"github.com/bobacgo/admin-service/apps/repo/model"
	"github.com/bobacgo/admin-service/apps/service"
	"github.com/bobacgo/admin-service/pkg/kit/hs/response"
)

type UserHandler struct {
	svc *service.Service
}

func NewUserHandler(svc *service.Service) *UserHandler {
	return &UserHandler{svc: svc}
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req = new(dto.LoginReq)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		response.JSON(w, response.Resp{
			Code: ErrCodeParam,
			Msg:  err.Error(),
		})
		return
	}

	row, err := h.svc.User.Get(r.Context(), &dto.GetUserReq{
		Account: req.Account,
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			response.JSON(w, response.Resp{
				Code: ErrCodeUsernameOrPassword,
				Msg:  "username or password error",
			})
			return
		}
		slog.Error("get error", "req", req, "err", err)
		response.JSON(w, response.Resp{
			Code: ErrCodeServer,
			Msg:  err.Error(),
		})
		return
	}

	if req.Password != row.Password {
		response.JSON(w, response.Resp{
			Code: ErrCodeUsernameOrPassword,
			Msg:  "username or password error",
		})
		return
	}

	if row.Status != 1 {
		response.JSON(w, response.Resp{
			Code: ErrCodeUserDisabled,
			Msg:  "username or password error",
		})
		return
	}

	row.LoginAt = time.Now().Unix()
	row.LoginIp = r.Host
	if err = h.svc.User.Update(r.Context(), row); err != nil {
		slog.Error("Update error", "err", err)
	}

	response.JSON(w, response.Resp{
		Code: OK,
		Msg:  "ok",
		Data: map[string]string{"token": "xxxxx"},
	})
}

func (h *UserHandler) Logout(w http.ResponseWriter, r *http.Request) {
	response.JSON(w, response.Resp{
		Code: OK,
		Msg:  "ok",
	})
}

func (h *UserHandler) GetInfo(w http.ResponseWriter, r *http.Request) {
	response.JSON(w, response.Resp{
		Code: 0,
		Msg:  "ok",
		Data: map[string]string{"username": "admin"},
	})
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		response.JSON(w, response.Resp{
			Code: ErrCodeParam,
			Msg:  err.Error(),
		})
		return
	}

	user.RegisterIp = r.Host
	user.RegisterAt = time.Now().Unix()

	if err := h.svc.User.Create(r.Context(), &user); err != nil {
		slog.Error("Create error", "user", user, "err", err)
		response.JSON(w, response.Resp{
			Code: ErrCodeServer,
			Msg:  err.Error(),
		})
		return
	}
	response.JSON(w, response.Resp{
		Code: OK,
		Msg:  "success",
	})
	return
}

func (h *UserHandler) List(w http.ResponseWriter, r *http.Request) {
	var req = new(dto.UserListReq)
	// if err := json.NewDecoder(r.Body).Decode(req); err != nil {
	// 	response.JSON(w, response.Resp{
	// 		Code: ErrCodeParam,
	// 		Msg:  err.Error(),
	// 	})
	// 	return
	// }

	slog.Info("List", "req", req)
	rows, total, err := h.svc.User.List(r.Context(), req)
	if err != nil {
		slog.Error("List error", "req", req, "err", err)
		response.JSON(w, response.Resp{
			Code: ErrCodeServer,
			Msg:  err.Error(),
		})
		return
	}

	response.JSON(w, response.Resp{
		Code: OK,
		Msg:  "success",
		Data: dto.NewPageResp(total, rows),
	})
}

func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		response.JSON(w, response.Resp{
			Code: ErrCodeParam,
			Msg:  err.Error(),
		})
		return
	}
	if err := h.svc.User.Update(r.Context(), &user); err != nil {
		slog.Error("Update error", "user", user, "err", err)
		response.JSON(w, response.Resp{
			Code: ErrCodeServer,
			Msg:  err.Error(),
		})
		return
	}
	response.JSON(w, response.Resp{
		Code: OK,
		Msg:  "success",
	})
}

func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	ids := r.URL.Query().Get("ids")
	if ids == "" {
		response.JSON(w, response.Resp{
			Code: ErrCodeParam,
			Msg:  "ids is empty",
		})
		return
	}
	if err := h.svc.User.Delete(r.Context(), ids); err != nil {
		slog.Error("Delete error", "ids", ids, "err", err)
		response.JSON(w, response.Resp{
			Code: ErrCodeServer,
			Msg:  err.Error(),
		})
		return
	}
	response.JSON(w, response.Resp{
		Code: OK,
		Msg:  "success",
	})
}
