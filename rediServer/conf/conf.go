//conf包提供全局配置
package conf

const (
	//服务监听的IP+Port
	ADDRESS = "0.0.0.0:6399"
	//持久化数据库名称
	REDI_FLIE_NAME = "dump.rdb"
	//定时保存的时间间隔，单位：s
	DURATION = 1
)