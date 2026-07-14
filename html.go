package main

const panelHTML = `<!doctype html>
<html lang="zh-CN">
<head>
<meta charset="utf-8"/>
<meta name="viewport" content="width=device-width,initial-scale=1"/>
<title>Grok Manager</title>
<style>
:root{
  --bg0:#f5f7fb;--bg1:#ffffff;--bg2:#f8fafc;--bg3:#f1f5f9;
  --line:#e8ecf2;--line2:#dbe2ea;
  --text:#0f172a;--muted:#64748b;--faint:#94a3b8;
  --ok:#059669;--ok-bg:#ecfdf5;--ok-bd:#a7f3d0;
  --bad:#dc2626;--bad-bg:#fef2f2;--bad-bd:#fecaca;
  --warn:#d97706;--warn-bg:#fffbeb;--warn-bd:#fde68a;
  --info:#2563eb;--info-bg:#eff6ff;--info-bd:#bfdbfe;
  --accent:#2563eb;--accent2:#1d4ed8;--accent-bg:#eff6ff;
  --pay:#7c3aed;--pay-bg:#f5f3ff;--pay-bd:#ddd6fe;
  --radius:12px;--radius-sm:8px;--shadow:0 1px 2px rgba(15,23,42,.04),0 4px 16px rgba(15,23,42,.04);
  --font:ui-sans-serif,system-ui,-apple-system,"Segoe UI",Roboto,"PingFang SC","Microsoft YaHei",sans-serif;
  --mono:ui-monospace,SFMono-Regular,Consolas,"Cascadia Mono",monospace;
}
*{box-sizing:border-box}
html,body{margin:0;min-height:100%;background:var(--bg0);color:var(--text);font-family:var(--font)}
button,input,textarea,select{font:inherit}
button{border:0;cursor:pointer;font-weight:600;transition:background .15s ease,opacity .15s ease,box-shadow .15s ease,color .15s ease}
button:active:not(:disabled){transform:translateY(1px)}
button:disabled{opacity:.45;cursor:not-allowed}
a{color:var(--accent)}

.app{max-width:1280px;margin:0 auto;padding:16px 16px 40px}

/* Header */
.topbar{
  display:flex;flex-wrap:wrap;gap:14px 20px;align-items:flex-start;justify-content:space-between;
  margin-bottom:14px;padding:16px 18px;border:1px solid var(--line);border-radius:14px;
  background:var(--bg1);box-shadow:var(--shadow);
}
.brand{display:flex;gap:12px;align-items:center;min-width:0}
.logo{
  width:42px;height:42px;border-radius:10px;flex:0 0 auto;
  background:var(--accent);display:grid;place-items:center;
  font-weight:800;font-size:14px;color:#fff;letter-spacing:-.02em;
}
.brand h1{margin:0;font-size:17px;font-weight:700;letter-spacing:-.02em;color:var(--text)}
.brand p{margin:2px 0 0;color:var(--muted);font-size:12.5px;line-height:1.4}
.brand .ver{display:inline-flex;align-items:center;gap:6px;margin-top:6px;flex-wrap:wrap}
.chip{
  display:inline-flex;align-items:center;gap:5px;padding:3px 9px;border-radius:999px;
  font-size:11px;font-weight:600;border:1px solid var(--line);background:var(--bg2);color:var(--muted);
}
.chip-accent{background:var(--accent-bg);border-color:var(--info-bd);color:var(--accent)}
.chip-ok{background:var(--ok-bg);border-color:var(--ok-bd);color:var(--ok)}
.chip-warn{background:var(--warn-bg);border-color:var(--warn-bd);color:var(--warn)}
.chip-bad{background:var(--bad-bg);border-color:var(--bad-bd);color:var(--bad)}
.chip-info{background:var(--info-bg);border-color:var(--info-bd);color:var(--info)}
.top-actions{display:flex;flex-wrap:wrap;gap:8px;align-items:end}
.field{display:flex;flex-direction:column;gap:5px;min-width:0}
.field>span{font-size:11px;color:var(--muted);font-weight:600}
.field input,.field textarea,.field select,
input[type=text],input[type=password],input[type=number],textarea,select{
  background:var(--bg1);border:1px solid var(--line2);color:var(--text);
  border-radius:var(--radius-sm);padding:8px 11px;outline:none;min-width:0;
  transition:border-color .15s ease,box-shadow .15s ease;
}
.field input:focus,input:focus,textarea:focus,select:focus{
  border-color:#93c5fd;box-shadow:0 0 0 3px rgba(37,99,235,.12);
}
#mgmtKey{min-width:240px;width:min(320px,70vw);font-family:var(--mono);font-size:12.5px;background:var(--bg2)}
/* Hide browser native password reveal (Edge/IE) so our toggle is the only control */
#mgmtKey::-ms-reveal,#mgmtKey::-ms-clear{display:none}
#mgmtKey::-webkit-credentials-auto-fill-button{visibility:hidden;pointer-events:none;position:absolute;right:0}

/* Nav tabs */
.nav{
  display:flex;gap:4px;flex-wrap:wrap;padding:5px;margin-bottom:14px;
  background:var(--bg1);border:1px solid var(--line);border-radius:12px;
  position:sticky;top:8px;z-index:20;box-shadow:var(--shadow);
}
.nav button{
  padding:8px 13px;border-radius:8px;background:transparent;color:var(--muted);font-size:13px;
}
.nav button:hover{background:var(--bg3);color:var(--text)}
.nav button.on{
  background:var(--accent-bg);color:var(--accent);box-shadow:none;font-weight:700;
}
.nav button .badge{
  display:inline-block;margin-left:5px;padding:0 6px;border-radius:999px;
  font-size:10px;background:var(--bg3);color:var(--muted);min-width:16px;text-align:center;
}
.nav button.on .badge{background:#dbeafe;color:var(--accent)}

/* Panels */
.panel{display:none}
.panel.on{display:block;animation:fade .15s ease}
@keyframes fade{from{opacity:0;transform:translateY(3px)}to{opacity:1;transform:none}}

.card{
  background:var(--bg1);border:1px solid var(--line);border-radius:var(--radius);
  padding:16px 18px;margin-bottom:12px;box-shadow:var(--shadow);
}
.card-hd{
  display:flex;flex-wrap:wrap;gap:8px 12px;align-items:center;justify-content:space-between;
  margin-bottom:12px;
}
.card-hd h2{margin:0;font-size:15px;font-weight:700;color:var(--text)}
.card-hd .sub{color:var(--muted);font-size:12px}
.hint{font-size:12.5px;color:var(--muted);line-height:1.55;margin:0}
.hint code,.path code,code{
  font-family:var(--mono);font-size:11.5px;background:var(--bg3);
  border:1px solid var(--line);padding:1px 6px;border-radius:5px;color:#334155;
}
.path{
  font-family:var(--mono);font-size:11px;color:#475569;word-break:break-all;
  line-height:1.45;margin-top:8px;
}
.muted{color:var(--muted)}
.faint{color:var(--faint)}

/* Banner */
.banner{
  border-radius:10px;padding:10px 12px;font-size:13px;line-height:1.5;
  border:1px solid transparent;margin:0;
}
.banner + .banner,.banner + .path,.banner + .row{margin-top:10px}
.banner-ok{background:var(--ok-bg);border-color:var(--ok-bd);color:#065f46}
.banner-warn{background:var(--warn-bg);border-color:var(--warn-bd);color:#92400e}
.banner-bad{background:var(--bad-bg);border-color:var(--bad-bd);color:#991b1b}
.banner-info{background:var(--info-bg);border-color:var(--info-bd);color:#1e40af}

/* Grid / form */
.row{display:flex;flex-wrap:wrap;gap:10px;align-items:end}
.grid-2{display:grid;grid-template-columns:1.4fr .9fr;gap:12px}
@media (max-width:900px){.grid-2{grid-template-columns:1fr}}
label.field{flex:0 1 auto}
label.field.grow,label.grow{flex:1 1 240px}
label.check{
  display:inline-flex;flex-direction:row;align-items:center;gap:8px;
  color:var(--text);font-size:13px;padding:8px 10px;border-radius:8px;
  background:var(--bg2);border:1px solid var(--line);cursor:pointer;user-select:none;
}
label.check:hover{border-color:var(--line2);background:var(--bg3)}
label.check input{width:15px;height:15px;accent-color:var(--accent);margin:0}
textarea{
  width:100%;min-height:140px;resize:vertical;font-family:var(--mono);font-size:12px;line-height:1.45;
  background:var(--bg2);
}
input[type=number]{width:100px}
.actions{display:flex;flex-wrap:wrap;gap:8px;margin-top:12px}

/* Buttons */
.btn,.btn-ghost,.btn-ok,.btn-warn,.btn-danger,.btn-soft{
  border-radius:9px;padding:8px 14px;font-size:13px;
}
.btn{background:var(--accent);color:#fff}
.btn:hover:not(:disabled){background:var(--accent2)}
.btn-ok{background:var(--ok);color:#fff}
.btn-ok:hover:not(:disabled){background:#047857}
.btn-warn{background:var(--warn);color:#fff}
.btn-danger{background:var(--bad);color:#fff}
.btn-danger:hover:not(:disabled){background:#b91c1c}
.btn-ghost{background:var(--bg1);color:var(--text);border:1px solid var(--line2)}
.btn-ghost:hover:not(:disabled){background:var(--bg3)}
.btn-soft{background:var(--accent-bg);color:var(--accent);border:1px solid var(--info-bd)}
.btn-soft:hover:not(:disabled){background:#dbeafe}
.btn-sm{padding:4px 9px;font-size:12px;border-radius:7px}

/* Stats — light cards, not black blocks */
.stats{display:grid;grid-template-columns:repeat(auto-fit,minmax(108px,1fr));gap:8px}
.stat{
  background:var(--bg2);border:1px solid var(--line);border-radius:10px;padding:11px 12px 10px;
  position:relative;overflow:hidden;
}
.stat::before{
  content:"";position:absolute;inset:0 auto 0 0;width:3px;border-radius:3px 0 0 3px;background:var(--line2);
}
.stat .n{font-size:21px;font-weight:700;letter-spacing:-.03em;line-height:1.1;font-variant-numeric:tabular-nums;color:var(--text)}
.stat .l{font-size:11.5px;color:var(--muted);margin-top:3px;font-weight:500}
.stat.ok::before{background:var(--ok)}.stat.ok .n{color:var(--ok)}
.stat.bad::before{background:var(--bad)}.stat.bad .n{color:var(--bad)}
.stat.warn::before{background:var(--warn)}.stat.warn .n{color:var(--warn)}
.stat.info::before{background:var(--info)}.stat.info .n{color:var(--info)}
.stat.pay::before{background:var(--pay)}.stat.pay .n{color:var(--pay)}

.bar{
  height:6px;background:var(--bg3);border-radius:999px;overflow:hidden;
  border:1px solid var(--line);margin-top:12px;
}
.bar>i{
  display:block;height:100%;width:0;background:var(--accent);
  transition:width .25s ease;
}

/* Tables */
.table-wrap{
  overflow:auto;max-height:420px;border:1px solid var(--line);border-radius:10px;
  background:var(--bg1);margin-top:10px;
}
.table-wrap.sm{max-height:260px}
table{width:100%;border-collapse:collapse;font-size:12.5px}
th,td{padding:9px 11px;border-bottom:1px solid var(--line);text-align:left;vertical-align:top}
th{
  position:sticky;top:0;z-index:1;background:var(--bg2);color:var(--muted);
  font-size:11px;font-weight:700;letter-spacing:.02em;text-transform:uppercase;
}
tr:hover td{background:#f8fafc}
td.mono,th.mono{font-family:var(--mono);font-size:11.5px}

.tag{
  display:inline-flex;align-items:center;padding:2px 8px;border-radius:999px;
  font-size:11px;font-weight:700;border:1px solid transparent;white-space:nowrap;
}
.tag-ok{background:var(--ok-bg);color:var(--ok);border-color:var(--ok-bd)}
.tag-bad,.tag-del{background:var(--bad-bg);color:var(--bad);border-color:var(--bad-bd)}
.tag-skip,.tag-rate{background:var(--warn-bg);color:var(--warn);border-color:var(--warn-bd)}
.tag-keep{background:var(--info-bg);color:var(--info);border-color:var(--info-bd)}
.tag-pay{background:var(--pay-bg);color:var(--pay);border-color:var(--pay-bd)}

/* Filter chips */
.filters{display:flex;flex-wrap:wrap;gap:6px;margin:4px 0 2px}
.filters button{
  padding:6px 11px;border-radius:999px;font-size:12px;
  background:var(--bg2);color:var(--muted);border:1px solid var(--line);
}
.filters button:hover{color:var(--text);border-color:var(--line2);background:var(--bg3)}
.filters button.on{background:var(--accent-bg);color:var(--accent);border-color:var(--info-bd);font-weight:700}
.filters button .fc{
  display:inline-block;margin-left:2px;min-width:1.2em;padding:0 6px;border-radius:999px;
  font-size:11px;font-weight:700;font-variant-numeric:tabular-nums;
  background:var(--bg3);color:var(--muted);line-height:1.5;
}
.filters button.on .fc{background:#dbeafe;color:var(--accent)}
.filters button .fc.zero{opacity:.45}

/* Log */
.log{
  font-family:var(--mono);font-size:11.5px;color:#475569;white-space:pre-wrap;
  max-height:180px;overflow:auto;background:var(--bg2);border:1px solid var(--line);
  border-radius:10px;padding:10px 12px;margin-top:10px;line-height:1.5;
}

/* Overview hero */
.hero-grid{display:grid;grid-template-columns:1.2fr .8fr;gap:12px}
@media (max-width:900px){.hero-grid{grid-template-columns:1fr}}
.kv{display:grid;gap:8px}
.kv-row{
  display:flex;justify-content:space-between;gap:10px;padding:10px 12px;
  background:var(--bg2);border:1px solid var(--line);border-radius:8px;font-size:12.5px;
}
.kv-row span:first-child{color:var(--muted)}
.kv-row span:last-child{font-weight:600;text-align:right;word-break:break-all;color:var(--text)}
.pill{
  display:inline-block;padding:2px 8px;border-radius:999px;background:var(--bg3);
  font-size:11px;margin-right:6px;border:1px solid var(--line);color:var(--muted);
}
.divider{height:1px;background:var(--line);margin:14px 0}
.foot{margin-top:8px;text-align:center;color:var(--faint);font-size:11.5px}
.toast{
  position:fixed;right:16px;bottom:16px;z-index:50;max-width:min(420px,92vw);
  padding:12px 14px;border-radius:10px;background:var(--bg1);border:1px solid var(--line2);
  color:var(--text);box-shadow:0 8px 28px rgba(15,23,42,.12);font-size:13px;line-height:1.45;
  transform:translateY(12px);opacity:0;pointer-events:none;transition:.2s ease;
}
.toast.show{transform:none;opacity:1}
.toast.err{border-color:var(--bad-bd);background:var(--bad-bg);color:#991b1b}
.toast.ok{border-color:var(--ok-bd);background:var(--ok-bg);color:#065f46}

/* UX polish */
.stat.clickable{cursor:pointer;transition:border-color .15s,box-shadow .15s,transform .1s}
.stat.clickable:hover{border-color:var(--line2);box-shadow:0 2px 8px rgba(15,23,42,.06)}
.stat.clickable:active{transform:scale(.98)}
.stat.on{border-color:var(--info-bd);background:var(--accent-bg);box-shadow:0 0 0 1px var(--info-bd)}
.policy-row{display:flex;flex-wrap:wrap;gap:6px;margin:0 0 4px}
.policy-chip{
  display:inline-flex;align-items:center;gap:5px;padding:4px 10px;border-radius:999px;
  font-size:11.5px;font-weight:600;background:var(--bg2);border:1px solid var(--line);color:var(--muted);
}
.policy-chip b{color:var(--text);font-weight:700}
.policy-chip.w{background:var(--warn-bg);border-color:var(--warn-bd);color:#92400e}
.policy-chip.b{background:var(--bad-bg);border-color:var(--bad-bd);color:#991b1b}
.policy-chip.p{background:var(--pay-bg);border-color:var(--pay-bd);color:#5b21b6}
.policy-chip.i{background:var(--info-bg);border-color:var(--info-bd);color:#1e40af}
.toolbar{
  display:flex;flex-wrap:wrap;gap:8px;align-items:center;justify-content:space-between;
  margin-top:12px;padding:10px 12px;background:var(--bg2);border:1px solid var(--line);border-radius:10px;
}
.toolbar .grp{display:flex;flex-wrap:wrap;gap:6px;align-items:center}
.toolbar .sep{width:1px;height:22px;background:var(--line2);margin:0 2px}
.toolbar .lbl{font-size:11px;font-weight:700;color:var(--faint);text-transform:uppercase;letter-spacing:.04em;margin-right:2px}
.pager{
  display:flex;flex-wrap:wrap;gap:8px;align-items:center;justify-content:space-between;
  margin:10px 0 0;padding:0 2px;
}
.pager .info{font-size:12px;color:var(--muted)}
.pager .btns{display:flex;gap:6px;align-items:center}
.remain{font-variant-numeric:tabular-nums;font-weight:600}
.remain.urgent{color:var(--bad)}
.remain.soon{color:var(--warn)}
.remain.ok{color:var(--ok)}
.id-cell{font-family:var(--mono);font-size:11px;word-break:break-all;max-width:220px;line-height:1.35}
.id-cell .short{display:inline}
.empty{
  text-align:center;color:var(--muted);padding:36px 16px;font-size:13px;line-height:1.6;
}
.empty strong{display:block;color:var(--text);font-size:14px;margin-bottom:4px}
.recheck-card{
  display:flex;flex-wrap:wrap;gap:10px 16px;align-items:center;justify-content:space-between;
  margin-top:10px;padding:12px 14px;border-radius:10px;
  background:linear-gradient(135deg,#fffbeb 0%,#fff7ed 100%);
  border:1px solid var(--warn-bd);
}
.recheck-card .t{font-size:13px;font-weight:700;color:#92400e}
.recheck-card .d{font-size:12px;color:#a16207;margin-top:2px;line-height:1.45}
.recheck-card.running{background:linear-gradient(135deg,#eff6ff 0%,#e0f2fe 100%);border-color:var(--info-bd)}
.recheck-card.running .t,.recheck-card.running .d{color:#1e40af}
.spin{display:inline-block;width:12px;height:12px;border:2px solid currentColor;border-right-color:transparent;border-radius:50%;animation:spin .7s linear infinite;vertical-align:-1px;margin-right:4px}
@keyframes spin{to{transform:rotate(360deg)}}
.sel-count{
  display:inline-flex;align-items:center;gap:4px;padding:3px 8px;border-radius:999px;
  font-size:11px;font-weight:700;background:var(--accent-bg);color:var(--accent);border:1px solid var(--info-bd);
}
.sel-count:empty,.sel-count.zero{display:none}
.table-wrap.tall{max-height:min(560px,62vh)}
.table-wrap.mid{max-height:min(420px,52vh)}

/* Shared layout helpers */
.sec-title{margin:0 0 2px;font-size:15px;font-weight:700}
.sec-sub{margin:0;color:var(--muted);font-size:12px;line-height:1.4}
.checks{display:flex;flex-wrap:wrap;gap:8px;margin-top:10px}
.form-grid{display:grid;grid-template-columns:repeat(auto-fit,minmax(140px,1fr));gap:10px;margin-top:12px}
.form-grid .grow{grid-column:span 2}
@media (max-width:640px){.form-grid .grow{grid-column:span 1}}
.quick-grid{display:grid;grid-template-columns:repeat(4,1fr);gap:10px;margin-bottom:12px}
@media (max-width:1000px){.quick-grid{grid-template-columns:repeat(2,1fr)}}
@media (max-width:520px){.quick-grid{grid-template-columns:1fr}}
.qcard{
  display:block;padding:14px 14px 12px;border-radius:12px;border:1px solid var(--line);
  background:var(--bg1);box-shadow:var(--shadow);cursor:pointer;text-align:left;width:100%;
  transition:border-color .15s,box-shadow .15s,transform .1s;
}
.qcard:hover{border-color:#bfdbfe;box-shadow:0 4px 14px rgba(37,99,235,.08);transform:translateY(-1px)}
.qcard .k{font-size:12px;font-weight:700;color:var(--muted);margin-bottom:6px}
.qcard .v{font-size:20px;font-weight:800;letter-spacing:-.03em;font-variant-numeric:tabular-nums;color:var(--text)}
.qcard .s{font-size:11.5px;color:var(--faint);margin-top:4px;line-height:1.35}
.qcard.warn .v{color:var(--warn)}.qcard.bad .v{color:var(--bad)}.qcard.ok .v{color:var(--ok)}.qcard.info .v{color:var(--info)}
.pipeline{display:flex;flex-wrap:wrap;gap:8px;align-items:center;margin:12px 0 4px}
.pstep{
  display:inline-flex;align-items:center;gap:6px;padding:7px 12px;border-radius:999px;
  background:var(--bg2);border:1px solid var(--line);font-size:12px;font-weight:600;color:var(--text);
}
.pstep i{display:inline-grid;place-items:center;width:18px;height:18px;border-radius:50%;
  background:var(--accent);color:#fff;font-size:10px;font-style:normal;font-weight:800}
.parr{color:var(--faint);font-size:14px}
.danger-zone{
  margin-top:12px;padding:10px 12px;border-radius:10px;border:1px dashed var(--bad-bd);background:var(--bad-bg);
}
.danger-zone .t{font-size:12px;font-weight:700;color:#991b1b;margin-bottom:8px}
.log-hd{display:flex;justify-content:space-between;align-items:center;margin-top:10px;gap:8px}
.log-hd .sub{font-size:11px;color:var(--muted)}
.mono-sm{font-family:var(--mono);font-size:11px;color:#475569}
.adv-row{max-width:360px;overflow:hidden;text-overflow:ellipsis;white-space:nowrap;color:var(--muted);font-size:12px}
td.nowrap{white-space:nowrap}
.nav button .badge:empty,.nav button .badge.zero{display:none}
.stat-sm .n{font-size:18px}
.topbar{backdrop-filter:saturate(1.2) blur(6px)}
.nav{backdrop-filter:saturate(1.1) blur(6px)}
@media (max-width:720px){
  .app{padding:10px 10px 32px}
  .topbar{padding:12px 14px}
  #mgmtKey{min-width:0;width:100%}
  .top-actions{width:100%}
  .id-cell{max-width:140px}
  .nav button{padding:7px 10px;font-size:12px}
}
</style>
</head>
<body>
<div class="app">
  <header class="topbar">
    <div class="brand">
      <div class="logo">GM</div>
      <div>
        <h1>Grok Manager</h1>
        <div class="ver">
          <span class="chip chip-accent">v<span id="ver">1.1.3</span></span>
          <span class="chip" id="jobState">待命</span>
          <span class="chip chip-info" id="hdrVault">库 0</span>
          <span class="chip chip-warn" id="hdrBan">隔离 0</span>
        </div>
      </div>
    </div>
    <div class="top-actions">
      <label class="field" for="mgmtKey">
        <span>管理密钥</span>
        <input id="mgmtKey" type="password" placeholder="密钥" autocomplete="off" spellcheck="false"/>
      </label>
      <button class="btn-ghost btn-sm" type="button" id="mgmtKeyToggle" title="显示/隐藏密钥">显示</button>
      <button class="btn-ghost" type="button" onclick="saveKey()">保存</button>
      <button class="btn-soft" type="button" onclick="boot()">刷新</button>
      <button class="btn-ghost btn-sm" type="button" onclick="doBackup()">备份</button>
    </div>
  </header>

  <nav class="nav" id="mainNav">
    <button type="button" class="on" data-tab="overview" onclick="switchTab('overview',this)">总览</button>
    <button type="button" data-tab="sso" onclick="switchTab('sso',this)">SSO</button>
    <button type="button" data-tab="scan" onclick="switchTab('scan',this)">测活 <span class="badge zero" id="navCand">0</span></button>
    <button type="button" data-tab="vault" onclick="switchTab('vault',this)">历史库 <span class="badge zero" id="navVault">0</span></button>
    <button type="button" data-tab="autoban" onclick="switchTab('autoban',this)">隔离 <span class="badge zero" id="navBan">0</span></button>
    <button type="button" data-tab="schedule" onclick="switchTab('schedule',this)">定时</button>
  </nav>

  <!-- OVERVIEW -->
  <section class="panel on" id="tab-overview">
    <div class="quick-grid">
      <button type="button" class="qcard info" onclick="switchTab('scan')">
        <div class="k">测活</div>
        <div class="v" id="ovQScan">0</div>
        <div class="s" id="ovQScanSub">—</div>
      </button>
      <button type="button" class="qcard ok" onclick="switchTab('vault')">
        <div class="k">历史库</div>
        <div class="v" id="ovQVault">0</div>
        <div class="s" id="ovQVaultSub">—</div>
      </button>
      <button type="button" class="qcard warn" onclick="switchTab('autoban')">
        <div class="k">隔离</div>
        <div class="v" id="ovQBan">0</div>
        <div class="s" id="ovQBanSub">—</div>
      </button>
      <button type="button" class="qcard" onclick="switchTab('schedule')">
        <div class="k">定时</div>
        <div class="v" id="ovQSch" style="font-size:16px;padding-top:4px">—</div>
        <div class="s" id="ovQSchSub">—</div>
      </button>
    </div>

    <div class="card">
      <div class="card-hd">
        <div>
          <h2>测活摘要</h2>
          <div class="sub" id="ovScanSub">—</div>
        </div>
        <button class="btn btn-sm" type="button" onclick="switchTab('scan')">测活</button>
      </div>
      <div class="stats">
        <div class="stat info"><div class="n" id="sTotal">0</div><div class="l">总数</div></div>
        <div class="stat ok"><div class="n" id="sOK">0</div><div class="l">健康</div></div>
        <div class="stat bad"><div class="n" id="s401">0</div><div class="l">401</div></div>
        <div class="stat bad"><div class="n" id="s403">0</div><div class="l">403</div></div>
        <div class="stat pay"><div class="n" id="s402">0</div><div class="l">402</div></div>
        <div class="stat warn"><div class="n" id="s429">0</div><div class="l">429</div></div>
        <div class="stat ok"><div class="n" id="sVaultMatch">0</div><div class="l">401 有库</div></div>
        <div class="stat warn"><div class="n" id="sVaultMiss">0</div><div class="l">401 无库</div></div>
        <div class="stat bad"><div class="n" id="sCand">0</div><div class="l">候选</div></div>
        <div class="stat"><div class="n" id="sDone">0</div><div class="l">完成</div></div>
        <div class="stat"><div class="n" id="sKeep">0</div><div class="l">保留</div></div>
        <div class="stat bad"><div class="n" id="sErr">0</div><div class="l">错误</div></div>
      </div>
      <div class="bar"><i id="bar"></i></div>
      <!-- hidden placeholders for JS -->
      <div id="persistBanner" class="banner banner-info" style="display:none"></div>
      <div id="persistPaths" class="path" style="display:none"></div>
      <div id="ovPaths" class="path" style="display:none"></div>
      <div id="scanPaths" class="path" style="display:none"></div>
      <span id="ovSso" style="display:none"></span>
      <span id="ovVault" style="display:none"></span>
      <span id="ovBan" style="display:none"></span>
      <span id="ovSch" style="display:none"></span>
    </div>
  </section>

  <!-- SSO -->
  <section class="panel" id="tab-sso">
    <div class="card">
      <div class="card-hd">
        <h2>SSO 导入</h2>
        <button class="btn-ghost btn-sm" type="button" onclick="switchTab('vault');loadVault(true)">历史库</button>
      </div>
      <div class="row">
        <label class="field grow">
          <span>文件</span>
          <input id="ssoFile" type="file" accept=".txt,.csv,text/plain" onchange="onSSOFile(event)"/>
        </label>
        <button class="btn-ghost" type="button" style="margin-top:18px" onclick="previewSSO()">预览</button>
        <button class="btn-soft" type="button" style="margin-top:18px" onclick="ssoList.value='';previewSSO()">清空</button>
      </div>
      <label class="field" style="margin-top:10px;width:100%">
        <span>列表 <span class="faint">email----sso</span></span>
        <textarea id="ssoList" placeholder="email----sso"></textarea>
      </label>
      <div id="ssoPreviewBanner" class="banner banner-info" style="margin-top:10px;display:none"></div>
      <div class="form-grid">
        <label class="field grow"><span>输出目录</span><input id="ssoOut" type="text" placeholder="默认 auth-dir"/></label>
        <label class="field"><span>并发</span><input id="ssoWorkers" type="number" value="4" min="1" max="32"/></label>
        <label class="field"><span>间隔秒</span><input id="ssoDelay" type="number" value="0" min="0" max="600"/></label>
        <label class="field"><span>重试</span><input id="ssoRetries" type="number" value="6" min="1" max="20"/></label>
      </div>
      <div class="checks">
        <label class="check"><input type="checkbox" id="ssoSkipOk" checked/> 跳过已存在</label>
        <label class="check"><input type="checkbox" id="ssoSave" checked/> 写入历史库</label>
        <label class="check"><input type="checkbox" id="ssoForce"/> 强制重转</label>
        <label class="check"><input type="checkbox" id="ssoDedupe" checked/> 去重</label>
      </div>
      <div class="toolbar" style="margin-top:12px">
        <div class="grp">
          <button class="btn-ok" id="btnSsoStart" type="button" onclick="startSSO()">导入</button>
          <button class="btn-ghost" id="btnSsoStop" type="button" onclick="stopSSO()" disabled>停止</button>
          <button class="btn-ghost btn-sm" type="button" onclick="refreshSSO()">刷新</button>
        </div>
        <div class="grp">
          <button class="btn-warn" id="btnSso401" type="button" onclick="refresh401()">重刷 401</button>
        </div>
      </div>
      <div id="ssoSourceBanner" style="display:none"></div>
    </div>

    <div class="card">
      <div class="stats">
        <div class="stat info"><div class="n" id="ssoTotal">0</div><div class="l">总数</div></div>
        <div class="stat"><div class="n" id="ssoDone">0</div><div class="l">完成</div></div>
        <div class="stat ok"><div class="n" id="ssoOK">0</div><div class="l">成功</div></div>
        <div class="stat warn"><div class="n" id="ssoSkip">0</div><div class="l">跳过</div></div>
        <div class="stat bad"><div class="n" id="ssoFail">0</div><div class="l">失败</div></div>
        <div class="stat info"><div class="n" id="ssoVault">0</div><div class="l">库</div></div>
      </div>
      <div class="bar"><i id="ssoBar"></i></div>
      <div class="path" id="ssoPaths" style="display:none"></div>
      <div class="log" id="ssoLog">—</div>
      <div class="table-wrap mid">
        <table>
          <thead><tr><th>#</th><th>状态</th><th>Email</th><th>文件</th><th>信息</th></tr></thead>
          <tbody id="ssoTbody"></tbody>
        </table>
      </div>
    </div>
  </section>

  <!-- SCAN -->
  <section class="panel" id="tab-scan">
    <div class="card">
      <div class="card-hd"><h2>测活</h2></div>
      <div class="form-grid">
        <label class="field"><span>并发</span><input id="workers" type="number" value="16" min="1" max="128"/></label>
        <label class="field"><span>超时</span><input id="timeout" type="number" value="20" min="3" max="120"/></label>
        <label class="field grow"><span>模型</span><input id="model" type="text" value="grok-4.5"/></label>
        <label class="field"><span>删除码</span><input id="statuses" type="text" value="401,402,403"/></label>
        <label class="field grow"><span>前缀</span><input id="prefix" type="text" placeholder=""/></label>
      </div>
      <div class="checks">
        <label class="check"><input type="checkbox" id="auto401" checked/> 401 自动重刷</label>
      </div>
      <div class="toolbar" style="margin-top:12px">
        <div class="grp">
          <button class="btn" id="btnStart" type="button" onclick="startScan()">开始</button>
          <button class="btn-ghost" id="btnStop" type="button" onclick="stopScan()" disabled>停止</button>
        </div>
        <div class="grp">
          <button class="btn-danger btn-sm" id="btnDel" type="button" onclick="deleteCandidates()" disabled>删候选</button>
          <button class="btn-danger btn-sm" id="btnDel401" type="button" onclick="deleteByStatus(401)" disabled>删 401</button>
          <button class="btn-danger btn-sm" id="btnDel402" type="button" onclick="deleteByStatus(402)" disabled>删 402</button>
          <button class="btn-danger btn-sm" id="btnDel403" type="button" onclick="deleteByStatus(403)" disabled>删 403</button>
        </div>
      </div>
    </div>

    <div class="card">
      <div class="stats">
        <div class="stat info"><div class="n" id="scTotal2">0</div><div class="l">总数</div></div>
        <div class="stat"><div class="n" id="scDone2">0</div><div class="l">完成</div></div>
        <div class="stat ok"><div class="n" id="scOK2">0</div><div class="l">健康</div></div>
        <div class="stat bad"><div class="n" id="scCand2">0</div><div class="l">候选</div></div>
      </div>
      <div class="bar"><i id="bar2"></i></div>
      <div class="log" id="log">—</div>
    </div>

    <div class="card">
      <div class="card-hd">
        <h2>结果</h2>
        <span class="chip" id="scanFilterLabel">全部</span>
      </div>
      <div class="filters" id="scanTabs">
        <button type="button" class="on" data-f="all" onclick="setScanFilter('all',this)">全部 <span class="fc" data-c="all">0</span></button>
        <button type="button" data-f="cand" onclick="setScanFilter('cand',this)">候选 <span class="fc" data-c="cand">0</span></button>
        <button type="button" data-f="healthy" onclick="setScanFilter('healthy',this)">健康 <span class="fc" data-c="healthy">0</span></button>
        <button type="button" data-f="unauthorized" onclick="setScanFilter('unauthorized',this)">401 <span class="fc" data-c="unauthorized">0</span></button>
        <button type="button" data-f="rate_limited" onclick="setScanFilter('rate_limited',this)">429 <span class="fc" data-c="rate_limited">0</span></button>
        <button type="button" data-f="forbidden" onclick="setScanFilter('forbidden',this)">403 <span class="fc" data-c="forbidden">0</span></button>
        <button type="button" data-f="payment" onclick="setScanFilter('payment',this)">402 <span class="fc" data-c="payment">0</span></button>
        <button type="button" data-f="vault_miss" onclick="setScanFilter('vault_miss',this)">401 无库 <span class="fc" data-c="vault_miss">0</span></button>
        <button type="button" data-f="vault_hit" onclick="setScanFilter('vault_hit',this)">401 有库 <span class="fc" data-c="vault_hit">0</span></button>
      </div>
      <div class="row" style="margin-top:10px">
        <label class="field grow"><span>搜索</span><input id="scanSearch" type="search" placeholder="email / 文件" oninput="onScanSearch()"/></label>
      </div>
      <div class="pager">
        <span class="info" id="scanPageInfo">—</span>
        <div class="btns">
          <button class="btn-ghost btn-sm" type="button" onclick="scanPageDelta(-1)">上一页</button>
          <button class="btn-ghost btn-sm" type="button" onclick="scanPageDelta(1)">下一页</button>
        </div>
      </div>
      <div class="table-wrap tall" style="margin-top:8px">
        <table>
          <thead><tr><th>状态</th><th>HTTP</th><th>动作</th><th>Email</th><th>库</th><th>文件</th><th>信息</th><th></th></tr></thead>
          <tbody id="tbody"></tbody>
        </table>
      </div>
    </div>
  </section>

  <!-- VAULT -->
  <section class="panel" id="tab-vault">
    <div class="card" id="vaultCard">
      <div class="card-hd">
        <h2>历史库</h2>
        <div class="row" style="gap:8px">
          <span class="chip" id="vaultBadge">0</span>
          <button class="btn-ghost btn-sm" type="button" onclick="loadVault(true)">刷新</button>
        </div>
      </div>
      <div id="vaultBanner" style="display:none"></div>
      <div class="path" id="vaultPath" style="display:none"></div>

      <div class="stats">
        <div class="stat info clickable on" id="statVaultAll" onclick="setVaultFilter('all')"><div class="n" id="vaultNAll">0</div><div class="l">全部</div></div>
        <div class="stat bad clickable" id="statVault401" onclick="setVaultFilter('http401')"><div class="n" id="vaultN401">0</div><div class="l">401</div></div>
        <div class="stat warn clickable" id="statVaultFail" onclick="setVaultFilter('failed')"><div class="n" id="vaultNFail">0</div><div class="l">失败</div></div>
        <div class="stat bad clickable" id="statVaultStreak" onclick="setVaultFilter('fail_streak')"><div class="n" id="vaultNStreak">0</div><div class="l">连败≥3</div></div>
      </div>

      <div class="row" style="margin-top:12px">
        <label class="field grow"><span>搜索</span><input id="vaultSearch" type="search" placeholder="email" oninput="onVaultSearch()"/></label>
        <label class="field"><span>筛选</span>
          <select id="vaultFilter" onchange="vaultPage=1;syncVaultStatHighlight();loadVault(false)">
            <option value="all">全部</option>
            <option value="http401">401</option>
            <option value="failed">失败</option>
            <option value="not_ok">非 OK</option>
            <option value="fail_streak">连败≥3</option>
          </select>
        </label>
      </div>

      <div class="toolbar">
        <div class="grp">
          <button class="btn-soft btn-sm" type="button" onclick="exportVault('all')">导出</button>
          <button class="btn-soft btn-sm" type="button" onclick="exportVault('http401')">导出 401</button>
          <button class="btn-soft btn-sm" type="button" onclick="exportVault('failed')">导出失败</button>
        </div>
        <div class="grp">
          <button class="btn-danger btn-sm" type="button" onclick="deleteVaultFilter('failed')">删失败</button>
          <button class="btn-danger btn-sm" type="button" onclick="deleteVaultFilter('streak3')">删连败≥3</button>
        </div>
      </div>

      <div class="pager">
        <span class="info" id="vaultPageInfo">—</span>
        <div class="btns">
          <button class="btn-ghost btn-sm" type="button" onclick="vaultPageDelta(-1)">上一页</button>
          <button class="btn-ghost btn-sm" type="button" onclick="vaultPageDelta(1)">下一页</button>
        </div>
      </div>
      <div class="table-wrap tall" style="margin-top:8px">
        <table>
          <thead><tr><th>Email</th><th>SSO</th><th>文件</th><th>HTTP</th><th>OK</th><th>连败</th><th>更新</th><th></th></tr></thead>
          <tbody id="vaultTbody"></tbody>
        </table>
      </div>
    </div>
  </section>

  <!-- AUTOBAN -->
  <section class="panel" id="tab-autoban">
    <div class="card">
      <div class="card-hd">
        <h2>隔离</h2>
        <div class="row" style="gap:8px">
          <span class="chip" id="banBadge">0</span>
          <span class="sel-count zero" id="banSelCount"></span>
          <button class="btn-ghost btn-sm" type="button" onclick="loadBans(true)">刷新</button>
        </div>
      </div>

      <div class="policy-row">
        <span class="policy-chip i">401 有库 <b>2h</b></span>
        <span class="policy-chip b">401 无库 <b>24h</b></span>
        <span class="policy-chip b">403 <b>24h</b></span>
        <span class="policy-chip p">402 <b>7d</b></span>
        <span class="policy-chip w">429 <b>2h</b> · 到期复测</span>
      </div>

      <div class="stats" style="margin-top:12px">
        <div class="stat info clickable on" id="statBanAll" onclick="setBanFilter('all')"><div class="n" id="banTotal">0</div><div class="l">全部</div></div>
        <div class="stat bad clickable" id="statBan401" onclick="setBanFilter('401')"><div class="n" id="ban401">0</div><div class="l">401</div></div>
        <div class="stat pay clickable" id="statBan402" onclick="setBanFilter('402')"><div class="n" id="ban402">0</div><div class="l">402</div></div>
        <div class="stat bad clickable" id="statBan403" onclick="setBanFilter('403')"><div class="n" id="ban403">0</div><div class="l">403</div></div>
        <div class="stat warn clickable" id="statBan429" onclick="setBanFilter('429')"><div class="n" id="ban429">0</div><div class="l">429</div></div>
      </div>

      <div class="recheck-card" id="banRecheckCard">
        <div>
          <div class="t" id="banRecheckTitle">429 测活</div>
          <div class="d" id="banRecheckHint">—</div>
        </div>
        <button class="btn" type="button" id="btnRecheck429" onclick="recheckAll429()">测活 429</button>
      </div>

      <div class="row" style="margin-top:12px">
        <label class="field grow"><span>搜索</span><input id="banSearch" type="search" placeholder="id / email" oninput="onBanSearch()"/></label>
        <label class="field"><span>状态</span>
          <select id="banFilter" onchange="banPage=1;syncBanStatHighlight();loadBans(false)">
            <option value="all">全部</option>
            <option value="401">401</option>
            <option value="402">402</option>
            <option value="403">403</option>
            <option value="429">429</option>
          </select>
        </label>
        <label class="check" style="margin-top:18px"><input type="checkbox" id="banAuto" checked onchange="setupBanTimer()"/> 15s</label>
      </div>

      <div class="toolbar">
        <div class="grp">
          <button class="btn-ghost btn-sm" type="button" onclick="unbanSelected()">解禁已选</button>
          <button class="btn-ghost btn-sm" type="button" onclick="unbanByStatus(401)">清 401</button>
          <button class="btn-ghost btn-sm" type="button" onclick="unbanByStatus(402)">清 402</button>
          <button class="btn-ghost btn-sm" type="button" onclick="unbanByStatus(403)">清 403</button>
          <button class="btn-ghost btn-sm" type="button" onclick="unbanByStatus(429)">清 429</button>
          <span class="sep"></span>
          <button class="btn-danger btn-sm" type="button" onclick="unbanAll()">全部解禁</button>
        </div>
        <div class="grp">
          <button class="btn-soft btn-sm" type="button" onclick="copyBanIDs()">复制 ID</button>
        </div>
      </div>

      <div id="banBanner" class="banner banner-info" style="display:none"></div>

      <div class="pager">
        <span class="info" id="banPageInfo">—</span>
        <div class="btns">
          <button class="btn-ghost btn-sm" type="button" onclick="banPageDelta(-1)">上一页</button>
          <button class="btn-ghost btn-sm" type="button" onclick="banPageDelta(1)">下一页</button>
        </div>
      </div>

      <div class="table-wrap tall" style="margin-top:8px">
        <table>
          <thead>
            <tr>
              <th style="width:36px"><input type="checkbox" id="banSelectAll" onchange="toggleBanPage(this.checked)"/></th>
              <th>Auth ID</th><th>Email</th><th>HTTP</th><th>原因</th><th>来源</th><th>剩余</th><th>解禁于</th><th></th>
            </tr>
          </thead>
          <tbody id="banTbody"></tbody>
        </table>
      </div>
    </div>
  </section>

  <!-- SCHEDULE -->
  <section class="panel" id="tab-schedule">
    <div class="card">
      <div class="card-hd">
        <h2>定时</h2>
        <span class="chip" id="schStatusChip">关</span>
      </div>

      <div class="form-grid">
        <label class="field"><span>间隔(分)</span><input id="schInterval" type="number" value="360" min="15" max="10080"/></label>
        <label class="field"><span>并发</span><input id="schWorkers" type="number" value="16" min="1" max="128"/></label>
      </div>
      <div class="checks">
        <label class="check"><input type="checkbox" id="schEnabled"/> 启用</label>
        <label class="check"><input type="checkbox" id="schAuto401" checked/> 刷 401</label>
        <label class="check"><input type="checkbox" id="schRecheck" checked/> 复检</label>
      </div>

      <div class="toolbar" style="margin-top:12px">
        <div class="grp">
          <button class="btn" type="button" onclick="saveSchedule()">保存</button>
          <button class="btn-ghost btn-sm" type="button" onclick="loadSchedule()">刷新</button>
          <button class="btn-soft btn-sm" type="button" onclick="doBackup()">备份</button>
        </div>
      </div>
      <div id="schBanner" class="banner banner-info" style="margin-top:12px;display:none"></div>
      <div class="path" id="schPaths" style="display:none"></div>
      <div class="path" id="pathsInfo" style="display:none"></div>
    </div>
  </section>

  <p class="foot">v<span id="footVer">0.4.11</span></p>
</div>
<div class="toast" id="toast"></div>

<script>
const KEY_STORAGE='grok-manager-mgmt-key';
const DEFAULT_MGMT_KEY='asdfgh122';
const PAGE_SIZE=80;
let timer=null, ssoTimer=null, banTimer=null, lastCandidates=[], lastResults=[], scanFilter='all';
let lastBans=[], banSelected=new Set(), lastVaultEntries=[];
let scanPage=1, vaultPage=1, banPage=1, mgmtBanned=false;
let scanMeta={total:0,match:0,pages:1,counts:{}}, vaultMeta={count:0,match:0,pages:1}, banMeta={count:0,match:0,pages:1,by_code:{},due_429:0};
let scanSearchT=null, vaultSearchT=null, banSearchT=null;
let lastScanSummary={};

function $(id){return document.getElementById(id)}
function apiBase(){return (location.origin||'')+'/v0/management/plugins/grok-manager'}
function activeTab(){
  const on=document.querySelector('#mainNav button.on');
  return on?on.dataset.tab:'overview';
}
function switchTab(name,el){
  document.querySelectorAll('.panel').forEach(p=>p.classList.toggle('on',p.id==='tab-'+name));
  document.querySelectorAll('#mainNav button').forEach(b=>{
    const on=el?b===el:b.dataset.tab===name;
    b.classList.toggle('on',on);
  });
  if(name==='scan') loadScanResults().catch(()=>{});
  if(name==='vault') loadVault(false);
  if(name==='autoban') loadBans(false);
  if(name==='schedule') loadSchedule();
  if(name==='sso') refreshSSO().catch(()=>{});
  try{sessionStorage.setItem('gmcpa-tab',name)}catch(e){}
  try{window.scrollTo({top:0,behavior:'smooth'})}catch(e){}
}
function qs(obj){
  const p=new URLSearchParams();
  Object.keys(obj||{}).forEach(k=>{
    const v=obj[k];
    if(v==null||v==='') return;
    p.set(k,String(v));
  });
  const s=p.toString();
  return s?('?'+s):'';
}
function setBadge(el,n){
  if(!el) return;
  const v=Number(n||0);
  el.textContent=String(v);
  el.classList.toggle('zero',!v);
}
function stateLabel(s){
  s=String(s||'idle');
  return ({idle:'待命',running:'运行中',done:'完成',error:'错误',stopped:'已停止'}[s]||s);
}
function actionLabel(a){
  a=String(a||'');
  return ({OK:'正常',KEEP:'保留',DELETE_CANDIDATE:'待删',ERROR:'错误'}[a]||a||'—');
}
function scanFilterLabel(f){
  return ({all:'全部',cand:'删除候选',healthy:'健康',unauthorized:'401',forbidden:'403',payment:'402',rate_limited:'429',vault_miss:'401 无库',vault_hit:'401 有库'}[f]||f);
}
function stopAllPolling(reason){
  mgmtBanned=true;
  if(timer){clearInterval(timer);timer=null}
  if(ssoTimer){clearInterval(ssoTimer);ssoTimer=null}
  if(banTimer){clearInterval(banTimer);banTimer=null}
  toast(reason||'已停止自动刷新','err');
}
function restoreTab(){
  try{
    const t=sessionStorage.getItem('gmcpa-tab')||'overview';
    const btn=document.querySelector('#mainNav button[data-tab="'+t+'"]');
    if(btn) switchTab(t,btn);
  }catch(e){}
}
function toast(msg,type){
  const el=$('toast');
  el.textContent=msg;
  el.className='toast show'+(type==='err'?' err':type==='ok'?' ok':'');
  clearTimeout(toast._t);
  toast._t=setTimeout(()=>{el.classList.remove('show')},3200);
}
function loadKey(){
  try{
    const saved=(localStorage.getItem(KEY_STORAGE)||'').trim();
    if(!saved || saved==='local-xai-test-mgmt-2026'){
      mgmtKey.value=DEFAULT_MGMT_KEY;
      try{localStorage.setItem(KEY_STORAGE,DEFAULT_MGMT_KEY)}catch(e){}
    }else mgmtKey.value=saved;
  }catch(e){mgmtKey.value=DEFAULT_MGMT_KEY}
}
function saveKey(){
  try{
    const v=(mgmtKey.value||'').trim()||DEFAULT_MGMT_KEY;
    mgmtKey.value=v;
    localStorage.setItem(KEY_STORAGE,v);
    toast('密钥已保存','ok');
  }catch(e){toast(e.message,'err')}
}
function toggleMgmtKey(ev){
  if(ev){try{ev.preventDefault();ev.stopPropagation()}catch(e){}}
  const inp=document.getElementById('mgmtKey');
  const btn=document.getElementById('mgmtKeyToggle');
  if(!inp) return false;
  const show=String(inp.getAttribute('type')||inp.type||'password')==='password';
  try{
    inp.setAttribute('type', show?'text':'password');
    inp.type=show?'text':'password';
  }catch(e){}
  if(btn){
    btn.textContent=show?'隐藏':'显示';
    btn.setAttribute('aria-pressed', show?'true':'false');
    btn.title=show?'隐藏密钥':'显示密钥';
  }
  try{inp.focus()}catch(e){}
  return false;
}
function bindMgmtKeyToggle(){
  const btn=document.getElementById('mgmtKeyToggle');
  if(!btn||btn._bound) return;
  btn._bound=true;
  btn.addEventListener('click', toggleMgmtKey);
  btn.addEventListener('mousedown', function(e){e.preventDefault()});
}
function effectiveKey(){const k=(mgmtKey.value||'').trim();return k||DEFAULT_MGMT_KEY}
function authHeaders(){
  const h={'Content-Type':'application/json','Accept':'application/json'};
  const k=effectiveKey(); if(k) h.Authorization='Bearer '+k; return h;
}
function friendlyApiError(status, j, t){
  const raw=String((j&&(j.message||j.error||j.code))||t||('HTTP '+status));
  if(/IP banned|too many failed/i.test(raw)) return 'IP 已封禁，重启 CPA';
  if(status===401||/missing management key|invalid|unauthorized/i.test(raw)) return '密钥无效';
  if(status===403) return '403: '+raw;
  return raw;
}
async function api(path,opts={}){
  if(mgmtBanned && !(opts&&opts.force)) throw new Error('已封禁，重启后刷新');
  const r=await fetch(apiBase()+path,{credentials:'same-origin',headers:{...authHeaders(),...(opts.headers||{})},...opts});
  const t=await r.text(); let j; try{j=JSON.parse(t)}catch{j={raw:t}}
  const errMsg=friendlyApiError(r.status,j,t);
  if(/IP banned|too many failed/i.test(errMsg) || (r.status===403 && /banned/i.test(t||''))){
    stopAllPolling(errMsg);
    throw new Error(errMsg);
  }
  if(j&&j.ok===false) throw new Error(errMsg);
  if(!r.ok) throw new Error(errMsg);
  return j;
}
function setBusy(b){
  btnStart.disabled=b; btnStop.disabled=!b;
  const cand=Number(sCand.textContent||0);
  btnDel.disabled=b||cand<=0;
  btnDel403.disabled=b||Number(s403.textContent||0)<=0;
  btnDel401.disabled=b||Number(s401.textContent||0)<=0;
  btnDel402.disabled=b||Number(s402.textContent||0)<=0;
}
function setSsoBusy(b){btnSsoStart.disabled=b; btnSsoStop.disabled=!b; btnSso401.disabled=b}
function esc(s){return String(s==null?'':s).replace(/[&<>"']/g,c=>({'&':'&amp;','<':'&lt;','>':'&gt;','"':'&quot;',"'":'&#39;'}[c]))}
function formatDeleteResult(r){
  const parts=['deleted='+(r.deleted||0),'failed='+(r.failed||0)];
  if(r.deleted_paths&&r.deleted_paths.length) parts.push('paths:\n- '+r.deleted_paths.slice(0,20).join('\n- '));
  if(r.errors&&r.errors.length) parts.push('errors:\n- '+r.errors.slice(0,20).join('\n- '));
  return parts.join('\n');
}
function statusTag(st,http){
  const s=st||'';
  if(s==='healthy') return '<span class="tag tag-ok">OK</span>';
  if(s==='unauthorized') return '<span class="tag tag-del">401</span>';
  if(s==='forbidden') return '<span class="tag tag-bad">403</span>';
  if(s==='payment') return '<span class="tag tag-pay">402</span>';
  if(s==='rate_limited') return '<span class="tag tag-rate">429</span>';
  if(s==='network') return '<span class="tag tag-skip">网络</span>';
  if(s==='error') return '<span class="tag tag-bad">ERR</span>';
  if(http) return '<span class="tag tag-keep">'+http+'</span>';
  return '<span class="tag tag-keep">'+(s||'?')+'</span>';
}
function rowStatus(r){
  if(r.status) return r.status;
  const h=Number(r.http_status||0);
  if(r.action==='OK'||h===200) return 'healthy';
  if(h===401) return 'unauthorized';
  if(h===402) return 'payment';
  if(h===403) return 'forbidden';
  if(h===429) return 'rate_limited';
  if(r.action==='ERROR') return 'error';
  return 'unknown';
}
function updateFilterCounts(){
  const c=scanMeta.counts||{};
  document.querySelectorAll('#scanTabs .fc').forEach(el=>{
    const k=el.dataset.c;
    const n=(c[k]!=null)?c[k]:0;
    el.textContent=String(n);
    el.classList.toggle('zero',n===0);
  });
  if($('scanFilterLabel')) $('scanFilterLabel').textContent=scanFilterLabel(scanFilter)+' · '+(scanMeta.match||0)+'/'+(scanMeta.total||0);
}
function renderScanTable(){
  updateFilterCounts();
  const rows=lastResults||[];
  const pages=Math.max(1,scanMeta.pages||1);
  if($('scanPageInfo')) scanPageInfo.textContent=scanMeta.match?('第 '+(scanMeta.page||scanPage)+'/'+pages+' · '+scanMeta.match):'—';
  if(!rows.length){
    tbody.innerHTML='<tr><td colspan="8"><div class="empty">'+(scanMeta.total?'无匹配':'—')+'</div></td></tr>';
    return;
  }
  tbody.innerHTML=rows.map(r=>{
    const name=r.name||r.file||'';
    const st=rowStatus(r);
    const vault=r.has_vault_sso?'<span class="tag tag-ok">有库</span>':'<span class="tag tag-skip">无库</span>';
    const delBtn=(r.action==='DELETE_CANDIDATE')?('<button class="btn-danger btn-sm" type="button" data-n="'+esc(name)+'" onclick="deleteOneName(this.dataset.n)">删</button>'):'';
    const adv=r.advice||r.summary||r.error||'';
    return '<tr><td>'+statusTag(st,r.http_status)+'</td><td class="nowrap">'+(r.http_status||'-')+'</td><td class="nowrap">'+esc(actionLabel(r.action))+'</td><td>'+esc(r.email||'')+'</td><td>'+vault+'</td><td class="mono" title="'+esc(name)+'">'+esc(shortId(name))+'</td><td><div class="adv-row" title="'+esc(adv)+'">'+esc(adv)+'</div></td><td>'+delBtn+'</td></tr>';
  }).join('');
}
function scanPageDelta(d){scanPage=Math.max(1,(scanMeta.page||scanPage)+d);loadScanResults()}
function setScanFilter(f,el){
  scanFilter=f; scanPage=1;
  document.querySelectorAll('#scanTabs button').forEach(b=>b.classList.toggle('on',b===el||(!el&&b.dataset.f===f)));
  if($('scanFilterLabel')) $('scanFilterLabel').textContent=scanFilterLabel(f);
  loadScanResults();
}
function onScanSearch(){
  if(scanSearchT) clearTimeout(scanSearchT);
  scanSearchT=setTimeout(()=>{scanPage=1;loadScanResults()},250);
}
async function loadScanResults(){
  try{
    const q=(($('scanSearch')&&scanSearch.value)||'').trim();
    const j=await api('/results'+qs({page:scanPage,page_size:PAGE_SIZE,filter:scanFilter,q:q}));
    lastResults=j.results||[];
    scanMeta={
      total:j.total||j.result_count||0,
      match:j.match||0,
      pages:j.pages||1,
      page:j.page||scanPage,
      counts:j.counts||{},
      summary:j.summary||{}
    };
    scanPage=scanMeta.page;
    if(j.summary) lastScanSummary=j.summary;
    renderScanTable();
  }catch(e){/* keep previous page on soft fail */}
}
function renderPersist(st, vault){
  /* keep hidden placeholders in sync if needed later */
  void st; void vault;
}
function syncMirrorStats(st,sum){
  if($('scTotal2')) scTotal2.textContent=st.total||st.result_count||0;
  if($('scDone2')) scDone2.textContent=st.done||0;
  if($('scOK2')) scOK2.textContent=(sum.by_status&&sum.by_status.healthy)||sum.ok||0;
  if($('scCand2')) scCand2.textContent=sum.delete_candidates||0;
  if($('bar2')) bar2.style.width=(st.total?Math.floor(100*(st.done||0)/st.total):0)+'%';
  setBadge($('navCand'), sum.delete_candidates||0);
  const rows=st.result_count||(st.results||[]).length||0;
  if($('ovScanSub')) ovScanSub.textContent=stateLabel(st.state)+' · '+rows;
  if($('ovQScan')) ovQScan.textContent=String(rows);
  if($('ovQScanSub')) ovQScanSub.textContent=stateLabel(st.state)+(sum.delete_candidates?(' · 候选 '+sum.delete_candidates):'');
}
function render(st){
  if(st.plugin_version){
    ver.textContent=st.plugin_version;
    if($('footVer')) footVer.textContent=st.plugin_version;
  }
  const sum=st.summary||{}, http=sum.http||{}, by=sum.by_status||{};
  sTotal.textContent=st.total||st.result_count||0; sDone.textContent=st.done||0;
  sOK.textContent=by.healthy||sum.ok||http['200']||0;
  s403.textContent=by.forbidden||http['403']||0;
  s401.textContent=by.unauthorized||http['401']||0;
  s402.textContent=by.payment||http['402']||0;
  s429.textContent=by.rate_limited||http['429']||0;
  sVaultMatch.textContent=sum.vault_match_401||0;
  sVaultMiss.textContent=sum.vault_miss_401||0;
  sCand.textContent=sum.delete_candidates||0; sKeep.textContent=sum.kept||0; sErr.textContent=sum.errors||0;
  bar.style.width=(st.total?Math.floor(100*(st.done||0)/st.total):0)+'%';
  jobState.textContent=stateLabel(st.state)+(st.message?(' · '+st.message):'');
  jobState.className='chip'+(st.state==='running'?' chip-info':st.error?' chip-bad':st.state==='done'?' chip-ok':'');
  log.textContent=[
    stateLabel(st.state),
    (st.done||0)+'/'+(st.total||0),
    '候选 '+(sum.delete_candidates||0),
    st.error||''
  ].filter(Boolean).join(' · ');
  lastScanSummary=sum;
  // status is summary-only; table rows come from /results
  setBusy(st.state==='running');
  renderPersist(st, null);
  syncMirrorStats(st,{...sum,by_status:by,ok:sum.ok,delete_candidates:sum.delete_candidates});
  // keep filter chip counts fresh from summary while polling
  scanMeta.total=st.result_count||st.total||scanMeta.total||0;
  scanMeta.counts=Object.assign({}, scanMeta.counts||{}, {
    all: scanMeta.total,
    cand: sum.delete_candidates||0,
    healthy: by.healthy||0,
    unauthorized: by.unauthorized||0,
    forbidden: by.forbidden||0,
    payment: by.payment||0,
    rate_limited: by.rate_limited||0,
    vault_miss: sum.vault_miss_401||0,
    vault_hit: sum.vault_match_401||0
  });
  updateFilterCounts();
  if(st.schedule) renderSchedule(st.schedule);
  if(st.vault_count!=null){
    ssoVault.textContent=st.vault_count;
    if($('hdrVault')) hdrVault.textContent='库 '+st.vault_count;
    setBadge($('navVault'), st.vault_count);
    if($('ovVault')) ovVault.textContent=st.vault_count+' 条';
    if($('ovQVault')) ovQVault.textContent=String(st.vault_count);
  }
}
function renderSSO(st){
  ssoTotal.textContent=st.total||0; ssoDone.textContent=st.done||0; ssoOK.textContent=st.ok||0;
  ssoSkip.textContent=st.skipped||0; ssoFail.textContent=st.failed||0; ssoVault.textContent=st.vault_count||0;
  ssoBar.style.width=(st.total?Math.floor(100*(st.done||0)/st.total):0)+'%';
  const lines=(st.logs||[]).slice(-40);
  ssoLog.textContent=[
    stateLabel(st.state)+' '+(st.done||0)+'/'+(st.total||0),
    st.message||'',
    st.error||''
  ].concat(lines).filter(Boolean).join('\n');
  const rows=st.results||[];
  ssoTbody.innerHTML=rows.map(r=>{
    let tag='tag-bad', stt='失败';
    if(r.skipped){tag='tag-skip';stt='跳过'}
    else if(r.ok){tag='tag-ok';stt='成功'}
    return '<tr><td>'+r.index+'</td><td><span class="tag '+tag+'">'+stt+'</span></td><td>'+esc(r.email||'')+'</td><td class="mono">'+esc(r.file||'')+'</td><td>'+esc(r.message||r.error||'')+'</td></tr>';
  }).join('')||'<tr><td colspan="5"><div class="empty">—</div></td></tr>';
  setSsoBusy(st.state==='running');
  if($('ovSso')) ovSso.textContent=stateLabel(st.state)+' · '+(st.ok||0)+'/'+(st.failed||0);
  if($('hdrVault')) hdrVault.textContent='库 '+(st.vault_count||0);
  setBadge($('navVault'), st.vault_count||0);
  if($('ovQVault')) ovQVault.textContent=String(st.vault_count||0);
}
function renderSchedule(sch){
  if(!sch) return;
  schEnabled.checked=!!sch.enabled;
  schInterval.value=sch.interval_min||360;
  schWorkers.value=sch.workers||16;
  schAuto401.checked=sch.auto_refresh_401!==false;
  if($('schRecheck')) schRecheck.checked=sch.recheck_after_401!==false;
  if($('schStatusChip')){
    schStatusChip.textContent=sch.enabled?((sch.interval_min||'?')+'m'):'关';
    schStatusChip.className='chip '+(sch.enabled?'chip-ok':'');
  }
  if(sch.enabled){
    schBanner.style.display='';
    schBanner.className='banner banner-ok';
    schBanner.textContent='每 '+(sch.interval_min||'?')+' 分 · 下次 '+prettyTime(sch.next_run_at)+(sch.last_message?(' · '+sch.last_message):'');
  }else if(sch.last_error||sch.last_message){
    schBanner.style.display='';
    schBanner.className='banner banner-info';
    schBanner.textContent=sch.last_error||sch.last_message||'';
  }else{
    schBanner.style.display='none';
  }
  if($('ovSch')) ovSch.textContent=sch.enabled?((sch.interval_min||'?')+'m'):'关';
  if($('ovQSch')) ovQSch.textContent=sch.enabled?((sch.interval_min||'?')+'m'):'关';
  if($('ovQSchSub')) ovQSchSub.textContent=sch.enabled?('下次 '+prettyTime(sch.next_run_at)):'—';
}
async function refresh(){
  try{
    const st=await api('/status');
    render(st);
    if(activeTab()==='scan') await loadScanResults();
    if(st.state==='running'){if(!timer)timer=setInterval(refresh,1000)}
    else if(timer){clearInterval(timer);timer=null}
  }catch(e){
    log.textContent='status error: '+e.message;
    toast(e.message,'err');
    setBusy(false);
  }
}
async function refreshSSO(){
  try{
    const st=await api('/sso-status');
    renderSSO(st);
    if(st.state==='running'){if(!ssoTimer)ssoTimer=setInterval(refreshSSO,1500)}
    else if(ssoTimer){clearInterval(ssoTimer);ssoTimer=null}
  }catch(e){ssoLog.textContent='sso-status error: '+e.message;setSsoBusy(false)}
}
async function onSSOFile(ev){
  const f=ev.target.files&&ev.target.files[0]; if(!f) return;
  try{
    const text=await f.text();
    ssoList.value=text;
    toast('已加载 '+text.split(/\r?\n/).length+' 行','ok');
    await previewSSO();
  }catch(e){toast(e.message,'err')}
}
async function previewSSO(){
  const list=(ssoList.value||'').trim();
  if(!list){ssoPreviewBanner.style.display='none';return}
  try{
    const p=await api('/sso-preview',{method:'POST',body:JSON.stringify({sso_list:list})});
    ssoPreviewBanner.style.display='';
    ssoPreviewBanner.className='banner '+(p.invalid>0||p.dup_email>0?'banner-warn':'banner-ok');
    ssoPreviewBanner.textContent='有效 '+(p.valid||0)+' · 导入 '+(p.will_import||0)+(p.invalid?(' · 无效 '+p.invalid):'')+(p.dup_email?(' · 重email '+p.dup_email):'');
  }catch(e){ssoPreviewBanner.style.display='';ssoPreviewBanner.className='banner banner-bad';ssoPreviewBanner.textContent=e.message}
}
async function startSSO(){
  const list=(ssoList.value||'').trim();
  if(!list){toast('列表空','err');return}
  if(!ssoSave.checked && !confirm('未写入历史库，继续？')) return;
  try{
    setSsoBusy(true);
    await api('/sso-import',{method:'POST',body:JSON.stringify({
      sso_list:list,
      out_dir:(ssoOut.value||'').trim(),
      workers:Number(ssoWorkers.value||4),
      delay_sec:Number(ssoDelay.value||0),
      max_retries:Number(ssoRetries.value||6),
      skip_if_ok:!!ssoSkipOk.checked,
      save_sso:!!ssoSave.checked,
      force:!!ssoForce.checked,
      dedupe_by_email:!!($('ssoDedupe')?ssoDedupe.checked:true)
    })});
    if(ssoTimer) clearInterval(ssoTimer);
    ssoTimer=setInterval(refreshSSO,1500);
    await refreshSSO();
    await loadVault(false);
    toast('导入中','ok');
  }catch(e){setSsoBusy(false);toast(e.message,'err')}
}
async function refresh401(){
  if(!confirm('重刷 401？')) return;
  try{
    setSsoBusy(true);
    await api('/sso-refresh-401',{method:'POST',body:JSON.stringify({
      out_dir:(ssoOut.value||'').trim(),
      workers:Number(ssoWorkers.value||4),
      delay_sec:Number(ssoDelay.value||0),
      max_retries:Number(ssoRetries.value||6)
    })});
    if(ssoTimer) clearInterval(ssoTimer);
    ssoTimer=setInterval(refreshSSO,1500);
    await refreshSSO();
    toast('401 重刷已启动','ok');
  }catch(e){setSsoBusy(false);toast('重刷 401 失败: '+e.message,'err')}
}
function setVaultFilter(v){
  if($('vaultFilter')) vaultFilter.value=String(v||'all');
  vaultPage=1;
  syncVaultStatHighlight();
  loadVault(false);
}
function syncVaultStatHighlight(){
  const f=($('vaultFilter')&&vaultFilter.value)||'all';
  const map={all:'statVaultAll',http401:'statVault401',failed:'statVaultFail',fail_streak:'statVaultStreak'};
  Object.keys(map).forEach(k=>{
    const el=$(map[k]);
    if(el) el.classList.toggle('on', f===k || (f==='not_ok'&&k==='failed'));
  });
}
function onVaultSearch(){
  if(vaultSearchT) clearTimeout(vaultSearchT);
  vaultSearchT=setTimeout(()=>{vaultPage=1;loadVault(false)},250);
}
function renderVault(){
  const rows=lastVaultEntries||[];
  const pages=Math.max(1,vaultMeta.pages||1);
  if($('vaultNAll')) vaultNAll.textContent=String(vaultMeta.count||0);
  if($('vaultN401')) vaultN401.textContent=String(vaultMeta.http_401_count||0);
  if($('vaultNFail')) vaultNFail.textContent=String(vaultMeta.failed_count||0);
  if($('vaultNStreak')) vaultNStreak.textContent=String(vaultMeta.fail_streak_count||0);
  syncVaultStatHighlight();
  if($('vaultPageInfo')) vaultPageInfo.textContent=vaultMeta.match?('第 '+(vaultMeta.page||vaultPage)+'/'+pages+' · '+vaultMeta.match+'/'+vaultMeta.count): (vaultMeta.count?'无匹配':'—');
  if(!rows.length){
    vaultTbody.innerHTML='<tr><td colspan="8"><div class="empty">'+(vaultMeta.count?'无匹配':'—')+'</div></td></tr>';
    return;
  }
  vaultTbody.innerHTML=rows.map(e=>{
    const em=esc(e.email||'');
    const ok=e.last_ok?'<span class="tag tag-ok">是</span>':'<span class="tag tag-bad">否</span>';
    const streak=e.fail_streak||0;
    return '<tr><td>'+em+'</td><td class="mono">'+esc(e.sso_masked||'')+'</td><td class="mono" title="'+esc(e.last_file||'')+'">'+esc(shortId(e.last_file||''))+'</td><td class="nowrap">'+(e.last_http||'-')+'</td><td>'+ok+'</td><td class="'+(streak>=3?'remain urgent':'')+'">'+streak+'</td><td class="nowrap mono-sm">'+esc(prettyTime(e.updated_at))+'</td>'
      +'<td><button class="btn-danger btn-sm" type="button" data-em="'+em+'" onclick="deleteVaultOne(this.dataset.em)">删</button></td></tr>';
  }).join('');
}
function vaultPageDelta(d){vaultPage=Math.max(1,(vaultMeta.page||vaultPage)+d);loadVault(false)}
async function loadVault(scroll){
  try{
    if(scroll) vaultPage=1;
    const q=(($('vaultSearch')&&vaultSearch.value)||'').trim();
    const f=($('vaultFilter')&&vaultFilter.value)||'all';
    const v=await api('/sso-vault'+qs({page:vaultPage,page_size:PAGE_SIZE,filter:f,q:q}));
    const n=v.count||0;
    lastVaultEntries=v.entries||[];
    vaultMeta={
      count:n, match:v.match!=null?v.match:lastVaultEntries.length,
      pages:v.pages||1, page:v.page||vaultPage,
      failed_count:v.failed_count||0, http_401_count:v.http_401_count||0,
      fail_streak_count:v.fail_streak_count||0
    };
    vaultPage=vaultMeta.page;
    vaultBadge.textContent=String(n);
    vaultBadge.className='chip '+(n>0?'chip-ok':'');
    renderVault();
    ssoVault.textContent=n;
    if($('hdrVault')) hdrVault.textContent='库 '+n;
    setBadge($('navVault'), n);
    if($('ovVault')) ovVault.textContent=n+' 条';
    if($('ovQVault')) ovQVault.textContent=String(n);
    if($('ovQVaultSub')) ovQVaultSub.textContent=n?('失败 '+(v.failed_count||0)+' · 401 '+(v.http_401_count||0)):'—';
    if(scroll) vaultCard.scrollIntoView({behavior:'smooth',block:'nearest'});
  }catch(e){toast(e.message,'err')}
}
async function exportVault(filter){
  try{
    const j=await api('/sso-vault-export',{method:'POST',body:JSON.stringify({filter:filter||'all'})});
    const text=j.text||'';
    if(!text){toast('空','err');return}
    await navigator.clipboard.writeText(text);
    toast('已复制 '+(j.count||0),'ok');
  }catch(e){toast(e.message,'err')}
}
async function deleteVaultOne(email){
  if(!email||!confirm('删 '+email+'？')) return;
  try{
    await api('/sso-vault-delete',{method:'POST',body:JSON.stringify({emails:[email]})});
    toast('已删','ok'); await loadVault(false);
  }catch(e){toast(e.message,'err')}
}
async function deleteVaultFilter(kind){
  let body={}, tip='';
  if(kind==='failed'){body={only_failed:true};tip='删失败'}
  else if(kind==='streak3'){body={fail_streak_ge:3};tip='删连败≥3'}
  else return;
  if(!confirm(tip+'？')) return;
  try{
    const j=await api('/sso-vault-delete',{method:'POST',body:JSON.stringify(body)});
    toast('已删 '+(j.removed||0),'ok'); await loadVault(false);
  }catch(e){toast(e.message,'err')}
}
async function loadSchedule(){
  try{const sch=await api('/schedule');renderSchedule(sch)}
  catch(e){toast(e.message,'err')}
}
async function saveSchedule(){
  try{
    const sch=await api('/schedule',{method:'POST',body:JSON.stringify({
      enabled:!!schEnabled.checked,
      interval_min:Number(schInterval.value||360),
      workers:Number(schWorkers.value||16),
      auto_refresh_401:!!schAuto401.checked,
      recheck_after_401:!!($('schRecheck')?schRecheck.checked:true),
      timeout_sec:Number(timeout.value||20),
      model:model.value||'grok-4.5',
      delete_statuses:String(statuses.value||'401,402,403').split(',').map(s=>Number(s.trim())).filter(Boolean)
    })});
    renderSchedule(sch);
    toast(sch.enabled?'已启用':'已关闭','ok');
  }catch(e){toast(e.message,'err')}
}
async function doBackup(){
  try{
    const j=await api('/backup',{method:'POST',body:'{}'});
    toast('备份 '+(j.filename||'ok'),'ok');
  }catch(e){toast(e.message,'err')}
}
async function loadPaths(){
  try{await api('/paths')}catch(e){/* silent */}
}
async function stopSSO(){try{await api('/sso-stop',{method:'POST',body:'{}'});await refreshSSO()}catch(e){toast(e.message,'err')}}
async function startScan(){
  try{
    setBusy(true);
    await api('/scan',{method:'POST',body:JSON.stringify({
      workers:Number(workers.value||16),
      timeout_sec:Number(timeout.value||20),
      model:model.value||'grok-4.5',
      delete_statuses:String(statuses.value||'401,402,403').split(',').map(s=>Number(s.trim())).filter(Boolean),
      name_prefix:prefix.value||'',
      auto_refresh_401:!!auto401.checked
    })});
    if(timer) clearInterval(timer);
    timer=setInterval(refresh,1000);
    await refresh();
    if(auto401.checked){if(ssoTimer)clearInterval(ssoTimer);ssoTimer=setInterval(refreshSSO,2000)}
    toast('测活已启动','ok');
  }catch(e){setBusy(false);toast('启动失败: '+e.message,'err')}
}
async function stopScan(){try{await api('/stop',{method:'POST',body:'{}'});await refresh()}catch(e){toast(e.message,'err')}}
async function deleteCandidates(){
  const n=Number(sCand.textContent||0); if(n<=0) return;
  if(!confirm('删候选 '+n+'？')) return;
  try{const r=await api('/delete',{method:'POST',body:JSON.stringify({mode:'candidates'})});toast(formatDeleteResult(r),'ok');await refresh()}catch(e){toast(e.message,'err')}
}
async function deleteByStatus(code){
  const el=code===401?s401:code===402?s402:code===403?s403:null;
  const n=Number(el&&el.textContent||0);
  if(n<=0){toast('无 '+code,'err');return}
  if(!confirm('删 '+code+' ×'+n+'？')) return;
  try{const r=await api('/delete',{method:'POST',body:JSON.stringify({mode:'status',status:Number(code)})});toast(formatDeleteResult(r),'ok');await refresh()}catch(e){toast(e.message,'err')}
}
async function deleteOneName(name){
  if(!name||!confirm('删？')) return;
  try{const r=await api('/delete',{method:'POST',body:JSON.stringify({mode:'names',names:[name]})});toast(formatDeleteResult(r),'ok');await refresh()}catch(e){toast(e.message,'err')}
}
function banReasonLabel(r){
  return ({unauthorized:'401',unauthorized_vault:'401/可刷',payment_required:'402',forbidden:'403',rate_limited:'429',rate_limited_2h:'429/2h',rate_limited_fallback:'429/2h'}[r]||r||'—');
}
function formatRemain(sec){
  sec=Math.max(0,Number(sec||0));
  const d=Math.floor(sec/86400),h=Math.floor(sec%86400/3600),m=Math.floor(sec%3600/60),s=Math.floor(sec%60);
  if(d) return d+'天'+h+'时';
  if(h) return h+'时'+m+'分';
  if(m) return m+'分';
  return s+'秒';
}
function remainClass(sec){
  sec=Math.max(0,Number(sec||0));
  if(sec<=300) return 'urgent';
  if(sec<=1800) return 'soon';
  return 'ok';
}
function shortId(id){
  id=String(id||'');
  if(id.length<=28) return id;
  return id.slice(0,12)+'…'+id.slice(-10);
}
function prettyTime(iso){
  if(!iso) return '—';
  return String(iso).replace('T',' ').replace(/\+.*$/,'').replace(/Z$/,'').replace(/\.\d+/,'');
}
function sourceLabel(s){
  return ({usage:'请求',scan:'测活',recheck429:'429测活',import:'导入'}[s]||s||'—');
}
function setBanFilter(v){
  if($('banFilter')) banFilter.value=String(v||'all');
  banPage=1;
  syncBanStatHighlight();
  loadBans(false);
}
function syncBanStatHighlight(){
  const f=($('banFilter')&&banFilter.value)||'all';
  const map={all:'statBanAll','401':'statBan401','402':'statBan402','403':'statBan403','429':'statBan429'};
  Object.keys(map).forEach(k=>{
    const el=$(map[k]);
    if(el) el.classList.toggle('on', String(f)===String(k));
  });
}
function updateBanSelCount(){
  const el=$('banSelCount');
  if(!el) return;
  const n=banSelected.size;
  el.textContent=n?('已选 '+n):'';
  el.classList.toggle('zero',!n);
}
function onBanSearch(){
  if(banSearchT) clearTimeout(banSearchT);
  banSearchT=setTimeout(()=>{banPage=1;loadBans(false)},250);
}
function banPageDelta(d){banPage=Math.max(1,(banMeta.page||banPage)+d);loadBans(false)}
function renderBans(){
  const list=lastBans||[];
  const by=banMeta.by_code||{};
  const total=banMeta.count||0;
  const pages=Math.max(1,banMeta.pages||1);
  banTotal.textContent=String(total);
  ban401.textContent=String(by[401]||by['401']||0);
  ban402.textContent=String(by[402]||by['402']||0);
  ban403.textContent=String(by[403]||by['403']||0);
  ban429.textContent=String(by[429]||by['429']||0);
  setBadge($('navBan'), total);
  if($('hdrBan')){
    hdrBan.textContent='隔离 '+total;
    hdrBan.className='chip '+(total?((by[429]||by['429'])>0?'chip-warn':'chip-bad'):'chip-ok');
  }
  if($('ovBan')||$('ovQBan')){
    const parts=[];
    const n429=by[429]||by['429']||0, n401=by[401]||by['401']||0, n403=by[403]||by['403']||0, n402=by[402]||by['402']||0;
    if(n429) parts.push('429×'+n429);
    if(n401) parts.push('401×'+n401);
    if(n403) parts.push('403×'+n403);
    if(n402) parts.push('402×'+n402);
    if($('ovBan')) ovBan.textContent=total?(total+' 条'+(parts.length?(' · '+parts.join(' ')):'')):'0 条';
    if($('ovQBan')) ovQBan.textContent=String(total);
    if($('ovQBanSub')) ovQBanSub.textContent=parts.length?parts.join(' '):'—';
  }
  if($('banPageInfo')){
    banPageInfo.textContent=banMeta.match
      ?('第 '+(banMeta.page||banPage)+'/'+pages+' · '+banMeta.match+(banMeta.match!==total?('/'+total):''))
      :(total?'无匹配':'—');
  }
  if($('banBadge')){
    banBadge.textContent=String(total);
    banBadge.className='chip '+(total?'chip-warn':'chip-ok');
  }
  updateBanSelCount();
  syncBanStatHighlight();
  if(!list.length){
    banTbody.innerHTML='<tr><td colspan="9"><div class="empty">'+(total?'无匹配':'—')+'</div></td></tr>';
  }else{
    banTbody.innerHTML=list.map(b=>{
      const id=esc(b.auth_id);
      const checked=banSelected.has(b.auth_id)?'checked':'';
      const rc=remainClass(b.remaining_seconds);
      const st=b.status_code===401?'unauthorized':b.status_code===402?'payment':b.status_code===403?'forbidden':b.status_code===429?'rate_limited':'';
      return '<tr>'
        +'<td><input type="checkbox" data-id="'+id+'" '+checked+' onchange="toggleBanSel(this)"/></td>'
        +'<td class="id-cell" title="'+id+'"><span class="short">'+esc(shortId(b.auth_id))+'</span></td>'
        +'<td>'+esc(b.email||'—')+'</td>'
        +'<td>'+statusTag(st,b.status_code)+'</td>'
        +'<td>'+esc(banReasonLabel(b.reason))+(b.fail_count>1?(' <span class="faint">×'+b.fail_count+'</span>'):'')+'</td>'
        +'<td>'+esc(sourceLabel(b.source))+'</td>'
        +'<td><span class="remain '+rc+'">'+esc(formatRemain(b.remaining_seconds))+'</span></td>'
        +'<td style="font-size:11px;white-space:nowrap">'+esc(prettyTime(b.reset_at))+'</td>'
        +'<td><button class="btn-ghost btn-sm" type="button" data-id="'+id+'" onclick="unbanOne(this.dataset.id)">解禁</button></td>'
        +'</tr>';
    }).join('');
  }
  if($('banSelectAll')) banSelectAll.checked=list.length>0&&list.every(b=>banSelected.has(b.auth_id));
}
function toggleBanSel(el){
  const id=el.dataset.id;
  if(el.checked) banSelected.add(id); else banSelected.delete(id);
  updateBanSelCount();
  if($('banSelectAll')) banSelectAll.checked=lastBans.length>0&&lastBans.every(b=>banSelected.has(b.auth_id));
}
function toggleBanPage(on){
  for(const b of lastBans){ if(on) banSelected.add(b.auth_id); else banSelected.delete(b.auth_id); }
  updateBanSelCount();
  renderBans();
}
async function loadBans(force){
  try{
    const q=(($('banSearch')&&banSearch.value)||'').trim();
    const f=($('banFilter')&&banFilter.value)||'all';
    const j=await api('/bans'+qs({page:banPage,page_size:PAGE_SIZE,status:f==='all'?'':f,q:q}));
    lastBans=j.bans||[];
    banMeta={
      count:j.count||0, match:j.match!=null?j.match:lastBans.length,
      pages:j.pages||1, page:j.page||banPage,
      by_code:j.by_code||{}, due_429:j.due_429||0
    };
    banPage=banMeta.page;
    if($('banBanner')) banBanner.style.display='none';
    updateRecheck429Hint(j.recheck_429);
    renderBans();
  }catch(e){
    banBadge.textContent='!';
    banBadge.className='chip chip-bad';
    if(force) toast(e.message,'err');
  }
}
function updateRecheck429Hint(rc){
  const hint=$('banRecheckHint');
  const title=$('banRecheckTitle');
  const card=$('banRecheckCard');
  const n429=banMeta.by_code? (banMeta.by_code[429]||banMeta.by_code['429']||0) : 0;
  const nDue=banMeta.due_429||0;
  if(!rc){
    if(title) title.textContent='429 测活'+(n429?(' · '+n429):'');
    if(hint) hint.textContent=n429?('到期复测 · 仍429 +2h'+(nDue?(' · 待测 '+nDue):'')):'—';
    if(card) card.classList.remove('running');
    if($('btnRecheck429')){ btnRecheck429.disabled=false; btnRecheck429.textContent=n429?('测活 429 ('+n429+')'):'测活 429'; }
    return;
  }
  if(rc.running){
    if(card) card.classList.add('running');
    if(title) title.innerHTML='<span class="spin"></span>测活中';
    if(hint) hint.textContent=rc.mode==='expiry'?'到期复测…':'…';
    if($('btnRecheck429')){ btnRecheck429.disabled=true; btnRecheck429.textContent='…'; }
    return;
  }
  if(card) card.classList.remove('running');
  if(title) title.textContent='429 测活'+(n429?(' · '+n429):'');
  const parts=[];
  if(rc.last_run) parts.push(prettyTime(rc.last_run));
  if(rc.message) parts.push(rc.message);
  if(nDue) parts.push('待测 '+nDue);
  if(hint) hint.textContent=parts.length?parts.join(' · '):(n429?('到期复测 · 仍429 +2h'):'—');
  if($('btnRecheck429')){
    btnRecheck429.disabled=false;
    btnRecheck429.textContent=n429?('测活 429 ('+n429+')'):'测活 429';
  }
}
async function recheckAll429(){
  const n=banMeta.by_code? (banMeta.by_code[429]||banMeta.by_code['429']||0) : 0;
  if(!n){toast('无 429','err');return}
  if(!confirm('测活 '+n+' 条 429？')) return;
  updateRecheck429Hint({running:true});
  try{
    const j=await api('/bans-recheck-429',{method:'POST',body:'{}'});
    await loadBans(true);
    updateRecheck429Hint(j);
    toast(j.message||'完成','ok');
  }catch(e){
    updateRecheck429Hint({running:false});
    toast(e.message,'err');
  }
}
function setupBanTimer(){
  if(banTimer){clearInterval(banTimer);banTimer=null}
  if(mgmtBanned) return;
  if($('banAuto')&&banAuto.checked) banTimer=setInterval(()=>loadBans(false),15000);
}
async function unbanOne(id){
  if(!confirm('解禁？')) return;
  try{
    await api('/unban',{method:'POST',body:JSON.stringify({auth_id:id})});
    toast('已解禁','ok');
    await loadBans(true);
  }catch(e){toast(e.message,'err')}
}
async function unbanSelected(){
  const ids=[...banSelected];
  if(!ids.length){toast('未选','err');return}
  if(!confirm('解禁 '+ids.length+'？')) return;
  try{
    await api('/unban',{method:'POST',body:JSON.stringify({auth_ids:ids})});
    banSelected.clear();
    toast('已解禁 '+ids.length,'ok');
    await loadBans(true);
  }catch(e){toast(e.message,'err')}
}
async function unbanByStatus(code){
  const n=banMeta.by_code? (banMeta.by_code[code]||banMeta.by_code[String(code)]||0) : 0;
  if(!n){toast('无 '+code,'err');return}
  if(!confirm('清 '+code+' ×'+n+'？')) return;
  try{
    await api('/unban',{method:'POST',body:JSON.stringify({status:code})});
    toast('已清 '+code,'ok');
    await loadBans(true);
  }catch(e){toast(e.message,'err')}
}
async function unbanAll(){
  const n=banMeta.count||0;
  if(!n){toast('空','err');return}
  if(!confirm('全部解禁 '+n+'？')) return;
  try{
    await api('/unban-all',{method:'POST',body:'{}'});
    banSelected.clear();
    toast('已解禁','ok');
    await loadBans(true);
  }catch(e){toast(e.message,'err')}
}
async function copyBanIDs(){
  if(!lastBans.length){toast('无数据','err');return}
  const ids=lastBans.map(b=>b.auth_id).join('\n');
  try{await navigator.clipboard.writeText(ids);toast('已复制 '+lastBans.length,'ok')}
  catch(e){toast(e.message,'err')}
}
async function boot(){
  loadKey();
  bindMgmtKeyToggle();
  restoreTab();
  mgmtBanned=false;
  setupBanTimer();
  await Promise.all([refresh(), refreshSSO(), loadVault(false), loadSchedule(), loadBans(false), loadPaths().catch(()=>{})]);
  if(activeTab()==='scan') await loadScanResults().catch(()=>{});
}
bindMgmtKeyToggle();
boot();
</script>
</body>
</html>`
