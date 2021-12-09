package evaluator

import (
	"fmt"
	"monkeyinterpreter/object"
)

var builtins = map[string]*object.Builtin{
	"len": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			l := len(args)
			if l != 1 {
				return newError("wrong number of arguments. got=%d, want=1", l)
			}
			argument := args[0]
			switch argument.(type) {
			case *object.String:
				s := argument.(*object.String)
				return &object.Integer{Value: int64(len(s.Value))}
			case *object.Array:
				elements := argument.(*object.Array).Elements
				return &object.Integer{Value: int64(len(elements))}
			default:
				return newError("argument to `len` not supported, got %s", argument.Type())
			}
		},
	},

	"first": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if l := len(args); l != 1 {
				return newError("wrong number of arguments. got=%d, want=1", l)
			}
			if typ := args[0].Type(); typ != object.ARRAY_OBJ {
				return newError("argument to `first` must be Array, got %s", typ)
			}

			array := args[0]
			elements := array.(*object.Array).Elements
			size := len(elements)
			if size == 0 {
				return nil
			}
			return elements[0]
		},
	},

	"last": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if l := len(args); l != 1 {
				return newError("wrong number of arguments. got=%d, want=1", l)
			}
			if typ := args[0].Type(); typ != object.ARRAY_OBJ {
				return newError("argument to `last` must be Array, got %s", typ)
			}

			array := args[0]
			elements := array.(*object.Array).Elements
			size := len(elements)
			if size == 0 {
				return nil
			}
			return elements[size-1]
		},
	},

	"rest": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if l := len(args); l != 1 {
				return newError("wrong number of arguments. got=%d, want=2", l)
			}
			if typ := args[0].Type(); typ != object.ARRAY_OBJ {
				return newError("argument to `rest` must be Array, got %s", typ)
			}

			array := args[0]
			elements := array.(*object.Array).Elements
			size := len(elements)
			if size == 0 {
				return nil
			}
			return &object.Array{Elements: elements[1:size]}
		},
	},

	"push": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {

			if l := len(args); l != 2 {
				return newError("wrong number of arguments. got=%d, want=2", l)
			}
			if typ := args[0].Type(); typ != object.ARRAY_OBJ {
				return newError("first argument to `push` must be Array, got %s", typ)
			}

			array := args[0]
			elements := array.(*object.Array).Elements
			elements = append(elements, args[1])
			return &object.Array{Elements: elements}
		},
	},

	"put": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			for _, arg := range args {
				fmt.Println(arg.Inspect())
			}
			return NULL
		},
	},
}
