#!/usr/bin/env node

import fs from 'fs';
import path from 'path';
import { fileURLToPath } from 'url';

const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);

const PACKAGE_JSON_PATH = path.join(__dirname, '../package.json');
const VERSION_CONFIG_PATH = path.join(__dirname, '../src/config/version.json');

/**
 * 解析版本号
 * @param {string} version
 * @returns {object}
 */
function parseVersion(version) {
  const match = version.match(/^(\d+)\.(\d+)\.(\d+)$/);
  if (!match) {
    throw new Error(`Invalid version format: ${version}`);
  }

  return {
    major: parseInt(match[1], 10),
    minor: parseInt(match[2], 10),
    patch: parseInt(match[3], 10)
  };
}

/**
 * 递增版本号
 * @param {object} version
 * @returns {object}
 */
function incrementVersion(version) {
  let { major, minor, patch } = version;

  patch++;

  // 当patch达到10时，进入下一个minor版本
  if (patch >= 10) {
    minor++;
    patch = 0;
  }

  return { major, minor, patch };
}

/**
 * 版本对象转字符串
 * @param {object} version
 * @returns {string}
 */
function versionToString(version) {
  return `${version.major}.${version.minor}.${version.patch}`;
}

/**
 * 更新版本
 */
function updateVersion() {
  try {
    // 读取package.json
    const packageJson = JSON.parse(fs.readFileSync(PACKAGE_JSON_PATH, 'utf8'));
    const currentVersion = packageJson.version;

    console.log(`当前版本: ${currentVersion}`);

    // 解析并递增版本
    const parsedVersion = parseVersion(currentVersion);
    const newVersionObj = incrementVersion(parsedVersion);
    const newVersion = versionToString(newVersionObj);

    console.log(`新版本: ${newVersion}`);

    // 更新package.json
    packageJson.version = newVersion;
    fs.writeFileSync(PACKAGE_JSON_PATH, JSON.stringify(packageJson, null, 2) + '\n');

    // 创建版本配置目录
    const configDir = path.dirname(VERSION_CONFIG_PATH);
    if (!fs.existsSync(configDir)) {
      fs.mkdirSync(configDir, { recursive: true });
    }

    // 生成版本配置文件
    const buildTime = new Date().toISOString();
    const versionConfig = {
      version: newVersion,
      buildTime: buildTime,
      buildTimestamp: Date.now(),
      buildDate: new Date().toLocaleDateString('zh-CN', {
        year: 'numeric',
        month: '2-digit',
        day: '2-digit',
        hour: '2-digit',
        minute: '2-digit',
        second: '2-digit'
      })
    };

    fs.writeFileSync(VERSION_CONFIG_PATH, JSON.stringify(versionConfig, null, 2) + '\n');

    console.log(`版本配置已生成: ${VERSION_CONFIG_PATH}`);
    console.log(`构建时间: ${versionConfig.buildDate}`);

    return versionConfig;
  } catch (error) {
    console.error('版本更新失败:', error.message);
    process.exit(1);
  }
}

// 如果直接运行此脚本
if (import.meta.url === `file://${process.argv[1]}`) {
  updateVersion();
}

export { updateVersion };
