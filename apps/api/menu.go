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

type MenuHandler struct {
	svc *service.Service
}

func NewMenuHandler(svc *service.Service) *MenuHandler {
	return &MenuHandler{svc: svc}
}

func (h *MenuHandler) Tree(w http.ResponseWriter, r *http.Request) {
	menuTree, err := h.svc.Menu.Tree(r.Context())
	if err != nil {
		slog.Error("Tree error", "err", err)
		response.JSON(w, response.Resp{
			Code: ErrCodeServer,
			Msg:  err.Error(),
		})
		return
	}

	response.JSON(w, response.Resp{
		Code: OK,
		Msg:  "success",
		Data: map[string]any{
			"list": menuTree,
		},
	})
}

func (h *MenuHandler) GetList(w http.ResponseWriter, r *http.Request) {
	req := &dto.MenuListReq{}

	page := r.URL.Query().Get("page")
	pageSize := r.URL.Query().Get("page_size")
	req.Path = r.URL.Query().Get("path")
	req.Name = r.URL.Query().Get("name")

	req.Page, _ = strconv.Atoi(page)
	req.PageSize, _ = strconv.Atoi(pageSize)

	list, total, err := h.svc.Menu.List(r.Context(), req)
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

func (h *MenuHandler) Get(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idVal, _ := strconv.ParseInt(id, 10, 64)

	menu, err := h.svc.Menu.Get(r.Context(), &dto.GetMenuReq{
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
		Data: menu,
	})
}

func (h *MenuHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req *dto.MenuCreateReq
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

	if err := h.svc.Menu.Create(r.Context(), req); err != nil {
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

func (h *MenuHandler) Update(w http.ResponseWriter, r *http.Request) {
	req := new(dto.MenuUpdateReq)
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

	if err := h.svc.Menu.Update(r.Context(), req); err != nil {
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

func (h *MenuHandler) Delete(w http.ResponseWriter, r *http.Request) {
	ids := r.URL.Query().Get("ids")
	if ids == "" {
		response.JSON(w, response.Resp{
			Code: ErrCodeParam,
			Msg:  "ids is empty",
		})
		return
	}
	if err := h.svc.Menu.Delete(r.Context(), ids); err != nil {
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

const menuJson = `
 [
  {
    "path": "/mgr",
    "name": "manger",
    "component": "LAYOUT",
    "redirect": "/mgr/user",
    "meta": {
      "title": {
        "zh_CN": "系统管理",
        "en_US": "System Manger"
      },
      "icon": "view-list"
    },
    "children": [
      {
        "path": "user",
        "name": "User Manger",
        "component": "/mgr/user/index",
        "meta": {
          "title": {
            "zh_CN": "用户管理",
            "en_US": "User Manger"
          }
        }
      },
      {
        "path": "card",
        "name": "ListCard",
        "component": "/list/card/index",
        "meta": {
          "title": {
            "zh_CN": "卡片列表页",
            "en_US": "Card List"
          }
        }
      },
      {
        "path": "filter",
        "name": "ListFilter",
        "component": "/list/filter/index",
        "meta": {
          "title": {
            "zh_CN": "筛选列表页",
            "en_US": "Filter List"
          }
        }
      },
      {
        "path": "tree",
        "name": "ListTree",
        "component": "/list/tree/index",
        "meta": {
          "title": {
            "zh_CN": "树状筛选列表页",
            "en_US": "Tree List"
          }
        }
      }
    ]
  },
  {
    "path": "/list",
    "name": "list",
    "component": "LAYOUT",
    "redirect": "/list/base",
    "meta": {
      "title": {
        "zh_CN": "列表页",
        "en_US": "List"
      },
      "icon": "view-list"
    },
    "children": [
      {
        "path": "base",
        "name": "ListBase",
        "component": "/list/base/index",
        "meta": {
          "title": {
            "zh_CN": "基础列表页",
            "en_US": "Base List"
          }
        }
      },
      {
        "path": "card",
        "name": "ListCard",
        "component": "/list/card/index",
        "meta": {
          "title": {
            "zh_CN": "卡片列表页",
            "en_US": "Card List"
          }
        }
      },
      {
        "path": "filter",
        "name": "ListFilter",
        "component": "/list/filter/index",
        "meta": {
          "title": {
            "zh_CN": "筛选列表页",
            "en_US": "Filter List"
          }
        }
      },
      {
        "path": "tree",
        "name": "ListTree",
        "component": "/list/tree/index",
        "meta": {
          "title": {
            "zh_CN": "树状筛选列表页",
            "en_US": "Tree List"
          }
        }
      }
    ]
  },
  {
    "path": "/form",
    "name": "form",
    "component": "LAYOUT",
    "redirect": "/form/base",
    "meta": {
      "title": {
        "zh_CN": "表单页",
        "en_US": "Form"
      },
      "icon": "edit-1"
    },
    "children": [
      {
        "path": "base",
        "name": "FormBase",
        "component": "/form/base/index",
        "meta": {
          "title": {
            "zh_CN": "基础表单页",
            "en_US": "Base Form"
          }
        }
      },
      {
        "path": "step",
        "name": "FormStep",
        "component": "/form/step/index",
        "meta": {
          "title": {
            "zh_CN": "分步表单页",
            "en_US": "Step Form"
          }
        }
      }
    ]
  },
  {
    "path": "/detail",
    "name": "detail",
    "component": "LAYOUT",
    "redirect": "/detail/base",
    "meta": {
      "title": {
        "zh_CN": "详情页",
        "en_US": "Detail"
      },
      "icon": "layers"
    },
    "children": [
      {
        "path": "base",
        "name": "DetailBase",
        "component": "/detail/base/index",
        "meta": {
          "title": {
            "zh_CN": "基础详情页",
            "en_US": "Base Detail"
          }
        }
      },
      {
        "path": "advanced",
        "name": "DetailAdvanced",
        "component": "/detail/advanced/index",
        "meta": {
          "title": {
            "zh_CN": "多卡片详情页",
            "en_US": "Card Detail"
          }
        }
      },
      {
        "path": "deploy",
        "name": "DetailDeploy",
        "component": "/detail/deploy/index",
        "meta": {
          "title": {
            "zh_CN": "数据详情页",
            "en_US": "Data Detail"
          }
        }
      },
      {
        "path": "secondary",
        "name": "DetailSecondary",
        "component": "/detail/secondary/index",
        "meta": {
          "title": {
            "zh_CN": "二级详情页",
            "en_US": "Secondary Detail"
          }
        }
      }
    ]
  },
  {
    "path": "/frame",
    "name": "Frame",
    "component": "Layout",
    "redirect": "/frame/doc",
    "meta": {
      "icon": "internet",
      "title": {
        "zh_CN": "外部页面",
        "en_US": "External"
      }
    },
    "children": [
      {
        "path": "doc",
        "name": "Doc",
        "component": "IFrame",
        "meta": {
          "frameSrc": "https://tdesign.tencent.com/starter/docs/vue-next/get-started",
          "title": {
            "zh_CN": "使用文档（内嵌）",
            "en_US": "Documentation(IFrame)"
          }
        }
      },
      {
        "path": "TDesign",
        "name": "TDesign",
        "component": "IFrame",
        "meta": {
          "frameSrc": "https://tdesign.tencent.com/vue-next/getting-started",
          "title": {
            "zh_CN": "TDesign 文档（内嵌）",
            "en_US": "TDesign (IFrame)"
          }
        }
      },
      {
        "path": "TDesign2",
        "name": "TDesign2",
        "component": "IFrame",
        "meta": {
          "frameSrc": "https://tdesign.tencent.com/vue-next/getting-started",
          "frameBlank": true,
          "title": {
            "zh_CN": "TDesign 文档（外链",
            "en_US": "TDesign Doc(Link)"
          }
        }
      }
    ]
  }
]
`