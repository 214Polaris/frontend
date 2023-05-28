(function(){var t={7252:function(t,e,i){"use strict";var s=i(9242),o=i(3396);function r(t,e){const i=(0,o.up)("router-link"),s=(0,o.up)("router-view");return(0,o.wg)(),(0,o.iD)(o.HY,null,[(0,o.Wm)(i,{to:"/"}),(0,o.Wm)(s)],64)}var n=i(89);const a={},l=(0,n.Z)(a,[["render",r]]);var d=l;const c={class:"login-container"},u={class:"login-form"},g=(0,o._)("h2",null,"用户登录",-1),m={class:"form-group"},p=(0,o._)("label",null,"用户名",-1),h={key:0,class:"invalid-feedback"},_={class:"form-group"},w=(0,o._)("label",null,"密码",-1),f=["disabled"],v={class:"form-group mt-3"},b={class:"register-form"},y=(0,o._)("h2",null,"用户注册",-1),k={class:"form-group"},x=(0,o._)("label",null,"用户名",-1),C={key:0,class:"invalid-feedback"},I={class:"form-group"},U=(0,o._)("label",null,"密码",-1),D={class:"form-group"},O=(0,o._)("label",null,"确认密码",-1),P={key:0,class:"invalid-feedback"},V=["disabled"],j={class:"form-group mt-3"};function L(t,e,i,r,n,a){return(0,o.wg)(),(0,o.iD)("div",c,[(0,o.wy)((0,o._)("div",u,[g,(0,o._)("form",null,[(0,o._)("div",m,[p,(0,o.wy)((0,o._)("input",{type:"text",class:"form-control","onUpdate:modelValue":e[0]||(e[0]=t=>n.username=t),placeholder:"请输入用户名"},null,512),[[s.nr,n.username,void 0,{trim:!0}]]),n.isUsernameValid?(0,o.kq)("",!0):((0,o.wg)(),(0,o.iD)("div",h," 用户名长度应在5到12个字符之间。 "))]),(0,o._)("div",_,[w,(0,o.wy)((0,o._)("input",{type:"password",class:"form-control","onUpdate:modelValue":e[1]||(e[1]=t=>n.password=t),placeholder:"请输入密码"},null,512),[[s.nr,n.password,void 0,{trim:!0}]])]),(0,o._)("button",{type:"button",class:"btn btn-primary",onClick:e[2]||(e[2]=(...t)=>a.login&&a.login(...t)),disabled:n.isButtonDisabled}," 登录 ",8,f),(0,o._)("div",v,[(0,o.Uk)(" 没有账号？"),(0,o._)("a",{href:"/#/login",onClick:e[3]||(e[3]=(...t)=>a.toggle&&a.toggle(...t))},"点击注册")])])],512),[[s.F8,!n.isRegisterFormVisible]]),(0,o.wy)((0,o._)("div",b,[y,(0,o._)("form",null,[(0,o._)("div",k,[x,(0,o.wy)((0,o._)("input",{type:"text",class:"form-control","onUpdate:modelValue":e[4]||(e[4]=t=>n.registerUsername=t),placeholder:"请输入用户名"},null,512),[[s.nr,n.registerUsername,void 0,{trim:!0}]]),n.isRegisterUsernameValid?(0,o.kq)("",!0):((0,o.wg)(),(0,o.iD)("div",C," 用户名长度应在5到12个字符之间。 "))]),(0,o._)("div",I,[U,(0,o.wy)((0,o._)("input",{type:"password",class:"form-control","onUpdate:modelValue":e[5]||(e[5]=t=>n.registerPassword=t),placeholder:"请输入密码"},null,512),[[s.nr,n.registerPassword,void 0,{trim:!0}]])]),(0,o._)("div",D,[O,(0,o.wy)((0,o._)("input",{type:"password",class:"form-control","onUpdate:modelValue":e[6]||(e[6]=t=>n.confirmPassword=t),placeholder:"请确认密码"},null,512),[[s.nr,n.confirmPassword,void 0,{trim:!0}]]),n.isConfirmPasswordValid?(0,o.kq)("",!0):((0,o.wg)(),(0,o.iD)("div",P," 两次输入的密码不一致。 "))]),(0,o._)("button",{type:"button",class:"btn btn-primary",onClick:e[7]||(e[7]=(...t)=>a.register&&a.register(...t)),disabled:n.isRegisterButtonDisabled}," 注册 ",8,V),(0,o._)("div",j,[(0,o.Uk)(" 已有账号？"),(0,o._)("a",{href:"/#/login",onClick:e[8]||(e[8]=(...t)=>a.toggle&&a.toggle(...t))},"点击登录")])])],512),[[s.F8,n.isRegisterFormVisible]])])}i(7658);var S=i(4161),R=i(5410),T=i.n(R),z={data(){return{username:"",password:"",isUsernameValid:!0,isButtonDisabled:!1,isRegisterFormVisible:!1,registerUsername:"",registerPassword:"",confirmPassword:"",isRegisterUsernameValid:!0,isConfirmPasswordValid:!0,isRegisterButtonDisabled:!1}},methods:{toggle(){this.isRegisterFormVisible=!this.isRegisterFormVisible,this.username="",this.password="",this.isUsernameValid=!0,this.isButtonDisabled=!1,this.registerUsername="",this.registerPassword="",this.confirmPassword="",this.isRegisterUsernameValid=!0,this.isConfirmPasswordValid=!0,this.isRegisterButtonDisabled=!1},login(){T().stringify({username:"",password:""});this.username.length<5||this.username.length>12?this.isUsernameValid=!1:this.isUsernameValid=!0,this.isUsernameValid&&this.password.length>0&&(this.isButtonDisabled=!0,(0,S.Z)({headers:{"Content-Type":"application/x-www-form-urlencoded"},method:"post",url:"/Login",data:{user:this.username,pass:this.password}}).then((t=>{200===t.code&&(this.$store.commit("SET_TOKEN",t.data.token),this.$store.commit("GET_USER",t.data.user),console.log("登录成功！"),this.$router.push({name:"/"})),422===t.code&&alert("您输入的账号或密码错误")})).catch((function(t){console.log(t)})))},register(){this.registerUsername.length<5||this.registerUsername.length>12?this.isRegisterUsernameValid=!1:this.isRegisterUsernameValid=!0,this.registerPassword!==this.confirmPassword?this.isConfirmPasswordValid=!1:this.isConfirmPasswordValid=!0,this.isRegisterUsernameValid&&this.isConfirmPasswordValid&&(this.isRegisterButtonDisabled=!0,S.Z.post("/Register",{user:this.registerUsername,pass:this.registerPassword}).then((t=>{200===t.code&&(this.$store.commit("SET_TOKEN",t.data.token),this.$store.commit("GET_USER",t.data.user),this.$message({message:"注册成功",type:"success"}),this.$router.replace({name:"/"})),422===t.code&&alert("用户已存在")})).catch((function(t){console.log(t)})))}}};const N=(0,n.Z)(z,[["render",L]]);var $=N,E=i(7139);const Z={class:"hello"};function A(t,e,i,s,r,n){return(0,o.wg)(),(0,o.iD)("div",Z,[(0,o._)("h1",null,(0,E.zw)(r.msg),1),(0,o._)("h1",null,"site : "+(0,E.zw)(r.site),1),(0,o._)("h1",null,"url : "+(0,E.zw)(r.url),1),(0,o._)("h3",null,(0,E.zw)(n.details()),1),((0,o.wg)(!0),(0,o.iD)(o.HY,null,(0,o.Ko)(r.ydata,(t=>((0,o.wg)(),(0,o.iD)("h1",{key:t},(0,E.zw)(t),1)))),128)),((0,o.wg)(!0),(0,o.iD)(o.HY,null,(0,o.Ko)(r.xdata,(t=>((0,o.wg)(),(0,o.iD)("h3",{key:t},(0,E.zw)(t),1)))),128))])}var B={name:"apidata",data(){return{msg:"hello,xxbandy！",site:"bgops",url:"https://xxbandy.github.io",xdata:null,ydata:null}},methods:{details:function(){return this.site}},mounted(){S.Z.get("http://localhost:8080/v1/line").then((t=>(this.xdata=t.data.legend_data,this.ydata=t.data.xAxis_data)))}};const Q=(0,n.Z)(B,[["render",A]]);var q=Q;const K={class:"main-container"},H={class:"button"},F={class:"item"},W=["src"],M={key:0,class:"page"},Y=["onClick"],G={class:"product-list"},J=["src","alt"];function X(t,e,i,r,n,a){const l=(0,o.up)("router-link"),d=(0,o.up)("Recommend");return(0,o.wg)(),(0,o.iD)("div",{class:"wholepage",onMouseover:e[4]||(e[4]=t=>a.stopInv()),onMouseout:e[5]||(e[5]=t=>a.runInv()),onTouchstart:e[6]||(e[6]=t=>a.stopInv()),onTouchend:e[7]||(e[7]=t=>a.runInv())},[(0,o._)("div",K,[(0,o._)("header",null,[(0,o.wy)((0,o._)("input",{type:"text","onUpdate:modelValue":e[0]||(e[0]=t=>n.searchQuery=t),placeholder:"搜索商品"},null,512),[[s.nr,n.searchQuery]]),(0,o._)("button",{class:"button",onClick:e[1]||(e[1]=(...t)=>a.search&&a.search(...t))},"搜索"),(0,o._)("button",H,[(0,o.Wm)(l,{to:"/activity"},{default:(0,o.w5)((()=>[(0,o.Uk)("登录")])),_:1})])])]),(0,o._)("div",F,[(0,o._)("img",{src:n.dataList[n.currentIndex]},null,8,W)]),this.dataList.length>1?((0,o.wg)(),(0,o.iD)("div",M,[(0,o._)("ul",null,[(0,o._)("li",{onClick:e[2]||(e[2]=t=>a.gotoPage(a.prevIndex))},"<"),((0,o.wg)(!0),(0,o.iD)(o.HY,null,(0,o.Ko)(n.dataList,((t,e)=>((0,o.wg)(),(0,o.iD)("li",{onClick:t=>a.gotoPage(e),class:(0,E.C_)({current:n.currentIndex==e}),key:e},(0,E.zw)(e+1),11,Y)))),128)),(0,o._)("li",{onClick:e[3]||(e[3]=t=>a.gotoPage(a.nextIndex))},">")])])):(0,o.kq)("",!0),(0,o._)("div",G,[((0,o.wg)(!0),(0,o.iD)(o.HY,null,(0,o.Ko)(n.products,(t=>((0,o.wg)(),(0,o.iD)("div",{key:t.id,class:"product"},[(0,o._)("img",{src:t.image,alt:t.introduce},null,8,J),(0,o._)("h3",null,(0,E.zw)(t.introduce),1),(0,o._)("p",null,"价格: "+(0,E.zw)(t.price),1)])))),128))]),(0,o.Wm)(d)],32)}const tt={class:"total"},et={class:"layer"},it=["src"],st={class:"introduce"},ot={class:"price"};function rt(t,e,i,s,r,n){const a=(0,o.up)("router-link");return(0,o.wg)(),(0,o.iD)("div",tt,[(0,o._)("div",et,[((0,o.wg)(!0),(0,o.iD)(o.HY,null,(0,o.Ko)(r.goodsList,((t,e)=>((0,o.wg)(),(0,o.iD)("div",{class:"recommend",key:e},[(0,o.Wm)(a,{to:{name:"Details",params:{goodId:t.id}}},{default:(0,o.w5)((()=>[(0,o._)("ul",null,[(0,o._)("li",null,[(0,o._)("img",{src:t.image},null,8,it)]),(0,o._)("li",st,(0,E.zw)(t.introduce),1),(0,o._)("li",ot,"￥"+(0,E.zw)(t.price),1)])])),_:2},1032,["to"])])))),128))])])}var nt=JSON.parse('[{"image":"https://gw.alicdn.com/imgextra/i4/2214249724330/O1CN014NLq0f1hrBnw9wxHk_!!2-item_pic.png_.webp","introduce":"SCI发文章/EI会议评职称北大中文核心期刊省级CN普刊论文投稿翻译","price":50,"id":0,"counter":0},{"image":"https://gw.alicdn.com/imgextra/i2/2957472787/O1CN01Z4CRii1WSUfUgrPxF_!!0-item_pic.jpg_Q75.jpg_.webp","introduce":"【论文 硕士便宜】MBA本科硕博士专科查重复率官网论文检测报告","price":50,"id":1,"counter":0},{"image":"https://gd1.alicdn.com/imgextra/i1/60075666/O1CN01k1iUIj1rj4oG1zWZv_!!60075666.jpg","introduce":"【Bulygames】爆炸小黄人 中文小黄人版爆炸猫 休闲聚会欢乐桌游","price":198,"id":2,"counter":0},{"image":"https://gd3.alicdn.com/imgextra/i4/1840088272/O1CN01BMH9Ba2Ayd4gvCUTP_!!1840088272.jpg","introduce":"初音未来痛包现货玩偶周边2023新品可爱体系列软笑容收纳miku正版","price":109,"id":3,"counter":0},{"image":"https://g-search3.alicdn.com/img/bao/uploaded/i4/i2/617953624/O1CN011P2bbi1cdpvvloSQD_!!617953624.jpg_460x460Q90.jpg_.webp","introduce":"2020新款t恤女夏短袖宽松韩范学生百搭上衣简约大码bf风休闲半袖","price":78,"id":4,"counter":0},{"image":"https://g-search1.alicdn.com/img/bao/uploaded/i4/imgextra/i3/738710194/O1CN012rS4Zb1DItQk5AKni_!!0-saturn_solar.jpg_460x460Q90.jpg_.webp","introduce":"蓝牙耳机双耳头戴式跑步挂耳挂脖不入耳可插卡一体式骨传导耳麦","price":299,"id":5,"counter":0},{"image":"https://g-search3.alicdn.com/img/bao/uploaded/i4/i4/2201482912275/O1CN01PMAtgd1SfzeZ3V8W3_!!0-item_pic.jpg_460x460Q90.jpg_.webp","introduce":"Type-c耳机适用小米8/9华为荣耀高音质入耳式耳塞语音线控带麦","price":59,"id":6,"counter":0},{"image":"https://g-search3.alicdn.com/img/bao/uploaded/i4/i3/835893472/O1CN01lyBYFJ1bWDk3UjIEo_!!0-item_pic.jpg_460x460Q90.jpg_.webp","introduce":"骨传导蓝牙耳机不入耳无线跑步运动型双耳头戴挂耳挂脖式","price":399,"id":7,"counter":0},{"image":"https://g-search1.alicdn.com/img/bao/uploaded/i4/i4/38522192/O1CN014WSqOl1S3yk3s1IxK_!!38522192.jpg_460x460Q90.jpg_.webp","introduce":"ESSONIO骨传导耳机不入耳跑步骑行运动挂耳式自带内存无线蓝牙","price":1789,"id":8,"counter":0},{"image":"https://gd3.alicdn.com/imgextra/i1/820688566/O1CN01RsqS762D9HUSC555T_!!820688566.jpg","introduce":"老胡的店碧蓝档案千年阿拜斯三一logo白洲梓苹果iPhone14手机壳套","price":35.8,"id":9,"counter":0},{"image":"https://gd3.alicdn.com/imgextra/i4/0/O1CN01KLQOhX1fzlqfZNHV2_!!0-item_pic.jpg","introduce":"初音未来 美国洛杉矶演唱会 MIKUNOPOLIS in LOS ANGELES (DVD)","price":15,"id":9,"counter":0}]'),at={name:"recommend",data(){return{goodsList:nt}}};const lt=(0,n.Z)(at,[["render",rt],["__scopeId","data-v-1cf84334"]]);var dt=lt,ct={data(){return{searchQuery:"",products:[],dataList:["https://i1.mifile.cn/a4/xmad_15535933141925_ulkYv.jpg","https://i1.mifile.cn/a4/xmad_15532384207972_iJXSx.jpg","https://i1.mifile.cn/a4/xmad_15517939170939_oiXCK.jpg"],currentIndex:0,timer:null}},mounted(){this.fetchProducts(),setTimeout((()=>{this.runInv()}))},computed:{prevIndex(){return 0===this.currentIndex?this.dataList.length-1:this.currentIndex-1},nextIndex(){return this.currentIndex===this.dataList.length-1?0:this.currentIndex+1}},methods:{stopInv(){clearInterval(this.timer)},runInv(){this.timer=setInterval((()=>{this.gotoPage(this.nextIndex)}),3e3)},gotoPage(t){this.currentIndex=t},fetchProducts(){S.Z.get("/api/products").then((t=>{this.products=t.data})).catch((t=>{console.error(t)}))},search(){S.Z.get("/api/products",{params:{searchQuery:this.searchQuery}}).then((t=>{this.products=t.data})).catch((t=>{console.error(t)}))}},components:{Recommend:dt}};const ut=(0,n.Z)(ct,[["render",X],["__scopeId","data-v-00fc7270"]]);var gt=ut,mt=i(65);const pt={user:window.sessionStorage.getItem("user"),token:window.sessionStorage.getItem("token"),goodsList:nt,is_login:!1},ht={SET_TOKEN:(t,e)=>{t.token=e,window.sessionStorage.setItem("token",e),t.is_login=!0},GET_USER:(t,e)=>{t.user=e,window.sessionStorage.setItem("user",e)},LOGOUT:t=>{t.token=null,t.user=null,window.sessionStorage.removeItem("token"),window.sessionStorage.removeItem("user")}},_t={};var wt=(0,mt.MT)({state:pt,mutations:ht,actions:_t});const ft={class:"pic"},vt=["src"],bt={class:"fonts"},yt={class:"titles"},kt={"slot:titles":""},xt={class:"price"},Ct={"slot:price":""},It={class:"num"},Ut=["disabled"];function Dt(t,e,i,s,r,n){return(0,o.wg)(),(0,o.iD)("div",null,[(0,o._)("div",ft,[(0,o._)("img",{src:this.$store.state.goodsList[r.good_id-1].image},null,8,vt)]),(0,o._)("div",bt,[(0,o._)("div",yt,[(0,o._)("template",kt,[(0,o._)("h3",null,(0,E.zw)(this.$store.state.goodsList[r.good_id-1].introduce),1)])]),(0,o._)("div",xt,[(0,o._)("template",Ct,[(0,o._)("h2",null,"￥"+(0,E.zw)(this.$store.state.goodsList[r.good_id-1].price),1)])]),(0,o._)("div",It,[(0,o._)("div",null,[(0,o._)("button",{onClick:e[0]||(e[0]=t=>n.add(r.good_id))},"+"),(0,o._)("span",null," 购买数量："+(0,E.zw)(this.$store.state.goodsList[r.good_id-1].counter),1),(0,o._)("button",{onClick:e[1]||(e[1]=t=>n.sub(r.good_id)),disabled:r.dis},"-",8,Ut)]),(0,o._)("button",{class:(0,E.C_)(["stopcart",{activeBgColor:1==r.isActive}]),onClick:e[2]||(e[2]=(...t)=>n.addToCart&&n.addToCart(...t))}," 加入购物车 ",2)])])])}var Ot={name:"Details",data(){return{isActive:!1,good_id:this.$route.params.goodId,dis:!0}},methods:{add(t){this.$store.state.goodsList[t-1].counter++,this.$store.state.goodsList[t-1].counter>=0&&(this.dis=!1)},sub(t){this.$store.state.goodsList[t-1].counter--,0==this.$store.state.goodsList[t-1].counter?this.dis=!0:this.dis=!1},goHome(){this.$router.push("/home")},addToCart(){if(this.$store.state.goodsList[this.good_id-1].counter>=1){this.isActive=!0;const t=this.$store.state.goodsList[this.good_id-1].id,e=this.$store.state.cart,i=e.find((e=>e.id===t));void 0===i?e.push(this.$store.state.goodsList[this.good_id-1]):document.querySelector(".stopcart").innerHTML="请勿重复添加"}}}};const Pt=(0,n.Z)(Ot,[["render",Dt],["__scopeId","data-v-3c07f770"]]);var Vt=Pt,jt=i(2483);const Lt=[{path:"/Details",name:"Details",component:Vt},{path:"/",name:"/",component:gt},{path:"/login",name:"login",component:$},{path:"/activity",name:"activity",component:q,meta:{requireAuth:!0}}],St=(0,jt.p7)({history:(0,jt.r5)(),routes:Lt});St.beforeEach(((t,e,i)=>{const s=wt.state.token;t.meta.requireAuth?s?i():(console.log("该页面需要登陆"),i({path:"/login"})):i()}));var Rt=St;S.Z.interceptors.request.use((t=>(null!=wt.state.Authorization&&""!=wt.state.Authorization?t.headers.Authorization=wt.state.Authorization:console.log("no token"),t)),(t=>(console.log("error in request"),Promise.reject(t))));var Tt=S.Z;Tt.defaults.baseURL="http://localhost:8080/",(0,s.ri)(d).use(wt).use(Rt).mount("#app")},4654:function(){}},e={};function i(s){var o=e[s];if(void 0!==o)return o.exports;var r=e[s]={exports:{}};return t[s].call(r.exports,r,r.exports,i),r.exports}i.m=t,function(){var t=[];i.O=function(e,s,o,r){if(!s){var n=1/0;for(c=0;c<t.length;c++){s=t[c][0],o=t[c][1],r=t[c][2];for(var a=!0,l=0;l<s.length;l++)(!1&r||n>=r)&&Object.keys(i.O).every((function(t){return i.O[t](s[l])}))?s.splice(l--,1):(a=!1,r<n&&(n=r));if(a){t.splice(c--,1);var d=o();void 0!==d&&(e=d)}}return e}r=r||0;for(var c=t.length;c>0&&t[c-1][2]>r;c--)t[c]=t[c-1];t[c]=[s,o,r]}}(),function(){i.n=function(t){var e=t&&t.__esModule?function(){return t["default"]}:function(){return t};return i.d(e,{a:e}),e}}(),function(){i.d=function(t,e){for(var s in e)i.o(e,s)&&!i.o(t,s)&&Object.defineProperty(t,s,{enumerable:!0,get:e[s]})}}(),function(){i.g=function(){if("object"===typeof globalThis)return globalThis;try{return this||new Function("return this")()}catch(t){if("object"===typeof window)return window}}()}(),function(){i.o=function(t,e){return Object.prototype.hasOwnProperty.call(t,e)}}(),function(){var t={143:0};i.O.j=function(e){return 0===t[e]};var e=function(e,s){var o,r,n=s[0],a=s[1],l=s[2],d=0;if(n.some((function(e){return 0!==t[e]}))){for(o in a)i.o(a,o)&&(i.m[o]=a[o]);if(l)var c=l(i)}for(e&&e(s);d<n.length;d++)r=n[d],i.o(t,r)&&t[r]&&t[r][0](),t[r]=0;return i.O(c)},s=self["webpackChunkshoppingmall_vue3"]=self["webpackChunkshoppingmall_vue3"]||[];s.forEach(e.bind(null,0)),s.push=e.bind(null,s.push.bind(s))}();var s=i.O(void 0,[998],(function(){return i(7252)}));s=i.O(s)})();
//# sourceMappingURL=app.8a2a7e37.js.map