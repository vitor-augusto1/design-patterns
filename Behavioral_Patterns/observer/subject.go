package main

type Subject interface {
  register(observer Observer)
  deRegister(observer Observer)
  notifyAll()
}
