package api

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/bobacgo/admin-service/pkg/kit-web/response"
	"github.com/bobacgo/admin-service/repo/dto"
	"github.com/bobacgo/admin-service/service"
)

type I18nHandler struct {
	svc *service.Service
}

func NewI18nHandler(svc *service.Service) *I18nHandler {
	return &I18nHandler{svc: svc}
}

func (h *I18nHandler) List(w http.ResponseWriter, r *http.Request) {
	req := &dto.I18nListReq{}

	list, err := h.svc.I18n.List(r.Context(), req)
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
		Data: map[string]any{"list": list, "total": len(list)},
	})
}

func (h *I18nHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req *dto.I18nCreateReq
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

	if err := h.svc.I18n.Create(r.Context(), req); err != nil {
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

func (h *I18nHandler) Update(w http.ResponseWriter, r *http.Request) {
	req := new(dto.I18nUpdateReq)
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

	if err := h.svc.I18n.Update(r.Context(), req); err != nil {
		slog.Error("Update error", "i18n", req, "err", err)
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

func (h *I18nHandler) Delete(w http.ResponseWriter, r *http.Request) {
	ids := r.URL.Query().Get("ids")
	if ids == "" {
		response.JSON(w, response.Resp{
			Code: ErrCodeParam,
			Msg:  "ids is empty",
		})
		return
	}
	if err := h.svc.I18n.Delete(r.Context(), ids); err != nil {
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
