# 📁 Upload模块完整流程总结

## 🔄 上传流程

### **1. 简单上传流程 (文件 ≤ 2MB)**

```mermaid
sequenceDiagram
    participant C as 客户端
    participant S as 服务器
    participant DB as 数据库
    participant FS as 文件系统

    C->>S: POST /api/v1/upload/simple
    Note over C,S: 文件数据 + MD5 + 用户信息
    
    S->>S: 1. 验证文件(大小/类型/扩展名)
    S->>S: 2. 计算文件MD5
    S->>DB: 3. 检查MD5是否存在(秒传检测)
    
    alt 文件已存在(秒传)
        DB-->>S: 返回已存在文件信息
        S-->>C: 立即返回文件URL
    else 新文件
        S->>S: 4. 生成UUID文件名
        S->>FS: 5. 保存文件到本地存储
        S->>DB: 6. 创建文件记录(IsPublic=true)
        S-->>C: 7. 返回文件访问URL
    end
```

**返回的URL格式**：`/files/preview/{fileId}/{filename.ext}`

### **2. 分片上传流程 (文件 > 2MB)**

```mermaid
sequenceDiagram
    participant C as 客户端
    participant S as 服务器
    participant DB as 数据库
    participant FS as 文件系统

    Note over C,S: 第一步：初始化分片上传
    C->>S: POST /api/v1/upload/chunk/init
    Note over C,S: 文件信息 + 总MD5 + 分片大小
    
    S->>DB: 检查文件MD5(秒传检测)
    alt 文件已存在
        S-->>C: 返回已存在文件ID
    else 新文件
        S->>S: 计算分片总数
        S->>DB: 创建文件记录
        S->>DB: 创建分片记录
        S-->>C: 返回文件ID和分片信息
    end

    Note over C,S: 第二步：并发上传分片
    loop 每个分片
        C->>S: POST /api/v1/upload/chunk
        Note over C,S: 文件ID + 分片索引 + 分片数据 + 分片MD5
        
        S->>DB: 检查分片是否已存在
        alt 分片已存在
            S-->>C: 返回成功(跳过重复上传)
        else 新分片
            S->>S: 验证分片MD5
            S->>FS: 保存分片到临时目录
            S->>DB: 更新分片状态
            S->>DB: 更新上传进度
            S-->>C: 返回上传进度
        end
    end

    Note over C,S: 第三步：合并分片
    C->>S: POST /api/v1/upload/chunk/merge
    S->>DB: 检查所有分片是否完成
    S->>FS: 合并分片到最终文件
    S->>DB: 更新文件状态为完成
    S->>FS: 清理临时分片文件
    S-->>C: 返回最终文件URL
```

## 🔗 URL生成流程

### **URL格式演变**

```typescript
// 旧格式 (仅UUID)
"/files/preview/67afd764-1e0b-4a54-a181-7b66fad62309"

// 新格式 (UUID + 文件名)
"/files/preview/67afd764-1e0b-4a54-a181-7b66fad62309/image.jpg"
```

### **URL生成逻辑**

```go
// 在 SimpleUpload 和 MergeChunks 中
FilePath: fmt.Sprintf("/files/preview/%s/%s", uploadFile.ID.String(), uploadFile.Filename)
```

**优势**：
- ✅ 保持UUID的唯一性和安全性
- ✅ URL中包含文件扩展名，便于识别文件类型
- ✅ 支持原始文件名显示
- ✅ 向后兼容，同时支持两种路由格式

## 🔐 文件权限控制流程

### **权限验证中间件流程**

```mermaid
flowchart TD
    A[请求文件: /files/preview/:fileId] --> B[获取文件ID]
    B --> C[查询文件信息]
    C --> D{文件存在?}
    D -->|否| E[返回404]
    D -->|是| F{检查临时访问令牌}
    F -->|有效| G[允许访问]
    F -->|无效/无| H{检查分享令牌}
    H -->|有效| I[允许访问]
    H -->|无效/无| J{文件是否公开?}
    J -->|是| K[允许访问]
    J -->|否| L{用户已登录?}
    L -->|否| M[返回401: 需要登录]
    L -->|是| N{用户有权限?}
    N -->|是| O[允许访问]
    N -->|否| P[返回403: 权限不足]
    
    G --> Q[记录下载统计]
    I --> Q
    K --> Q
    O --> Q
    Q --> R[返回文件内容]
```

### **权限层级 (优先级从高到低)**

1. **临时访问令牌** - 系统生成的临时链接
2. **分享令牌** - 用户创建的分享链接  
3. **公开文件** - 任何人都可访问
4. **私有文件权限** - 需要用户登录并验证权限

### **文件访问权限类型**

