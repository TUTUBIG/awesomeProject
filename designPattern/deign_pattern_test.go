package designPattern

import "testing"

func TestStrategyDesignDemo(t *testing.T) {
	t.Run("strategy", func(t *testing.T) {
		StrategyDesignDemo()
	})
	t.Run("factory", func(t *testing.T) {
		FactoryDesignDemo()
	})
	t.Run("abstract factory", func(t *testing.T) {
		AbstractFactoryDesignDemo()
	})
	t.Run("decorator", func(t *testing.T) {
		DecoratorDesignDemo()
	})
	t.Run("observer", func(t *testing.T) {
		ObserverDesignDemo()
	})
	t.Run("singleton", func(t *testing.T) {
		SingletonDesignDemo()
	})
}
