<html>
<head>
<meta http-equiv="content-type" content="text/html;charset=utf-8"/>
<meta name="viewport" content="width=device-width,initial-scale=1.0">
<title>所有记录</title>
<head>
<body>
<h1>所有记录</h1>
{{range .}} <p> {{.Id}} &nbsp; <a target=blank href="{{.Id}}/{{.Id}}.html">{{.Title}}</a> &nbsp; {{.MTime}}</p>{{end}}
</body>
</html>