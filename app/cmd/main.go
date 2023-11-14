package main

import (
	"fmt"

	"github.com/DiLRandI/go-workspace/libraries/libone/calculation"
	"github.com/DiLRandI/go-workspace/libraries/libtwo/logger"
)

func main() {
	logger.SetLevel("debug")

	logger.Debug("Starting the application...")

	logger.Info("Adding two numbers...")
	result := calculation.Add(1, 2)
	logger.Info(fmt.Sprintf("Result: %d", result))

	logger.Info("Subtracting two numbers...")
	result = calculation.Sub(1, 2)
	logger.Info(fmt.Sprintf("Result: %d", result))

	logger.Debug("Terminating the application...")

}
