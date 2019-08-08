package creator

import "github.com/micro-plat/hydra/conf"

type IPlatBinder interface {
	iplatBinder
	SetDB(*conf.DBConf)
	SetDBByName(name string, d *conf.DBConf)
}

type PlatBinder struct {
	*platBinder
}

func NewPlatBinder(params map[string]string, inputs map[string]*Input) *PlatBinder {
	return &PlatBinder{
		platBinder: newPlatBinder(params, inputs),
	}
}

func (b *PlatBinder) SetDB(d *conf.DBConf) {
	b.platBinder.SetVarConf("db", "db", d)
}

func (b *PlatBinder) SetDBByName(name string, d *conf.DBConf) {
	b.platBinder.SetVarConf("db", name, d)
}
func (b *PlatBinder) SetQueue(d *conf.QueueConf) {
	b.platBinder.SetVarConf("queue", "queue", d)
}
func (b *PlatBinder) SetQueueByName(name string, d *conf.QueueConf) {
	b.platBinder.SetVarConf("queue", name, d)
}
