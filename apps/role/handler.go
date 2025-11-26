package role

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/bobacgo/admin-service/apps/ecode"
	"github.com/bobacgo/admin-service/apps/repo/dto"
	"github.com/bobacgo/admin-service/pkg/kit/hs/response"
	"github.com/go-playground/validator/v10"
)

type RoleHandler struct {
	svc       *RoleService
	Validator *validator.Validate
}

func NewRoleHandler(svc *RoleService, v *validator.Validate) *RoleHandler {
	return &RoleHandler{svc: svc, Validator: v}
}

func (h *RoleHandler) Get(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idVal, _ := strconv.ParseInt(id, 10, 64)

	role, err := h.svc.Get(r.Context(), &GetRoleReq{
		ID: idVal,
	})
	if err != nil {
		slog.Error("Get error", "err", err)
		response.JSON(w, response.Resp{
			Code: ecode.ErrCodeServer,
			Msg:  err.Error(),
		})
		return
	}
	response.JSON(w, response.Resp{
		Code: ecode.OK,
		Msg:  "success",
		Data: role,
	})
}

func (h *RoleHandler) List(w http.ResponseWriter, r *http.Request) {
	req := &RoleListReq{}

	page := r.URL.Query().Get("page")
	pageSize := r.URL.Query().Get("page_size")
	req.Code = r.URL.Query().Get("code")
	req.Status = r.URL.Query().Get("status")

	req.Page, _ = strconv.Atoi(page)
	req.PageSize, _ = strconv.Atoi(pageSize)

	list, total, err := h.svc.List(r.Context(), req)
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
		Data: dto.NewPageResp(total, list),
	})
}

func (h *RoleHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req *RoleCreateReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.JSON(w, response.Resp{
			Code: ecode.ErrCodeParam,
			Msg:  err.Error(),
		})
		return
	}

	if err := h.Validator.StructCtx(r.Context(), req); err != nil {
		response.JSON(w, response.Resp{
			Code: ecode.ErrCodeParam,
			Msg:  err.Error(),
		})
		return
	}

	if err := h.svc.Create(r.Context(), req); err != nil {
		slog.Error("Create error", "req", req, "err", err)
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

func (h *RoleHandler) Update(w http.ResponseWriter, r *http.Request) {
	req := new(RoleUpdateReq)
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.JSON(w, response.Resp{
			Code: ecode.ErrCodeParam,
			Msg:  err.Error(),
		})
		return
	}

	if err := h.Validator.StructCtx(r.Context(), req); err != nil {
		response.JSON(w, response.Resp{
			Code: ecode.ErrCodeParam,
			Msg:  err.Error(),
		})
		return
	}

	if err := h.svc.Update(r.Context(), req); err != nil {
		slog.Error("Update error", "req", req, "err", err)
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

func (h *RoleHandler) Delete(w http.ResponseWriter, r *http.Request) {
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
