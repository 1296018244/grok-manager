# grok-manager

`grok-manager` 是 [CLIProxyAPI](https://github.com/router-for-me/CLIProxyAPI)（CPA）原生插件，面向 **xAI / Grok** 账号池的日常运维：

- **测活**：并发探测 CPA 中的 xAI 凭证是否可用
- **清理**：按状态批量删除失效凭证（401 / 402 / 403 等）
- **运行时隔离（Autoban）**：命中错误时自动隔离，调度器跳过坏号
- **定时扫描**：可配置周期测活与复检
- **分页结果 / 备份**：管理面板与轻量数据备份

> 本仓库是**公开发布版**。不包含 SSO Cookie 转 CPA 凭证、SSO 历史库等私有功能。

当前版本：**v1.0.0**

## 功能一览

| 模块 | 说明 |
| --- | --- |
| 测活 | 对 `xai` provider 凭证发起真实探测，汇总健康 / 401 / 402 / 403 / 429 |
| 清理 | 删除候选、按 HTTP 状态删除、按文件名删除 |
| 隔离 | `usage.handle` 写入隔离表；`scheduler.pick` 跳过被隔离凭证 |
| 429 策略 | 固定隔离 **2 小时**；到期自动复测，仍 429 则再 +2h |
| 定时 | 周期扫描，可选二次复检（无 SSO 自动重刷） |
| 面板 | CPA 管理 UI 内嵌「Grok Manager」页 |

### 隔离策略（默认）

| 上游状态 | 默认隔离时长 |
| --- | --- |
| `401` | 24 小时 |
| `402` | 7 天 |
| `403` | 24 小时 |
| `429` | 2 小时（到期复测，仍限流再 +2h） |

- 只处理 `xai` provider，不影响 Codex / Claude / Gemini 等。
- 隔离状态默认持久化到 `plugins/grok-manager/bans.json`（随 CPA 重启恢复）。

## 预编译二进制

Release 资源（Windows）：

```text
plugins/windows/amd64/grok-manager.dll
# 或带版本号
plugins/windows/amd64/grok-manager-v1.0.0.dll
```

Linux（amd64）建议自行构建 `.so`（见下方「构建」）。

## 安装

1. 将动态库放到 CPA 插件目录（按系统选择）：

```text
plugins/windows/amd64/grok-manager.dll
plugins/linux/amd64/grok-manager.so
```

2. 在 `config.yaml` 启用：

```yaml
plugins:
  enabled: true
  dir: plugins
  configs:
    grok-manager:
      enabled: true
```

3. 重启 CLIProxyAPI。日志中应出现类似：

```text
pluginhost: plugin registered plugin_id=grok-manager plugin_name=grok-manager version=1.0.0
```

4. 打开 CPA 管理面板，侧栏进入 **Grok Manager**。  
   面板请求走 Management API，需要配置有效的 **管理密钥**（与 CPA `remote-management.secret-key` 对应的明文密码）。

## 管理 API（需管理密钥）

基路径：

```text
/v0/management/plugins/grok-manager
```

| 方法 | 路径 | 说明 |
| --- | --- | --- |
| GET | `/status` | 测活任务摘要（不含大结果集） |
| GET | `/results` | 分页测活结果 |
| POST | `/scan` | 开始测活 |
| POST | `/stop` | 停止测活 |
| POST | `/delete` | 删除候选 / 按状态 / 按名 |
| GET/POST | `/schedule` | 读取 / 更新定时任务 |
| GET | `/bans` | 分页隔离列表 |
| POST | `/bans-recheck-429` | 对 429 隔离项复测 |
| POST | `/unban` | 解禁（单条 / 多条 / 按状态） |
| POST | `/unban-all` | 全部解禁 |
| POST | `/bans-import` | 导入隔离快照 |
| GET | `/paths` | 数据文件路径 |
| POST | `/backup` | 打包 scan / schedule / bans |

面板资源：

```text
/v0/management/plugins/grok-manager/panel
```

（具体前缀以你的 CPA 版本与反代配置为准。）

## 构建

需要 **Go 1.22+**、**CGO** 与 C 编译器（Windows 可用 MinGW，Linux 用 gcc）。

### Windows

```bat
build-windows.bat
```

产物默认输出到 `dist\grok-manager.dll`。

### Linux

```bash
chmod +x build.sh
./build.sh
```

兼容 Debian 12 / 官方镜像时可使用：

```bash
docker run --rm \
  -v "$PWD:/src" \
  -w /src \
  golang:1.24-bookworm \
  sh -c 'CGO_ENABLED=1 go build -buildmode=c-shared -trimpath -ldflags="-s -w" -o dist/grok-manager-linux-amd64.so .'
```

## 配置与数据目录

插件运行时数据（相对 CPA 工作目录）：

```text
plugins/grok-manager/last-scan.json
plugins/grok-manager/schedule.json
plugins/grok-manager/bans.json
```

## 与完整私有版的区别

| | 本仓库 `grok-manager` | 私有完整版 |
| --- | --- | --- |
| 测活 / 清理 / 隔离 / 定时 | ✅ | ✅ |
| SSO Cookie → CPA 凭证转换 | ❌ | ✅ |
| SSO 历史库 / 401 自动重刷 | ❌ | ✅ |

## 致谢

- 运行时隔离与调度跳过思路参考 [akihitohyh/xai-autoban](https://github.com/akihitohyh/xai-autoban)（MIT）
- 上游项目脉络：[ysxk/codex-429-autoban](https://github.com/ysxk/codex-429-autoban)
- CLIProxyAPI：[router-for-me/CLIProxyAPI](https://github.com/router-for-me/CLIProxyAPI)

## License

[MIT](LICENSE)
