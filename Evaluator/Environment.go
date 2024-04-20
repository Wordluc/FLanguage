package Evaluator

import (
	"FLanguage/Parser/Statements"
	"errors"
	"reflect"
)

type Environment struct {
	variables   map[string]IObject
	functions   map[string]*Statements.FuncDeclarationStatement
	externals   *Environment
	builtInFunc map[string]*BuiltInFuncObject
	builtInVar  map[string]IObject
}

func NewEnvironment() *Environment {
	return &Environment{
		variables:   make(map[string]IObject),
		functions:   make(map[string]*Statements.FuncDeclarationStatement),
		externals:   nil,
		builtInFunc: make(map[string]*BuiltInFuncObject),
		builtInVar:  make(map[string]IObject),
	}
}
func (v *Environment) AddVariable(name string, value IObject) error {
	if name == "_" {
		return nil
	}
	if _, exist := v.variables[name]; exist {
		return errors.New("variable already exists:" + name)
	}
	if _, existBuiltIn := v.builtInVar[name]; existBuiltIn {
		return errors.New("builtIn variable already exists:" + name)
	}
	v.variables[name] = value
	return nil
}

func (v *Environment) SetVariable(name string, value IObject) error {
	variable, exist := v.variables[name]
	if !exist {
		return errors.New("variable not defined,name:" + name)
	}

	if reflect.TypeOf(variable) != reflect.TypeOf(value) {
		return errors.New(name + " should have same type")
	}
	v.variables[name] = value
	return nil
}

func (v *Environment) GetVariable(name string) (IObject, error) {

	builtInVar, exist := v.builtInVar[name]
	if exist {
		return builtInVar, nil
	}
	variable, exist := v.variables[name]
	if exist {
		return variable, nil
	}
	if v.externals == nil {
		return nil, errors.New("variable not defined,name:" + name)
	}
	variable, existEx := v.externals.GetVariable(name)
	if existEx == nil {
		return variable, nil
	}
	return nil, errors.New("variable not defined,name:" + name)
}

func (v *Environment) GetFunction(name string) (*Statements.FuncDeclarationStatement, error) {
	funct, exist := v.functions[name]
	if !exist {
		if v.externals == nil {
			return nil, errors.New("variable not defined,name:" + name)
		}
		funct, existEx := v.externals.GetFunction(name)
		if existEx != nil {
			return nil, errors.New("function not defined,name:" + name)
		}
		return funct, nil
	}
	return funct, nil
}

func (v *Environment) AddFunction(name string, value *Statements.FuncDeclarationStatement) error {
	if _, ok := v.builtInFunc[name]; ok {
		return errors.New("function already exists as builtInFunc:" + name)
	}
	if _, ok := v.functions[name]; ok {
		return errors.New("function already exists:" + name)
	}
	v.functions[name] = value
	return nil

}

func (v *Environment) AddBuiltInFunc(name string, value *BuiltInFuncObject) error {
	_, exist := v.builtInFunc[name]
	if exist {
		return errors.New("function already exists:" + name)
	}
	v.builtInFunc[name] = value
	return nil
}

func (v *Environment) AddBuiltInVar(name string, value IObject) error {
	if _, ok := v.builtInVar[name]; ok {
		return errors.New("var already exists:" + name)
	}
	v.builtInVar[name] = value
	return nil
}

func (v *Environment) GetBuiltInFunc(name string) (*BuiltInFuncObject, error) {
	funct, exist := v.builtInFunc[name]
	if !exist {
		return nil, errors.New("function not defined")
	}
	return funct, nil
}
