package dimensions

import (
	"reflect"
)

type Setter interface {
	Set(...int) error
}

type Volumer interface {
	Volume() int
}

type Masser interface {
	Mass(int) int
}

type Dimensioner interface {
	Setter
	Volumer
	Masser
}

type Operator struct {
	worker Dimensioner
}

var allovedDimensions map[string]Dimensioner

func Init(dimensions map[string]Dimensioner) {
	allovedDimensions = dimensions
}

func New(name string) Operator {
	return Operator{
		worker: reflect.New(reflect.TypeOf(allovedDimensions[name]).Elem()).Interface().(Dimensioner),
	}
}

func (o *Operator) Set(d ...int) error {
	return o.worker.Set(d...)
}

func (o Operator) Volume() int {
	return o.worker.Volume()
}

func (o Operator) Mass(ro int) int {
	return o.worker.Mass(ro)
}
