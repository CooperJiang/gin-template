-- +migrate Up
-- 将用户表的ID字段从自增int改为UUID

-- 1. 创建新的用户表（使用UUID）
CREATE TABLE user_new (
    id CHAR(36) PRIMARY KEY,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL,
    deleted_at DATETIME NULL,
    username VARCHAR(50) NOT NULL,
    password VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL,
    avatar VARCHAR(255) DEFAULT '',
    bio VARCHAR(500) DEFAULT '',
    status INT DEFAULT 1,
    role INT DEFAULT 3,
    UNIQUE KEY idx_username (username),
    UNIQUE KEY idx_email (email),
    KEY idx_deleted_at (deleted_at)
);

-- 2. 如果旧表存在数据，可以选择性迁移（这里先清空，实际使用时需要根据情况调整）
-- INSERT INTO user_new (id, created_at, updated_at, username, password, email, avatar, bio, status, role)
-- SELECT 
--     LOWER(CONCAT(
--         SUBSTR(HEX(RANDOM_BYTES(4)), 1, 8), '-',
--         SUBSTR(HEX(RANDOM_BYTES(2)), 1, 4), '-',
--         '4', SUBSTR(HEX(RANDOM_BYTES(2)), 2, 3), '-',
--         CONCAT(SUBSTR('89ab', FLOOR(1 + RAND() * 4), 1), SUBSTR(HEX(RANDOM_BYTES(2)), 2, 3)), '-',
--         SUBSTR(HEX(RANDOM_BYTES(6)), 1, 12)
--     )) as id,
--     created_at, updated_at, username, password, email, avatar, bio, status, role
-- FROM user;

-- 3. 删除旧表
DROP TABLE IF EXISTS user;

-- 4. 重命名新表
ALTER TABLE user_new RENAME TO user;

-- +migrate Down
-- 回滚：将UUID改回自增int（注意：这会丢失数据）

-- 1. 创建旧的用户表结构
CREATE TABLE user_old (
    id INT AUTO_INCREMENT PRIMARY KEY,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL,
    deleted_at DATETIME NULL,
    username VARCHAR(50) NOT NULL,
    password VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL,
    avatar VARCHAR(255) DEFAULT '',
    bio VARCHAR(500) DEFAULT '',
    status INT DEFAULT 1,
    role INT DEFAULT 3,
    UNIQUE KEY idx_username (username),
    UNIQUE KEY idx_email (email),
    KEY idx_deleted_at (deleted_at)
);

-- 2. 删除UUID表
DROP TABLE IF EXISTS user;

-- 3. 重命名回原表名
ALTER TABLE user_old RENAME TO user; 