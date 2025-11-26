package menu

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

type MenuHandler struct {
	svc       *MenuService
	Validator *validator.Validate
}

func NewMenuHandler(svc *MenuService, v *validator.Validate) *MenuHandler {
	return &MenuHandler{svc: svc, Validator: v}
}

func (h *MenuHandler) Tree(w http.ResponseWriter, r *http.Request) {
	menuTree, err := h.svc.Tree(r.Context())
	if err != nil {
		slog.Error("Tree error", "err", err)
		response.JSON(w, response.Resp{
			Code: ecode.ErrCodeServer,
			Msg:  err.Error(),
		})
		return
	}

	response.JSON(w, response.Resp{
		Code: ecode.OK,
		Msg:  "success",
		Data: map[string]any{
			"list": menuTree,
		},
	})
}

func (h *MenuHandler) GetList(w http.ResponseWriter, r *http.Request) {
	req := &MenuListReq{}

	page := r.URL.Query().Get("page")
	pageSize := r.URL.Query().Get("page_size")
	req.Path = r.URL.Query().Get("path")
	req.Name = r.URL.Query().Get("name")

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

func (h *MenuHandler) Get(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idVal, _ := strconv.ParseInt(id, 10, 64)

	menu, err := h.svc.Get(r.Context(), &GetMenuReq{
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
		Data: menu,
	})
}

func (h *MenuHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req *MenuCreateReq
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

func (h *MenuHandler) Update(w http.ResponseWriter, r *http.Request) {
	req := new(MenuUpdateReq)
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

func (h *MenuHandler) Delete(w http.ResponseWriter, r *http.Request) {
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
