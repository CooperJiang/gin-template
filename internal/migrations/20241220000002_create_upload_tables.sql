-- +migrate Up
-- 创建文件上传相关表

-- 1. 创建文件上传记录表
CREATE TABLE upload_file (
    id CHAR(36) PRIMARY KEY,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL,
    deleted_at DATETIME NULL,
    filename VARCHAR(255) NOT NULL COMMENT '原始文件名',
    stored_name VARCHAR(255) NOT NULL UNIQUE COMMENT '存储的文件名(UUID)',
    file_path VARCHAR(500) NOT NULL COMMENT '文件存储路径',
    file_size BIGINT NOT NULL COMMENT '文件大小(字节)',
    mime_type VARCHAR(100) NOT NULL COMMENT '文件MIME类型',
    extension VARCHAR(20) NOT NULL COMMENT '文件扩展名',
    md5_hash VARCHAR(32) NOT NULL COMMENT '文件MD5哈希',
    upload_status INT DEFAULT 1 COMMENT '上传状态: 1-进行中 2-完成 3-失败',
    chunk_total INT DEFAULT 1 COMMENT '总分片数',
    chunk_uploaded INT DEFAULT 0 COMMENT '已上传分片数',
    user_id CHAR(36) NOT NULL COMMENT '上传用户ID',
    uploaded_at DATETIME NULL COMMENT '上传完成时间',
    KEY idx_md5_hash (md5_hash),
    KEY idx_user_id (user_id),
    KEY idx_deleted_at (deleted_at),
    KEY idx_upload_status (upload_status)
);

-- 2. 创建分片上传信息表
CREATE TABLE chunk_info (
    id CHAR(36) PRIMARY KEY,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL,
    deleted_at DATETIME NULL,
    file_id CHAR(36) NOT NULL COMMENT '文件ID',
    chunk_index INT NOT NULL COMMENT '分片索引',
    chunk_size BIGINT NOT NULL COMMENT '分片大小',
    chunk_path VARCHAR(500) NOT NULL COMMENT '分片存储路径',
    md5_hash VARCHAR(32) NOT NULL COMMENT '分片MD5哈希',
    is_uploaded BOOLEAN DEFAULT FALSE COMMENT '是否已上传',
    KEY idx_file_id (file_id),
    KEY idx_file_chunk (file_id, chunk_index),
    KEY idx_deleted_at (deleted_at),
    UNIQUE KEY uk_file_chunk (file_id, chunk_index)
);

-- +migrate Down
-- 删除文件上传相关表

DROP TABLE IF EXISTS chunk_info;
DROP TABLE IF EXISTS upload_file; 