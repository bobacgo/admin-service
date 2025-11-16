import type { PageReq } from './model';

export interface MenuMeta {
  title?: Record<string, string>;
  [key: string]: any;
}

export interface MenuItem {
  id: number;
  parent_id?: number;
  path: string;
  name: string;
  component?: string;
  redirect?: string;
  meta?: MenuMeta;
  icon?: string;
  sort?: number;
  children?: MenuItem[];
  created_at?: number;
  updated_at?: number;
}

export interface MenuListReq extends PageReq {
  path?: string;
  name?: string;
}

export interface MenuCreateReq {
  path: string;
  name: string;
  component?: string;
  redirect?: string;
  meta?: string;
  icon?: string;
}

export interface MenuUpdateReq {
  id: number;
  path?: string;
  name?: string;
  component?: string;
  redirect?: string;
  meta?: string;
  icon?: string;
}
