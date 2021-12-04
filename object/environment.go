package object

type Environment interface {
	Get(name string) (Object, bool)
	Set(name string, val Object) Object
}

type environment struct {
	store    map[string]Object
	builtins map[string]Object
	outer    Environment
}

func NewEnvironment() Environment {
	return &environment{
		store: make(map[string]Object),
		outer: nil,
	}
}

func (e *environment) Get(name string) (Object, bool) {
	obj, exists := e.store[name]
	if !exists && e.outer != nil {
		obj, exists = e.outer.Get(name)
	}
	if builtin, ok := e.builtins[name]; ok {
		return builtin, ok
	}
	return obj, exists
}

func (e *environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}

func NewEnclosedEnvironment(outer Environment) Environment {
	return &environment{
		store: make(map[string]Object),
		outer: outer,
	}
}