| 文件状态 | 访问条件 | 说明 |
|---------|---------|------|
| **公开文件** | 无需验证 | `IsPublic = true`，任何人可访问 |
| **私有文件** | 需要权限 | `IsPublic = false`，需要登录且有权限 |
| **分享文件** | 分享令牌 | 通过分享链接访问，可设置密码和有效期 |
| **临时访问** | 临时令牌 | 系统生成的临时访问，有使用次数和时间限制 |

## 💾 数据库模型关系

### **核心表结构**

```sql
-- 主文件表
upload_file {
    id: UUID (主键)
    filename: STRING (原始文件名)
    stored_name: STRING (UUID存储名)
    file_path: STRING (文件路径)
    file_size: INT64 (文件大小)
    mime_type: STRING (MIME类型)
    extension: STRING (文件扩展名)
    md5_hash: STRING (文件MD5，用于秒传)
    is_public: BOOLEAN (是否公开，默认true)
    download_count: INT (下载次数统计)
    user_id: STRING (上传者ID)
    upload_status: INT (上传状态)
    uploaded_at: TIMESTAMP (上传完成时间)
}

-- 分片信息表
chunk_info {
    id: UUID (主键)
    file_id: STRING (文件ID，外键)
    chunk_index: INT (分片索引)
    chunk_size: INT64 (分片大小)
    chunk_path: STRING (分片路径)
    md5_hash: STRING (分片MD5)
    is_uploaded: BOOLEAN (是否已上传)
}

-- 文件分享表
file_share {
    id: UUID (主键)
    file_id: STRING (文件ID，外键)
    share_token: STRING (分享令牌，唯一)
    share_name: STRING (分享名称)
    expires_at: TIMESTAMP (过期时间)
    access_password: STRING (访问密码)
    download_limit: INT (下载次数限制)
    download_count: INT (已下载次数)
    created_by: STRING (创建者ID)
    is_active: BOOLEAN (是否激活)
}

-- 文件权限表
file_permission {
    id: UUID (主键)
    file_id: STRING (文件ID，外键)
    user_id: STRING (用户ID，外键)
    permission: STRING (权限类型: read/download/manage)
    granted_by: STRING (授权者ID)
    expires_at: TIMESTAMP (权限过期时间)
    is_active: BOOLEAN (是否激活)
}

-- 临时访问表
temporary_access {
    id: UUID (主键)
    file_id: STRING (文件ID，外键)
    access_token: STRING (访问令牌，唯一)
    purpose: STRING (用途说明)
    expires_at: TIMESTAMP (过期时间，必须)
    usage_limit: INT (使用次数限制)
    usage_count: INT (已使用次数)
    created_by: STRING (创建者)
    is_active: BOOLEAN (是否激活)
}
```

## 🚀 技术特性

### **已实现功能**

✅ **双模式上传**：简单上传 + 分片上传  
✅ **秒传检测**：基于MD5哈希避免重复上传  
✅ **断点续传**：支持暂停和恢复上传  
✅ **进度跟踪**：实时显示上传进度  
✅ **文件去重**：相同文件只存储一份  
✅ **权限控制**：多层级权限验证  
✅ **安全存储**：UUID文件名防止遍历攻击  
✅ **统计功能**：下载次数和访问记录  

### **文件存储策略**

- **存储路径**：`./uploads/` (最终文件)
- **临时路径**：`./uploads/tmp/` (分片临时存储)
- **命名规则**：`{UUID}.{原始扩展名}`
- **访问方式**：通过文件ID而非直接路径访问

### **安全特性**

🔒 **访问控制**：基于权限的文件访问  
🔒 **令牌验证**：分享和临时访问令牌  
🔒 **文件验证**：MIME类型和扩展名验证  
🔒 **MD5校验**：确保文件完整性  
🔒 **路径隔离**：UUID防止路径遍历  

## 🎯 使用示例

### **前端上传代码**

```typescript
// 使用全局文件上传组件
<GlobalFileUpload 
  :multiple="true"
  :auto-upload="true"
  @success="handleUploadSuccess"
  @error="handleUploadError"
/>

// 处理上传成功
const handleUploadSuccess = (response: any) => {
  console.log('文件URL:', response.filePath)
  // 输出: /files/preview/uuid/filename.jpg
}
```

### **后端文件访问**

```bash
# 直接访问公开文件
GET /files/preview/67afd764-1e0b-4a54-a181-7b66fad62309/image.jpg

# 带分享令牌访问
GET /files/preview/67afd764-1e0b-4a54-a181-7b66fad62309/image.jpg?share_token=abc123

# 下载文件
GET /files/download/67afd764-1e0b-4a54-a181-7b66fad62309/image.jpg
```

---

**总结**：Upload模块提供了完整的文件上传、存储、访问控制解决方案，支持多种上传模式和灵活的权限管理，确保文件的安全性和可访问性。 