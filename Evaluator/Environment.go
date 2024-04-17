package Evaluator

import (
	"FLanguage/Parser/Statements"
	"errors"
	"reflect"
)

type Environment struct {
	variables map[string]IObject
	functions map[string]*Statements.FuncDeclarationStatement
	innerFunc map[string]InnerFuncObject
	externals *Environment
}

func NewEnvironment() *Environment {
	return &Environment{
		variables: make(map[string]IObject),
		functions: make(map[string]*Statements.FuncDeclarationStatement),
		innerFunc: make(map[string]InnerFuncObject),
		externals: nil,
	}
}
func (v *Environment) AddVariable(name string, value IObject) error {
	_, exist := v.variables[name]
	if exist {
		return errors.New("variable already exists:" + name)
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
	variable, exist := v.variables[name]
	if !exist {
		variable, existEx := v.externals.GetVariable(name)
		if existEx != nil {
			return nil, errors.New("variable not defined")
		}
		return variable, nil
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

func (v *Environment) SetFunction(name string, value *Statements.FuncDeclarationStatement) error {
	_, exist := v.functions[name]
	if exist {
		return errors.New("function already exists:" + name)
	}
	v.functions[name] = value
	return nil

}

func (v *Environment) SetInnerFunc(name string, value InnerFuncObject) error {
	_, exist := v.innerFunc[name]
	if exist {
		return errors.New("function already exists:" + name)
	}
	v.innerFunc[name] = value
	return nil
}

func (v *Environment) GetInnerFunc(name string) (*InnerFuncObject, error) {
	funct, exist := v.innerFunc[name]
	if !exist {
		return nil, errors.New("function not defined")
	}
	return &funct, nil
}
