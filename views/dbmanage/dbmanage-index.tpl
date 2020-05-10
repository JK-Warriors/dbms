<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf-8">
<title>{{config "String" "globaltitle" ""}}</title>
{{template "inc/meta.tpl" .}}
</head>
<body class="sticky-header">
<section> 
  {{template "inc/left.tpl" .}}
  <!-- main content start-->
  <div class="main-content" >
    <!-- header section start-->
    <div class="header-section">
      <!--toggle button start-->
      <a class="toggle-btn"><i class="fa fa-bars"></i></a>
      <!--toggle button end-->
      <!--search start-->
      <!--search end-->
      {{template "inc/user-info.tpl" .}} 
    </div>
    <!-- header section end-->
    <!-- page heading start-->
    
    <div class="page-heading">
      <!--<h3> 组织管理 {{template "users/nav.tpl" .}}</h3>-->
      <ul class="breadcrumb pull-left">
        <li> <a href="/config/db/manage">配置</a> </li>
        <li class="active"> 数据库配置 </li>
      </ul>
    </div>
    <!-- page heading end-->
    <!--body wrapper start-->
    <div class="wrapper">
      <div class="row">
        <div class="col-sm-12">
          <div class="searchdiv">
            <div class="search-form">
              <div class="form-inline">
                <div class="form-group">
                  <form action="/config/db/manage" method="get">
                    <select name="dbtype" class="form-control">
                      <option value="">数据库类型</option>
                      <option value="1" {{if eq "1" .condArr.dbtype}}selected{{end}}>Oracle</option>
                      <option value="2" {{if eq "2" .condArr.dbtype}}selected{{end}}>MySQL</option>
                      <option value="3" {{if eq "3" .condArr.dbtype}}selected{{end}}>SQLServer</option>
                    </select>
                    <input type="text" class="form-control" name="host" placeholder="请输入IP" value="{{.condArr.host}}"/>
                    <input type="text" class="form-control" name="alias" placeholder="请输入别名" value="{{.condArr.alias}}"/>
                    <button type="submit" class="btn btn-primary"><i class="fa fa-search"></i>搜索</button>
                    <a href="/config/db/manage" class="btn btn-default" type="submit"> <i class="fa fa-reset"></i> 重置 </a>
                  </form>
                </div>
              </div>
            </div>
          </div>
          <section class="panel">
            <header class="panel-heading"> 数据库列表 / 总数：{{.countDb}}
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
                        <th>数据库类型</th>
                        <th>数据库IP</th>
                        <th>端口</th>
                        <th>别名</th>
                        <th>角色</th>
                        <th>实例名</th>
                        <th>数据库名</th>
                        <th>用户名</th>
                        <th>状态</th>
                        <th>操作</th>
                      </tr>
                    </thead>
                    <tbody>
                    
                    {{range $k,$v := .dbconf}}
                    <tr>
                      <td>{{getDBtype $v.Dbtype}}</td>
                      <td>{{$v.Host}}</td>
                      <td>{{$v.Port}}</a></td>
                      <td>{{$v.Alias}}</a></td>
                      <td>{{if eq 1 $v.Role}}主{{else}}备{{end}}</td>
                      <td>{{$v.InstanceName}}</td>
                      <td>{{$v.Dbname}}</td>
                      <td>{{$v.Username}}</td>
                      <td>{{if eq 1 $v.Status}}激活{{else}}禁用{{end}}</td>
                      <td><div class="btn-group">
                          <button type="button" class="btn btn-primary dropdown-toggle" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false"> 操作<span class="caret"></span> </button>
                          <ul class="dropdown-menu">
                            <li><a href="/config/db/edit/{{$v.Id}}">编辑</a></li>
                            <li role="separator" class="divider"></li>
                            {{if eq 1 $v.Status}}
                            <li><a href="javascript:;" class="js-dbconfig-single" data-id="{{$v.Id}}" data-status="2">禁用</a></li>
                            {{else}}
                            <li><a href="javascript:;" class="js-dbconfig-single" data-id="{{$v.Id}}" data-status="1">激活</a></li>
                            {{end}}
                            <li role="separator" class="divider"></li>
                            <li><a href="javascript:;" class="js-dbconfig-delete" data-op="delete" data-id="{{$v.Id}}">删除</a></li>
                          </ul>
                        </div></td>
                    </tr>
                    {{end}}
                    </tbody>
                    
                  </table>
                </form>
                {{template "inc/page.tpl" .}}
				      </section>
            </div>
          </section>
        </div>
      </div>
    </div>
    <!--body wrapper end-->
    <!--footer section start-->
    {{template "inc/foot-info.tpl" .}}
    <!--footer section end-->
  </div>
  <!-- main content end-->
</section>

{{template "inc/foot.tpl" .}}
<script>
    $('.js-dbconfig-single').on('click', function(){
    	var that = $(this);
    	var status = that.attr('data-status')
    	var id = that.attr('data-id');
      $.post('/config/db/ajax/status', { status: status, id: id },function(data){
        dialogInfo(data.message)
        if (data.code) {
          that.attr('data-status', status == 2 ? 1 : 2).text(status == 2 ? '激活' : '禁用').parents('td').prev('td').text(status == 2 ? '禁用' : '激活');
        } else {
          
        }
        setTimeout(function(){ $('#dialogInfo').modal('hide'); }, 1000);
      },'json');
    }); 
	
	$('.js-dbconfig-delete').on('click', function(){
		var that = $(this);
		var id = that.attr('data-id');

		layer.confirm('您确定要删除吗？', {
			btn: ['确定','取消'] //按钮
			,title:"提示"
		}, function(index){
			layer.close(index);
			
			$.post('/config/db/ajax/delete', {ids:id},function(data){
				dialogInfo(data.message)
				if (data.code) {
					setTimeout(function(){ window.location.reload() }, 1000);
				} else {
					setTimeout(function(){ $('#dialogInfo').modal('hide'); }, 1000);
				}
			},'json');
		});
		
	});

</script>
</body>
</html>
