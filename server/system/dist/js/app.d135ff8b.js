(function(){"use strict";var e={9328:function(e,t,i){var s=i(9242),o=i(3396);function r(e,t){const i=(0,o.up)("router-link"),s=(0,o.up)("router-view");return(0,o.wg)(),(0,o.iD)(o.HY,null,[(0,o.Wm)(i,{to:"/"}),(0,o.Wm)(s)],64)}var n=i(89);const a={},l=(0,n.Z)(a,[["render",r]]);var d=l;const c={class:"login-container"},u={class:"login-form"},g=(0,o._)("h2",null,"用户登录",-1),m={class:"form-group"},p=(0,o._)("label",null,"用户名",-1),h={key:0,class:"invalid-feedback"},_={class:"form-group"},f=(0,o._)("label",null,"密码",-1),w={class:"form-group mt-3"},v={class:"register-form"},b=(0,o._)("h2",null,"用户注册",-1),k={class:"form-group"},y=(0,o._)("label",null,"用户名",-1),x={key:0,class:"invalid-feedback"},C={class:"form-group"},I=(0,o._)("label",null,"密码",-1),U={class:"form-group"},P=(0,o._)("label",null,"确认密码",-1),O={key:0,class:"invalid-feedback"},S={class:"form-group mt-3"};function D(e,t,i,r,n,a){const l=(0,o.up)("el-button");return(0,o.wg)(),(0,o.iD)("div",c,[(0,o.wy)((0,o._)("div",u,[g,(0,o._)("form",null,[(0,o._)("div",m,[p,(0,o.wy)((0,o._)("input",{type:"text",class:"form-control","onUpdate:modelValue":t[0]||(t[0]=e=>n.username=e),placeholder:"请输入用户名",autocomplete:"off"},null,512),[[s.nr,n.username,void 0,{trim:!0}]]),n.isUsernameValid?(0,o.kq)("",!0):((0,o.wg)(),(0,o.iD)("div",h," 用户名长度应在5到12个字符之间。 "))]),(0,o._)("div",_,[f,(0,o.wy)((0,o._)("input",{type:"password",class:"form-control","onUpdate:modelValue":t[1]||(t[1]=e=>n.password=e),placeholder:"请输入密码",autocomplete:"off"},null,512),[[s.nr,n.password,void 0,{trim:!0}]])]),(0,o.Wm)(l,{type:"primary",class:"btn btn-primary",onClick:a.login,loading:n.isSending},{default:(0,o.w5)((()=>[(0,o.Uk)(" 登录 ")])),_:1},8,["onClick","loading"]),(0,o._)("div",w,[(0,o.Uk)(" 没有账号？"),(0,o._)("a",{href:"/#/login",onClick:t[2]||(t[2]=(...e)=>a.toggle&&a.toggle(...e))},"点击注册")]),(0,o.Wm)(l,{type:"primary",size:"small",onClick:a.back},{default:(0,o.w5)((()=>[(0,o.Uk)("返回主页")])),_:1},8,["onClick"])])],512),[[s.F8,!n.isRegisterFormVisible]]),(0,o.wy)((0,o._)("div",v,[b,(0,o._)("form",null,[(0,o._)("div",k,[y,(0,o.wy)((0,o._)("input",{type:"text",class:"form-control","onUpdate:modelValue":t[3]||(t[3]=e=>n.registerUsername=e),placeholder:"请输入用户名",autocomplete:"off"},null,512),[[s.nr,n.registerUsername,void 0,{trim:!0}]]),n.isRegisterUsernameValid?(0,o.kq)("",!0):((0,o.wg)(),(0,o.iD)("div",x," 用户名长度应在5到12个字符之间。 "))]),(0,o._)("div",C,[I,(0,o.wy)((0,o._)("input",{type:"password",class:"form-control","onUpdate:modelValue":t[4]||(t[4]=e=>n.registerPassword=e),placeholder:"请输入密码",autocomplete:"off"},null,512),[[s.nr,n.registerPassword,void 0,{trim:!0}]])]),(0,o._)("div",U,[P,(0,o.wy)((0,o._)("input",{type:"password",class:"form-control","onUpdate:modelValue":t[5]||(t[5]=e=>n.confirmPassword=e),placeholder:"请确认密码",autocomplete:"off"},null,512),[[s.nr,n.confirmPassword,void 0,{trim:!0}]]),n.isConfirmPasswordValid?(0,o.kq)("",!0):((0,o.wg)(),(0,o.iD)("div",O," 两次输入的密码不一致。 "))]),(0,o.Wm)(l,{type:"primary",class:"btn btn-primary",onClick:a.register,loading:n.isSending},{default:(0,o.w5)((()=>[(0,o.Uk)(" 注册 ")])),_:1},8,["onClick","loading"]),(0,o._)("div",S,[(0,o.Uk)(" 已有账号？"),(0,o._)("a",{href:"/#/login",onClick:t[6]||(t[6]=(...e)=>a.toggle&&a.toggle(...e))},"点击登录")]),(0,o.Wm)(l,{type:"primary",size:"small",onClick:a.back},{default:(0,o.w5)((()=>[(0,o.Uk)("返回主页")])),_:1},8,["onClick"])])],512),[[s.F8,n.isRegisterFormVisible]])])}var V=i(4161),$=i(1008),L={data(){return{isSending:!1,username:"",password:"",isUsernameValid:!0,isButtonDisabled:!1,isRegisterFormVisible:!1,registerUsername:"",registerPassword:"",confirmPassword:"",isRegisterUsernameValid:!0,isConfirmPasswordValid:!0,isRegisterButtonDisabled:!1}},methods:{back(){this.$router.replace({name:"MainPage"})},toggle(){this.isRegisterFormVisible=!this.isRegisterFormVisible,this.username="",this.password="",this.isUsernameValid=!0,this.isButtonDisabled=!1,this.registerUsername="",this.registerPassword="",this.confirmPassword="",this.isRegisterUsernameValid=!0,this.isConfirmPasswordValid=!0,this.isRegisterButtonDisabled=!1},login(){this.username.length<5||this.username.length>12?this.isUsernameValid=!1:this.isUsernameValid=!0,this.isUsernameValid&&this.password.length>0&&(this.isSending=!0,(0,V.Z)({headers:{"Content-Type":"application/x-www-form-urlencoded"},method:"post",url:"/api/Login",data:{user:this.username,pass:this.password}}).then((e=>{200===e.data.code&&(this.$store.commit("SET_TOKEN",e.data.token),this.$store.commit("GET_USER",e.data.user),console.log("登录成功！"),this.$router.replace({name:"MainPage"}))})).catch((function(){$.T$.alert("您输入的账号或密码错误").then((()=>{location.reload(),this.$router.go(0)}))})))},register(){this.registerUsername.length<5||this.registerUsername.length>12?this.isRegisterUsernameValid=!1:this.isRegisterUsernameValid=!0,this.registerPassword!==this.confirmPassword?this.isConfirmPasswordValid=!1:this.isConfirmPasswordValid=!0,this.isRegisterUsernameValid&&this.isConfirmPasswordValid&&(this.isSending=!0,(0,V.Z)({headers:{"Content-Type":"application/x-www-form-urlencoded"},method:"post",url:"/api/Register",data:{user:this.registerUsername,pass:this.registerPassword}}).then((e=>{200===e.data.code&&(this.$store.commit("SET_TOKEN",e.data.token),this.$store.commit("GET_USER",e.data.user),this.$message({message:"注册成功",type:"success"}),this.$router.replace({name:"MainPage"}))})).catch((function(){$.T$.alert("用户已存在！").then((()=>{location.reload(),this.$router.go(0)}))})))}}};const j=(0,n.Z)(L,[["render",D]]);var T=j,R=i(7139);const N={class:"hello"};function Z(e,t,i,s,r,n){return(0,o.wg)(),(0,o.iD)("div",N,[(0,o._)("h1",null,(0,R.zw)(r.msg),1),(0,o._)("h1",null,"site : "+(0,R.zw)(r.site),1),(0,o._)("h1",null,"url : "+(0,R.zw)(r.url),1),(0,o._)("h3",null,(0,R.zw)(n.details()),1),((0,o.wg)(!0),(0,o.iD)(o.HY,null,(0,o.Ko)(r.ydata,(e=>((0,o.wg)(),(0,o.iD)("h1",{key:e},(0,R.zw)(e),1)))),128)),((0,o.wg)(!0),(0,o.iD)(o.HY,null,(0,o.Ko)(r.xdata,(e=>((0,o.wg)(),(0,o.iD)("h3",{key:e},(0,R.zw)(e),1)))),128))])}var z={name:"apidata",data(){return{msg:"hello,xxbandy！",site:"bgops",url:"https://xxbandy.github.io",xdata:null,ydata:null}},methods:{details:function(){return this.site}},mounted(){V.Z.get("http://localhost:8080/v1/line").then((e=>(this.xdata=e.data.legend_data,this.ydata=e.data.xAxis_data)))}};const E=(0,n.Z)(z,[["render",Z]]);var W=E;const q={class:"item"},M=["src"],H={key:0,class:"page"},K=["onClick"],A={class:"product-list"},Q=["src","alt"];function B(e,t,i,r,n,a){const l=(0,o.up)("router-link"),d=(0,o.up)("Recommend");return(0,o.wg)(),(0,o.iD)(o.HY,null,[(0,o._)("div",{class:"wholepage",onMouseover:t[5]||(t[5]=e=>a.stopInv()),onMouseout:t[6]||(t[6]=e=>a.runInv()),onTouchstart:t[7]||(t[7]=e=>a.stopInv()),onTouchend:t[8]||(t[8]=e=>a.runInv())},[(0,o._)("header",null,[(0,o.wy)((0,o._)("input",{type:"text","onUpdate:modelValue":t[0]||(t[0]=e=>n.searchQuery=e),placeholder:"搜索商品"},null,512),[[s.nr,n.searchQuery]]),(0,o._)("button",{class:"button",onClick:t[1]||(t[1]=(...e)=>a.search&&a.search(...e))},"搜索"),(0,o._)("button",{class:"button",onClick:t[2]||(t[2]=(...e)=>a.login&&a.login(...e))},[(0,o.Wm)(l,{to:"/login"},{default:(0,o.w5)((()=>[(0,o.Uk)("登录")])),_:1})])]),(0,o._)("div",q,[(0,o._)("img",{src:n.dataList[n.currentIndex]},null,8,M)]),this.dataList.length>1?((0,o.wg)(),(0,o.iD)("div",H,[(0,o._)("ul",null,[(0,o._)("li",{onClick:t[3]||(t[3]=e=>a.gotoPage(a.prevIndex))},"<"),((0,o.wg)(!0),(0,o.iD)(o.HY,null,(0,o.Ko)(n.dataList,((e,t)=>((0,o.wg)(),(0,o.iD)("li",{onClick:e=>a.gotoPage(t),class:(0,R.C_)({current:n.currentIndex==t}),key:t},(0,R.zw)(t+1),11,K)))),128)),(0,o._)("li",{onClick:t[4]||(t[4]=e=>a.gotoPage(a.nextIndex))},">")])])):(0,o.kq)("",!0),(0,o._)("div",A,[((0,o.wg)(!0),(0,o.iD)(o.HY,null,(0,o.Ko)(n.products,(e=>((0,o.wg)(),(0,o.iD)("div",{key:e.id,class:"product"},[(0,o._)("img",{src:e.image,alt:e.introduce},null,8,Q),(0,o._)("h3",null,(0,R.zw)(e.introduce),1),(0,o._)("p",null,"价格: "+(0,R.zw)(e.price),1)])))),128))])],32),(0,o.Wm)(d)],64)}const F={class:"total"},Y={class:"layer"},G=["src"],J={class:"introduce"},X={class:"price"};function ee(e,t,i,s,r,n){const a=(0,o.up)("router-link");return(0,o.wg)(),(0,o.iD)("div",F,[(0,o._)("div",Y,[((0,o.wg)(!0),(0,o.iD)(o.HY,null,(0,o.Ko)(r.goodsList,((e,t)=>((0,o.wg)(),(0,o.iD)("div",{class:"recommend",key:t},[(0,o.Wm)(a,{to:{name:"Details",params:{goodId:e.id}}},{default:(0,o.w5)((()=>[(0,o._)("ul",null,[(0,o._)("li",null,[(0,o._)("img",{src:e.image},null,8,G)]),(0,o._)("li",J,(0,R.zw)(e.introduce),1),(0,o._)("li",X,"￥"+(0,R.zw)(e.price),1)])])),_:2},1032,["to"])])))),128))])])}var te=JSON.parse('[{"image":"https://gw.alicdn.com/imgextra/i4/2214249724330/O1CN014NLq0f1hrBnw9wxHk_!!2-item_pic.png_.webp","introduce":"SCI发文章/EI会议评职称北大中文核心期刊省级CN普刊论文投稿翻译","price":50,"id":0,"counter":0},{"image":"https://gw.alicdn.com/imgextra/i2/2957472787/O1CN01Z4CRii1WSUfUgrPxF_!!0-item_pic.jpg_Q75.jpg_.webp","introduce":"【论文 硕士便宜】MBA本科硕博士专科查重复率官网论文检测报告","price":50,"id":1,"counter":0},{"image":"https://gd1.alicdn.com/imgextra/i1/60075666/O1CN01k1iUIj1rj4oG1zWZv_!!60075666.jpg","introduce":"【Bulygames】爆炸小黄人 中文小黄人版爆炸猫 休闲聚会欢乐桌游","price":198,"id":2,"counter":0},{"image":"https://gd3.alicdn.com/imgextra/i4/1840088272/O1CN01BMH9Ba2Ayd4gvCUTP_!!1840088272.jpg","introduce":"初音未来痛包现货玩偶周边2023新品可爱体系列软笑容收纳miku正版","price":109,"id":3,"counter":0},{"image":"https://g-search3.alicdn.com/img/bao/uploaded/i4/i2/617953624/O1CN011P2bbi1cdpvvloSQD_!!617953624.jpg_460x460Q90.jpg_.webp","introduce":"2020新款t恤女夏短袖宽松韩范学生百搭上衣简约大码bf风休闲半袖","price":78,"id":4,"counter":0},{"image":"https://g-search1.alicdn.com/img/bao/uploaded/i4/imgextra/i3/738710194/O1CN012rS4Zb1DItQk5AKni_!!0-saturn_solar.jpg_460x460Q90.jpg_.webp","introduce":"蓝牙耳机双耳头戴式跑步挂耳挂脖不入耳可插卡一体式骨传导耳麦","price":299,"id":5,"counter":0},{"image":"https://g-search3.alicdn.com/img/bao/uploaded/i4/i4/2201482912275/O1CN01PMAtgd1SfzeZ3V8W3_!!0-item_pic.jpg_460x460Q90.jpg_.webp","introduce":"Type-c耳机适用小米8/9华为荣耀高音质入耳式耳塞语音线控带麦","price":59,"id":6,"counter":0},{"image":"https://g-search3.alicdn.com/img/bao/uploaded/i4/i3/835893472/O1CN01lyBYFJ1bWDk3UjIEo_!!0-item_pic.jpg_460x460Q90.jpg_.webp","introduce":"骨传导蓝牙耳机不入耳无线跑步运动型双耳头戴挂耳挂脖式","price":399,"id":7,"counter":0},{"image":"https://g-search1.alicdn.com/img/bao/uploaded/i4/i4/38522192/O1CN014WSqOl1S3yk3s1IxK_!!38522192.jpg_460x460Q90.jpg_.webp","introduce":"ESSONIO骨传导耳机不入耳跑步骑行运动挂耳式自带内存无线蓝牙","price":1789,"id":8,"counter":0},{"image":"https://gd3.alicdn.com/imgextra/i1/820688566/O1CN01RsqS762D9HUSC555T_!!820688566.jpg","introduce":"老胡的店碧蓝档案千年阿拜斯三一logo白洲梓苹果iPhone14手机壳套","price":35.8,"id":9,"counter":0},{"image":"https://gd3.alicdn.com/imgextra/i4/0/O1CN01KLQOhX1fzlqfZNHV2_!!0-item_pic.jpg","introduce":"初音未来 美国洛杉矶演唱会 MIKUNOPOLIS in LOS ANGELES (DVD)","price":15,"id":10,"counter":0}]'),ie={name:"recommend",data(){return{goodsList:te}}};const se=(0,n.Z)(ie,[["render",ee],["__scopeId","data-v-1cf84334"]]);var oe=se,re={data(){return{searchQuery:"",products:[],dataList:["https://i1.mifile.cn/a4/xmad_15535933141925_ulkYv.jpg","https://i1.mifile.cn/a4/xmad_15532384207972_iJXSx.jpg","https://i1.mifile.cn/a4/xmad_15517939170939_oiXCK.jpg"],currentIndex:0,timer:null}},mounted(){this.fetchProducts(),setTimeout((()=>{this.runInv()}))},computed:{prevIndex(){return 0===this.currentIndex?this.dataList.length-1:this.currentIndex-1},nextIndex(){return this.currentIndex===this.dataList.length-1?0:this.currentIndex+1}},methods:{login(){this.$router.replace({name:"login"})},stopInv(){clearInterval(this.timer)},runInv(){this.timer=setInterval((()=>{this.gotoPage(this.nextIndex)}),3e3)},gotoPage(e){this.currentIndex=e},fetchProducts(){V.Z.get("/api/products").then((e=>{this.products=e.data})).catch((e=>{console.error(e)}))},search(){$.T$.alert("当前没有连接到后端").then((()=>{this.$router.replace({name:"login"})}))}},components:{Recommend:oe}};const ne=(0,n.Z)(re,[["render",B],["__scopeId","data-v-3a897cf4"]]);var ae=ne,le=i(65);const de={user:window.sessionStorage.getItem("user"),token:window.sessionStorage.getItem("token"),goodsList:te,is_login:!1},ce={SET_TOKEN:(e,t)=>{e.token=t,window.sessionStorage.setItem("token",t),e.is_login=!0},GET_USER:(e,t)=>{e.user=t,window.sessionStorage.setItem("user",t)},LOGOUT:e=>{e.token=null,e.user=null,window.sessionStorage.removeItem("token"),window.sessionStorage.removeItem("user")}},ue={};var ge=(0,le.MT)({state:de,mutations:ce,actions:ue});const me={key:0},pe={class:"pic"},he=["src"],_e={class:"fonts"},fe={class:"titles"},we={class:"price"},ve={class:"num"},be=["disabled"];function ke(e,t,i,s,r,n){const a=(0,o.up)("slot:titles"),l=(0,o.up)("slot:price");return this.$store.state.goodsList[r.good_id]?((0,o.wg)(),(0,o.iD)("div",me,[(0,o._)("div",pe,[(0,o._)("img",{src:this.$store.state.goodsList[r.good_id].image},null,8,he)]),(0,o._)("div",_e,[(0,o._)("div",fe,[(0,o.Wm)(a,null,{default:(0,o.w5)((()=>[(0,o._)("h3",null,(0,R.zw)(this.$store.state.goodsList[r.good_id].introduce),1)])),_:1})]),(0,o._)("div",we,[(0,o.Wm)(l,null,{default:(0,o.w5)((()=>[(0,o._)("h2",null,"￥"+(0,R.zw)(this.$store.state.goodsList[r.good_id].price),1)])),_:1})]),(0,o._)("div",ve,[(0,o._)("div",null,[(0,o._)("button",{onClick:t[0]||(t[0]=e=>n.add(r.good_id))},"+"),(0,o._)("span",null," 购买数量："+(0,R.zw)(this.$store.state.goodsList[r.good_id].counter),1),(0,o._)("button",{onClick:t[1]||(t[1]=e=>n.sub(r.good_id)),disabled:r.dis},"-",8,be)]),(0,o._)("button",{class:(0,R.C_)(["stopcart",{activeBgColor:1==r.isActive}]),onClick:t[2]||(t[2]=(...e)=>n.addToCart&&n.addToCart(...e))}," 加入购物车 ",2)])])])):(0,o.kq)("",!0)}i(7658);var ye={name:"Details",data(){return{good_id:this.$route.params.goodId,isActive:!1,dis:!0}},methods:{add(e){this.$store.state.goodsList[e].counter++,this.$store.state.goodsList[e].counter>=0&&(this.dis=!1)},sub(e){this.$store.state.goodsList[e].counter--,0==this.$store.state.goodsList[e].counter?this.dis=!0:this.dis=!1},goHome(){this.$router.push("/home")},addToCart(){if(this.$store.state.goodsList[this.good_id].counter>=1){this.isActive=!0;const e=this.$store.state.goodsList[this.good_id].id,t=this.$store.state.cart,i=t.find((t=>t.id===e));void 0===i?t.push(this.$store.state.goodsList[this.good_id]):document.querySelector(".stopcart").innerHTML="请勿重复添加"}}}};const xe=(0,n.Z)(ye,[["render",ke],["__scopeId","data-v-009b135c"]]);var Ce=xe,Ie=i(2483);const Ue=[{path:"/MainPage/Details/:goodId",name:"Details",meta:{requireAuth:!0},component:Ce},{path:"/",name:"MainPage",component:ae},{path:"/login",name:"login",component:T},{path:"/activity",name:"activity",component:W,meta:{requireAuth:!0}}],Pe=(0,Ie.p7)({history:(0,Ie.r5)(),routes:Ue});Pe.beforeEach(((e,t,i)=>{const s=ge.state.token;e.meta.requireAuth?s?i():(window.console.log("该页面需要登陆"),i({path:"/login"})):i()}));var Oe=Pe;V.Z.defaults.baseURL="http://localhost:8080/",V.Z.interceptors.request.use((e=>{const t=sessionStorage.getItem("token");return e.headers["token"]=t,e}),(e=>Promise.error(e)));var Se=V.Z,De=i(50),Ve=e=>{e.use($.ZP,{locale:De.Z})};const $e=(0,s.ri)(d);$e.config.globalProperties.$axios=Se,Ve($e),$e.use(ge).use($.ZP).use(Oe).mount("#app")}},t={};function i(s){var o=t[s];if(void 0!==o)return o.exports;var r=t[s]={id:s,loaded:!1,exports:{}};return e[s].call(r.exports,r,r.exports,i),r.loaded=!0,r.exports}i.m=e,function(){var e=[];i.O=function(t,s,o,r){if(!s){var n=1/0;for(c=0;c<e.length;c++){s=e[c][0],o=e[c][1],r=e[c][2];for(var a=!0,l=0;l<s.length;l++)(!1&r||n>=r)&&Object.keys(i.O).every((function(e){return i.O[e](s[l])}))?s.splice(l--,1):(a=!1,r<n&&(n=r));if(a){e.splice(c--,1);var d=o();void 0!==d&&(t=d)}}return t}r=r||0;for(var c=e.length;c>0&&e[c-1][2]>r;c--)e[c]=e[c-1];e[c]=[s,o,r]}}(),function(){i.n=function(e){var t=e&&e.__esModule?function(){return e["default"]}:function(){return e};return i.d(t,{a:t}),t}}(),function(){i.d=function(e,t){for(var s in t)i.o(t,s)&&!i.o(e,s)&&Object.defineProperty(e,s,{enumerable:!0,get:t[s]})}}(),function(){i.g=function(){if("object"===typeof globalThis)return globalThis;try{return this||new Function("return this")()}catch(e){if("object"===typeof window)return window}}()}(),function(){i.o=function(e,t){return Object.prototype.hasOwnProperty.call(e,t)}}(),function(){i.nmd=function(e){return e.paths=[],e.children||(e.children=[]),e}}(),function(){var e={143:0};i.O.j=function(t){return 0===e[t]};var t=function(t,s){var o,r,n=s[0],a=s[1],l=s[2],d=0;if(n.some((function(t){return 0!==e[t]}))){for(o in a)i.o(a,o)&&(i.m[o]=a[o]);if(l)var c=l(i)}for(t&&t(s);d<n.length;d++)r=n[d],i.o(e,r)&&e[r]&&e[r][0](),e[r]=0;return i.O(c)},s=self["webpackChunkshoppingmall_vue3"]=self["webpackChunkshoppingmall_vue3"]||[];s.forEach(t.bind(null,0)),s.push=t.bind(null,s.push.bind(s))}();var s=i.O(void 0,[998],(function(){return i(9328)}));s=i.O(s)})();
//# sourceMappingURL=app.d135ff8b.js.map