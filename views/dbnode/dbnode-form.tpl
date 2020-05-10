<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf-8">
<title>{{config "String" "globaltitle" ""}}</title>
{{template "inc/meta.tpl" .}}
</head><body class="sticky-header">
<section> {{template "inc/left.tpl" .}}
  <!-- main content start-->
  <div class="main-content" >
    <!-- header section start-->
    <div class="header-section">
      <!--toggle button start-->
      <a class="toggle-btn"><i class="fa fa-bars"></i></a> {{template "inc/user-info.tpl" .}} </div>
    <!-- header section end-->
    <!-- page heading start-->
    <div class="page-heading">
      <!-- <h3> 组织管理 {{template "users/nav.tpl" .}}</h3>-->
      <ul class="breadcrumb pull-left">
        <li> <a href="/config/group/manage">配置</a> </li>
        <li> <a href="/config/node/manage">节点配置</a> </li>
        <li class="active"> {{if gt .osconf.Id 0}}编辑{{else}}新增{{end}}节点 </li>
      </ul>
      <div class="pull-right"><a href="/config/node/add" class="btn btn-success">+添加节点</a></div>
    </div>
    <!-- page heading end-->
    <!--body wrapper start-->
    <div class="wrapper">
      <div class="row">
        <div class="col-lg-12">
          <section class="panel">
            <header class="panel-heading"> {{.title}} </header>
            <div class="panel-body">
              <form class="form-horizontal adminex-form" id="dbnode-form">
                <header><b> 基本信息 </b></header>
                <div class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label"><span>*</span>所属数据库</label>
                  <div class="col-sm-10">
                    <select id="db_id" name="db_id" class="form-control">
                      {{range $k,$v := .dbconf}}
                        <option value="{{$v.Id}}" {{if eq $.osconf.DbId $v.Id}}selected{{end}}>{{$v.Id}}</option>
                      {{end}}
                    </select>
                  </div>
                </div>
                <div class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label"><span>*</span>主机类型</label>
                  <div class="col-sm-10">
                    <select id="os_type" name="os_type" class="form-control">
                      <option value="">请选择类型</option>
                      <option value="1" {{if eq 1 .osconf.Ostype}}selected{{end}}>Linux</option>
                      <option value="2" {{if eq 2 .osconf.Ostype}}selected{{end}}>Windows</option>
                      <option value="3" {{if eq 3 .osconf.Ostype}}selected{{end}}>AIX</option>
                    </select>
                  </div>
                </div>
                <div class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label"><span>*</span>主机IP</label>
                  <div class="col-sm-10">
                    <input type="text" id="host" name="host"  value="{{.osconf.Host}}" class="form-control">
                  </div>
                </div>
                <div class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label"><span>*</span>协议</label>
                  <div class="col-sm-10">
                    <select id="protocol" name="protocol" class="form-control">
                      <option value="1" {{if eq "1" .osconf.Protocol}}selected{{end}}>ssh</option>
                      <option value="2" {{if eq "2" .osconf.Protocol}}selected{{end}}>snmp</option>
                      <option value="3" {{if eq "3" .osconf.Protocol}}selected{{end}}>telnet</option>
                    </select>
                  </div>
                </div>
                <div class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label"><span>*</span>端口</label>
                  <div class="col-sm-10">
                    <input type="text" id="port" name="port"  value="{{.osconf.Port}}" class="form-control">
                  </div>
                </div>
                <div class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label"><span>*</span>主机名</label>
                  <div class="col-sm-10">
                    <input type="text" id="nodename" name="nodename"  value="{{.osconf.Nodename}}" class="form-control">
                  </div>
                </div>
                <div class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label"><span>*</span>用户名</label>
                  <div class="col-sm-10">
                    <input type="text" id="username" name="username"  value="{{.osconf.Username}}" class="form-control" placeholder="请填写用户名">
                  </div>
                </div>
                <div class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label"><span>*</span>密码</label>
                  <div class="col-sm-10">
                    <input type="password" id="password" name="password"  value="{{.osconf.Password}}" class="form-control" placeholder="请填写密码">
                  </div>
                </div>
                <div class="form-group">
                  <label class="col-lg-2 col-sm-2 control-label"></label>
                  <div class="col-lg-10">
                    <input type="hidden" id="id" name="id" value="{{.osconf.Id}}">
                    <button type="button" onclick="checkConnect()" class="btn btn-primary">连接测试</button>
                    <button type="submit" class="btn btn-primary">提 交</button>
                  </div>
                </div>
              </form>
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
<script src="/static/js/jquery-ui-1.10.3.min.js"></script>
<script>
    $(function() {// 初始化内容
        
    });  

    
    $('#dbnode-form').validate({
        ignore:'',        
		    rules : {
			      username:{required: true},
			      role:{required: true},
        },
        messages : {
			      username:{required: '请填写用户名'},
			      role:{required: '请选择角色'}, 
        },
        submitHandler:function(form) {
            $(form).ajaxSubmit({
                type:'POST',
                dataType:'json',
                success:function(data) {
                    dialogInfo(data.message)
                    if (data.code) {
                       		setTimeout(function(){window.location.href="/config/node/manage"}, 1000);
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
