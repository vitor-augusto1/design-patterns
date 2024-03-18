package main

import (
	"fmt"
	"os"
	"strconv"
)

type Strategy interface {
  executeStrategy(a, b float64) (float64, error)
}

type ConcreteStrategyAdd struct {}
func (ccSA ConcreteStrategyAdd) executeStrategy(a, b float64) (float64, error) {
  return a + b, nil
}

type ConcreteStrategySubtract struct {}
func (ccSS ConcreteStrategySubtract) executeStrategy(a, b float64) (float64, error) {
  return a - b, nil
}

type ConcreteStrategyMultiply struct {}
func (ccSM ConcreteStrategyMultiply) executeStrategy(a, b float64) (float64, error) {
  return a * b, nil
}


type Context struct {
  strategy Strategy
}

func initContext() *Context {
  return &Context{}
}

func (c *Context) setStrategy(stg Strategy) {
  c.strategy = stg
}

func (c *Context) executeStrategy(a, b float64) (float64, error) {
  return c.strategy.executeStrategy(a, b)
}

func main() {
  ctxx := initContext()

  if len(os.Args) == 1 {
    fmt.Print("Please give me more arguments")
    os.Exit(1)
  }

  args := os.Args
  fst, _ := strconv.ParseFloat(args[1], 64)
  sec, _ := strconv.ParseFloat(args[2], 64)
  action := args[3]

  if action == "sum" {
    ctxx.setStrategy(ConcreteStrategyAdd{})
  }

  if action == "sub" {
    ctxx.setStrategy(ConcreteStrategySubtract{})
  }

  if action == "mul" {
    ctxx.setStrategy(ConcreteStrategyMultiply{})
  }

  result, _ := ctxx.executeStrategy(fst, sec)
  fmt.Println("Result :: ", result)
}
