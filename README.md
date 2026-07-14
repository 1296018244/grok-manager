# grok-manager

[CLIProxyAPI](https://github.com/router-for-me/CLIProxyAPI)（CPA）原生插件，面向 **xAI / Grok** 账号池运维。

当前版本：**v1.1.0**（完整版，含 SSO → CPA）

## 功能

| 模块 | 说明 |
| --- | --- |
| 测活 | 并发探测 `xai` 凭证，汇总健康 / 401 / 402 / 403 / 429 |
| 清理 | 按候选 / HTTP 状态 / 文件名删除 |
| 运行时隔离 | `usage.handle` 写入隔离表；`scheduler.pick` 跳过坏号 |
| 429 策略 | 固定 **2 小时**硬顶；到期复测，仍限流再 +2h |
| 邮箱主键 | 隔离按 email 去重；usage 无邮箱时从 auth 反填 |
| 定时 | 周期扫描 / 复检 / 可选 401 自动从 vault 重刷 |
| **SSO 转换** | SSO Cookie → CPA xAI OAuth 凭证 |
| **SSO 历史库** | vault 持久化、预览、导出、401 重刷 |
| 面板 | CPA 管理 UI 内嵌 **Grok Manager** |

### 默认隔离时长

| 上游状态 | 时长 |
| --- | --- |
| 401 | 24h（vault 有 SSO 时更短，方便自动重刷） |
| 402 | 7 天 |
| 403 | 24 小时 |
| 429 | 2 小时（到期复测） |

只处理 `xai` provider。
<img width="1780" height="651" alt="image" src="https://github.com/user-attachments/assets/5bced33c-90a0-4462-8b2a-469933c40d91" />


## 预编译二进制

Release（Windows）：

```text
plugins/windows/amd64/grok-manager.dll
# 或
plugins/windows/amd64/grok-manager-v1.1.0.dll
```

Linux amd64 见下方构建。

## 安装

1. 将动态库放入 CPA 插件目录：

```text
plugins/windows/amd64/grok-manager.dll
plugins/linux/amd64/grok-manager.so
```

2. 配置：

```yaml
plugins:
  enabled: true
  dir: plugins
  configs:
    grok-manager:
      enabled: true
```

3. 重启 CPA，日志应出现 `version=1.1.0`。

4. 管理面板侧栏进入 **Grok Manager**（需管理密钥）。

## 管理 API

基路径：`/v0/management/plugins/grok-manager`

主要接口：测活 `/scan`、结果 `/results`、隔离 `/bans`、定时 `/schedule`、  
SSO `/sso-import` `/sso-vault` `/sso-refresh-401`、备份 `/backup` 等。详见面板与源码路由注册。

## 构建

需要 **Go 1.22+**、**CGO**。

### Windows

```bat
build-windows.bat
```

### Linux

```bash
chmod +x build.sh && ./build.sh
```

## 数据目录

```text
plugins/grok-manager/last-scan.json
plugins/grok-manager/schedule.json
plugins/grok-manager/bans.json
plugins/grok-manager/sso-vault.json
plugins/grok-manager/last-sso-import.json
```

## 友链

学 AI 上 L 站！[L 站链接](https://linux.do)

## 致谢

- 运行时隔离思路参考 [akihitohyh/xai-autoban](https://github.com/akihitohyh/xai-autoban)（MIT）
- CLIProxyAPI：[router-for-me/CLIProxyAPI](https://github.com/router-for-me/CLIProxyAPI)

## License

[MIT](LICENSE)
