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
                      <td>{{$v.InstanceName}}</td>
                      <td>{{$v.Dbname}}</td>
                      <td>{{$v.Username}}</td>
                      <td>{{if eq 1 $v.Status}}激活{{else}}禁用{{end}}</td>
                      <td><div class="btn-group">
                          <button type="button" class="btn btn-primary">巡检</button>
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

</script>
</body>
</html>
