-- Create database and tables for admin-service (MySQL)
-- Adjust DB name via `DB_NAME` env var; default used in code: admin_db

CREATE DATABASE IF NOT EXISTS `admin_db` DEFAULT CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci;
USE `admin_db`;

-- users
CREATE TABLE IF NOT EXISTS `users` (
  `id` BIGINT NOT NULL AUTO_INCREMENT,
  `account` VARCHAR(255) NOT NULL,
  `password` VARCHAR(255) NOT NULL,
  `phone` VARCHAR(255),
  `email` VARCHAR(255),
  `status` TINYINT,
  `register_at` BIGINT,
  `register_ip` VARCHAR(255),
  `login_at` BIGINT,
  `login_ip` VARCHAR(255),
  `created_at` BIGINT NOT NULL DEFAULT 0,
  `updated_at` BIGINT NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
CREATE UNIQUE INDEX `uq_users_account` ON `users` (`account`);

INSERT INTO `users` (`account`, `password`, `status`, `created_at`, `updated_at`) VALUES
('admin', 'admin', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP());

-- roles
CREATE TABLE IF NOT EXISTS `roles` (
  `id` BIGINT NOT NULL AUTO_INCREMENT,
  `code` VARCHAR(255) NOT NULL,
  `description` VARCHAR(255),
  `status` TINYINT,
  `created_at` BIGINT NOT NULL DEFAULT 0,
  `updated_at` BIGINT NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
CREATE UNIQUE INDEX `uq_roles_code` ON `roles` (`code`);

INSERT INTO `roles` (`code`, `description`, `status`, `created_at`, `updated_at`) VALUES
('super_admin', 'all power', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP());

-- user_roles (association)
CREATE TABLE IF NOT EXISTS `user_roles` (
  `id` BIGINT NOT NULL AUTO_INCREMENT,
  `r1_id` BIGINT,
  `r2_id` BIGINT,
  `created_at` BIGINT NOT NULL DEFAULT 0,
  `updated_at` BIGINT NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
CREATE UNIQUE INDEX `uq_user_roles_r1_r2` ON `user_roles` (`r1_id`, `r2_id`);
INSERT INTO `user_roles` (`r1_id`, `r2_id`, `created_at`, `updated_at`) VALUES (1, 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP());

-- menus
CREATE TABLE IF NOT EXISTS `menus` (
  `id` BIGINT NOT NULL AUTO_INCREMENT,
  `parent_id` BIGINT,
  `path` VARCHAR(255) NOT NULL,
  `name` VARCHAR(255) NOT NULL,
  `component` VARCHAR(255),
  `redirect` VARCHAR(255),
  `meta` VARCHAR(1024),
  `icon` VARCHAR(255),
  `sort` BIGINT,
  `created_at` BIGINT NOT NULL DEFAULT 0,
  `updated_at` BIGINT NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
CREATE UNIQUE INDEX `uq_menus_path` ON `menus` (`path`);
CREATE INDEX `idx_sys_menu_parent_id` ON `menus` (`parent_id`);
CREATE INDEX `idx_sys_menu_path` ON `menus` (`path`);
CREATE INDEX `idx_sys_menu_name` ON `menus` (`name`);

-- menu_roles (association)
CREATE TABLE IF NOT EXISTS `menu_roles` (
  `id` BIGINT NOT NULL AUTO_INCREMENT,
  `r1_id` BIGINT,
  `r2_id` BIGINT,
  `created_at` BIGINT NOT NULL DEFAULT 0,
  `updated_at` BIGINT NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
CREATE UNIQUE INDEX `uq_menu_roles_r1_r2` ON `menu_roles` (`r1_id`, `r2_id`);

-- i18n
CREATE TABLE IF NOT EXISTS `i18n` (
  `id` BIGINT NOT NULL AUTO_INCREMENT,
  `class` VARCHAR(255) NOT NULL,
  `lang` VARCHAR(255) NOT NULL,
  `key` VARCHAR(255) NOT NULL,
  `value` VARCHAR(1024) NOT NULL,
  `created_at` BIGINT NOT NULL DEFAULT 0,
  `updated_at` BIGINT NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
CREATE UNIQUE INDEX `uq_i18n_key_lang` ON `i18n` (`key`, `lang`);

-- seed menus from original db.sql (kept minimal here)
INSERT INTO `menus` (`parent_id`, `path`, `name`, `component`, `redirect`, `meta`, `icon`, `sort`, `created_at`, `updated_at`) VALUES
(0, '/mgr', 'manger', 'LAYOUT', '/mgr/user', '{"title":{"zh_CN":"系统管理","en_US":"System Manger"},"icon":"view-list"}', 'view-list', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(0, '/list', 'list', 'LAYOUT', '/list/base', '{"title":{"zh_CN":"列表页","en_US":"List"},"icon":"view-list"}', 'view-list', 2, UNIX_TIMESTAMP(), UNIX_TIMESTAMP());

-- End of migration
