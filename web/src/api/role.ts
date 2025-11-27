import { request } from '@/utils/request';
import type { PageResp } from './model/model';
import type { Role, RoleListReq, RoleCreateReq, RoleUpdateReq } from './model/roleModel';

const Api = {
  RoleList: '/role/list',
  RoleOne: '/role/one',
  RoleCreate: '/role',
  RoleUpdate: '/role',
  RoleDelete: '/role',
};

export function getRoleList(params: RoleListReq) {
  return request.get<PageResp<Role>>({
    url: Api.RoleList,
    params,
  });
}

export function getRole(id: number) {
  return request.get<Role>({
    url: Api.RoleOne,
    params: { id },
  });
}

export function addRole(data: RoleCreateReq) {
  return request.post<Role>({
    url: Api.RoleCreate,
    data,
  });
}

export function updateRole(data: RoleUpdateReq) {
  return request.put<Role>({
    url: Api.RoleUpdate,
    data,
  });
}

export function deleteRole(ids: number[]) {
  const params = { ids: ids.join(',') };
  return request.delete({
    url: Api.RoleDelete,
    params,
  });
}
