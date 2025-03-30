/**
 * 模板应用前端JS
 */

// 应用配置
window.CONFIG = {
    SYSTEM: {
        STORAGE_KEYS: {
            AUTH_TOKEN: 'token',
            USER_INFO: 'user'
        },
        API_BASE_URL: '/api'
    }
};

// 应用状态
window.appState = {
    isLoggedIn: false,
    token: null,
    user: null
};

/**
 * 发送API请求
 * @param {string} url - API路径
 * @param {Object} options - 请求选项
 * @returns {Promise} - 请求Promise
 */
async function apiRequest(url, options = {}) {
    const defaultOptions = {
        headers: {
            'Content-Type': 'application/json'
        }
    };

    // 如果已登录，添加认证头
    if (window.appState.token) {
        defaultOptions.headers.Authorization = `Bearer ${window.appState.token}`;
    }

    const response = await fetch(`${window.CONFIG.SYSTEM.API_BASE_URL}${url}`, {
        ...defaultOptions,
        ...options,
        headers: {
            ...defaultOptions.headers,
            ...options.headers
        }
    });

    const data = await response.json();
    
    if (!response.ok) {
        throw new Error(data.message || '请求失败');
    }
    
    return data;
}

/**
 * 检查认证状态
 */
function checkAuthState() {
    // 检查本地存储中的令牌
    const token = localStorage.getItem(window.CONFIG.SYSTEM.STORAGE_KEYS.AUTH_TOKEN);
    const user = localStorage.getItem(window.CONFIG.SYSTEM.STORAGE_KEYS.USER_INFO);
    
    if (token && user) {
        try {
            window.appState.token = token;
            window.appState.user = JSON.parse(user);
            window.appState.isLoggedIn = true;
            updateLoginButton();
        } catch (error) {
            console.error('无法解析用户信息', error);
            logout();
        }
    }
}

/**
 * 更新登录按钮状态
 */
function updateLoginButton() {
    const loginBtn = document.getElementById('loginBtn');
    if (!loginBtn) return;
    
    if (window.appState.isLoggedIn) {
        loginBtn.textContent = '退出';
        loginBtn.onclick = logout;
    } else {
        loginBtn.textContent = '登录';
        loginBtn.onclick = showLoginForm;
    }
}

/**
 * 登出
 */
function logout() {
    localStorage.removeItem(window.CONFIG.SYSTEM.STORAGE_KEYS.AUTH_TOKEN);
    localStorage.removeItem(window.CONFIG.SYSTEM.STORAGE_KEYS.USER_INFO);
    window.appState.token = null;
    window.appState.user = null;
    window.appState.isLoggedIn = false;
    updateLoginButton();
}

/**
 * 显示登录表单
 */
function showLoginForm() {
    // 实现登录表单显示逻辑
    console.log('显示登录表单');
}

// 页面加载时初始化
document.addEventListener('DOMContentLoaded', function() {
    checkAuthState();
}); 