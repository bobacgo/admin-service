-- Insert menus converted from menu.json attachment
USE `admin_db`;

-- Insert top-level menus (id chosen by AUTO_INCREMENT). Use ON DUPLICATE KEY to be idempotent by `path` unique index.
INSERT INTO `menus` (parent_id, path, name, component, redirect, meta, icon, sort, created_at, updated_at)
VALUES
  (0, '/mgr', 'manger', 'LAYOUT', '/mgr/user', '{"title":{"zh_CN":"系统管理","en_US":"System Manger"},"icon":"view-list"}', 'view-list', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
  (0, '/list', 'list', 'LAYOUT', '/list/base', '{"title":{"zh_CN":"列表页","en_US":"List"},"icon":"view-list"}', 'view-list', 2, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
  (0, '/form', 'form', 'LAYOUT', '/form/base', '{"title":{"zh_CN":"表单页","en_US":"Form"},"icon":"edit-1"}', 'edit-1', 3, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
  (0, '/detail', 'detail', 'LAYOUT', '/detail/base', '{"title":{"zh_CN":"详情页","en_US":"Detail"},"icon":"layers"}', 'layers', 4, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
  (0, '/frame', 'Frame', 'Layout', '/frame/doc', '{"icon":"internet","title":{"zh_CN":"外部页面","en_US":"External"}}', 'internet', 5, UNIX_TIMESTAMP(), UNIX_TIMESTAMP())
ON DUPLICATE KEY UPDATE updated_at = VALUES(updated_at);

-- Insert children for /mgr
INSERT INTO `menus` (parent_id, path, name, component, redirect, meta, icon, sort, created_at, updated_at)
VALUES
  ((SELECT id FROM (SELECT id FROM menus WHERE path = '/mgr' LIMIT 1) AS p), 'user', 'User Manger', '/mgr/user/index', '', '{"title":{"zh_CN":"用户管理","en_US":"User Manger"}}', 'view-list', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
  ((SELECT id FROM (SELECT id FROM menus WHERE path = '/mgr' LIMIT 1) AS p), 'card', 'ListCard', '/list/card/index', '', '{"title":{"zh_CN":"卡片列表页","en_US":"Card List"}}', 'view-list', 2, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
  ((SELECT id FROM (SELECT id FROM menus WHERE path = '/mgr' LIMIT 1) AS p), 'filter', 'ListFilter', '/list/filter/index', '', '{"title":{"zh_CN":"筛选列表页","en_US":"Filter List"}}', 'view-list', 3, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
  ((SELECT id FROM (SELECT id FROM menus WHERE path = '/mgr' LIMIT 1) AS p), 'tree', 'ListTree', '/list/tree/index', '', '{"title":{"zh_CN":"树状筛选列表页","en_US":"Tree List"}}', 'view-list', 4, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
  ((SELECT id FROM (SELECT id FROM menus WHERE path = '/mgr' LIMIT 1) AS p), 'menu', 'Menu Manger', '/mgr/menu/index', '', '{"title":{"zh_CN":"菜单管理","en_US":"Menu Manger"}}', 'view-list', 5, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
  ((SELECT id FROM (SELECT id FROM menus WHERE path = '/mgr' LIMIT 1) AS p), 'role', 'Role Manger', '/mgr/role/index', '', '{"title":{"zh_CN":"角色管理","en_US":"Role Manger"}}', 'user', 6, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
  ((SELECT id FROM (SELECT id FROM menus WHERE path = '/mgr' LIMIT 1) AS p), 'i18n', 'I18n Manger', '/mgr/i18n/index', '', '{"title":{"zh_CN":"多语言管理","en_US":"I18n Manger"}}', 'language', 7, UNIX_TIMESTAMP(), UNIX_TIMESTAMP())
ON DUPLICATE KEY UPDATE updated_at = VALUES(updated_at);

-- Insert children for /list
INSERT INTO `menus` (parent_id, path, name, component, redirect, meta, icon, sort, created_at, updated_at)
VALUES
  ((SELECT id FROM (SELECT id FROM menus WHERE path = '/list' LIMIT 1) AS p), 'base', 'ListBase', '/list/base/index', '', '{"title":{"zh_CN":"基础列表页","en_US":"Base List"}}', 'view-list', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
  ((SELECT id FROM (SELECT id FROM menus WHERE path = '/list' LIMIT 1) AS p), 'card', 'ListCard', '/list/card/index', '', '{"title":{"zh_CN":"卡片列表页","en_US":"Card List"}}', 'view-list', 2, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
  ((SELECT id FROM (SELECT id FROM menus WHERE path = '/list' LIMIT 1) AS p), 'filter', 'ListFilter', '/list/filter/index', '', '{"title":{"zh_CN":"筛选列表页","en_US":"Filter List"}}', 'view-list', 3, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
  ((SELECT id FROM (SELECT id FROM menus WHERE path = '/list' LIMIT 1) AS p), 'tree', 'ListTree', '/list/tree/index', '', '{"title":{"zh_CN":"树状筛选列表页","en_US":"Tree List"}}', 'view-list', 4, UNIX_TIMESTAMP(), UNIX_TIMESTAMP())
ON DUPLICATE KEY UPDATE updated_at = VALUES(updated_at);

-- Insert children for /form
INSERT INTO `menus` (parent_id, path, name, component, redirect, meta, icon, sort, created_at, updated_at)
VALUES
  ((SELECT id FROM (SELECT id FROM menus WHERE path = '/form' LIMIT 1) AS p), 'base', 'FormBase', '/form/base/index', '', '{"title":{"zh_CN":"基础表单页","en_US":"Base Form"}}', 'edit-1', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
  ((SELECT id FROM (SELECT id FROM menus WHERE path = '/form' LIMIT 1) AS p), 'step', 'FormStep', '/form/step/index', '', '{"title":{"zh_CN":"分步表单页","en_US":"Step Form"}}', 'edit-1', 2, UNIX_TIMESTAMP(), UNIX_TIMESTAMP())
ON DUPLICATE KEY UPDATE updated_at = VALUES(updated_at);

-- Insert children for /detail
INSERT INTO `menus` (parent_id, path, name, component, redirect, meta, icon, sort, created_at, updated_at)
VALUES
  ((SELECT id FROM (SELECT id FROM menus WHERE path = '/detail' LIMIT 1) AS p), 'base', 'DetailBase', '/detail/base/index', '', '{"title":{"zh_CN":"基础详情页","en_US":"Base Detail"}}', 'layers', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
  ((SELECT id FROM (SELECT id FROM menus WHERE path = '/detail' LIMIT 1) AS p), 'advanced', 'DetailAdvanced', '/detail/advanced/index', '', '{"title":{"zh_CN":"多卡片详情页","en_US":"Card Detail"}}', 'layers', 2, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
  ((SELECT id FROM (SELECT id FROM menus WHERE path = '/detail' LIMIT 1) AS p), 'deploy', 'DetailDeploy', '/detail/deploy/index', '', '{"title":{"zh_CN":"数据详情页","en_US":"Data Detail"}}', 'layers', 3, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
  ((SELECT id FROM (SELECT id FROM menus WHERE path = '/detail' LIMIT 1) AS p), 'secondary', 'DetailSecondary', '/detail/secondary/index', '', '{"title":{"zh_CN":"二级详情页","en_US":"Secondary Detail"}}', 'layers', 4, UNIX_TIMESTAMP(), UNIX_TIMESTAMP())
ON DUPLICATE KEY UPDATE updated_at = VALUES(updated_at);

-- Insert children for /frame
INSERT INTO `menus` (parent_id, path, name, component, redirect, meta, icon, sort, created_at, updated_at)
VALUES
  ((SELECT id FROM (SELECT id FROM menus WHERE path = '/frame' LIMIT 1) AS p), 'doc', 'Doc', 'IFrame', '/frame/doc', '{"frameSrc":"https://tdesign.tencent.com/starter/docs/vue-next/get-started","title":{"zh_CN":"使用文档（内嵌）","en_US":"Documentation(IFrame)"}}', 'layers', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
  ((SELECT id FROM (SELECT id FROM menus WHERE path = '/frame' LIMIT 1) AS p), 'TDesign', 'TDesign', 'IFrame', '/frame/tdesign', '{"frameSrc":"https://tdesign.tencent.com/vue-next/getting-started","title":{"zh_CN":"TDesign 文档（内嵌）","en_US":"TDesign (IFrame)"}}', 'layers', 2, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
  ((SELECT id FROM (SELECT id FROM menus WHERE path = '/frame' LIMIT 1) AS p), 'TDesign2', 'TDesign2', 'IFrame', '/frame/tdesign2', '{"frameSrc":"https://tdesign.tencent.com/vue-next/getting-started","frameBlank":true,"title":{"zh_CN":"TDesign 文档（外链）","en_US":"TDesign Doc(Link)"}}', 'layers', 3, UNIX_TIMESTAMP(), UNIX_TIMESTAMP())
ON DUPLICATE KEY UPDATE updated_at = VALUES(updated_at);

-- End of insert script
