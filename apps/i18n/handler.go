package i18n

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

type I18nHandler struct {
	svc       *I18nService
	Validator *validator.Validate
}

func NewI18nHandler(svc *I18nService, v *validator.Validate) *I18nHandler {
	return &I18nHandler{svc: svc, Validator: v}
}

func (h *I18nHandler) Get(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idVal, _ := strconv.ParseInt(id, 10, 64)

	i18n, err := h.svc.Get(r.Context(), &GetI18nReq{
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
		Data: i18n,
	})
}

func (h *I18nHandler) List(w http.ResponseWriter, r *http.Request) {
	req := &I18nListReq{}

	page := r.URL.Query().Get("page")
	pageSize := r.URL.Query().Get("page_size")
	req.Class = r.URL.Query().Get("class")
	req.Lang = r.URL.Query().Get("lang")
	req.Key = r.URL.Query().Get("key")

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

func (h *I18nHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req *I18nCreateReq
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

func (h *I18nHandler) Update(w http.ResponseWriter, r *http.Request) {
	req := new(I18nUpdateReq)
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

func (h *I18nHandler) Delete(w http.ResponseWriter, r *http.Request) {
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
