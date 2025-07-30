package main

import (
	"go.uber.org/zap"
	"main/logger"
	"main/shared"
	"time"
)

func main() {
	logger.Info("Starting application", zap.String("version", "1.0.0"))
	
	shared.InitialiseAppSettings(".")
	
	// Example of different log levels and structured logging
	logger.Info("Application initialized successfully")
	
	// Simulate some work with logging examples
	for i := 0; i < 5; i++ {
		logger.Debug("Processing item", 
			zap.Int("item_number", i),
			zap.Duration("processing_time", time.Millisecond*100))
		
		time.Sleep(time.Millisecond * 100)
		
		if i == 3 {
			logger.Warn("Item processing took longer than expected", 
				zap.Int("item_number", i),
				zap.String("reason", "network_delay"))
		}
	}
	
	logger.Info("Application completed successfully", 
		zap.Duration("total_runtime", time.Since(time.Now().Add(-500*time.Millisecond))))
}
