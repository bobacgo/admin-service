-- 创建索引
CREATE INDEX idx_sys_menu_parent_id ON menus (parent_id);
CREATE INDEX idx_sys_menu_path ON menus (path);
CREATE INDEX idx_sys_menu_name ON menus (NAME);

-- 插入顶级菜单数据
INSERT INTO menus (parent_id, path, NAME, component, redirect, meta, icon, sort)
VALUES
    (0, '/mgr', 'manger', 'LAYOUT', '/mgr/user', '{"title":{"zh_CN":"系统管理","en_US":"System Manger"},"icon":"view-list"}', 'view-list', 1),
    (0, '/list', 'list', 'LAYOUT', '/list/base', '{"title":{"zh_CN":"列表页","en_US":"List"},"icon":"view-list"}', 'view-list', 2),
    (0, '/form', 'form', 'LAYOUT', '/form/base', '{"title":{"zh_CN":"表单页","en_US":"Form"},"icon":"edit-1"}', 'edit-1', 3),
    (0, '/detail', 'detail', 'LAYOUT', '/detail/base', '{"title":{"zh_CN":"详情页","en_US":"Detail"},"icon":"layers"}', 'layers', 4),
    (0, '/frame', 'Frame', 'Layout', '/frame/doc', '{"icon":"internet","title":{"zh_CN":"外部页面","en_US":"External"}}', 'internet', 5);
-- 插入系统管理子菜单
INSERT INTO menus (parent_id, path, NAME, component, redirect, meta, icon, sort)
VALUES
    (1, 'user', 'User Manger', '/mgr/user/index','', '{"title":{"zh_CN":"用户管理","en_US":"User Manger"}}', 'view-list', 1),
    (1, 'card', 'ListCard', '/list/card/index','', '{"title":{"zh_CN":"卡片列表页","en_US":"Card List"}}', 'view-list', 2),
    (1, 'filter', 'ListFilter', '/list/filter/index','', '{"title":{"zh_CN":"筛选列表页","en_US":"Filter List"}}', 'view-list', 3),
    (1, 'tree', 'ListTree', '/list/tree/index','', '{"title":{"zh_CN":"树状筛选列表页","en_US":"Tree List"}}', 'view-list', 4);
-- 插入列表页子菜单
INSERT INTO menus (parent_id, path, NAME, component, redirect, meta, icon, sort)
VALUES
    (2, 'base', 'ListBase', '/list/base/index','', '{"title":{"zh_CN":"基础列表页","en_US":"Base List"}}', 'view-list', 1),
    (2, 'card', 'ListCard', '/list/card/index','', '{"title":{"zh_CN":"卡片列表页","en_US":"Card List"}}', 'view-list', 2),
    (2, 'filter', 'ListFilter', '/list/filter/index','', '{"title":{"zh_CN":"筛选列表页","en_US":"Filter List"}}', 'view-list', 3),
    (2, 'tree', 'ListTree', '/list/tree/index','', '{"title":{"zh_CN":"树状筛选列表页","en_US":"Tree List"}}', 'view-list', 4);
-- 插入表单页子菜单
INSERT INTO menus (parent_id, path, NAME, component, redirect, meta, icon, sort)
VALUES
    (3, 'base', 'FormBase', '/form/base/index','', '{"title":{"zh_CN":"基础表单页","en_US":"Base Form"}}', 'edit-1', 1),
    (3, 'step', 'FormStep', '/form/step/index','', '{"title":{"zh_CN":"分步表单页","en_US":"Step Form"}}', 'edit-1', 2);
-- 插入详情页子菜单
INSERT INTO menus (parent_id, path, NAME, component, redirect, meta, icon, sort)
VALUES
    (4, 'base', 'DetailBase', '/detail/base/index','', '{"title":{"zh_CN":"基础详情页","en_US":"Base Detail"}}', 'layers', 1),
    (4, 'advanced', 'DetailAdvanced', '/detail/advanced/index','', '{"title":{"zh_CN":"多卡片详情页","en_US":"Card Detail"}}', 'layers', 2),
    (4, 'deploy', 'DetailDeploy', '/detail/deploy/index','', '{"title":{"zh_CN":"数据详情页","en_US":"Data Detail"}}', 'layers', 3),
    (4, 'secondary', 'DetailSecondary', '/detail/secondary/index','', '{"title":{"zh_CN":"二级详情页","en_US":"Secondary Detail"}}', 'layers', 4);
-- 插入外部页面子菜单
INSERT INTO menus (parent_id, path, NAME, component, redirect, meta, icon, sort)
VALUES
    (5, 'doc', 'Doc', 'IFrame', '/frame/doc','{"frameSrc":"https://tdesign.tencent.com/starter/docs/vue-next/get-started","title":{"zh_CN":"使用文档（内嵌）","en_US":"Documentation(IFrame)"}}', 'layers', 1),
    (5, 'TDesign', 'TDesign', 'IFrame', '/frame/tdesign','{"frameSrc":"https://tdesign.tencent.com/vue-next/getting-started","title":{"zh_CN":"TDesign 文档（内嵌）","en_US":"TDesign (IFrame)"}}', 'layers', 2),
    (5, 'TDesign2', 'TDesign2', 'IFrame', '/frame/tdesign2','{"frameSrc":"https://tdesign.tencent.com/vue-next/getting-started","frameBlank":true,"title":{"zh_CN":"TDesign 文档（外链）","en_US":"TDesign Doc(Link)"}}', 'layers', 3);