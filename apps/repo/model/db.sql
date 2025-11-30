-- 用户表
CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    account VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    phone VARCHAR(255),
    email VARCHAR(255),
    status INT,
    register_at INT,
    register_ip VARCHAR(255),
    login_at INT,
    login_ip VARCHAR(255),
    created_at INT DEFAULT CURRENT_TIMESTAMP,
    updated_at INT DEFAULT CURRENT_TIMESTAMP
);
CREATE UNIQUE INDEX IF NOT EXISTS uq_account ON users(account);

INSERT INTO users (account, password, status) VALUES ('admin', 'admin', 1);

-- 角色表
CREATE TABLE IF NOT EXISTS roles (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    code VARCHAR(255) NOT NULL ,
    description VARCHAR(255),
    status INT,
    created_at INT DEFAULT CURRENT_TIMESTAMP,
    updated_at INT DEFAULT CURRENT_TIMESTAMP
);
CREATE UNIQUE INDEX IF NOT EXISTS uq_code ON roles(code);

INSERT INTO roles (code, description, status) VALUES ('super_admin', 'all power', 1);

-- 菜单表
CREATE TABLE IF NOT EXISTS menus (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    path VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    component VARCHAR(255),
    redirect VARCHAR(255),
    meta VARCHAR(255),
    icon VARCHAR(255),
    created_at INT DEFAULT CURRENT_TIMESTAMP,
    updated_at INT DEFAULT CURRENT_TIMESTAMP
);
CREATE UNIQUE INDEX IF NOT EXISTS uq_path ON menus(path);

-- 多语言翻译表
CREATE TABLE IF NOT EXISTS i18n (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    class VARCHAR(255) NOT NULL,
    lang VARCHAR(255) NOT NULL,
    key VARCHAR(255) NOT NULL,
    value VARCHAR(255) NOT NULL,
    created_at INT DEFAULT CURRENT_TIMESTAMP,
    updated_at INT DEFAULT CURRENT_TIMESTAMP
);
CREATE UNIQUE INDEX IF NOT EXISTS uq_key_lang ON i18n(key, lang);