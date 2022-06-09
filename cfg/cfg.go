package cfg

import (
	"gopkg.in/yaml.v3"
	"os"
)

var C *Cfg = NewCfg()

const (
	PRoot    = "tunnel"
	yamlFile = "tunnel.yaml"
)

type Cfg struct {
	//ClusterName string `yaml:"cluster_name"`
	Code              string `yaml:"code"`
	AliyunAk          string `yaml:"aliyun_ak"`
	AliyunSk          string `yaml:"aliyun_sk"`
	AliyunSlsEndpoint string `yaml:"aliyun_sls_endpoint"`
	AliyunSlsLogstore string `yaml:"aliyun_sls_logstore"`

	ZkAddr    string `yaml:"zk_addr"`
	ZkTimeout int    `yaml:"zk_timeout"`

	Cluster string `yaml:"cluster"`
	Group   string `yaml:"group"`

	TaskLoadType string     `yaml:"task_load_type"`
	Tasks        []*CfgTask `yaml:"tasks"`
}

type CfgTask struct {
	Id       string                 `yaml:"id"`
	Addr     string                 `yaml:"addr"`
	Port     int                    `yaml:"port"`
	Username string                 `yaml:"username"`
	Password string                 `yaml:"password"`
	Schemas  []CfgSchema            `yaml:"schemas"`
	Sink     map[string]interface{} `yaml:"sink"`

	FlushBufferSize int64 `yaml:"flush_buffer_size"`
	FlushInterval   int   `yaml:"flush_interval"`
}

type CfgSchema struct {
	Name   string     `yaml:"name"`
	Tables []CfgTable `yaml:"tables"`
}

type CfgTable struct {
	Name string `yaml:"name"`
}

func NewCfg() *Cfg {
	return loadFromYaml()
}

func loadFromYaml() *Cfg {
	cfg := Cfg{}
	f, _ := os.Open("./" + yamlFile)
	err := yaml.NewDecoder(f).Decode(&cfg)
	if err != nil {
		panic(err)
	}
	return &cfg
}
