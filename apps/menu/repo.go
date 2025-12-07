package menu

import (
	"context"
	"time"

	"github.com/bobacgo/admin-service/apps/repo"
	"github.com/bobacgo/admin-service/apps/repo/data"
	"github.com/bobacgo/admin-service/apps/repo/model"
	. "github.com/bobacgo/orm"
)

type MenuRepo struct {
	clt *data.Client
}

func NewMenuRepo(clt *data.Client) *MenuRepo {
	return &MenuRepo{clt: clt}
}

func (r *MenuRepo) Find(ctx context.Context) ([]*Menu, error) {
	list := make([]*Menu, 0)
	if err := SELECT2(&list).FROM(MenuTable).ORDER_BY(repo.DESC(model.Id)).Query(ctx, r.clt.DB); err != nil {
		return nil, err
	}
	return list, nil
}

func (r *MenuRepo) FindOne(ctx context.Context, id int64) (*Menu, error) {
	row := new(Menu)
	if err := SELECT1(row).FROM(MenuTable).WHERE(M{repo.AND(model.Id): id}).Query(ctx, r.clt.DB); err != nil {
		return nil, err
	}
	return row, nil
}

func (r *MenuRepo) Create(ctx context.Context, row *Menu) error {
	id, err := INSERT(row).INTO(MenuTable).Omit(model.Id).Exec(ctx, r.clt.DB)
	row.ID = id
	return err
}

func (r *MenuRepo) Update(ctx context.Context, row *Menu) error {
	_, err := UPDATE(MenuTable).SET1(row).WHERE(M{repo.AND(model.Id): row.ID}).Omit(model.Id).Exec(ctx, r.clt.DB)
	return err
}

func (r *MenuRepo) Delete(ctx context.Context, ids string) error {
	_, err := DELETE().FROM(MenuTable).WHERE(M{repo.AND_IN(model.Id): ids}).Exec(ctx, r.clt.DB)
	return err
}

// RemoveRoleIdFromAllMenus 从所有菜单的role_ids中移除该角色
func (r *MenuRepo) RemoveRoleIdFromAllMenus(ctx context.Context, roleId string) error {
	// 获取所有包含该角色的菜单
	var menus []*Menu
	where := M{repo.AND_LIKE("role_ids"): "%" + roleId + "%"}
	if err := SELECT2(&menus).FROM(MenuTable).WHERE(where).Query(ctx, r.clt.DB); err != nil {
		return err
	}

	// 从每个菜单的role_ids中移除该角色
	for _, menu := range menus {
		menu.RemoveRoleId(roleId)
		menu.UpdatedAt = time.Now().Unix()
		if err := r.Update(ctx, menu); err != nil {
			return err
		}
	}

	return nil
}

// AddRoleIdToMenus 将角色添加到指定菜单的role_ids中
func (r *MenuRepo) AddRoleIdToMenus(ctx context.Context, roleId string, menuIds []int64) error {
	for _, menuId := range menuIds {
		menu, err := r.FindOne(ctx, menuId)
		if err != nil {
			return err
		}

		menu.AddRoleId(roleId)
		menu.UpdatedAt = time.Now().Unix()
		if err := r.Update(ctx, menu); err != nil {
			return err
		}
	}

	return nil
}

// GetMenuIdsByRoleId 根据角色ID获取菜单ID列表
func (r *MenuRepo) GetMenuIdsByRoleId(ctx context.Context, roleId string) ([]int64, error) {
	var menus []*Menu
	where := M{repo.AND_LIKE("role_ids"): "%" + roleId + "%"}
	if err := SELECT2(&menus).FROM(MenuTable).WHERE(where).Query(ctx, r.clt.DB); err != nil {
		return nil, err
	}

	var ids []int64
	for _, menu := range menus {
		// 再次确认role_ids中确实包含该角色（因为LIKE可能会匹配到部分字符串）
		if menu.HasRoleId(roleId) {
			ids = append(ids, menu.ID)
		}
	}
	return ids, nil
}
