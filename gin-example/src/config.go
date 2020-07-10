//{
//    "config_version" : 1.0,
//    "elastic_addr" : "http:172.17.0.xx:9000",
//    "readTimeout" : 500,
//    "writeTimeout" : 500,
//    "plugins":[
//        {
//            "name" : "plugin_a",
//            "enable" : true,
//            "port" : 8000,
//
//		}
//	]
//}

package config

type SystemConfig struct {
	ElasticAddr  string         `json:"elastic_addr"`
	ReadTimeout  int            `json:"readTimeout"`
	WriteTimeout int            `json:"writeTimeout"`
	Plugins      []PluginConfig `json:"plugins"`
}

type PluginConfig struct {
	Name         string `json:"name"`
	Enable       bool   `json:"enable"`
	Port         int    `json:"port"`
	ReadTimeout  int    `json:"readTimeout"`
	WriteTimeout int    `json:"writeTimeout"`
}
