# Grok Manager v1.1.3

Single public build — no private fork.

## Changes
- Management key show/hide toggle (text button, works on Edge/Chrome)
- Hide browser native password reveal so only one control remains
- UI title unified as Grok Manager
- Keeps hard isolation from v1.1.2: after 429/ban, scheduler.pick returns Handled:true and never falls back to free-for-all banned pool

## Assets
- grok-manager-linux-amd64.so / grok-manager.so
- grok-manager-linux-arm64.so (if built)
- grok-manager-windows-amd64.dll / grok-manager.dll (CI)

## Install
Place .so / .dll under CLIProxyAPI plugins/ and enable grok-manager only.
