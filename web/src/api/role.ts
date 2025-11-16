import { request } from '@/utils/request';
import type { PageResp } from './model/model';
import type { Role, RoleListReq, RoleCreateReq, RoleUpdateReq } from './model/roleModel';

const Api = {
  RoleList: '/role/list',
  RoleGet: '/role/get',
  RoleAdd: '/role/add',
  RoleUpdate: '/role/update',
  RoleDelete: '/role/delete',
};

export function getRoleList(params: RoleListReq) {
  return request.get<PageResp<Role>>({
    url: Api.RoleList,
    params,
  });
}

export function getRole(id: number) {
  return request.get<Role>({
    url: Api.RoleGet,
    params: { id },
  });
}

export function addRole(data: RoleCreateReq) {
  return request.post({
    url: Api.RoleAdd,
    data,
  });
}

export function updateRole(data: RoleUpdateReq) {
  return request.put({
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
