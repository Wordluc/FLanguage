package Evaluator

import (
	"FLanguage/Parser/Statements"
	"errors"
	"reflect"
)

type Environment struct {
	variables map[string]IObject
	functions map[string]*Statements.FuncDeclarationStatement
	externals *Environment
	innerFunc map[string]*InnerFuncObject
	innerVar  map[string]*IObject
}

func NewEnvironment() *Environment {
	return &Environment{
		variables: make(map[string]IObject),
		functions: make(map[string]*Statements.FuncDeclarationStatement),
		externals: nil,
		innerFunc: make(map[string]*InnerFuncObject),
		innerVar:  make(map[string]*IObject),
	}
}

func (v *Environment) AddVariable(name string, value IObject) error {
	if _, existInner := v.variables[name]; existInner {
		return errors.New("variable already exists:" + name)
	}
	if _, existInner := v.innerVar[name]; existInner {
		return errors.New("inner variable already exists:" + name)
	}
	v.variables[name] = value
	return nil
}

func (v *Environment) SetVariable(name string, value IObject) error {
	variable, exist := v.variables[name]
	if !exist {
		return errors.New("variable not defined")
	}

	if reflect.TypeOf(variable) != reflect.TypeOf(value) {
		return errors.New("should have same type")
	}
	v.variables[name] = value
	return nil
}

func (v *Environment) GetVariable(name string) (IObject, error) {
	innerVar, exist := v.innerVar[name]
	if exist {
		return *innerVar, nil
	}
	variable, exist := v.variables[name]
	if exist {
		return variable, nil
	}
	variable, existEx := v.externals.GetVariable(name)
	if existEx != nil {
		return nil, errors.New("variable not defined")
	}
	return variable, nil
}

func (v *Environment) GetFunction(name string) (*Statements.FuncDeclarationStatement, error) {
	funct, exist := v.functions[name]
	if !exist {
		funct, existEx := v.externals.GetFunction(name)
		if existEx != nil {
			return nil, errors.New("function not defined")
		}
		return funct, nil
	}
	return funct, nil
}

func (v *Environment) AddFunction(name string, value *Statements.FuncDeclarationStatement) error {
	if _, ok := v.innerFunc[name]; ok {
		return errors.New("function already exists as innerFunc:" + name)
	}
	if _, ok := v.functions[name]; ok {
		return errors.New("function already exists:" + name)
	}
	v.functions[name] = value
	return nil

}

func (v *Environment) AddInnerFunc(name string, value *InnerFuncObject) error {
	_, exist := v.innerFunc[name]
	if exist {
		return errors.New("function already exists:" + name)
	}
	v.innerFunc[name] = value
	return nil
}

func (v *Environment) AddInnerVar(name string, value IObject) error {
	if _, ok := v.innerVar[name]; ok {
		return errors.New("var already exists:" + name)
	}
	v.innerVar[name] = &value
	return nil
}

func (v *Environment) GetInnerFunc(name string) (*InnerFuncObject, error) {
	funct, exist := v.innerFunc[name]
	if !exist {
		return nil, errors.New("function not defined")
	}
	return funct, nil
}
