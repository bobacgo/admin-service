import { request } from '@/utils/request';
import type { PageResp } from './model/model';
import type { MenuItem, MenuListReq, MenuCreateReq, MenuUpdateReq } from './model/menuModel';

const Api = {
  MenuList: '/menu/list',
  MenuOne: '/menu/one',
  MenuTree: '/menu/tree',
  MenuCreate: '/menu',
  MenuUpdate: '/menu',
  MenuDelete: '/menu',
};

export function getMenuList(params: MenuListReq) {
  return request.get<PageResp<MenuItem>>({
    url: Api.MenuList,
    params,
  });
}

export function getMenu(id: number) {
  return request.get<MenuItem>({
    url: Api.MenuOne,
    params: { id },
  });
}

export function getMenuTree() {
  return request.get<MenuItem[]>({
    url: Api.MenuTree,
  });
}

export function addMenu(data: MenuCreateReq) {
  return request.post<MenuItem>({
    url: Api.MenuCreate,
    data,
  });
}

export function updateMenu(data: MenuUpdateReq) {
  return request.put<MenuItem>({
    url: Api.MenuUpdate,
    data,
  });
}

export function deleteMenu(ids: number[]) {
  const params = { ids: ids.join(',') };
  return request.delete({
    url: Api.MenuDelete,
    params,
  });
}
