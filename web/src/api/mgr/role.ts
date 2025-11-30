import { request } from '@/utils/request';
import type { PageResp, PageReq } from '../model';

export interface Role {
  id: number;
  code: string;
  description?: string;
  status?: number;
  created_at?: number;
  updated_at?: number;
}

export interface RoleListReq extends PageReq {
  code?: string;
  status?: string; // comma separated
}

export interface RoleCreateReq {
  code: string;
  description?: string;
  status?: number;
}

export interface RoleUpdateReq {
  id: number;
  code?: string;
  description?: string;
  status?: number;
}

const Api = {
  RoleList: '/role/list',
  RoleOne: '/role/one',
  RoleCreate: '/role',
  RoleUpdate: '/role',
  RoleDelete: '/role',
  RolePermissions: '/role/permissions',
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

export function saveRolePermissions(roleId: number, menuIds: number[]) {
  return request.post({
    url: Api.RolePermissions,
    data: { role_id: roleId, menu_ids: menuIds },
  });
}

export interface RolePermissionsResp {
  menu_ids: number[];
}

export function getRolePermissions(roleId: number) {
  return request.get<RolePermissionsResp>({
    url: Api.RolePermissions,
    params: { role_id: roleId },
  });
}
