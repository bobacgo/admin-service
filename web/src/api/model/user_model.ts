import { PageReq } from './model';

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
  roleIds?: number[];
}

export interface UserUpdateReq {
  id: number;
  account?: string;
  email?: string;
  phone?: string;
  status?: number;
  roleIds?: number[];
}