package haproxy

import (
	"strings"
	"text/template"
)

func HeaderName(v string) string {
	v = strings.TrimSpace(v)
	if v == "" {
		return ""
	}
	index := strings.Index(v, " ")
	if index < 0 {
		return ""
	}
	return v[:index]
}

func HostName(v string) string {
	v = strings.TrimSpace(v)
	if v == "" {
		return ""
	}
	if strings.HasPrefix(v, "*") {
		return "hdr_end(host) -i " + v[1:]
	}
	return "hdr(host) -i " + v
}

var (
	funcMap = template.FuncMap{
		"header_name": HeaderName,
		"host_name":   HostName,
	}

	haproxyTemplate = template.Must(template.New("haproxy-config").Funcs(funcMap).Parse(`
# HAProxy configuration generated by https://github.com/appscode/voyager
# DO NOT EDIT!

global
	daemon
	stats socket /tmp/haproxy
	server-state-file global
	server-state-base /var/state/haproxy/
	maxconn 4000
	# log using a syslog socket
	log /dev/log local0 info
	log /dev/log local0 notice
	tune.ssl.default-dh-param 2048
	ssl-default-bind-ciphers ECDHE-RSA-AES128-GCM-SHA256:ECDHE-ECDSA-AES128-GCM-SHA256:ECDHE-RSA-AES256-GCM-SHA384:ECDHE-ECDSA-AES256-GCM-SHA384:DHE-RSA-AES128-GCM-SHA256:DHE-DSS-AES128-GCM-SHA256:kEDH+AESGCM:ECDHE-RSA-AES128-SHA256:ECDHE-ECDSA-AES128-SHA256:ECDHE-RSA-AES128-SHA:ECDHE-ECDSA-AES128-SHA:ECDHE-RSA-AES256-SHA384:ECDHE-ECDSA-AES256-SHA384:ECDHE-RSA-AES256-SHA:ECDHE-ECDSA-AES256-SHA:DHE-RSA-AES128-SHA256:DHE-RSA-AES128-SHA:DHE-DSS-AES128-SHA256:DHE-RSA-AES256-SHA256:DHE-DSS-AES256-SHA:DHE-RSA-AES256-SHA:!aNULL:!eNULL:!EXPORT:!DES:!RC4:!3DES:!MD5:!PSK

defaults
	log global

	# https://cbonte.github.io/haproxy-dconv/1.7/configuration.html#4.2-option%20abortonclose
	# https://github.com/appscode/voyager/pull/403
	{{- range $k, $v := .OptionsDefaults }}
	{{ if not $v }}no {{ end }}option {{ $k -}}
	{{ end }}

	# Timeout values
	{{- range $k, $v := .TimeoutDefaults }}
	timeout {{ $k }} {{ $v -}}
	{{ end }}

	# default traffic mode is http
	# mode is overwritten in case of tcp services
	mode http

{{- range $resolver := .DNSResolvers }}
{{ template "dns-resolver" $resolver -}}
{{ end }}

{{- if .Stats }}
{{ template "stats" .Stats -}}
{{ end -}}

{{- range $svc := .HTTPService }}
{{- template "http-frontend" $svc  }}
{{ template "http-backend" $svc  }}
{{ end -}}

{{- range $svc := .TCPService }}
{{ template "tcp-frontend" $svc }}
{{ template "tcp-backend" $svc }}
{{ end -}}

{{- if and (not .HTTPService) .DefaultBackend }}
{{ template "default-frontend" .SharedInfo }}
{{ end }}

{{- if .DefaultBackend }}
{{ template "default-backend" .SharedInfo }}
{{ end }}
`))
	_ = template.Must(haproxyTemplate.New("dns-resolver").Funcs(funcMap).Parse(`
resolvers {{ .Name }}
	{{- range $index, $ns := .NameServer }}
	nameserver dns{{ $index }} {{ $ns -}}
	{{ end }}
	{{ if .Retries -}}
	resolve_retries {{ .Retries -}}
	{{ end -}}
	{{- range $event, $time := .Timeout }}
	timeout {{ $event }} {{ $time -}}
	{{ end -}}
	{{- range $status, $period := .Hold }}
	hold {{ $status }} {{ $period -}}
	{{ end -}}
`))

	_ = template.Must(haproxyTemplate.New("stats").Funcs(funcMap).Parse(`
listen stats
	bind *:{{ .Port }}
	mode http
	stats enable
	stats realm Haproxy\ Statistics
	stats uri /
	{{ if .Username }}stats auth {{ .Username }}:{{ .PassWord }}{{ end -}}
`))

	_ = template.Must(haproxyTemplate.New("default-frontend").Funcs(funcMap).Parse(`
frontend http-frontend
	bind *:80 {{ if .AcceptProxy }}accept-proxy{{ end }}

	mode http
	option httplog
	option forwardfor

	default_backend {{ .DefaultBackend.Name }}
`))

	_ = template.Must(haproxyTemplate.New("default-backend").Funcs(funcMap).Parse(`
backend {{ .DefaultBackend.Name }}
	{{ if .Sticky }}cookie SERVERID insert indirect nocache{{- end -}}
	{{- range $rule := .DefaultBackend.BackendRules }}
	{{ $rule -}}
	{{ end }}

	{{- range $rule := .DefaultBackend.RewriteRules }}
	reqrep {{ $rule -}}
	{{ end }}

	{{- range $index, $rule := .DefaultBackend.HeaderRules }}
	acl ___header_x_{{ $index }}_exists req.hdr({{ $rule | header_name }}) -m found
	http-request add-header {{ $rule }} unless ___header_x_{{ $index }}_exists
	{{- end }}

	{{- range $e := .DefaultBackend.Endpoints }}
	{{- if $e.ExternalName }}
	{{- if $e.UseDNSResolver }}
	server {{ $e.Name }} {{ $e.ExternalName }}:{{ $e.Port -}} {{ if $e.DNSResolver }} {{ if $e.CheckHealth }} check {{ end }} resolvers {{ $e.DNSResolver }} resolve-prefer ipv4 {{ end -}} {{ if $e.TLSOption }} {{ $e.TLSOption }} {{ end -}}
	{{- else if not $.DefaultBackend.BackendRules }}
	acl https ssl_fc
	http-request redirect location https://{{$e.ExternalName}}:{{ $e.Port }} code 301 if https
	http-request redirect location http://{{$e.ExternalName}}:{{ $e.Port }} code 301 unless https
	{{ end -}}
	{{- else }}
	server {{ $e.Name }} {{ $e.IP }}:{{ $e.Port -}} {{ if $e.Weight }} weight {{ $e.Weight }}{{ end -}} {{ if $.Sticky }} cookie {{ $e.Name }}{{ end -}} {{ if $e.TLSOption }} {{ $e.TLSOption }} {{ end -}}
	{{ end -}}
	{{ end -}}
`))

	_ = template.Must(haproxyTemplate.New("http-frontend").Funcs(funcMap).Parse(`
frontend {{ .FrontendName }}
	{{ if .UsesSSL -}}
	bind *:{{ .Port }} {{ if .AcceptProxy }}accept-proxy{{ end }} ssl no-sslv3 no-tlsv10 no-tls-tickets crt /etc/ssl/private/haproxy/ alpn http/1.1
	# Mark all cookies as secure
	rsprep ^Set-Cookie:\ (.*) Set-Cookie:\ \1;\ Secure
	# Add the HSTS header with a 6 month max-age
	rspadd  Strict-Transport-Security:\ max-age=15768000
	{{ else -}}
	bind *:{{ .Port }} {{ if .AcceptProxy }}accept-proxy{{ end }}
	{{ end }}
	mode http
	option httplog
	option forwardfor

	{{- range $path := .Paths }}
	{{ if  and (or (eq $.Port 80) (eq $.Port 443)) (not $.NodePort) }}
	{{ if $path.Host }}acl host_acl_{{ $path.Backend.Name }} {{ $path.Host | host_name }}{{ end }}
	{{- end }}
	{{ if $path.Host }}acl host_acl_{{ $path.Backend.Name }} {{ $path.Host | host_name }}{{ if $.NodePort }}:{{ $.NodePort }}{{ else }}:{{ $.Port }}{{ end }}{{ end }}
	{{ if $path.Path }}acl url_acl_{{ $path.Backend.Name }} path_beg {{ $path.Path }}{{ end }}
	use_backend {{ $path.Backend.Name }} {{ if or $path.Host $path.Path }}if {{ end }}{{ if $path.Host }}host_acl_{{ $path.Backend.Name }}{{ end }}{{ if $path.Path }} url_acl_{{ $path.Backend.Name }}{{ end -}}
	{{ end }}
	{{ if .DefaultBackend }}
	default_backend {{ .DefaultBackend.Name }}
	{{ end -}}
`))

	_ = template.Must(haproxyTemplate.New("http-backend").Funcs(funcMap).Parse(`
{{- range $path := .Paths }}
backend {{ $path.Backend.Name }}
	{{ if $.Sticky }}cookie SERVERID insert indirect nocache{{- end -}}
	{{- range $rule := $path.Backend.BackendRules }}
	{{ $rule -}}
	{{ end }}

	{{- range $rule := $path.Backend.RewriteRules }}
	reqrep {{ $rule -}}
	{{ end }}

	{{- range $index, $rule := $path.Backend.HeaderRules }}
	acl ___header_x_{{ $index }}_exists req.hdr({{ $rule | header_name }}) -m found
	http-request add-header {{ $rule }} unless ___header_x_{{ $index }}_exists
	{{- end }}

	{{- range $e := $path.Backend.Endpoints }}
	{{- if $e.ExternalName }}
	{{- if $e.UseDNSResolver }}
	server {{ $e.Name }} {{ $e.ExternalName }}:{{ $e.Port -}} {{ if $e.DNSResolver }} {{ if $e.CheckHealth }} check {{ end }} resolvers {{ $e.DNSResolver }} resolve-prefer ipv4 {{ end -}} {{ if $e.TLSOption }} {{ $e.TLSOption }} {{ end -}}
	{{- else if not $path.Backend.BackendRules }}
	http-request redirect location {{ if $.UsesSSL }}https://{{ else }}http://{{ end }}{{$e.ExternalName}}:{{ $e.Port }} code 301
	{{- end }}
	{{- else }}
	server {{ $e.Name }} {{ $e.IP }}:{{ $e.Port -}} {{ if $e.Weight }} weight {{ $e.Weight }} {{ end -}} {{ if $.Sticky }} cookie {{ $e.Name }} {{ end -}} {{ if $e.TLSOption }} {{ $e.TLSOption }} {{ end -}}
	{{ end -}}
	{{ end }}
{{ end -}}
`))

	_ = template.Must(haproxyTemplate.New("tcp-frontend").Funcs(funcMap).Parse(`
frontend {{ .FrontendName }}
	bind *:{{ .Port }} {{ if .AcceptProxy }}accept-proxy{{ end }} {{ if .SecretName }}ssl no-sslv3 no-tlsv10 no-tls-tickets crt /etc/ssl/private/haproxy/{{ .SecretName }}.pem{{ end }} {{ if .ALPNOptions }}{{ .ALPNOptions }}{{ end }}
	mode tcp
	default_backend {{ .Backend.Name -}}
`))

	_ = template.Must(haproxyTemplate.New("tcp-backend").Funcs(funcMap).Parse(`
backend {{ .Backend.Name }}
	mode tcp

	{{- range $rule := .Backend.BackendRules }}
	{{ $rule }}
	{{ end -}}

	{{- if $.Sticky }}
	stick-table type ip size 100k expire 30m
	stick on src
	{{ end -}}

	{{- range $e := .Backend.Endpoints }}
	{{- if $e.ExternalName }}
	server {{ $e.Name }} {{ $e.ExternalName }}:{{ $e.Port -}} {{ if $e.DNSResolver }} {{ if $e.CheckHealth }} check{{ end }} resolvers {{ $e.DNSResolver }} resolve-prefer ipv4{{ end -}} {{ if $e.TLSOption }} {{ $e.TLSOption }} {{ end -}}
	{{- else }}
	server {{ $e.Name }} {{ $e.IP }}:{{ $e.Port -}} {{ if $e.Weight }} weight {{ $e.Weight }}{{ end -}} {{ if $e.TLSOption }} {{ $e.TLSOption }} {{ end -}}
	{{ end -}}
	{{ end -}}
`))
)
