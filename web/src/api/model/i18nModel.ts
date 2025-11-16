import type { PageReq } from './model';

export interface I18nItem {
  id: number;
  class?: string;
  lang: string;
  key: string;
  value: string;
  created_at?: number;
  updated_at?: number;
}

export interface I18nListReq extends PageReq {
  class?: string;
  lang?: string;
  key?: string;
}

export interface I18nCreateReq {
  class?: string;
  lang: string;
  key: string;
  value: string;
}

export interface I18nUpdateReq {
  id: number;
  class?: string;
  lang?: string;
  value?: string;
}
