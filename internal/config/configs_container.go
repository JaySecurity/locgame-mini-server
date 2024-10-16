package config

import (
	"reflect"

	"locgame-mini-server/pkg/log"
)

type Container struct {
	values map[reflect.Type]reflect.Value
}

func NewConfigsContainer() *Container {
	c := new(Container)
	c.values = make(map[reflect.Type]reflect.Value)
	return c
}

func (c *Container) Register(constructor interface{}) {
	constructValue := reflect.ValueOf(constructor)
	constructType := constructValue.Type()

	if _, exists := c.values[constructType.Out(0)]; exists {
		log.Fatalf("%v already registered", constructType)
	}

	if constructType == nil {
		log.Fatal("can't provide an untyped nil")
	}
	if constructType.Kind() != reflect.Func {
		log.Fatalf("must provide constructor function, got %v (type %v)", constructor, constructType)
	}

	c.values[constructType.Out(0)] = constructValue
}

func (c *Container) Get(configType reflect.Type) reflect.Value {
	funcType := configType
	if funcType == nil {
		log.Fatal("can't invoke an untyped nil")
	}

	value, ok := c.values[funcType]

	if ok {
		values := value.Call(nil)
		return values[0]
	}

	log.Fatal("Can't get config for type:", configType)
	return reflect.Zero(configType)
}
