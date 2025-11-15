-- Migration: replace NULLs with defaults and set NOT NULL + DEFAULT on all tables
USE `admin_db`;

-- USERS: replace NULLs
UPDATE `users` SET `phone` = '' WHERE `phone` IS NULL;
UPDATE `users` SET `email` = '' WHERE `email` IS NULL;
UPDATE `users` SET `register_ip` = '' WHERE `register_ip` IS NULL;
UPDATE `users` SET `login_ip` = '' WHERE `login_ip` IS NULL;
UPDATE `users` SET `status` = 0 WHERE `status` IS NULL;
UPDATE `users` SET `register_at` = 0 WHERE `register_at` IS NULL;
UPDATE `users` SET `login_at` = 0 WHERE `login_at` IS NULL;
UPDATE `users` SET `created_at` = COALESCE(`created_at`, 0);
UPDATE `users` SET `updated_at` = COALESCE(`updated_at`, 0);

ALTER TABLE `users`
  MODIFY `account` VARCHAR(255) NOT NULL DEFAULT '',
  MODIFY `password` VARCHAR(255) NOT NULL DEFAULT '',
  MODIFY `phone` VARCHAR(255) NOT NULL DEFAULT '',
  MODIFY `email` VARCHAR(255) NOT NULL DEFAULT '',
  MODIFY `status` TINYINT NOT NULL DEFAULT 0,
  MODIFY `register_at` BIGINT NOT NULL DEFAULT 0,
  MODIFY `register_ip` VARCHAR(255) NOT NULL DEFAULT '',
  MODIFY `login_at` BIGINT NOT NULL DEFAULT 0,
  MODIFY `login_ip` VARCHAR(255) NOT NULL DEFAULT 0,
  MODIFY `created_at` BIGINT NOT NULL DEFAULT 0,
  MODIFY `updated_at` BIGINT NOT NULL DEFAULT 0;

-- ROLES
UPDATE `roles` SET `code` = '' WHERE `code` IS NULL;
UPDATE `roles` SET `description` = '' WHERE `description` IS NULL;
UPDATE `roles` SET `status` = 0 WHERE `status` IS NULL;
UPDATE `roles` SET `created_at` = COALESCE(`created_at`, 0);
UPDATE `roles` SET `updated_at` = COALESCE(`updated_at`, 0);

ALTER TABLE `roles`
  MODIFY `code` VARCHAR(255) NOT NULL DEFAULT '',
  MODIFY `description` VARCHAR(255) NOT NULL DEFAULT '',
  MODIFY `status` TINYINT NOT NULL DEFAULT 0,
  MODIFY `created_at` BIGINT NOT NULL DEFAULT 0,
  MODIFY `updated_at` BIGINT NOT NULL DEFAULT 0;

-- USER_ROLES
UPDATE `user_roles` SET `r1_id` = 0 WHERE `r1_id` IS NULL;
UPDATE `user_roles` SET `r2_id` = 0 WHERE `r2_id` IS NULL;
UPDATE `user_roles` SET `created_at` = COALESCE(`created_at`, 0);
UPDATE `user_roles` SET `updated_at` = COALESCE(`updated_at`, 0);

ALTER TABLE `user_roles`
  MODIFY `r1_id` BIGINT NOT NULL DEFAULT 0,
  MODIFY `r2_id` BIGINT NOT NULL DEFAULT 0,
  MODIFY `created_at` BIGINT NOT NULL DEFAULT 0,
  MODIFY `updated_at` BIGINT NOT NULL DEFAULT 0;

-- MENUS
UPDATE `menus` SET `parent_id` = 0 WHERE `parent_id` IS NULL;
UPDATE `menus` SET `path` = '' WHERE `path` IS NULL;
UPDATE `menus` SET `name` = '' WHERE `name` IS NULL;
UPDATE `menus` SET `component` = '' WHERE `component` IS NULL;
UPDATE `menus` SET `redirect` = '' WHERE `redirect` IS NULL;
UPDATE `menus` SET `meta` = '' WHERE `meta` IS NULL;
UPDATE `menus` SET `icon` = '' WHERE `icon` IS NULL;
UPDATE `menus` SET `sort` = 0 WHERE `sort` IS NULL;
UPDATE `menus` SET `created_at` = COALESCE(`created_at`, 0);
UPDATE `menus` SET `updated_at` = COALESCE(`updated_at`, 0);

ALTER TABLE `menus`
  MODIFY `parent_id` BIGINT NOT NULL DEFAULT 0,
  MODIFY `path` VARCHAR(255) NOT NULL DEFAULT '',
  MODIFY `name` VARCHAR(255) NOT NULL DEFAULT '',
  MODIFY `component` VARCHAR(255) NOT NULL DEFAULT '',
  MODIFY `redirect` VARCHAR(255) NOT NULL DEFAULT '',
  MODIFY `meta` VARCHAR(1024) NOT NULL DEFAULT '',
  MODIFY `icon` VARCHAR(255) NOT NULL DEFAULT '',
  MODIFY `sort` BIGINT NOT NULL DEFAULT 0,
  MODIFY `created_at` BIGINT NOT NULL DEFAULT 0,
  MODIFY `updated_at` BIGINT NOT NULL DEFAULT 0;

-- MENU_ROLES
UPDATE `menu_roles` SET `r1_id` = 0 WHERE `r1_id` IS NULL;
UPDATE `menu_roles` SET `r2_id` = 0 WHERE `r2_id` IS NULL;
UPDATE `menu_roles` SET `created_at` = COALESCE(`created_at`, 0);
UPDATE `menu_roles` SET `updated_at` = COALESCE(`updated_at`, 0);

ALTER TABLE `menu_roles`
  MODIFY `r1_id` BIGINT NOT NULL DEFAULT 0,
  MODIFY `r2_id` BIGINT NOT NULL DEFAULT 0,
  MODIFY `created_at` BIGINT NOT NULL DEFAULT 0,
  MODIFY `updated_at` BIGINT NOT NULL DEFAULT 0;

-- I18N
UPDATE `i18n` SET `class` = '' WHERE `class` IS NULL;
UPDATE `i18n` SET `lang` = '' WHERE `lang` IS NULL;
UPDATE `i18n` SET `key` = '' WHERE `key` IS NULL;
UPDATE `i18n` SET `value` = '' WHERE `value` IS NULL;
UPDATE `i18n` SET `created_at` = COALESCE(`created_at`, 0);
UPDATE `i18n` SET `updated_at` = COALESCE(`updated_at`, 0);

ALTER TABLE `i18n`
  MODIFY `class` VARCHAR(255) NOT NULL DEFAULT '',
  MODIFY `lang` VARCHAR(255) NOT NULL DEFAULT '',
  MODIFY `key` VARCHAR(255) NOT NULL DEFAULT '',
  MODIFY `value` VARCHAR(1024) NOT NULL DEFAULT '',
  MODIFY `created_at` BIGINT NOT NULL DEFAULT 0,
  MODIFY `updated_at` BIGINT NOT NULL DEFAULT 0;

-- End migration
