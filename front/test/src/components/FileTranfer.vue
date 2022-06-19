<template>
    <div  style="width: 90%;left: 5%;position:relative;">
     <!-- <div> -->
    <img style="width: 15%" src="../assets/tranfer.png">
    <br>
    <Input v-model="selectKey" size="large" placeholder="请输入你想查询的文件名..."  style="width: 60%; margin: 0.675rem;" @on-enter="select_file"></Input> 
    <Button type="primary" @click="select_file">查询</Button>
    <br/>
   
    <Upload
        :on-success="getData"
        multiple
        type="drag"
        action="/api/uploadFile"
        >
        <div style="padding: 1.25rem ;background-color:coral;">
            <Icon type="ios-cloud-upload" size="100" style="color: #8F8F9E"></Icon>
            <p class="p1">点击或将文件拖拽到这里上传</p>
        </div>
    </Upload>
    <br>
    <Table ref='tableData' stripe border :context="self" :columns="columns1" :data="data1"></Table>
    <Page     :total="tablePage.total"  :current="tablePage.pageIndex"   :page-size-opts="itemsPerPages" 
        show-sizer  
        show-total  
        style="text-align: right;margin: 0.675rem;"
        @on-change="handlePage" @on-page-size-change='handlePageSize'
        ></Page>
    </div>
</template>

<script>

import { Button } from 'iview';
export default {
    data () {
    let _this=this;
    return {
            selectKey:"",
            self: this,
            tablePage:{
                total:100,
                pageIndex:1,
                pageSize:10,
              
            },
            itemsPerPages:[5,10,20,30,40],
            columns1: [
                {
                    title: '文件名',
                    align:'center',
                    key: 'filename',
                    render :function (h, params) {
                        //console.log(params)
                                return h('div',[h(Button,{
                                            props: {
                                                // type: 'success',
                                                size: 'small',
                                              
                                            },
                                            style:{
                                                marginRight: '0.3125rem',
                                                // color: '#515a6e',
                                                // fontSize: '14px',
                                                overflow: 'hidden',
                                                // textOverflow: 'ellipsis',
                                                // whiteSpace: 'nowrap',
                                                //display: 'block'     ,//设置样式，超过文字省略号显示
                                                cursor: 'pointer'
                                            },
                                            on: {
                                                click:()=> {
                                                   // console.log(_this)
                                                    var value=params.row.filename
                                                    _this.download_file(value)
                                                }
                                            }
                                },params.row.filename)])
                            }
                },
                {
                    title: '操作',
                    key: 'action',
                    align:'center',
                    render :function (h, params) {
                                return h('div',[h(Button,{
                                     props: {
                                                type: 'error',
                                                size: 'small'
                                            },
                                            on: {
                                                click: () => {
                                                    console.log(params)
                                                    var index=params.index
                                                    var value=params.row.filename
                                                   _this.delete_file( index,value)
                                                }
                                            }
                                },"删除")])
                            }
                }
            ],
            data1: [
                {
                    filename: '图片ssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssss.png',
                },
                {
                    filename: '视频.mp4',
                },
                {
                    filename: '1.txt',
                },
                {
                    filename: '身份证.jpg',
                }
            ]
        }
    },
    mounted: function () {
        window.onresize = event => {
            this.$refs.tableData.handleResize();       
        }
        this.getData()
    },
    methods: {
        download_file(fileName) {
            var a = document.createElement("a"); //创建一个<a></a>标签
            a.href = "/data/"+fileName; // 给a标签的href属性值加上地址，注意，这里是绝对路径，不用加 点.
            a.download = fileName; //
            a.style.display = "none"; // 障眼法藏起来a标签
            document.body.appendChild(a); // 将a标签追加到文档对象中
            a.click(); // 模拟点击了a标签，会触发a标签的href的读取，浏览器就会自动下载了
            a.remove(); // 一次性的，用完就删除a标签
        },
        delete_file (index,value) {
            let _this=this;
            let param={
                fileName:value
            };
            _this.axios.delete('/api/deleteFile',
                {
                    params: param
                }) .then(function (response) {
                   // console.log(response)
                    if(response.status==200){
                        _this.getData()
                    }else{
                        _this.$Modal.error({
                            title: '删除',
                            content: response.msg
                        })
                    }
                })
                .catch(function (error) {
                    console.log(error);
            });
            
        },
        select_file(){
            this._data.tablePage.pageIndex = 1;
            this.getData();
        },
        getData(){
            let _this=this;
            //获取数据
            let param={
                key:this._data.selectKey,
                pageIndex:this._data.tablePage.pageIndex,
                pageSize:this._data.tablePage.pageSize
            }
           // console.log(_this)
            _this.axios.get('/api/files',
                {
                    params: param
                })   
                .then(function (response) {
                    if(response.status==200){
                      //  console.log(response)
                            let _data=[]
                          //  console.log(response)
                            response.data.data.forEach (function (item, intex) {
                            var model = {
                                    filename:item
                                }
                            _data.push(model);
                        })
                        _this._data.data1=_data;
                        _this._data.tablePage.total=response.data.total
                    }else{
                        _this.$Modal.info({
                            title: '查询数据',
                            content: "获取失败"
                        })
                    }
                })
                .catch(function (error) {
                    console.log(error);
                }
            );
        },
        handlePage(value){
            this._data.tablePage.pageIndex = value;
            this.getData();
        },
        handlePageSize(value){
            this._data.tablePage.pageIndex = 1;
            this._data.tablePage.pageSize = value;
            this.getData();
        }
    }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1, h2 {
  font-weight: normal;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 0.675rem;
}
a {
  color: #42b983;
}

.p1{
    font-family: "宋体","仿宋",sans-serif;/*若电脑不支持宋体，则用仿宋，若不支持仿宋，则在sans-serif中找*/
    font-weight: bold;
    font-size: 150%;
    font-style: italic;
    color: black;/*字体颜色*/
    opacity: 0.7;/*字体的透明度：1：默认样式，0：全透明*/
}
</style>
