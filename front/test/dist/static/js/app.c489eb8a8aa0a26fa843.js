webpackJsonp([1],{"+skl":function(e,s){},"AE+t":function(e,s,t){e.exports=t.p+"static/img/tranfer.37330ac.png"},NHnr:function(e,s,t){"use strict";Object.defineProperty(s,"__esModule",{value:!0});var a=t("7+uW"),n={render:function(){var e=this.$createElement,s=this._self._c||e;return s("div",{attrs:{id:"app"}},[s("meta",{attrs:{name:"viewport",content:"width=device-width, initial-scale=1"}}),this._v(" "),s("router-view")],1)},staticRenderFns:[]};var i=t("VU/8")({name:"App"},n,!1,function(e){t("fq6c")},null,null).exports,l=t("/ocq"),o=t("BTaQ"),r=t.n(o),c={data:function(){var e=this;return{selectKey:"",self:this,tablePage:{total:100,pageIndex:1,pageSize:10},itemsPerPages:[5,10,20,30,40],columns1:[{title:"文件名",align:"center",key:"filename",render:function(s,t){return s("div",[s(o.Button,{props:{size:"small"},style:{marginRight:"0.3125rem",overflow:"hidden",cursor:"pointer"},on:{click:function(){var s=t.row.filename;e.download_file(s)}}},t.row.filename)])}},{title:"操作",key:"action",align:"center",render:function(s,t){return s("div",[s(o.Button,{props:{type:"error",size:"small"},on:{click:function(){console.log(t);var s=t.index,a=t.row.filename;e.delete_file(s,a)}}},"删除")])}}],data1:[{filename:"图片ssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssss.png"},{filename:"视频.mp4"},{filename:"1.txt"},{filename:"身份证.jpg"}]}},mounted:function(){var e=this;window.onresize=function(s){e.$refs.tableData.handleResize()},this.getData()},methods:{download_file:function(e){var s=document.createElement("a");s.href="/data/"+e,s.download=e,s.style.display="none",document.body.appendChild(s),s.click(),s.remove()},delete_file:function(e,s){var t=this,a={fileName:s};t.axios.delete("/api/deleteFile",{params:a}).then(function(e){200==e.status?t.getData():t.$Modal.error({title:"删除",content:e.msg})}).catch(function(e){console.log(e)})},select_file:function(){this._data.tablePage.pageIndex=1,this.getData()},getData:function(){var e=this,s={key:this._data.selectKey,pageIndex:this._data.tablePage.pageIndex,pageSize:this._data.tablePage.pageSize};e.axios.get("/api/files",{params:s}).then(function(s){if(200==s.status){var t=[];s.data.data.forEach(function(e,s){var a={filename:e};t.push(a)}),e._data.data1=t,e._data.tablePage.total=s.data.total}else e.$Modal.info({title:"查询数据",content:"获取失败"})}).catch(function(e){console.log(e)})},handlePage:function(e){this._data.tablePage.pageIndex=e,this.getData()},handlePageSize:function(e){this._data.tablePage.pageIndex=1,this._data.tablePage.pageSize=e,this.getData()}}},d={render:function(){var e=this,s=e.$createElement,a=e._self._c||s;return a("div",{staticStyle:{width:"90%",left:"5%",position:"relative"}},[a("img",{staticStyle:{width:"15%"},attrs:{src:t("AE+t")}}),e._v(" "),a("br"),e._v(" "),a("Input",{staticStyle:{width:"60%",margin:"0.675rem"},attrs:{size:"large",placeholder:"请输入你想查询的文件名..."},on:{"on-enter":e.select_file},model:{value:e.selectKey,callback:function(s){e.selectKey=s},expression:"selectKey"}}),e._v(" "),a("Button",{attrs:{type:"primary"},on:{click:e.select_file}},[e._v("查询")]),e._v(" "),a("br"),e._v(" "),a("Upload",{attrs:{"on-success":e.getData,multiple:"",type:"drag",action:"/api/uploadFile"}},[a("div",{staticStyle:{padding:"1.25rem","background-color":"coral"}},[a("Icon",{staticStyle:{color:"#8F8F9E"},attrs:{type:"ios-cloud-upload",size:"100"}}),e._v(" "),a("p",{staticClass:"p1"},[e._v("点击或将文件拖拽到这里上传")])],1)]),e._v(" "),a("br"),e._v(" "),a("Table",{ref:"tableData",attrs:{stripe:"",border:"",context:e.self,columns:e.columns1,data:e.data1}}),e._v(" "),a("Page",{staticStyle:{"text-align":"right",margin:"0.675rem"},attrs:{total:e.tablePage.total,current:e.tablePage.pageIndex,"page-size-opts":e.itemsPerPages,"show-sizer":"","show-total":""},on:{"on-change":e.handlePage,"on-page-size-change":e.handlePageSize}})],1)},staticRenderFns:[]};var u=t("VU/8")(c,d,!1,function(e){t("bErE")},"data-v-6678260a",null).exports;a.default.use(l.a);var p=new l.a({routes:[{path:"/",name:"FileTranfer",component:u}]}),f=t("mtWM"),g=t.n(f),m=t("aLYK");t("+skl");a.default.use(m.a,g.a),a.default.use(r.a),a.default.config.productionTip=!1,new a.default({el:"#app",router:p,components:{App:i},template:"<App/>"})},bErE:function(e,s){},fq6c:function(e,s){}},["NHnr"]);
//# sourceMappingURL=app.c489eb8a8aa0a26fa843.js.map