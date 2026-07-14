# grok-manager v1.1.0

**完整公开版**：在测活 / 清理 / 运行时隔离基础上，开放 SSO Cookie → CPA 转换与 SSO 历史库。

## 新增 / 完整能力

- SSO Cookie 批量转 CPA xAI OAuth 凭证
- SSO vault 持久化、分页、导出、删除
- 401 从 vault 自动重刷
- 定时管线：扫描 → 复检 → 可选 401 重刷
- 429 硬顶 2h + 到期复测
- 隔离 email 主键 + auth 反填邮箱

## 安装（Windows）

```text
plugins/windows/amd64/grok-manager.dll
# 或
plugins/windows/amd64/grok-manager-v1.1.0.dll
```

```yaml
plugins:
  enabled: true
  dir: plugins
  configs:
    grok-manager:
      enabled: true
```

重启后日志：`plugin_id=grok-manager version=1.1.0`
