(this.webpackJsonpregi=this.webpackJsonpregi||[]).push([[0],{347:function(e,t,n){},494:function(e,t){},533:function(e,t,n){"use strict";n.r(t);var c=n(0),r=n.n(c),a=n(13),s=n.n(a),o=(n(347),n(293)),u=n(602),i=n(609),j=n(92),b=n.n(j),f=n(203),l=n(307),p=n(158),h=n(139),d=n(46),O=n(603),x=n(604),g=n(599),v=n(600),y=n(601),m=n(25);function w(){return(w=Object(p.a)(b.a.mark((function e(){var t;return b.a.wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.next=2,fetch("/api/enable",{method:"POST"});case 2:return t=e.sent,console.log(t),e.abrupt("return",t);case 5:case"end":return e.stop()}}),e)})))).apply(this,arguments)}function k(){return(k=Object(p.a)(b.a.mark((function e(){var t;return b.a.wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.next=2,fetch("/api/disable",{method:"POST"});case 2:return t=e.sent,console.log(t),e.abrupt("return",t);case 5:case"end":return e.stop()}}),e)})))).apply(this,arguments)}function P(){return(P=Object(p.a)(b.a.mark((function e(){var t;return b.a.wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.next=2,fetch("/api/enabled");case 2:return t=e.sent,e.abrupt("return",t.json());case 4:case"end":return e.stop()}}),e)})))).apply(this,arguments)}var S=function(e){var t=Object(c.useState)(!1),n=Object(l.a)(t,2),r=n[0],a=n[1],s=function(){(function(){return P.apply(this,arguments)})().then((function(e){console.log(e),a(e.enabled)}))};return Object(c.useEffect)(s,[]),Object(m.jsxs)(h.a,{children:[Object(m.jsx)(d.a,{onClick:function(){r?function(){return k.apply(this,arguments)}().then((function(e){return s()})).catch((function(e){console.log(e),s()})):function(){return w.apply(this,arguments)}().then((function(e){return s()})).catch((function(e){console.log(e),s()}))},label:r?"disable":"enable"})," :"]})},C=function(e){return Object(m.jsx)("div",{children:Object(m.jsx)(O.a,Object(f.a)(Object(f.a)({},e),{},{actions:Object(m.jsx)(S,{}),children:Object(m.jsxs)(x.a,{children:[Object(m.jsx)(g.a,{source:"id"}),Object(m.jsx)(g.a,{source:"path"}),Object(m.jsx)(g.a,{source:"request"}),Object(m.jsx)(v.a,{source:"status"}),Object(m.jsx)(g.a,{source:"body"}),Object(m.jsx)(g.a,{source:"body_golden"}),Object(m.jsx)(g.a,{source:"status_golden"}),Object(m.jsx)(y.a,{source:"created_at"})]})}))})},F=Object(o.a)("api"),T=function(){return Object(m.jsx)(u.a,{dataProvider:F,children:Object(m.jsx)(i.a,{name:"diffs",list:C})})},_=function(e){e&&e instanceof Function&&n.e(3).then(n.bind(null,612)).then((function(t){var n=t.getCLS,c=t.getFID,r=t.getFCP,a=t.getLCP,s=t.getTTFB;n(e),c(e),r(e),a(e),s(e)}))};s.a.render(Object(m.jsx)(r.a.StrictMode,{children:Object(m.jsx)(T,{})}),document.getElementById("root")),_()}},[[533,1,2]]]);
//# sourceMappingURL=main.2b6c5306.chunk.js.map