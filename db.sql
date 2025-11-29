-- ====================================================================
-- Admin Service Database Schema
-- ====================================================================
-- 数据库配置：通过环境变量 DB_NAME 调整，默认值：admin_db
-- 字符集：utf8mb4，支持完整的 Unicode 字符（包括 emoji）
-- 时间戳：使用 BIGINT 存储 Unix 时间戳（秒级）

CREATE DATABASE IF NOT EXISTS `admin_db` 
  DEFAULT CHARACTER SET = utf8mb4 
  COLLATE = utf8mb4_general_ci;

USE `admin_db`;

-- ====================================================================
-- 用户表 (users)
-- ====================================================================
CREATE TABLE IF NOT EXISTS `users` (
  `id` BIGINT NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  `account` VARCHAR(100) NOT NULL COMMENT '用户账号',
  `password` VARCHAR(255) NOT NULL COMMENT '用户密码（建议存储哈希值）',
  `phone` VARCHAR(20) DEFAULT NULL COMMENT '手机号码',
  `email` VARCHAR(100) DEFAULT NULL COMMENT '电子邮箱',
  `status` TINYINT NOT NULL DEFAULT 1 COMMENT '用户状态：1=正常 2=禁用',
  `register_at` BIGINT DEFAULT NULL COMMENT '注册时间（Unix时间戳）',
  `register_ip` VARCHAR(50) DEFAULT NULL COMMENT '注册IP地址（支持IPv6）',
  `login_at` BIGINT DEFAULT NULL COMMENT '最后登录时间（Unix时间戳）',
  `login_ip` VARCHAR(50) DEFAULT NULL COMMENT '最后登录IP地址',
  `created_at` BIGINT NOT NULL DEFAULT 0 COMMENT '创建时间（Unix时间戳）',
  `updated_at` BIGINT NOT NULL DEFAULT 0 COMMENT '更新时间（Unix时间戳）',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uq_users_account` (`account`),
  KEY `idx_users_status` (`status`),
  KEY `idx_users_email` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=1000 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户表';

-- 初始化管理员账户
INSERT INTO `users` (`account`, `password`, `status`, `created_at`, `updated_at`) VALUES
('admin', 'admin', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP());

-- ====================================================================
-- 角色表 (roles)
-- ====================================================================
CREATE TABLE IF NOT EXISTS `roles` (
  `id` BIGINT NOT NULL AUTO_INCREMENT COMMENT '角色ID',
  `code` VARCHAR(50) NOT NULL COMMENT '角色编码（唯一标识）',
  `description` VARCHAR(255) DEFAULT NULL COMMENT '角色描述',
  `status` TINYINT NOT NULL DEFAULT 1 COMMENT '角色状态：1=启用 2=禁用',
  `created_at` BIGINT NOT NULL DEFAULT 0 COMMENT '创建时间（Unix时间戳）',
  `updated_at` BIGINT NOT NULL DEFAULT 0 COMMENT '更新时间（Unix时间戳）',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uq_roles_code` (`code`),
  KEY `idx_roles_status` (`status`)
) ENGINE=InnoDB AUTO_INCREMENT=1000 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='角色表';

-- 初始化超级管理员角色
INSERT INTO `roles` (`code`, `description`, `status`, `created_at`, `updated_at`) VALUES
('super_admin', '超级管理员，拥有所有权限', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP());

