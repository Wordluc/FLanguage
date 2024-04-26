package Evaluator

import (
	"FLanguage/Parser/Statements"
	"errors"
	"reflect"
)

type Environment struct {
	variables   map[string]iObject
	functions   map[string]Statements.FuncDeclarationStatement
	externals   *Environment
	builtInFunc map[string]builtInFuncObject
	builtInVar  map[string]iObject
}

func NewEnvironment() *Environment {
	return &Environment{
		variables:   make(map[string]iObject),
		functions:   make(map[string]Statements.FuncDeclarationStatement),
		externals:   nil,
		builtInFunc: make(map[string]builtInFuncObject),
		builtInVar:  make(map[string]iObject),
	}
}
func (v *Environment) addVariable(name string, value iObject) error {
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

func (v *Environment) setVariable(name string, value iObject) error {
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

func (v *Environment) getVariable(name string) (iObject, error) {

	variable, exist := v.builtInVar[name]
	if exist {
		return variable, nil
	}
	variable, exist = v.variables[name]
	if exist {
		return variable, nil
	}
	if v.externals == nil {
		return nil, errors.New("variable not defined,name:" + name)
	}
	variable, existEx := v.externals.getVariable(name)
	if existEx == nil {
		return variable, nil
	}

	return nil, errors.New("variable not defined,name:" + name)
}

func (v *Environment) getFunction(name string) (Statements.FuncDeclarationStatement, error) {
	funct, exist := v.functions[name]
	if !exist {
		if v.externals != nil {
			variable, existEx := v.externals.getFunction(name)
			if existEx == nil {
				return variable, nil
			}
		}
	}
	return funct, nil
}

func (v *Environment) addFunction(name string, value Statements.FuncDeclarationStatement) error {
	if _, ok := v.builtInFunc[name]; ok {
		return errors.New("function already exists as builtInFunc:" + name)
	}
	if _, ok := v.functions[name]; ok {
		return errors.New("function already exists:" + name)
	}
	v.functions[name] = value
	return nil

}

func (v *Environment) addBuiltInFunc(name string, value builtInFuncObject) error {
	_, exist := v.builtInFunc[name]
	if exist {
		return errors.New("function already exists:" + name)
	}
	v.builtInFunc[name] = value
	return nil
}

func (v *Environment) addBuiltInVar(name string, value iObject) error {
	if _, ok := v.builtInVar[name]; ok {
		return errors.New("var already exists:" + name)
	}
	v.builtInVar[name] = value
	return nil
}

func (v *Environment) getBuiltInFunc(name string) (builtInFuncObject, error) {
	funct, exist := v.builtInFunc[name]
	if !exist {
		return builtInFuncObject{}, errors.New("function not defined")
	}
	return funct, nil
}
