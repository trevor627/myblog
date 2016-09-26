{{define "navbar"}}

<a class="navbar-brand" href="/">我的博客</a>
	<form class="navbar-form navbar-left" role="search">
		<div class="form-group">
			<input type="text" class="form-control" placeholder="Search">
		</div>
		<button type="submit" class="btn btn-default">Search</button>
	</form>
					
	<ul class="nav navbar-nav" role="tablist">
		<li {{if .IsHome}}class="active"{{end}}><a href="/">首页</a></li>
		<li {{if .IsCategory}}class="active"{{end}}><a href="/category">分类</a></li>
		<li {{if .IsTopic}}class="active"{{end}}><a href="/topic">文章</a></li>
	</ul>
					
	<form class="navbar-form navbar-right" role="tablist">
		{{if .IsLogin}}
			<li><a href="/login?exit=true" class="btn btn-primary active" role="button">logout</a></li>
		{{else}}
			<li><a href="/login" class="btn btn-primary active" role="button">login</a></li>
		{{end}}
	</form>

{{end}}