import { request } from '@/utils/request';
import type { PageResp } from './model/model';
import type { MenuItem, MenuListReq, MenuCreateReq, MenuUpdateReq } from './model/menuModel';

const Api = {
  MenuList: '/menu/list',
  MenuGet: '/menu/get',
  MenuAdd: '/menu/add',
  MenuUpdate: '/menu/update',
  MenuDelete: '/menu/delete',
};

export function getMenuList(params: MenuListReq) {
  return request.get<PageResp<MenuItem>>({
    url: Api.MenuList,
    params,
  });
}

export function getMenu(id: number) {
  return request.get<MenuItem>({
    url: Api.MenuGet,
    params: { id },
  });
}

export function addMenu(data: MenuCreateReq) {
  return request.post({
    url: Api.MenuAdd,
    data,
  });
}

export function updateMenu(data: MenuUpdateReq) {
  return request.put({
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
