# grok-manager v1.1.0

Full public build: live-check, cleanup, runtime isolation, plus SSO Cookie to CPA convert and SSO vault.

## Features

- SSO Cookie batch convert to CPA xAI OAuth credentials
- SSO vault persist / paginate / export / delete
- Auto refresh 401 from vault
- Schedule pipeline: scan -> recheck -> optional 401 refresh
- 429 hard-cap 2h + recheck on expiry
- Ban by email primary key + auth email backfill

## Install (Windows)

```text
plugins/windows/amd64/grok-manager.dll
# or
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

Restart CPA; log should show: `plugin_id=grok-manager version=1.1.0`
