import { request } from '@/utils/request';
import { IdsReq, PageResp, PageReq } from '../model';

export interface LoginResp {
    token: string;
}

export interface User {
    id: number;
    created_at: number;
    updated_at: number;
    account: string;
    phone: string;
    email: string;
    status: number;
    register_at: number;
    register_ip: string;
    login_at: number;
    login_ip: string;
    role_codes: string;
    operator?: string;
}

export interface UserListReq extends PageReq {
  keyword?: string;
  status?: number;
}

export interface UserAddReq {
  account: string;
  password: string;
  email?: string;
  phone?: string;
  status: number;
  role_codes?: string;
  operator?: string;
}

export interface UserUpdateReq {
  id: number;
  account?: string;
  email?: string;
  phone?: string;
  status?: number;
  role_codes?: string;
  operator?: string;
}

const Api = {
  Login: '/login',
  UserInfo: '/user-info',
  Logout: '/logout',
  UserList: '/user/list',
  UserOne: '/user/one',
  UserCreate: '/user',
  UserUpdate: '/user',
  UserDelete: '/user',
};

// 用户登录相关接口
export function PostLogin(req: Record<string, unknown>) {
  return request.post<LoginResp>({
    url: Api.Login,
    data: req,
  });
}

export function GetUserInfo() {
  return request.get<LoginResp>({
    url: Api.UserInfo,
  });
}

export function PostLogout() {
  return request.post({
    url: Api.Logout,
  });
}

// 用户管理相关接口

// 获取用户列表
export function getUserList(params: UserListReq) {
  return request.get<PageResp<User>>({
    url: Api.UserList,
    params,
  });
}

// 获取单个用户
export function getUser(id: number) {
  return request.get<User>({
    url: Api.UserOne,
    params: { id },
  });
}

// 添加用户
export function addUser(data: UserAddReq) {
  return request.post<User>({
    url: Api.UserCreate,
    data,
  });
}

// 更新用户
export function updateUser(data: UserUpdateReq) {
  return request.put<User>({
    url: Api.UserUpdate,
    data,
  });
}

// 删除用户
export function deleteUser(ids: number[]) {
  const params = { ids: ids.join(',') };
  return request.delete({
    url: Api.UserDelete,
    params,
  });
}