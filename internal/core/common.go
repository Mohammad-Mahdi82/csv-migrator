package core

type ILoader func(base int8, data *[][]string, cur int, loader IWorkLoader) error

type IWorkLoader func(records *[][]string, i int) (int, error)