-- ====================================================================
-- 用户角色关联表 (user_roles)
-- ====================================================================
CREATE TABLE IF NOT EXISTS `user_roles` (
  `id` BIGINT NOT NULL AUTO_INCREMENT COMMENT '关联ID',
  `user_id` BIGINT NOT NULL COMMENT '用户ID',
  `role_id` BIGINT NOT NULL COMMENT '角色ID',
  `created_at` BIGINT NOT NULL DEFAULT 0 COMMENT '创建时间（Unix时间戳）',
  `updated_at` BIGINT NOT NULL DEFAULT 0 COMMENT '更新时间（Unix时间戳）',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uq_user_roles` (`user_id`, `role_id`),
  KEY `idx_user_roles_user_id` (`user_id`),
  KEY `idx_user_roles_role_id` (`role_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户角色关联表';

-- 初始化管理员角色关联
INSERT INTO `user_roles` (`user_id`, `role_id`, `created_at`, `updated_at`) VALUES 
(1, 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP());

-- ====================================================================
-- 菜单表 (menus)
-- ====================================================================
CREATE TABLE IF NOT EXISTS `menus` (
  `id` BIGINT NOT NULL AUTO_INCREMENT COMMENT '菜单ID',
  `parent_id` BIGINT NOT NULL DEFAULT 0 COMMENT '父菜单ID，0表示顶级菜单',
  `path` VARCHAR(255) NOT NULL COMMENT '路由路径',
  `name` VARCHAR(100) NOT NULL COMMENT '路由名称',
  `component` VARCHAR(255) DEFAULT NULL COMMENT '组件路径',
  `redirect` VARCHAR(255) DEFAULT NULL COMMENT '重定向路径',
  `meta` VARCHAR(1024) DEFAULT NULL COMMENT '菜单元数据（JSON格式）',
  `icon` VARCHAR(50) DEFAULT NULL COMMENT '菜单图标',
  `sort` INT NOT NULL DEFAULT 0 COMMENT '排序号，数字越小越靠前',
  `status` TINYINT NOT NULL DEFAULT 1 COMMENT '菜单状态：1=显示 2=隐藏',
  `created_at` BIGINT NOT NULL DEFAULT 0 COMMENT '创建时间（Unix时间戳）',
  `updated_at` BIGINT NOT NULL DEFAULT 0 COMMENT '更新时间（Unix时间戳）',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uq_menus_path` (`path`),
  KEY `idx_menus_parent_id` (`parent_id`),
  KEY `idx_menus_name` (`name`),
  KEY `idx_menus_status_sort` (`status`, `sort`)
) ENGINE=InnoDB AUTO_INCREMENT=1000 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='菜单表';

-- ====================================================================
-- 菜单角色关联表 (menu_roles)
-- ====================================================================
CREATE TABLE IF NOT EXISTS `menu_roles` (
  `id` BIGINT NOT NULL AUTO_INCREMENT COMMENT '关联ID',
  `menu_id` BIGINT NOT NULL COMMENT '菜单ID',
  `role_id` BIGINT NOT NULL COMMENT '角色ID',
  `created_at` BIGINT NOT NULL DEFAULT 0 COMMENT '创建时间（Unix时间戳）',
  `updated_at` BIGINT NOT NULL DEFAULT 0 COMMENT '更新时间（Unix时间戳）',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uq_menu_roles` (`menu_id`, `role_id`),
  KEY `idx_menu_roles_menu_id` (`menu_id`),
  KEY `idx_menu_roles_role_id` (`role_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='菜单角色关联表';

-- ====================================================================
-- 国际化表 (i18n)
-- ====================================================================
CREATE TABLE IF NOT EXISTS `i18n` (
  `id` BIGINT NOT NULL AUTO_INCREMENT COMMENT '翻译ID',
  `class` VARCHAR(100) NOT NULL COMMENT '分类（如：menu, button, message）',
  `lang` VARCHAR(10) NOT NULL COMMENT '语言代码（如：zh_CN, en_US）',
  `key` VARCHAR(255) NOT NULL COMMENT '翻译键',
  `value` VARCHAR(1024) NOT NULL COMMENT '翻译值',
  `created_at` BIGINT NOT NULL DEFAULT 0 COMMENT '创建时间（Unix时间戳）',
  `updated_at` BIGINT NOT NULL DEFAULT 0 COMMENT '更新时间（Unix时间戳）',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uq_i18n_key_lang` (`key`, `lang`),
  KEY `idx_i18n_class` (`class`),
  KEY `idx_i18n_lang` (`lang`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='国际化翻译表';