import type { PageReq } from './model';

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
