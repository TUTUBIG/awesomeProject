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
}
