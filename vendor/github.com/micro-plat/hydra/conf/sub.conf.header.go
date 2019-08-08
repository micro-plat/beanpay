package conf

import "strings"

//Headers http头信息
type Headers map[string]string

//CrossDomainHeader 跨域配置
type CrossDomainHeader Headers

//NewHeader 构建http头信息
func NewHeader(kv ...string) Headers {
	h := map[string]string{}
	l := len(kv)
	for i := 0; i < len(kv)/2 && i < l-1; i++ {
		h[kv[i*2]] = kv[i*2+1]
	}
	return h
}

//WithCrossDomain 添加跨域头配置
func (imp Headers) WithCrossDomain(host ...string) Headers {
	imp["Access-Control-Allow-Origin"] = "*"
	if len(host) > 0 {
		imp["Access-Control-Allow-Origin"] = strings.Join(host, ",")
	}
	imp["Access-Control-Allow-Methods"] = "GET,POST,PUT,DELETE,PATCH,OPTIONS"
	imp["Access-Control-Allow-Credentials"] = "true"
	return imp
}
