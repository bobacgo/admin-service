package api

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/bobacgo/admin-service/apps/repo/dto"
	"github.com/bobacgo/admin-service/apps/service"
	"github.com/bobacgo/admin-service/pkg/kit/hs/response"
)

type RoleHandler struct {
	svc *service.Service
}

func NewRoleHandler(svc *service.Service) *RoleHandler {
	return &RoleHandler{svc: svc}
}

func (h *RoleHandler) Get(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idVal, _ := strconv.ParseInt(id, 10, 64)

	role, err := h.svc.Role.Get(r.Context(), &dto.GetRoleReq{
		ID: idVal,
	})
	if err != nil {
		slog.Error("Get error", "err", err)
		response.JSON(w, response.Resp{
			Code: ErrCodeServer,
			Msg:  err.Error(),
		})
		return
	}
	response.JSON(w, response.Resp{
		Code: OK,
		Msg:  "success",
		Data: role,
	})
}

func (h *RoleHandler) List(w http.ResponseWriter, r *http.Request) {
	req := &dto.RoleListReq{}

	page := r.URL.Query().Get("page")
	pageSize := r.URL.Query().Get("page_size")
	req.Code = r.URL.Query().Get("code")
	req.Status = r.URL.Query().Get("status")

	req.Page, _ = strconv.Atoi(page)
	req.PageSize, _ = strconv.Atoi(pageSize)

	list, total, err := h.svc.Role.List(r.Context(), req)
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
		Data: dto.NewPageResp(total, list),
	})
}

func (h *RoleHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req *dto.RoleCreateReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.JSON(w, response.Resp{
			Code: ErrCodeParam,
			Msg:  err.Error(),
		})
		return
	}

	if err := h.svc.Validator.StructCtx(r.Context(), req); err != nil {
		response.JSON(w, response.Resp{
			Code: ErrCodeParam,
			Msg:  err.Error(),
		})
		return
	}

	if err := h.svc.Role.Create(r.Context(), req); err != nil {
		slog.Error("Create error", "req", req, "err", err)
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

func (h *RoleHandler) Update(w http.ResponseWriter, r *http.Request) {
	req := new(dto.RoleUpdateReq)
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.JSON(w, response.Resp{
			Code: ErrCodeParam,
			Msg:  err.Error(),
		})
		return
	}

	if err := h.svc.Validator.StructCtx(r.Context(), req); err != nil {
		response.JSON(w, response.Resp{
			Code: ErrCodeParam,
			Msg:  err.Error(),
		})
		return
	}

	if err := h.svc.Role.Update(r.Context(), req); err != nil {
		slog.Error("Update error", "req", req, "err", err)
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

func (h *RoleHandler) Delete(w http.ResponseWriter, r *http.Request) {
	ids := r.URL.Query().Get("ids")
	if ids == "" {
		response.JSON(w, response.Resp{
			Code: ErrCodeParam,
			Msg:  "ids is empty",
		})
		return
	}
	if err := h.svc.Role.Delete(r.Context(), ids); err != nil {
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
