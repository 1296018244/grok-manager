package main

const panelHTML = `<!doctype html>
<html lang="zh-CN">
<head>
<meta charset="utf-8"/>
<meta name="viewport" content="width=device-width,initial-scale=1"/>
<title>Grok Manager</title>
<style>
:root{
  --bg:#eef2f9;
  --bg-glow:radial-gradient(1200px 500px at 10% -10%, rgba(59,130,246,.12), transparent 55%),
            radial-gradient(900px 400px at 100% 0%, rgba(124,58,237,.08), transparent 50%),
            #eef2f9;
  --card:#ffffff;
  --text:#0f172a;
  --muted:#64748b;
  --faint:#94a3b8;
  --line:rgba(148,163,184,.22);
  --line2:#d7e0ee;
  --blue:#3b82f6;
  --blue-deep:#2563eb;
  --blue-soft:#eff6ff;
  --blue-soft2:#dbeafe;
  --purple:#7c3aed;
  --purple-soft:#f5f3ff;
  --green:#10b981;
  --green-soft:#ecfdf5;
  --orange:#f97316;
  --orange-soft:#fff7ed;
  --red:#ef4444;
  --red-soft:#fef2f2;
  --amber:#f59e0b;
  --r:20px; --r-sm:14px; --r-pill:999px;
  --shadow:0 1px 2px rgba(15,23,42,.03), 0 10px 30px rgba(15,23,42,.05);
  --shadow-sm:0 1px 2px rgba(15,23,42,.04);
  --font:ui-sans-serif,system-ui,-apple-system,"Segoe UI",Roboto,"PingFang SC","Microsoft YaHei",sans-serif;
  --mono:ui-monospace,SFMono-Regular,Consolas,monospace;
  --ok:var(--green); --ok-bg:var(--green-soft); --ok-bd:#a7f3d0;
  --bad:var(--red); --bad-bg:var(--red-soft); --bad-bd:#fecaca;
  --warn:var(--amber); --warn-bg:#fffbeb; --warn-bd:#fde68a;
  --info:var(--blue-deep); --info-bg:var(--blue-soft); --info-bd:var(--blue-soft2);
  --accent:var(--blue-deep); --accent2:#1d4ed8; --accent-bg:var(--blue-soft);
  --pay:var(--purple); --pay-bg:var(--purple-soft); --pay-bd:#ddd6fe;
  --bg0:var(--bg); --bg1:var(--card); --bg2:#f8fafc; --bg3:#f1f5f9;
}
*{box-sizing:border-box}
html,body{
  margin:0;min-height:100%;
  background:var(--bg-glow);color:var(--text);
  font:14px/1.5 var(--font);-webkit-font-smoothing:antialiased;
}
button,input,textarea,select{font:inherit}
button{border:0;cursor:pointer;font-weight:600;transition:.16s ease;background:transparent}
button:active:not(:disabled){transform:translateY(1px)}
button:disabled{opacity:.42;cursor:not-allowed}
a{color:var(--blue-deep)}
.app{max-width:1180px;margin:0 auto;padding:18px 18px 48px}

/* Header */
.topbar{
  display:flex;flex-wrap:wrap;gap:14px 18px;align-items:center;justify-content:space-between;
  margin-bottom:14px;padding:14px 18px;border-radius:var(--r);background:rgba(255,255,255,.92);
  border:1px solid var(--line);box-shadow:var(--shadow);backdrop-filter:blur(10px);
}
.brand{display:flex;gap:12px;align-items:center}
.logo{
  width:42px;height:42px;border-radius:14px;display:grid;place-items:center;
  background:linear-gradient(145deg,#60a5fa 0%,#2563eb 55%,#1d4ed8 100%);
  color:#fff;font-weight:800;font-size:13px;letter-spacing:-.02em;
  box-shadow:0 8px 20px rgba(37,99,235,.32);
}
.brand h1{margin:0;font-size:16.5px;font-weight:750;letter-spacing:-.03em}
.brand .ver{display:flex;flex-wrap:wrap;gap:6px;margin-top:6px}
.chip{
  display:inline-flex;align-items:center;gap:4px;padding:3px 9px;border-radius:var(--r-pill);
  font-size:11px;font-weight:650;background:#f8fafc;border:1px solid var(--line);color:var(--muted);
}
.chip-accent{background:var(--blue-soft);border-color:var(--blue-soft2);color:var(--blue-deep)}
.chip-ok{background:var(--green-soft);border-color:#a7f3d0;color:var(--green)}
.chip-warn{background:#fffbeb;border-color:#fde68a;color:#d97706}
.chip-bad{background:var(--red-soft);border-color:#fecaca;color:var(--red)}
.chip-info{background:var(--blue-soft);border-color:var(--blue-soft2);color:var(--blue-deep)}
.top-actions{display:flex;flex-wrap:wrap;gap:8px;align-items:end}
.field{display:flex;flex-direction:column;gap:5px;min-width:0}
.field>span{font-size:11px;font-weight:650;color:var(--muted)}
.field input,.field textarea,.field select,
input[type=text],input[type=password],input[type=number],input[type=search],textarea,select{
  background:#fff;border:1px solid var(--line2);color:var(--text);
  border-radius:12px;padding:9px 12px;outline:none;min-width:0;
  transition:border-color .15s,box-shadow .15s;
}
.field input:focus,input:focus,textarea:focus,select:focus{
  border-color:#93c5fd;box-shadow:0 0 0 4px rgba(59,130,246,.12);
}
#mgmtKey{min-width:200px;width:min(280px,55vw);font-family:var(--mono);font-size:12px;background:#f8fafc}

/* Nav */
.nav{
  display:flex;gap:6px;padding:6px;margin-bottom:16px;border-radius:18px;
  background:rgba(255,255,255,.92);border:1px solid var(--line);box-shadow:var(--shadow);
  position:sticky;top:10px;z-index:40;backdrop-filter:blur(10px);
}
.nav button{
  flex:1;padding:11px 12px;border-radius:14px;color:var(--muted);font-size:13.5px;font-weight:650;
}
.nav button:hover{background:#f1f5f9;color:var(--text)}
.nav button.on{
  background:linear-gradient(180deg,#3b82f6,#2563eb);color:#fff;
  box-shadow:0 6px 18px rgba(37,99,235,.30);
}
.nav button .badge{
  display:inline-block;margin-left:5px;padding:0 7px;border-radius:var(--r-pill);
  font-size:10px;background:#e2e8f0;color:var(--muted);min-width:16px;text-align:center;
}
.nav button.on .badge{background:rgba(255,255,255,.22);color:#fff}
.nav button .badge:empty,.nav button .badge.zero{display:none}

.panel{display:none}
.panel.on{display:block;animation:fade .18s ease}
@keyframes fade{from{opacity:0;transform:translateY(6px)}to{opacity:1;transform:none}}

.card{
  background:var(--card);border:1px solid var(--line);border-radius:var(--r);
  padding:18px 20px;margin-bottom:14px;box-shadow:var(--shadow);
}
.card-hd{
  display:flex;flex-wrap:wrap;gap:8px 12px;align-items:center;justify-content:space-between;
  margin:0 0 16px;
}
.card-hd h2{margin:0;font-size:16px;font-weight:750;letter-spacing:-.02em}
.card-hd .sub{margin-top:3px;font-size:12.5px;color:var(--muted)}
.hint{font-size:12.5px;color:var(--muted);line-height:1.55;margin:0}
.hint code,code{font-family:var(--mono);font-size:11px;background:#f1f5f9;border:1px solid var(--line);padding:1px 5px;border-radius:6px}
.path{font-family:var(--mono);font-size:11px;color:#64748b;word-break:break-all;margin-top:8px}
.muted{color:var(--muted)}.faint{color:var(--faint)}
.help-line{margin:10px 0 0;font-size:12px;color:var(--faint);line-height:1.45}

.banner,.info-bar{
  display:flex;gap:10px;align-items:flex-start;
  border-radius:16px;padding:12px 14px;font-size:13px;line-height:1.55;
  border:1px solid transparent;margin:0;
}
.info-bar,.banner-info{background:linear-gradient(180deg,#f0f7ff,#eff6ff);border-color:#dbeafe;color:#1e40af}
.banner-ok{background:var(--green-soft);border-color:#a7f3d0;color:#065f46}
.banner-warn{background:#fffbeb;border-color:#fde68a;color:#92400e}
.banner-bad{background:var(--red-soft);border-color:#fecaca;color:#991b1b}
.info-ico{
  flex:0 0 auto;width:22px;height:22px;border-radius:50%;
  background:linear-gradient(180deg,#3b82f6,#2563eb);color:#fff;
  display:grid;place-items:center;font-size:12px;font-weight:800;margin-top:1px;
  box-shadow:0 2px 8px rgba(37,99,235,.25);
}

.row{display:flex;flex-wrap:wrap;gap:10px;align-items:end}
label.field{flex:0 1 auto}
label.field.grow,label.grow{flex:1 1 220px}
label.check{
  display:inline-flex;align-items:center;gap:7px;padding:8px 12px;border-radius:var(--r-pill);
  background:#f8fafc;border:1px solid var(--line);font-size:12.5px;color:var(--text);cursor:pointer;user-select:none;
}
label.check:hover{border-color:var(--line2);background:#f1f5f9}
label.check input{width:14px;height:14px;accent-color:var(--blue-deep);margin:0}
textarea{width:100%;min-height:120px;resize:vertical;font-family:var(--mono);font-size:12px;line-height:1.45;background:#f8fafc}
input[type=number]{width:96px}
.checks{display:flex;flex-wrap:wrap;gap:8px;margin-top:12px}
.form-grid{display:grid;grid-template-columns:repeat(auto-fit,minmax(130px,1fr));gap:10px;margin-top:4px}
.form-grid .grow{grid-column:span 2}
@media (max-width:640px){.form-grid .grow{grid-column:span 1}}

/* Buttons */
.btn,.btn-ghost,.btn-ok,.btn-warn,.btn-danger,.btn-soft,.pill-btn{
  display:inline-flex;align-items:center;justify-content:center;gap:7px;
  border-radius:var(--r-pill);padding:10px 16px;font-size:13px;font-weight:650;
  border:1px solid transparent;
}
.btn,.pill-btn.primary{
  background:linear-gradient(180deg,#3b82f6,#2563eb);color:#fff;
  box-shadow:0 6px 16px rgba(37,99,235,.28);
}
.btn:hover:not(:disabled),.pill-btn.primary:hover:not(:disabled){filter:brightness(1.03)}
.btn-ok{background:linear-gradient(180deg,#34d399,#10b981);color:#fff;box-shadow:0 4px 12px rgba(16,185,129,.25)}
.btn-warn{background:linear-gradient(180deg,#fbbf24,#f59e0b);color:#fff}
.btn-danger{background:linear-gradient(180deg,#f87171,#ef4444);color:#fff}
.btn-danger:hover:not(:disabled){filter:brightness(1.03)}
.btn-ghost,.pill-btn{
  background:#fff;color:var(--text);border-color:var(--line2);box-shadow:var(--shadow-sm);
}
.btn-ghost:hover:not(:disabled),.pill-btn:hover:not(:disabled){background:#f8fafc;border-color:#c5d0e0}
.btn-soft{background:var(--blue-soft);color:var(--blue-deep);border-color:var(--blue-soft2)}
.btn-soft:hover:not(:disabled){background:var(--blue-soft2)}
.btn-sm,.pill-btn.sm{padding:8px 12px;font-size:12px}
.btn-ico{width:15px;height:15px;display:inline-block;flex:0 0 auto}
.pill-btn.danger-soft{color:#b91c1c;border-color:#fecaca;background:#fff}
.pill-btn.danger-soft:hover:not(:disabled){background:var(--red-soft)}

.toolbar,.action-bar,.toolbar-row{
  display:flex;flex-wrap:wrap;gap:10px;align-items:center;justify-content:space-between;
  margin-top:14px;padding:12px 14px;
  background:linear-gradient(180deg,#fbfdff,#f6f9fc);
  border:1px solid var(--line);border-radius:16px;
}
.toolbar.stack,.action-bar.stack,.toolbar-row.stack{flex-direction:column;align-items:stretch}
.toolbar .grp,.action-bar .grp,.toolbar-row .grp{display:flex;flex-wrap:wrap;gap:8px;align-items:center}
.toolbar .lbl,.action-bar .lbl{font-size:10px;font-weight:700;color:var(--faint);text-transform:uppercase;letter-spacing:.04em}
.action-bar.danger{background:linear-gradient(180deg,#fff,#fff5f5);border-color:#fecaca}
.action-bar.danger .lbl{color:#b91c1c}
.action-bar.safe{background:linear-gradient(180deg,#fbfdff,#f6f9fc)}

/* Metric cards — horizontal icon + number like premium dashboards */
.metric-grid{
  display:grid;grid-template-columns:repeat(4,1fr);gap:12px;margin:4px 0 14px;
}
@media (max-width:1100px){.metric-grid[style*="repeat(5"]{grid-template-columns:repeat(3,1fr)!important}}
@media (max-width:960px){.metric-grid{grid-template-columns:repeat(2,1fr)!important}}
@media (max-width:520px){.metric-grid{grid-template-columns:1fr!important}}
.metric{
  position:relative;overflow:hidden;
  background:var(--card);border:1px solid var(--line);border-radius:20px;
  padding:18px 18px 16px;box-shadow:var(--shadow);min-height:118px;
}
.metric .row-m{display:flex;align-items:center;gap:14px}
.metric .mi{
  width:48px;height:48px;border-radius:16px;display:grid;place-items:center;
  background:#f1f5f9;color:var(--blue-deep);flex:0 0 auto;
  box-shadow:inset 0 1px 0 rgba(255,255,255,.8);
}
.metric .mi svg{width:24px;height:24px}
.metric .n{
  font-size:30px;font-weight:800;letter-spacing:-.04em;line-height:1;
  font-variant-numeric:tabular-nums;color:var(--blue-deep);
}
.metric .l{margin-top:5px;font-size:13px;color:var(--muted);font-weight:500}
.metric.m-blue .mi{background:linear-gradient(180deg,#eff6ff,#dbeafe);color:var(--blue-deep)}
.metric.m-blue .n{color:var(--blue-deep)}
.metric.m-purple .mi{background:linear-gradient(180deg,#f5f3ff,#ede9fe);color:var(--purple)}
.metric.m-purple .n{color:var(--purple)}
.metric.m-green .mi{background:linear-gradient(180deg,#ecfdf5,#d1fae5);color:var(--green)}
.metric.m-green .n{color:var(--green)}
.metric.m-orange .mi{background:linear-gradient(180deg,#fff7ed,#ffedd5);color:var(--orange)}
.metric.m-orange .n{color:var(--orange)}
.metric.m-red .mi{background:linear-gradient(180deg,#fef2f2,#fee2e2);color:var(--red)}
.metric.m-red .n{color:var(--red)}
.metric .wave{position:absolute;right:0;bottom:0;left:28%;height:46px;opacity:.7;pointer-events:none}
.metric.clickable{cursor:pointer;transition:transform .14s,box-shadow .16s,border-color .16s}
.metric.clickable:hover{transform:translateY(-2px);box-shadow:0 12px 28px rgba(15,23,42,.08);border-color:#c7d7f5}
.metric.on{border-color:#93c5fd;box-shadow:0 0 0 3px rgba(59,130,246,.12)}

/* legacy stats */
.stats{display:grid;grid-template-columns:repeat(auto-fit,minmax(108px,1fr));gap:10px}
.stat{
  background:var(--card);border:1px solid var(--line);border-radius:16px;padding:14px;
  position:relative;overflow:hidden;box-shadow:var(--shadow-sm);
}
.stat::before{content:"";position:absolute;left:0;top:0;bottom:0;width:3px;background:var(--line2)}
.stat .n{font-size:24px;font-weight:800;letter-spacing:-.03em;font-variant-numeric:tabular-nums}
.stat .l{font-size:12px;color:var(--muted);margin-top:4px}
.stat.ok::before{background:var(--green)}.stat.ok .n{color:var(--green)}
.stat.bad::before{background:var(--red)}.stat.bad .n{color:var(--red)}
.stat.warn::before{background:var(--amber)}.stat.warn .n{color:var(--amber)}
.stat.info::before{background:var(--blue-deep)}.stat.info .n{color:var(--blue-deep)}
.stat.pay::before{background:var(--purple)}.stat.pay .n{color:var(--purple)}
.stat.clickable{cursor:pointer;transition:transform .12s,border-color .15s}
.stat.clickable:hover{border-color:#c7d7f5;transform:translateY(-1px)}
.stat.on{border-color:#93c5fd;background:var(--blue-soft)}
.stat-sm .n{font-size:18px}

.progress-panel{
  display:flex;flex-wrap:wrap;gap:18px 22px;align-items:center;
  padding:18px 20px;border-radius:20px;background:var(--card);
  border:1px solid var(--line);box-shadow:var(--shadow);margin-bottom:14px;
}
.ring{
  --p:0;width:72px;height:72px;border-radius:50%;flex:0 0 auto;
  background:conic-gradient(#2563eb calc(var(--p)*1%), #e8eef7 0);
  display:grid;place-items:center;position:relative;
  box-shadow:inset 0 0 0 1px rgba(37,99,235,.06);
}
.ring::after{content:"";position:absolute;inset:8px;border-radius:50%;background:var(--card)}
.ring span{
  position:relative;z-index:1;font-size:13.5px;font-weight:800;color:var(--blue-deep);
  font-variant-numeric:tabular-nums;
}
.progress-meta{flex:1 1 220px;min-width:0}
.progress-meta .t{font-size:14.5px;font-weight:750;color:var(--text)}
.progress-meta .d{margin-top:4px;font-size:12.5px;color:var(--muted);line-height:1.45}
.bar,.progress-bar{
  height:9px;background:#e8eef7;border-radius:var(--r-pill);overflow:hidden;margin-top:12px;border:0;
}
.bar>i,.progress-bar>i{
  display:block;height:100%;width:0;
  background:linear-gradient(90deg,#60a5fa,#2563eb 60%,#1d4ed8);
  border-radius:var(--r-pill);transition:width .3s ease;
}
.log{
  font-family:var(--mono);font-size:11.5px;color:#64748b;white-space:pre-wrap;
  max-height:88px;overflow:auto;background:#f8fafc;border:1px solid var(--line);
  border-radius:12px;padding:10px 12px;margin-top:10px;line-height:1.5;
}

.filters{display:flex;flex-wrap:wrap;gap:8px;margin:2px 0 14px}
.filters button{
  display:inline-flex;align-items:center;gap:5px;
  padding:8px 13px;border-radius:var(--r-pill);font-size:12.5px;font-weight:650;
  background:#fff;color:var(--muted);border:1px solid var(--line);box-shadow:var(--shadow-sm);
}
.filters button:hover{color:var(--text);border-color:#c5d0e0;background:#f8fafc}
.filters button.on{background:var(--blue-soft);color:var(--blue-deep);border-color:#bfdbfe;box-shadow:none}
.filters button .fc{
  display:inline-block;margin-left:2px;min-width:1.1em;padding:0 6px;border-radius:var(--r-pill);
  font-size:11px;font-weight:750;font-variant-numeric:tabular-nums;
  background:#f1f5f9;color:var(--muted);line-height:1.5;
}
.filters button.on .fc{background:#dbeafe;color:var(--blue-deep)}
.filters button .fc.zero{opacity:.4}
.filters button.f-ok.on{background:var(--green-soft);color:var(--green);border-color:#a7f3d0}
.filters button.f-ok.on .fc{background:#d1fae5;color:var(--green)}
.filters button.f-bad.on{background:var(--red-soft);color:var(--red);border-color:#fecaca}
.filters button.f-bad.on .fc{background:#fee2e2;color:var(--red)}

.search-row{display:flex;flex-wrap:wrap;gap:10px;align-items:center;margin:0 0 12px}
.search-box{
  flex:1 1 240px;display:flex;align-items:center;gap:8px;
  padding:0 14px;height:44px;border-radius:var(--r-pill);
  background:#fff;border:1px solid var(--line2);box-shadow:var(--shadow-sm);
}
.search-box:focus-within{border-color:#93c5fd;box-shadow:0 0 0 4px rgba(59,130,246,.1)}
.search-box svg{flex:0 0 auto;color:var(--faint)}
.search-box input{flex:1;border:0;outline:none;background:transparent;padding:0;min-width:0;font-size:13px}
.search-box input:focus{box-shadow:none;border:0}

.table-wrap{
  overflow:auto;max-height:min(520px,56vh);border:1px solid var(--line);border-radius:16px;
  background:#fff;margin-top:6px;
}
.table-wrap.mid{max-height:min(380px,48vh)}
.table-wrap.sm{max-height:240px}
.table-wrap.tall{max-height:min(520px,56vh)}
table{width:100%;border-collapse:collapse;font-size:12.5px}
th,td{padding:11px 12px;border-bottom:1px solid #f1f5f9;text-align:left;vertical-align:middle}
th{
  position:sticky;top:0;z-index:1;background:#f8fafc;color:var(--muted);
  font-size:10.5px;font-weight:750;letter-spacing:.05em;text-transform:uppercase;
}
tr:last-child td{border-bottom:0}
tr:hover td{background:#f8fafc}
td.mono,th.mono{font-family:var(--mono);font-size:11.5px}
td.nowrap{white-space:nowrap}
.tag{
  display:inline-flex;align-items:center;padding:3px 8px;border-radius:var(--r-pill);
  font-size:11px;font-weight:750;border:1px solid transparent;white-space:nowrap;
}
.tag-ok{background:var(--green-soft);color:var(--green);border-color:#a7f3d0}
.tag-bad,.tag-del{background:var(--red-soft);color:var(--red);border-color:#fecaca}
.tag-skip,.tag-rate{background:#fffbeb;color:#d97706;border-color:#fde68a}
.tag-keep{background:var(--blue-soft);color:var(--blue-deep);border-color:#bfdbfe}
.tag-pay{background:var(--purple-soft);color:var(--purple);border-color:#ddd6fe}

.pager{display:flex;flex-wrap:wrap;gap:8px;align-items:center;justify-content:space-between;margin:10px 0 0}
.pager .info{font-size:12px;color:var(--muted)}
.pager .btns{display:flex;gap:6px}
.remain{font-variant-numeric:tabular-nums;font-weight:650}
.remain.urgent{color:var(--red)}.remain.soon{color:var(--amber)}.remain.ok{color:var(--green)}
.id-cell{font-family:var(--mono);font-size:11px;word-break:break-all;max-width:200px;line-height:1.35}
.empty{text-align:center;color:var(--muted);padding:48px 16px;font-size:13px}
.empty strong{display:block;color:var(--text);font-size:14px;margin-bottom:4px}
.row-actions{display:inline-flex;gap:6px;align-items:center}
.adv-row{max-width:280px;overflow:hidden;text-overflow:ellipsis;white-space:nowrap;color:var(--muted);font-size:12px}

.policy-row{display:flex;flex-wrap:wrap;gap:6px;margin:0 0 14px}
.policy-chip{
  display:inline-flex;align-items:center;gap:4px;padding:5px 10px;border-radius:var(--r-pill);
  font-size:11px;font-weight:650;background:#f8fafc;border:1px solid var(--line);color:var(--muted);
}
.policy-chip b{color:var(--text)}
.policy-chip.w{background:#fffbeb;border-color:#fde68a;color:#92400e}
.policy-chip.b{background:var(--red-soft);border-color:#fecaca;color:#991b1b}
.policy-chip.p{background:var(--purple-soft);border-color:#ddd6fe;color:#5b21b6}
.policy-chip.i{background:var(--blue-soft);border-color:#bfdbfe;color:#1e40af}

.recheck-card{
  display:flex;flex-wrap:wrap;gap:12px;align-items:center;justify-content:space-between;
  margin:0 0 14px;padding:14px 16px;border-radius:18px;
  background:linear-gradient(135deg,#fffbeb,#fff7ed);border:1px solid #fde68a;
}
.recheck-card .t{font-size:13.5px;font-weight:750;color:#92400e}
.recheck-card .d{font-size:12.5px;color:#a16207;margin-top:2px}
.recheck-card.running{background:linear-gradient(135deg,#eff6ff,#e0f2fe);border-color:#bfdbfe}
.recheck-card.running .t,.recheck-card.running .d{color:#1e40af}
.spin{display:inline-block;width:11px;height:11px;border:2px solid currentColor;border-right-color:transparent;border-radius:50%;animation:spin .7s linear infinite;vertical-align:-1px;margin-right:4px}
@keyframes spin{to{transform:rotate(360deg)}}
.sel-count{
  display:inline-flex;align-items:center;padding:3px 9px;border-radius:var(--r-pill);
  font-size:11px;font-weight:750;background:var(--blue-soft);color:var(--blue-deep);border:1px solid #bfdbfe;
}
.sel-count:empty,.sel-count.zero{display:none}

.foot{margin-top:10px;text-align:center;color:var(--faint);font-size:11px}
.toast{
  position:fixed;right:16px;bottom:16px;z-index:60;max-width:min(400px,92vw);
  padding:12px 14px;border-radius:16px;background:#fff;border:1px solid var(--line2);
  box-shadow:0 14px 40px rgba(15,23,42,.14);font-size:13px;
  transform:translateY(10px);opacity:0;pointer-events:none;transition:.2s ease;
}
.toast.show{transform:none;opacity:1}
.toast.err{border-color:#fecaca;background:var(--red-soft);color:#991b1b}
.toast.ok{border-color:#a7f3d0;background:var(--green-soft);color:#065f46}

.quick-grid{display:grid;grid-template-columns:repeat(4,1fr);gap:10px;margin-bottom:12px}
.qcard{display:block;padding:14px;border-radius:16px;border:1px solid var(--line);background:#fff;box-shadow:var(--shadow);cursor:pointer;text-align:left;width:100%}
.qcard .k{font-size:12px;font-weight:700;color:var(--muted)}.qcard .v{font-size:20px;font-weight:800}.qcard .s{font-size:11.5px;color:var(--faint)}
.hero-grid{display:grid;gap:12px}
.kv-row{display:flex;justify-content:space-between;padding:10px;background:#f8fafc;border-radius:10px;font-size:12.5px}
.divider{height:1px;background:var(--line);margin:14px 0}
.danger-zone{margin-top:12px;padding:10px 12px;border-radius:12px;border:1px dashed #fecaca;background:var(--red-soft)}
.details-fold{margin-top:12px;border:1px solid var(--line);border-radius:16px;background:#f8fafc;overflow:hidden}
.details-fold>summary{
  cursor:pointer;list-style:none;padding:12px 14px;font-size:13px;font-weight:650;color:var(--muted);user-select:none;
}
.details-fold>summary::-webkit-details-marker{display:none}
.details-fold[open]>summary{border-bottom:1px solid var(--line);color:var(--text);background:#fff}
.details-fold .inner{padding:14px}
.sec-title{margin:0 0 2px;font-size:15px;font-weight:700}
.sec-sub{margin:0;color:var(--muted);font-size:12px}

@media (max-width:720px){
  .app{padding:12px 12px 32px}
  .topbar{padding:12px 14px}
  #mgmtKey{min-width:0;width:100%}
  .top-actions{width:100%}
  .nav button{padding:10px 8px;font-size:12.5px}
  .metric .n{font-size:26px}
  .metric .mi{width:42px;height:42px;border-radius:14px}
}

/* compact density */
.app{padding:12px 12px 28px!important;max-width:1200px}
.topbar{padding:10px 14px!important;margin-bottom:10px!important}
.nav{padding:4px!important;margin-bottom:10px!important}
.nav button{padding:8px 10px!important;font-size:13px!important}
.card{padding:12px 14px!important;margin-bottom:10px!important;border-radius:16px!important}
.card-hd{margin:0 0 10px!important}
.card-hd h2{font-size:14.5px!important}
.card-hd .sub{display:none!important}
.metric-grid{gap:8px!important;margin:0 0 10px!important}
.metric{min-height:0!important;padding:12px 12px 10px!important;border-radius:14px!important}
.metric .mi{width:36px!important;height:36px!important;border-radius:11px!important}
.metric .mi svg{width:18px!important;height:18px!important}
.metric .n{font-size:22px!important}
.metric .l{font-size:11.5px!important;margin-top:2px!important}
.metric .wave{height:28px!important;opacity:.45!important}
.progress-panel{padding:10px 14px!important;margin-bottom:10px!important;border-radius:14px!important;gap:12px 16px!important}
.ring{width:52px!important;height:52px!important}
.ring::after{inset:6px!important}
.ring span{font-size:11.5px!important}
.progress-meta .t{font-size:13px!important}
.progress-meta .d{font-size:12px!important;margin-top:2px!important}
.bar,.progress-bar{height:6px!important;margin-top:8px!important}
.toolbar,.action-bar,.toolbar-row{
  margin-top:8px!important;padding:8px 10px!important;border-radius:12px!important;gap:6px!important;
  flex-wrap:wrap!important;align-items:center!important;
}
.toolbar .grp,.action-bar .grp,.toolbar-row .grp{gap:6px!important}
.btn,.btn-ghost,.btn-ok,.btn-warn,.btn-danger,.btn-soft,.pill-btn{padding:7px 12px!important;font-size:12.5px!important}
.btn-sm,.pill-btn.sm{padding:6px 10px!important;font-size:12px!important}
.filters{gap:6px!important;margin:0 0 8px!important}
.filters button{padding:6px 10px!important;font-size:12px!important}
.search-row{margin:0 0 8px!important;gap:8px!important}
.search-box{height:36px!important}
.table-wrap{margin-top:6px!important;border-radius:12px!important}
th,td{padding:8px 10px!important}
.pager{margin:6px 0 0!important}
.info-bar,.help-line,.policy-row{display:none!important}
.details-fold{margin-top:8px!important;border-radius:12px!important}
.details-fold>summary{padding:8px 12px!important;font-size:12.5px!important}
.details-fold .inner{padding:10px 12px!important}
.recheck-card{margin:0 0 8px!important;padding:10px 12px!important;border-radius:12px!important}
.checks{margin-top:8px!important;gap:6px!important}
label.check{padding:5px 10px!important;font-size:12px!important}
.form-grid{gap:8px!important}
.hd-inline{display:flex;flex-wrap:wrap;gap:8px;align-items:center;justify-content:space-between;margin-bottom:8px}
.hd-inline h2{margin:0;font-size:14.5px;font-weight:750}
.one-row{display:flex;flex-wrap:wrap;gap:6px;align-items:center}
.one-row.spread{justify-content:space-between}
.compact-bar{
  display:flex;flex-wrap:wrap;gap:6px;align-items:center;
  padding:8px 10px;border-radius:12px;background:linear-gradient(180deg,#fbfdff,#f6f9fc);
  border:1px solid var(--line);
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
          <span class="chip chip-accent">v<span id="ver">1.2.2</span></span>
          <span class="chip" id="jobState">待命</span>
          <span class="chip chip-info" id="hdrVault">库 0</span>
          <span class="chip chip-warn" id="hdrBan">隔离 0</span>
        </div>
      </div>
    </div>
    <div class="top-actions">
      <label class="field">
        <span>管理密钥</span>
        <input id="mgmtKey" type="password" placeholder="密钥" autocomplete="off"/>
      </label>
      <button class="btn-ghost" type="button" onclick="saveKey()">保存</button>
      <button class="btn-soft" type="button" onclick="boot()">刷新</button>
      <button class="btn-ghost btn-sm" type="button" onclick="doBackup()">备份</button>
    </div>
  </header>

  <nav class="nav" id="mainNav">
    <button type="button" class="on" data-tab="autoban" onclick="switchTab('autoban',this)">隔离 <span class="badge zero" id="navBan">0</span></button>
    <button type="button" data-tab="scan" onclick="switchTab('scan',this)">测活 <span class="badge zero" id="navCand">0</span></button>
    <button type="button" data-tab="sso" onclick="switchTab('sso',this)">入库 <span class="badge zero" id="navVault">0</span></button>
  </nav>

  <!-- OVERVIEW -->
  <section class="panel" id="tab-overview" style="display:none" aria-hidden="true">
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
        <h2 title="SSO 导入与历史库">入库</h2>
        <button class="btn-ghost btn-sm" type="button" onclick="document.getElementById('vaultCard').scrollIntoView({behavior:'smooth'});loadVault(true)">历史库</button>
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
      <div class="action-bar" style="margin-top:12px">
        <div class="grp">
          <button class="btn-ok" id="btnSsoStart" type="button" onclick="startSSO()">导入</button>
          <button class="btn-ghost" id="btnSsoStop" type="button" onclick="stopSSO()" disabled>停止</button>
          <button class="btn-ghost btn-sm" type="button" onclick="refreshSSO()">刷新</button>
        </div>
        <div class="grp">
          <button class="btn-warn btn-sm" id="btnSso401" type="button" onclick="refresh401()">重刷 401</button>
        </div>
      </div>
      <div id="ssoSourceBanner" style="display:none"></div>
    </div>

    <div class="card">
      <div class="metric-grid" style="grid-template-columns:repeat(3,1fr)">
        <div class="metric m-blue"><div class="row-m"><div class="mi"><svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8"><path d="M4 6h16M4 12h16M4 18h10"/></svg></div><div><div class="n" id="ssoTotal">0</div><div class="l">总数</div></div></div></div>
        <div class="metric m-purple"><div class="row-m"><div class="mi"><svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8"><circle cx="12" cy="12" r="9"/><path d="M8 12.5l2.5 2.5L16 9"/></svg></div><div><div class="n" id="ssoDone">0</div><div class="l">完成</div></div></div></div>
        <div class="metric m-green"><div class="row-m"><div class="mi"><svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8"><path d="M20 6L9 17l-5-5"/></svg></div><div><div class="n" id="ssoOK">0</div><div class="l">成功</div></div></div></div>
        <div class="metric m-orange"><div class="row-m"><div class="mi"><svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8"><path d="M5 12h14"/></svg></div><div><div class="n" id="ssoSkip">0</div><div class="l">跳过</div></div></div></div>
        <div class="metric m-red"><div class="row-m"><div class="mi"><svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8"><path d="M18 6L6 18M6 6l12 12"/></svg></div><div><div class="n" id="ssoFail">0</div><div class="l">失败</div></div></div></div>
        <div class="metric m-blue"><div class="row-m"><div class="mi"><svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8"><path d="M4 7h16v12H4z"/><path d="M8 7V5h8v2"/></svg></div><div><div class="n" id="ssoVault">0</div><div class="l">库</div></div></div></div>
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

      <div class="action-bar">
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

  <!-- SCAN -->
  <section class="panel" id="tab-scan">
    <div class="card">
      <div class="hd-inline">
        <h2 title="批量探测 xAI 凭证；结束后默认同步隔离">测活</h2>
        <details class="details-fold" style="margin:0;flex:0 0 auto">
          <summary style="padding:6px 10px">参数</summary>
          <div class="inner" style="min-width:min(520px,80vw)">
            <div class="form-grid">
              <label class="field"><span>并发</span><input id="workers" type="number" value="16" min="1" max="128"/></label>
              <label class="field"><span>超时</span><input id="timeout" type="number" value="20" min="3" max="120"/></label>
              <label class="field grow"><span>模型</span><input id="model" type="text" value="grok-4.5"/></label>
              <label class="field"><span>删除码</span><input id="statuses" type="text" value="401,402,403"/></label>
              <label class="field grow"><span>前缀</span><input id="prefix" type="text" placeholder="可选"/></label>
            </div>
            <div class="checks">
              <label class="check"><input type="checkbox" id="auto401" checked/> 401 重刷</label>
              <label class="check"><input type="checkbox" id="syncBans" checked/> 同步隔离</label>
              <label class="check"><input type="checkbox" id="unbanHealthy" checked/> 健康解禁</label>
            </div>
          </div>
        </details>
      </div>
      <div class="compact-bar one-row">
        <button class="btn btn-sm" id="btnStart" type="button" onclick="startScan()" title="开始测活">
          <svg class="btn-ico" viewBox="0 0 24 24" fill="currentColor"><path d="M8 5v14l11-7z"/></svg>开始
        </button>
        <button class="btn-ghost btn-sm" id="btnStop" type="button" onclick="stopScan()" disabled title="停止">停止</button>
        <button class="btn-ghost btn-sm" type="button" onclick="syncScanToBans()" title="用当前结果对账隔离">同步隔离</button>
        <button class="btn-soft btn-sm" type="button" onclick="switchTab('autoban');loadBans(true)" title="打开隔离页">隔离</button>
        <span style="width:1px;height:18px;background:var(--line2);margin:0 2px"></span>
        <button class="pill-btn danger-soft sm" id="btnDel" type="button" onclick="deleteCandidates()" disabled title="删除候选凭证文件">删候选</button>
        <button class="pill-btn danger-soft sm" id="btnDel401" type="button" onclick="deleteByStatus(401)" disabled title="删除 401 凭证">删401</button>
        <button class="pill-btn danger-soft sm" id="btnDel402" type="button" onclick="deleteByStatus(402)" disabled title="删除 402 凭证">删402</button>
        <button class="pill-btn danger-soft sm" id="btnDel403" type="button" onclick="deleteByStatus(403)" disabled title="删除 403 凭证">删403</button>
      </div>
    </div>

    <div class="card">
      <div class="metric-grid">
        <div class="metric m-blue">
          <div class="row-m">
            <div class="mi"><svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8"><path d="M12 3l8 4.5v9L12 21l-8-4.5v-9L12 3z"/><path d="M12 12l8-4.5M12 12v9M12 12L4 7.5"/></svg></div>
            <div><div class="n" id="scTotal2">0</div><div class="l">总数</div></div>
          </div>
        </div>
        <div class="metric m-purple">
          <div class="row-m">
            <div class="mi"><svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8"><circle cx="12" cy="12" r="9"/><path d="M8 12.5l2.5 2.5L16 9"/></svg></div>
            <div><div class="n" id="scDone2">0</div><div class="l">完成</div></div>
          </div>
        </div>
        <div class="metric m-green">
          <div class="row-m">
            <div class="mi"><svg viewBox="0 0 24 24" fill="currentColor"><path d="M12 2l7 3v6c0 5-3 8.5-7 10-4-1.5-7-5-7-10V5l7-3zm-1 13l6-6-1.4-1.4L11 12.2 8.4 9.6 7 11l4 4z"/></svg></div>
            <div><div class="n" id="scOK2">0</div><div class="l">健康</div></div>
          </div>
        </div>
        <div class="metric m-orange">
          <div class="row-m">
            <div class="mi"><svg viewBox="0 0 24 24" fill="currentColor"><path d="M12 3l10 18H2L12 3zm-1 7v5h2v-5h-2zm0 7v2h2v-2h-2z"/></svg></div>
            <div><div class="n" id="scCand2">0</div><div class="l">候选</div></div>
          </div>
        </div>
      </div>
      <div class="progress-panel" style="margin:0;box-shadow:none;border:1px solid var(--line)">
        <div class="ring" id="scanRing" style="--p:0"><span id="scanPct">0%</span></div>
        <div class="progress-meta">
          <div class="t">进度</div>
          <div class="d" id="log">待命</div>
          <div class="progress-bar"><i id="bar2"></i></div>
        </div>
      </div>
    </div>

    <div class="card">
      <div class="hd-inline">
        <div class="one-row">
          <h2 style="margin:0">结果</h2>
          <span class="chip" id="scanFilterLabel">—</span>
        </div>
        <div class="one-row">
          <span class="info" id="scanPageInfo" style="font-size:12px;color:var(--muted)">—</span>
          <button class="btn-ghost btn-sm" type="button" onclick="scanPageDelta(-1)">上页</button>
          <button class="btn-ghost btn-sm" type="button" onclick="scanPageDelta(1)">下页</button>
        </div>
      </div>
      <div class="filters" id="scanTabs">
        <button type="button" class="on" data-f="all" onclick="setScanFilter('all',this)">全部 <span class="fc" data-c="all">0</span></button>
        <button type="button" class="f-bad" data-f="cand" onclick="setScanFilter('cand',this)">候选 <span class="fc" data-c="cand">0</span></button>
        <button type="button" class="f-ok" data-f="healthy" onclick="setScanFilter('healthy',this)">健康 <span class="fc" data-c="healthy">0</span></button>
        <button type="button" data-f="unauthorized" onclick="setScanFilter('unauthorized',this)">401 <span class="fc" data-c="unauthorized">0</span></button>
        <button type="button" data-f="rate_limited" onclick="setScanFilter('rate_limited',this)">429 <span class="fc" data-c="rate_limited">0</span></button>
        <button type="button" data-f="forbidden" onclick="setScanFilter('forbidden',this)">403 <span class="fc" data-c="forbidden">0</span></button>
        <button type="button" data-f="payment" onclick="setScanFilter('payment',this)">402 <span class="fc" data-c="payment">0</span></button>
        <button type="button" data-f="vault_miss" onclick="setScanFilter('vault_miss',this)">401无库 <span class="fc" data-c="vault_miss">0</span></button>
        <button type="button" data-f="vault_hit" onclick="setScanFilter('vault_hit',this)">401有库 <span class="fc" data-c="vault_hit">0</span></button>
      </div>
      <div class="search-row">
        <div class="search-box">
          <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><circle cx="11" cy="11" r="7"/><path d="M20 20l-3-3"/></svg>
          <input id="scanSearch" type="search" placeholder="搜索 email / 文件..." oninput="onScanSearch()"/>
        </div>
      </div>
      <div class="table-wrap tall">
        <table>
          <thead><tr><th>状态</th><th>HTTP</th><th>动作</th><th>Email</th><th>库</th><th>文件</th><th>信息</th><th></th></tr></thead>
          <tbody id="tbody"></tbody>
        </table>
      </div>
    </div>

    <div class="card">
      <div class="hd-inline">
        <div class="one-row">
          <h2 style="margin:0" title="按间隔自动测活并同步隔离">定时</h2>
          <span class="chip" id="schStatusChip">关</span>
        </div>
        <div class="one-row">
          <button class="btn btn-sm" type="button" onclick="saveSchedule()">保存</button>
          <button class="btn-ghost btn-sm" type="button" onclick="loadSchedule()">刷新</button>
          <button class="btn-soft btn-sm" type="button" onclick="doBackup()">备份</button>
        </div>
      </div>
      <div class="one-row" style="gap:10px">
        <label class="field"><span>间隔(分)</span><input id="schInterval" type="number" value="360" min="15" max="10080"/></label>
        <label class="field"><span>并发</span><input id="schWorkers" type="number" value="16" min="1" max="128"/></label>
        <label class="check" style="margin-top:16px"><input type="checkbox" id="schEnabled"/> 启用</label>
        <label class="check" style="margin-top:16px"><input type="checkbox" id="schAuto401" checked/> 刷401</label>
        <label class="check" style="margin-top:16px"><input type="checkbox" id="schRecheck" checked/> 复检</label>
      </div>
      <div id="schBanner" class="banner banner-info" style="margin-top:8px;display:none"></div>
      <div class="path" id="schPaths" style="display:none"></div>
      <div class="path" id="pathsInfo" style="display:none"></div>
    </div>

  </section>

  <!-- AUTOBAN -->
  <section class="panel on" id="tab-autoban">
    <div class="card">
      <div class="hd-inline">
        <div class="one-row">
          <h2 style="margin:0" title="调度黑名单；解禁=恢复可用，删除=删凭证文件">隔离</h2>
          <span class="chip" id="banBadge">0</span>
          <span class="sel-count zero" id="banSelCount"></span>
          <details class="details-fold" style="margin:0">
            <summary style="padding:5px 10px">策略</summary>
            <div class="inner one-row" style="gap:6px">
              <span class="policy-chip i" style="display:inline-flex!important">401有库 <b>2h</b></span>
              <span class="policy-chip b" style="display:inline-flex!important">401无库 <b>24h</b></span>
              <span class="policy-chip b" style="display:inline-flex!important">403 <b>24h</b></span>
              <span class="policy-chip p" style="display:inline-flex!important">402 <b>7d</b></span>
              <span class="policy-chip w" style="display:inline-flex!important">429 <b>2h</b></span>
            </div>
          </details>
        </div>
        <div class="one-row">
          <button class="btn-ghost btn-sm" type="button" onclick="loadBans(true)">刷新</button>
          <button class="btn btn-sm" type="button" id="btnRecheck429" onclick="recheckAll429()" title="复测到期 429">复测429</button>
        </div>
      </div>
      <div class="recheck-card" id="banRecheckCard" style="display:none">
        <div>
          <div class="t" id="banRecheckTitle">429 测活</div>
          <div class="d" id="banRecheckHint">—</div>
        </div>
      </div>

      <div class="metric-grid" style="grid-template-columns:repeat(5,1fr)">
        <div class="metric m-blue clickable on" id="statBanAll" onclick="setBanFilter('all')">
          <div class="row-m"><div class="mi"><svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8"><path d="M4 6h16M4 12h16M4 18h10"/></svg></div><div><div class="n" id="banTotal">0</div><div class="l">全部</div></div></div>
        </div>
        <div class="metric m-red clickable" id="statBan401" onclick="setBanFilter('401')">
          <div class="row-m"><div class="mi"><svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8"><circle cx="12" cy="12" r="9"/><path d="M12 8v5M12 16h.01"/></svg></div><div><div class="n" id="ban401">0</div><div class="l">401</div></div></div>
        </div>
        <div class="metric m-purple clickable" id="statBan402" onclick="setBanFilter('402')">
          <div class="row-m"><div class="mi"><svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8"><rect x="3" y="6" width="18" height="12" rx="2"/><path d="M3 10h18"/></svg></div><div><div class="n" id="ban402">0</div><div class="l">402</div></div></div>
        </div>
        <div class="metric m-red clickable" id="statBan403" onclick="setBanFilter('403')">
          <div class="row-m"><div class="mi"><svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8"><path d="M12 3l9 16H3L12 3z"/><path d="M12 10v4M12 17h.01"/></svg></div><div><div class="n" id="ban403">0</div><div class="l">403</div></div></div>
        </div>
        <div class="metric m-orange clickable" id="statBan429" onclick="setBanFilter('429')">
          <div class="row-m"><div class="mi"><svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8"><circle cx="12" cy="12" r="9"/><path d="M12 7v5l3 2"/></svg></div><div><div class="n" id="ban429">0</div><div class="l">429</div></div></div>
        </div>
      </div>

      <div class="search-row">
        <div class="search-box">
          <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><circle cx="11" cy="11" r="7"/><path d="M20 20l-3-3"/></svg>
          <input id="banSearch" type="search" placeholder="搜索 id / email..." oninput="onBanSearch()"/>
        </div>
        <select id="banFilter" onchange="banPage=1;syncBanStatHighlight();loadBans(false)" style="height:36px;border-radius:999px;padding:0 12px">
          <option value="all">全部</option>
          <option value="401">401</option>
          <option value="402">402</option>
          <option value="403">403</option>
          <option value="429">429</option>
        </select>
        <label class="check"><input type="checkbox" id="banAuto" checked onchange="setupBanTimer()"/> 15s</label>
        <span class="info" id="banPageInfo" style="font-size:12px;color:var(--muted)">—</span>
        <button class="btn-ghost btn-sm" type="button" onclick="banPageDelta(-1)">上页</button>
        <button class="btn-ghost btn-sm" type="button" onclick="banPageDelta(1)">下页</button>
      </div>

      <div class="compact-bar one-row">
        <button class="btn-ghost btn-sm" type="button" onclick="unbanSelected()" title="仅解禁，保留文件">解禁已选</button>
        <button class="btn-ghost btn-sm" type="button" onclick="unbanByStatus(401)" title="仅解禁">解禁401</button>
        <button class="btn-ghost btn-sm" type="button" onclick="unbanByStatus(402)" title="仅解禁">解禁402</button>
        <button class="btn-ghost btn-sm" type="button" onclick="unbanByStatus(403)" title="仅解禁">解禁403</button>
        <button class="btn-ghost btn-sm" type="button" onclick="unbanByStatus(429)" title="仅解禁">解禁429</button>
        <button class="btn-ghost btn-sm" type="button" onclick="unbanAll()" title="仅解禁">全解禁</button>
        <span style="width:1px;height:18px;background:var(--line2)"></span>
        <button class="btn-soft btn-sm" type="button" onclick="pruneOrphanBans()" title="清掉文件已不存在的隔离记录">清幽灵</button>
        <button class="btn-soft btn-sm" type="button" onclick="copyBanIDs()">复制ID</button>
        <span style="width:1px;height:18px;background:var(--line2)"></span>
        <button class="pill-btn danger-soft sm" type="button" onclick="deleteBanSelected()" title="删凭证文件+去隔离">删已选</button>
        <button class="pill-btn danger-soft sm" type="button" onclick="deleteBanByStatus(403)" title="删全部403凭证文件">删403</button>
      </div>

      <div id="banBanner" class="banner banner-info" style="display:none"></div>

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

  <p class="foot">v<span id="footVer">1.2.2</span></p>
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
  return on?on.dataset.tab:'autoban';
}
function switchTab(name,el){
  // aliases after tab merge
  if(name==='vault'||name==='overview') name='sso';
  if(name==='schedule') name='scan';
  // only real panels: autoban / scan / sso (overview kept hidden for stats sinks)
  const panelName = (name==='sso'||name==='scan'||name==='autoban') ? name : 'autoban';
  document.querySelectorAll('.panel').forEach(p=>{
    if(p.id==='tab-overview'){ p.classList.remove('on'); return; }
    p.classList.toggle('on', p.id==='tab-'+panelName);
  });
  document.querySelectorAll('#mainNav button').forEach(b=>{
    const on=el?b===el:b.dataset.tab===panelName;
    b.classList.toggle('on',on);
  });
  if(panelName==='scan'){ loadScanResults().catch(()=>{}); loadSchedule().catch(()=>{}); }
  if(panelName==='autoban') loadBans(false);
  if(panelName==='sso'){ refreshSSO().catch(()=>{}); loadVault(false); }
  try{sessionStorage.setItem('gmcpa-tab',panelName)}catch(e){}
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
    let t=sessionStorage.getItem('gmcpa-tab')||'autoban';
    if(t==='overview'||t==='vault') t='sso';
    if(t==='schedule') t='scan';
    if(t!=='autoban'&&t!=='scan'&&t!=='sso') t='autoban';
    const btn=document.querySelector('#mainNav button[data-tab="'+t+'"]');
    if(btn) switchTab(t,btn);
    else switchTab('autoban');
  }catch(e){ try{switchTab('autoban')}catch(_){ } }
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
  const el=id=>{try{return $(id)||document.getElementById(id)}catch(e){return null}};
  const start=el('btnStart'), stop=el('btnStop');
  if(start) start.disabled=!!b;
  if(stop) stop.disabled=!b;
  const cand=Number((el('sCand')&&el('sCand').textContent)||(el('scCand2')&&el('scCand2').textContent)||0);
  const n401=Number((el('s401')&&el('s401').textContent)||0);
  const n402=Number((el('s402')&&el('s402').textContent)||0);
  const n403=Number((el('s403')&&el('s403').textContent)||0);
  if(el('btnDel')) el('btnDel').disabled=!!b||cand<=0;
  if(el('btnDel401')) el('btnDel401').disabled=!!b||n401<=0;
  if(el('btnDel402')) el('btnDel402').disabled=!!b||n402<=0;
  if(el('btnDel403')) el('btnDel403').disabled=!!b||n403<=0;
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
    const delBtn=name?('<button class="btn-danger btn-sm" type="button" data-n="'+esc(name)+'" onclick="deleteOneName(this.dataset.n)" title="删除凭证文件">删除</button>'):'';
    const adv=r.advice||r.summary||r.error||'';
    return '<tr><td>'+statusTag(st,r.http_status)+'</td><td class="nowrap">'+(r.http_status||'-')+'</td><td class="nowrap">'+esc(actionLabel(r.action))+'</td><td>'+esc(r.email||'')+'</td><td>'+vault+'</td><td class="mono" title="'+esc(name)+'">'+esc(shortId(name))+'</td><td><div class="adv-row" title="'+esc(adv)+'">'+esc(adv)+'</div></td><td style="white-space:nowrap">'+delBtn+'</td></tr>';
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
  const total=st.total||st.result_count||0;
  const done=st.done||0;
  const ok=(sum.by_status&&sum.by_status.healthy)||sum.ok||0;
  const cand=sum.delete_candidates||0;
  if($('scTotal2')) scTotal2.textContent=total;
  if($('scDone2')) scDone2.textContent=done;
  if($('scOK2')) scOK2.textContent=ok;
  if($('scCand2')) scCand2.textContent=cand;
  const pct=total?Math.floor(100*done/total):0;
  if($('bar2')) bar2.style.width=pct+'%';
  if($('scanRing')) scanRing.style.setProperty('--p', String(pct));
  if($('scanPct')) scanPct.textContent=pct+'%';
  setBadge($('navCand'), cand);
  const rows=st.result_count||(st.results||[]).length||0;
  if($('ovScanSub')) ovScanSub.textContent=stateLabel(st.state)+' · '+rows;
  if($('ovQScan')) ovQScan.textContent=String(rows);
  if($('ovQScanSub')) ovQScanSub.textContent=stateLabel(st.state)+(cand?(' · 候选 '+cand):'');
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
  const syncBits=[];
  if(st.scan_sync){
    syncBits.push('同步隔离 +'+(st.scan_sync.banned||0)+' 解禁'+(st.scan_sync.unbanned||0));
  }
  log.textContent=[
    stateLabel(st.state),
    (st.done||0)+'/'+(st.total||0),
    '候选 '+(sum.delete_candidates||0),
    ...syncBits,
    st.message||'',
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
      auto_refresh_401:!!auto401.checked,
      sync_to_bans:!($('syncBans')&&!syncBans.checked),
      unban_healthy:!($('unbanHealthy')&&!unbanHealthy.checked)
    })});
    if(timer) clearInterval(timer);
    timer=setInterval(refresh,1000);
    await refresh();
    if(auto401.checked){if(ssoTimer)clearInterval(ssoTimer);ssoTimer=setInterval(refreshSSO,2000)}
    toast('测活已启动','ok');
  }catch(e){setBusy(false);toast('启动失败: '+e.message,'err')}
}
async function stopScan(){try{await api('/stop',{method:'POST',body:'{}'});await refresh()}catch(e){toast(e.message,'err')}}
async function syncScanToBans(){
  const uh=!($('unbanHealthy')&&!unbanHealthy.checked);
  if(!confirm('用当前测活结果同步到隔离？\\n坏状态写入/续期；'+(uh?'健康自动解禁':'不自动解禁健康号')+'。\\n不会删除凭证文件。')) return;
  try{
    const j=await api('/bans-sync-scan',{method:'POST',body:JSON.stringify({unban_healthy:uh})});
    toast(j.message||j.sync&&j.sync.message||'已同步','ok');
    await refresh();
  }catch(e){toast(e.message,'err')}
}
async function deleteCandidates(){
  const n=Number((($('sCand')&&sCand.textContent)||($('scCand2')&&scCand2.textContent)||0)); if(n<=0){toast('无候选','err');return}
  if(!confirm('删除候选凭证 '+n+' 个文件？不可恢复。')) return;
  try{const r=await api('/delete',{method:'POST',body:JSON.stringify({mode:'candidates'})});toast(formatDeleteResult(r),'ok');await refresh()}catch(e){toast(e.message,'err')}
}
async function deleteByStatus(code){
  const el=code===401?($('s401')||null):code===402?($('s402')||null):code===403?($('s403')||null):null;
  let n=Number(el&&el.textContent||0);
  // fall back to scan filter counts when overview stats hidden
  if(!n && scanMeta.counts){
    const map={401:'unauthorized',402:'payment',403:'forbidden'};
    n=Number(scanMeta.counts[map[code]]||0);
  }
  if(n<=0){toast('无 '+code,'err');return}
  if(!confirm('删除全部 HTTP '+code+' 凭证文件（约 '+n+'）？不可恢复。')) return;
  try{const r=await api('/delete',{method:'POST',body:JSON.stringify({mode:'status',status:Number(code)})});toast(formatDeleteResult(r),'ok');await refresh()}catch(e){toast(e.message,'err')}
}
async function deleteOneName(name){
  if(!name||!confirm('删除凭证「'+name+'」？\\n将删除 auth 文件（不可恢复）。')) return;
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
        +'<td><div class="row-actions">'
        +'<button class="btn-ghost btn-sm" type="button" data-id="'+id+'" onclick="unbanOne(this.dataset.id)" title="仅解禁">解禁</button>'
        +'<button class="btn-danger btn-sm" type="button" data-id="'+id+'" onclick="deleteBanOne(this.dataset.id)" title="删文件+去隔离">删除</button>'
        +'</div></td>'
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
  const setBtn=function(txt, dis){
    if(!$('btnRecheck429')) return;
    btnRecheck429.disabled=!!dis;
    btnRecheck429.textContent=txt;
    btnRecheck429.title=nDue?('待测 '+nDue):'复测 429';
  };
  if(!rc){
    if(card){ card.style.display='none'; card.classList.remove('running'); }
    setBtn(n429?('复测429('+n429+')'):'复测429', false);
    return;
  }
  if(rc.running){
    if(card){ card.style.display='flex'; card.classList.add('running'); }
    if(title) title.innerHTML='<span class="spin"></span>复测中';
    if(hint) hint.textContent=rc.mode==='expiry'?'到期复测…':'进行中…';
    setBtn('…', true);
    return;
  }
  if(card){ card.style.display='none'; card.classList.remove('running'); }
  setBtn(n429?('复测429('+n429+')'):'复测429', false);
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
async function deleteBanOne(id){
  if(!confirm('删除凭证「'+id+'」？\n将删除 auth 文件，并移除隔离（不可恢复）。')) return;
  try{
    const j=await api('/bans-delete',{method:'POST',body:JSON.stringify({auth_id:id})});
    banSelected.delete(id);
    toast(j.message||'已删除','ok');
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
async function deleteBanSelected(){
  const ids=[...banSelected];
  if(!ids.length){toast('未选','err');return}
  if(!confirm('删除已选 '+ids.length+' 个凭证文件，并移除隔离？\n此操作不可恢复。')) return;
  try{
    const j=await api('/bans-delete',{method:'POST',body:JSON.stringify({auth_ids:ids})});
    banSelected.clear();
    toast(j.message||('已删除 '+ids.length),'ok');
    await loadBans(true);
  }catch(e){toast(e.message,'err')}
}
async function unbanByStatus(code){
  const n=banMeta.by_code? (banMeta.by_code[code]||banMeta.by_code[String(code)]||0) : 0;
  if(!n){toast('无 '+code,'err');return}
  if(!confirm('解禁全部 HTTP '+code+'（共 '+n+'）？\n仅解除隔离，不删除凭证文件。')) return;
  try{
    await api('/unban',{method:'POST',body:JSON.stringify({status:code})});
    toast('已解禁全部 '+code,'ok');
    await loadBans(true);
  }catch(e){toast(e.message,'err')}
}
async function deleteBanByStatus(code){
  const n=banMeta.by_code? (banMeta.by_code[code]||banMeta.by_code[String(code)]||0) : 0;
  if(!n){toast('无 '+code,'err');return}
  if(!confirm('删除全部 HTTP '+code+' 凭证文件，并移除隔离？\n共 '+n+' 条，不可恢复。')) return;
  try{
    const j=await api('/bans-delete',{method:'POST',body:JSON.stringify({status:code})});
    banSelected.clear();
    toast(j.message||('已删 '+code),'ok');
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
async function pruneOrphanBans(){
  if(!confirm('同步凭证：把「凭证文件已删除」的隔离记录清掉？\n（只清幽灵记录，不影响仍在的账号）')) return;
  try{
    const j=await api('/bans-prune',{method:'POST',body:'{}'});
    banSelected.clear();
    const n=j.removed||0;
    toast(j.message||('已移除 '+n),'ok');
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
  restoreTab();
  mgmtBanned=false;
  setupBanTimer();
  await Promise.all([refresh(), refreshSSO(), loadVault(false), loadSchedule(), loadBans(false), loadPaths().catch(()=>{})]);
  if(activeTab()==='scan') await loadScanResults().catch(()=>{});
}
boot();
</script>
</body>
</html>`
