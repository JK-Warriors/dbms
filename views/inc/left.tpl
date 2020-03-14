<!-- left side start-->
<div class="left-side sticky-left-side">
  <!--logo and iconic logo start-->
  <div class="logo"> <a href="/"><img src="/static/img/logo-left.png" alt="OPMS管理系统"></a> </div>
  <div class="logo-icon text-center"> <a href="/"><img src="/static/img/logo_icon.png" style="width:40px;" alt="OPMS管理系统"></a> </div>
  <!--logo and iconic logo end-->
  <div class="left-side-inner">
    <!-- visible to small devices only -->
    <div class="visible-xs hidden-sm hidden-md hidden-lg">
      <div class="media logged-user"> <img alt="{{.LoginUsername}}" src="{{getAvatar .LoginAvatar}}" class="media-object">
        <div class="media-body">
          <h4><a href="/user/show/{{.LoginUserid}}">{{.LoginUsername}}</a></h4>
          <span>OPMS系统</span> </div>
      </div>
      <h5 class="left-nav-title">控制台</h5>
      <ul class="nav nav-pills nav-stacked custom-nav">
        <li><a href="/user/profile"><i class="fa fa-user"></i> <span>个人设置</span></a></li>
        <li><a href="/logout"><i class="fa fa-sign-out"></i> <span>退出</span></a></li>
      </ul>
    </div>
    <!--sidebar nav start-->
    <!--<ul class="nav nav-pills nav-stacked custom-nav js-left-nav">-->
    <ul class="nav nav-pills nav-stacked custom-nav js-left-nav"  lay-filter="nav-side">
	<!--li><a href="/user/show/{{.LoginUserid}}"><i class="fa fa-home"></i> <span>我的主页</span></a></li-->


   
	<li><a href="/"><i class="fa fa-home"></i> <span>首页</span></a></li> 
  {{range $index, $elem := .leftNav}}
    {{if eq 0 $elem.Parentid}}
      <li class="layui-nav-item">
          <a class="" href="javascript:;">
              <i class="fa {{$elem.Icon}}"></i> <span>{{$elem.Name}}</span>
          </a>
          <dl class="layui-nav-child pp-nav-childs">
          {{range $i, $e := $.leftNav}}
              {{if and (eq $e.Parentid $elem.Id) (eq 1 $e.IsShow)}}
                <dd>
                  <a href="{{$e.Url}}" data-icon="{{$e.Icon}}" data-title="{{$e.Name}}"  class="pointer" data-id="{{$e.Id}}">
                    <i class="fa {{$e.Icon}}"></i><span>{{$e.Name}}</span>
                  </a>
                </dd>
              {{end}}
          {{end}}
          </dl>
      </li>
    {{end}}
  {{end}}
  </ul>
	 
    <!--sidebar nav end-->
  </div>
</div>
<!-- left side end-->


<script src="/static/js/layui/layui.js"></script>
<script>
    //JavaScript代码区域
    layui.use(['element','jquery','layer'], function(){

        //tips
        $(".pp-nav-childs").find('a').hover(function(){
            layer.tips($(this).attr('data-title'), $(this),{time:1000});
        });
    });
</script>
