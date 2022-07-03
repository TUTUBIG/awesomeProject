package designPattern

import (
	"fmt"
	"math/rand"
	"time"
)

type observer interface {
	update(price float32)
}

type Apple struct{}

func (a Apple) update(price float32) {
	fmt.Println("apple price: ", "$", price)
}

func NewApple() *Apple {
	return &Apple{}
}

type Tesla struct{}

func NewTesla() *Tesla {
	return &Tesla{}
}

func (t Tesla) update(price float32) {
	fmt.Println("tesla price: ", "$", price)

}

type StockMarket struct {
	observers []observer
}

func NewStockMarket() *StockMarket {
	return &StockMarket{observers: make([]observer, 0)}
}

func (sm *StockMarket) RegisterObserver(o ...observer) {
	sm.observers = append(sm.observers, o...)
}

func (sm *StockMarket) Run() {
	for i := 0; i <= 5; i++ {
		time.Sleep(time.Second)
		for i := range sm.observers {
			sm.observers[i].update(rand.Float32())
		}
	}
}

func ObserverDesignDemo() {
	stockMarket := NewStockMarket()
	appleStock, teslaStock := NewApple(), NewTesla()

	stockMarket.RegisterObserver(appleStock, teslaStock)
	stockMarket.Run()
}
