-- ====================================================================
-- 初始化菜单数据
-- ====================================================================
-- 一级菜单
INSERT INTO `menus` (`id`, `parent_id`, `path`, `name`, `component`, `redirect`, `meta`, `icon`, `sort`, `created_at`, `updated_at`) VALUES
(1, 0, '/mgr', 'manger', 'LAYOUT', '/mgr/user', '{"title":{"zh_CN":"系统管理","en_US":"System Manager"},"icon":"view-list"}', 'view-list', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(2, 0, '/list', 'list', 'LAYOUT', '/list/base', '{"title":{"zh_CN":"列表页","en_US":"List"},"icon":"view-list"}', 'view-list', 2, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(3, 0, '/form', 'form', 'LAYOUT', '/form/base', '{"title":{"zh_CN":"表单页","en_US":"Form"},"icon":"edit-1"}', 'edit-1', 3, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(4, 0, '/detail', 'detail', 'LAYOUT', '/detail/base', '{"title":{"zh_CN":"详情页","en_US":"Detail"},"icon":"layers"}', 'layers', 4, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(5, 0, '/frame', 'frame', 'LAYOUT', '/frame/doc', '{"title":{"zh_CN":"外部页面","en_US":"External"},"icon":"internet"}', 'internet', 5, UNIX_TIMESTAMP(), UNIX_TIMESTAMP());

-- 系统管理子菜单
INSERT INTO `menus` (`parent_id`, `path`, `name`, `component`, `redirect`, `meta`, `icon`, `sort`, `created_at`, `updated_at`) VALUES
(1, '/mgr/user', 'UserManager', '/mgr/user/index', '', '{"title":{"zh_CN":"用户管理","en_US":"User Manager"}}', 'user', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(1, '/mgr/role', 'RoleManager', '/mgr/role/index', '', '{"title":{"zh_CN":"角色管理","en_US":"Role Manager"}}', 'usergroup', 2, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(1, '/mgr/menu', 'MenuManager', '/mgr/menu/index', '', '{"title":{"zh_CN":"菜单管理","en_US":"Menu Manager"}}', 'menu-fold', 3, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(1, '/mgr/i18n', 'I18nManager', '/mgr/i18n/index', '', '{"title":{"zh_CN":"国际化管理","en_US":"I18n Manager"}}', 'global', 4, UNIX_TIMESTAMP(), UNIX_TIMESTAMP());

-- 列表页子菜单
INSERT INTO `menus` (`parent_id`, `path`, `name`, `component`, `redirect`, `meta`, `icon`, `sort`, `created_at`, `updated_at`) VALUES
(2, '/list/base', 'ListBase', '/list/base/index', '', '{"title":{"zh_CN":"基础列表页","en_US":"Base List"}}', 'view-list', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(2, '/list/card', 'ListCard', '/list/card/index', '', '{"title":{"zh_CN":"卡片列表页","en_US":"Card List"}}', 'view-module', 2, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(2, '/list/filter', 'ListFilter', '/list/filter/index', '', '{"title":{"zh_CN":"筛选列表页","en_US":"Filter List"}}', 'filter', 3, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(2, '/list/tree', 'ListTree', '/list/tree/index', '', '{"title":{"zh_CN":"树状筛选列表页","en_US":"Tree List"}}', 'tree', 4, UNIX_TIMESTAMP(), UNIX_TIMESTAMP());

-- 表单页子菜单
INSERT INTO `menus` (`parent_id`, `path`, `name`, `component`, `redirect`, `meta`, `icon`, `sort`, `created_at`, `updated_at`) VALUES
(3, '/form/base', 'FormBase', '/form/base/index', '', '{"title":{"zh_CN":"基础表单页","en_US":"Base Form"}}', 'edit', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(3, '/form/step', 'FormStep', '/form/step/index', '', '{"title":{"zh_CN":"分步表单页","en_US":"Step Form"}}', 'swap', 2, UNIX_TIMESTAMP(), UNIX_TIMESTAMP());

-- 详情页子菜单
INSERT INTO `menus` (`parent_id`, `path`, `name`, `component`, `redirect`, `meta`, `icon`, `sort`, `created_at`, `updated_at`) VALUES
(4, '/detail/base', 'DetailBase', '/detail/base/index', '', '{"title":{"zh_CN":"基础详情页","en_US":"Base Detail"}}', 'layers', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(4, '/detail/advanced', 'DetailAdvanced', '/detail/advanced/index', '', '{"title":{"zh_CN":"多卡片详情页","en_US":"Card Detail"}}', 'dashboard', 2, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(4, '/detail/deploy', 'DetailDeploy', '/detail/deploy/index', '', '{"title":{"zh_CN":"数据详情页","en_US":"Data Detail"}}', 'chart-bar', 3, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(4, '/detail/secondary', 'DetailSecondary', '/detail/secondary/index', '', '{"title":{"zh_CN":"二级详情页","en_US":"Secondary Detail"}}', 'jump', 4, UNIX_TIMESTAMP(), UNIX_TIMESTAMP());

-- 外部页面子菜单
INSERT INTO `menus` (`parent_id`, `path`, `name`, `component`, `redirect`, `meta`, `icon`, `sort`, `created_at`, `updated_at`) VALUES
(5, '/frame/doc', 'Doc', 'IFrame', '', '{"frameSrc":"https://tdesign.tencent.com/starter/docs/vue-next/get-started","title":{"zh_CN":"使用文档（内嵌）","en_US":"Documentation(IFrame)"}}', 'help-circle', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(5, '/frame/tdesign', 'TDesign', 'IFrame', '', '{"frameSrc":"https://tdesign.tencent.com/vue-next/getting-started","title":{"zh_CN":"TDesign 文档（内嵌）","en_US":"TDesign (IFrame)"}}', 'logo-github', 2, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(5, '/frame/tdesign-ext', 'TDesignExt', 'IFrame', '', '{"frameSrc":"https://tdesign.tencent.com/vue-next/getting-started","frameBlank":true,"title":{"zh_CN":"TDesign 文档（外链）","en_US":"TDesign Doc(Link)"}}', 'link', 3, UNIX_TIMESTAMP(), UNIX_TIMESTAMP());