package main

const DEFAULT_VHOST_TEMPLATE = `## Basic reverse proxy server ##
upstream {{.ID}}  {
{{range .Endpoints}}	server {{.}};
{{end}}}
server {
	listen       0.0.0.0:80;
	server_name {{.HostList}}
	access_log  /var/log/nginx/log/access.log  main;
	error_log  /var/log/nginx/log/error.log;
	root   /usr/share/nginx/html;
	index  index.html index.htm;
	location / {
		proxy_pass  http://{{.ID}};
		proxy_next_upstream error timeout invalid_header http_500 http_502 http_503 http_504;
		proxy_redirect off;
		proxy_buffering off;
		proxy_set_header        Host            $host;
		proxy_set_header        X-Real-IP       $remote_addr;
		proxy_set_header        X-Forwarded-For $proxy_add_x_forwarded_for;
	}
}
`
