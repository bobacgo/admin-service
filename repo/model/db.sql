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