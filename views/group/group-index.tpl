<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf-8" />
<title>{{config "String" "globaltitle" ""}}</title>
{{template "inc/meta.tpl" .}}
</head>
<body class="sticky-header">
<section>
  {{template "inc/left.tpl" .}}
  <!-- main content start-->
  <div class="main-content">
    <!-- header section start-->
    <div class="header-section">
      <a class="toggle-btn"><i class="fa fa-bars"></i></a>
      <!--search start-->
      <!--search end-->
      {{template "inc/user-info.tpl" .}}
    </div>
    <!-- header section end-->
    <!-- page heading start-->
    <div class="page-heading">
      <!-- <h3> 日志管理 </h3>-->
      <ul class="breadcrumb pull-left">
        <li><a href="/config/business/manage">配置</a></li>
        <li class="active">分组管理</li>
      </ul>
    </div>
    <!-- page heading end-->
    <!--body wrapper start-->
    <div class="wrapper">
      <div class="row">
        <div class="col-sm-12">
          <!-- 主体内容 开始 -->
            <div class="searchdiv">
              <div class="search-form">
                <div class="form-inline">
                  <div class="form-group">
                    <form action="/config/group/manage" method="get">
                    <input type="text" name="search_name" placeholder="请输入名称" class="form-control" value="{{.condArr.search_name}}"/>
                    <button class="btn btn-primary" type="submit"> <i class="fa fa-search"></i> 搜索 </button>
                    <a href="/config/group/manage" class="btn btn-default" type="submit"> <i class="fa fa-reset"></i> 重置 </a>
                    </form>
                  </div>
                </div>
              </div>
              <div class="pull-right">
                <a href="javascript:;" class="btn btn-success" id="add_group">
                  <i class="fa fa-plus"></i> 新增分组</a>
              </div>
            </div>

            <section class="panel">
              <header class="panel-heading"> 分组列表 / 总数：{{.countGp}}
                <span class="tools pull-right"><a href="javascript:;" class="fa fa-chevron-down"></a>
                <!--a href="javascript:;" class="fa fa-times"></a-->
                </span> 
              </header>
              <div class="panel-body">
                <section id="unseen">
                  <form id="user-form-list">
                    <table class="table table-bordered table-striped table-condensed">
                      <thead>
                        <tr>
                          <th>组名</th>
                          <th>操作</th>
                        </tr>
                      </thead>
                      <tbody>
                      {{range $k,$v := .gpconf}}
                        <tr>
                          <td>{{$v.GroupName}}</td>
                          <td>
                            <a href="javascript:;" class="table_btn table_btn_icon" onclick="edit_group(this)" data-id="{{$v.Id}}" data-name="{{$v.GroupName}}">
                              <i class="iconfont icon-btn_edit"></i>编辑
                            </a>
                            <a href="javascript:;" class="table_btn table_btn_icon" onclick="delete_group(this)" data-id="{{$v.Id}}">
                              <i class="iconfont icon-iconfontshanchu"></i>
                            </a>
                          </td>
                        </tr>
                      {{end}}
                      </tbody>
                    </table>
                  </form>
                  {{template "inc/page.tpl" .}}
                </section>
              </div>
            </section>
          <!-- 主体内容 结束 -->
        </div>
      </div>
    </div>
    <!--body wrapper end-->
    <!--footer section start-->
    {{template "inc/foot-info.tpl" .}}
    <!--footer section end-->
  </div>
  <!-- main content end-->
      
  <form id="group-form">
    <div id="group_box" class="layui_drm">
      <div class="layercontent">
        <!-- layer content start -->
        <div class="form-horizontal adminex-form">
          <div class="form-group">
            <label class="col-xs-2  control-label">分组名称</label>
            <div class="col-xs-10">
              <input type="hidden" id="group_id" name="group_id" value="" class="form-control"/>
              <input type="text" id="group_name" name="group_name" value="" class="form-control" placeholder="请填写分组名称"/>
            </div>
          </div>
        </div>
        <!-- layer content end -->
      </div>
      <!-- layer foot start -->
      <div class="layerfoot">
        <button type="submit" class="btn btn-primary">提 交</button>
      </div>
      <!-- layer content end -->
    </div>
  </form>
</section>

{{template "inc/foot.tpl" .}}    
<script>
    //layer
    $(function() {
      $('#add_group').click(function() {
        $("#group_id").attr("value",'');
        $("#group_name").attr("value",'');

        layer.open({
          type: 1,
          closeBtn: true,
          shift: 2,
          title: '新增分组',
          area: ['600px', '30%'],
          offset: ['180px'],
          shadeClose: true,
          content: $('#group_box')
        })
      })
    })

    
    function edit_group(e){
      var id = e.getAttribute("data-id");
      var group_name = e.getAttribute("data-name");

      $("#group_id").attr("value",id);
      $("#group_name").attr("value",bs_name);

      layer.open({
        type: 1,
        closeBtn: true,
        shift: 2,
        title: '编辑分组',
        area: ['600px', '30%'],
        offset: ['180px'],
        shadeClose: true,
        content: $('#business_box')
      })
    }


    function delete_group(e){
      var id = e.getAttribute("data-id");
      
      layer.confirm('您确定要删除吗？', {
        btn: ['确定','取消'] //按钮
        ,title:"提示"
      }, function(index){
        layer.close(index);
        
        $.post('/config/group/ajax/delete', {ids:id},function(data){
          dialogInfo(data.message)
          if (data.code) {
            setTimeout(function(){ window.location.reload() }, 1000);
          } else {
            setTimeout(function(){ $('#dialogInfo').modal('hide'); }, 1000);
          }
        },'json');
      });
    }

    
    $('#group-form').validate({
        ignore:'',        
		    rules : {
			      bs_name:{required: true},
        },
        messages : {
			      bs_name:{required: '请填写组名'},
        },

        submitHandler:function(form) {
            var id = $("#group_id").val()
            if(id == ""){
                target_url = "/config/group/add";
            }else{
                target_url = "/config/group/edit";
            }
            $(form).ajaxSubmit({
                type:'POST',
                url: target_url, 
                dataType:'json',
                success:function(data) {
                    dialogInfo(data.message)
                    if (data.code) {
                       setTimeout(function(){window.location.href="/config/group/manage"}, 1000);
                    } else {
                       setTimeout(function(){ $('#dialogInfo').modal('hide'); }, 1000);
                    }
                }
            });
        }

    });
</script>
</body>
</html>
