import { request } from '@/utils/request';
import { LoginResp, UserListReq, UserAddReq, UserUpdateReq, User } from './model/user_model';
import { IdsReq, PageResp } from './model/model';

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