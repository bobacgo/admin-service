import type { MenuListResult } from '@/api/model/permissionModel';
import { request } from '@/utils/request';

const Api = {
  MenuList: '/menu/tree',
};

export function getMenuList() {
  return request.get<MenuListResult>({
    url: Api.MenuList,
  });
}
