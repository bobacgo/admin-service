-- Insert only children for /mgr
USE `admin_db`;

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

-- End of /mgr children insert
