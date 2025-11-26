package user

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"time"

	"github.com/bobacgo/admin-service/apps/ecode"
	"github.com/bobacgo/admin-service/apps/repo/dto"

	"github.com/bobacgo/admin-service/pkg/kit/hs/response"
)

type UserHandler struct {
	svc *UserService
}

func NewUserHandler(svc *UserService) *UserHandler {
	return &UserHandler{svc: svc}
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req = new(LoginReq)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		response.JSON(w, response.Resp{
			Code: ecode.ErrCodeParam,
			Msg:  err.Error(),
		})
		return
	}

	row, err := h.svc.Get(r.Context(), &GetUserReq{
		Account: req.Account,
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			response.JSON(w, response.Resp{
				Code: ecode.ErrCodeUsernameOrPassword,
				Msg:  "username or password error",
			})
			return
		}
		slog.Error("get error", "req", req, "err", err)
		response.JSON(w, response.Resp{
			Code: ecode.ErrCodeServer,
			Msg:  err.Error(),
		})
		return
	}

	if req.Password != row.Password {
		response.JSON(w, response.Resp{
			Code: ecode.ErrCodeUsernameOrPassword,
			Msg:  "username or password error",
		})
		return
	}

	if row.Status != 1 {
		response.JSON(w, response.Resp{
			Code: ecode.ErrCodeUserDisabled,
			Msg:  "username or password error",
		})
		return
	}

	row.LoginAt = time.Now().Unix()
	row.LoginIp = r.Host
	if err = h.svc.Update(r.Context(), row); err != nil {
		slog.Error("Update error", "err", err)
	}

	response.JSON(w, response.Resp{
		Code: ecode.OK,
		Msg:  "ok",
		Data: map[string]string{"token": "xxxxx"},
	})
}

func (h *UserHandler) Logout(w http.ResponseWriter, r *http.Request) {
	response.JSON(w, response.Resp{
		Code: ecode.OK,
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
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		response.JSON(w, response.Resp{
			Code: ecode.ErrCodeParam,
			Msg:  err.Error(),
		})
		return
	}

	user.RegisterIp = r.Host
	user.RegisterAt = time.Now().Unix()

	if err := h.svc.Create(r.Context(), &user); err != nil {
		slog.Error("Create error", "user", user, "err", err)
		response.JSON(w, response.Resp{
			Code: ecode.ErrCodeServer,
			Msg:  err.Error(),
		})
		return
	}
	response.JSON(w, response.Resp{
		Code: ecode.OK,
		Msg:  "success",
	})
}

func (h *UserHandler) List(w http.ResponseWriter, r *http.Request) {
	var req = new(UserListReq)

	slog.Info("List", "req", req)
	rows, total, err := h.svc.List(r.Context(), req)
	if err != nil {
		slog.Error("List error", "req", req, "err", err)
		response.JSON(w, response.Resp{
			Code: ecode.ErrCodeServer,
			Msg:  err.Error(),
		})
		return
	}

	response.JSON(w, response.Resp{
		Code: ecode.OK,
		Msg:  "success",
		Data: dto.NewPageResp(total, rows),
	})
}

func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		response.JSON(w, response.Resp{
			Code: ecode.ErrCodeParam,
			Msg:  err.Error(),
		})
		return
	}
	if err := h.svc.Update(r.Context(), &user); err != nil {
		slog.Error("Update error", "user", user, "err", err)
		response.JSON(w, response.Resp{
			Code: ecode.ErrCodeServer,
			Msg:  err.Error(),
		})
		return
	}
	response.JSON(w, response.Resp{
		Code: ecode.OK,
		Msg:  "success",
	})
}

func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	ids := r.URL.Query().Get("ids")
	if ids == "" {
		response.JSON(w, response.Resp{
			Code: ecode.ErrCodeParam,
			Msg:  "ids is empty",
		})
		return
	}
	if err := h.svc.Delete(r.Context(), ids); err != nil {
		slog.Error("Delete error", "ids", ids, "err", err)
		response.JSON(w, response.Resp{
			Code: ecode.ErrCodeServer,
			Msg:  err.Error(),
		})
		return
	}
	response.JSON(w, response.Resp{
		Code: ecode.OK,
		Msg:  "success",
	})
}
