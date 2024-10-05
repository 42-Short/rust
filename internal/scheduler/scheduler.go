package scheduler

import (
	"fmt"
	"sort"
	"time"

	"github.com/42-Short/shortinette/pkg/logger"
	Short "github.com/42-Short/shortinette/pkg/short"
)

func Schedule(short Short.Short, startTime time.Time, moduleDuration time.Duration) (err error) {
	moduleList := []string{}
	for key := range short.Modules {
		moduleList = append(moduleList, key)
	}
	sort.Strings(moduleList)
	config, err := Short.GetConfig()
	if err != nil {
		return err
	}

	for _, moduleName := range moduleList {
		module := short.Modules[moduleName]
		now := time.Now()
		if now.Before(startTime) {
			time.Sleep(time.Until(startTime))
		}

		logger.Info.Printf("Starting module %s", moduleName)
		
		if err = short.StartModule(moduleName); err != nil {
			return fmt.Errorf("error starting module %s: %v", moduleName, err)
		}

		endTime := startTime.Add(moduleDuration)
		if now.Before(endTime) {
			time.Sleep(time.Until(endTime))
		}

		logger.Info.Printf("Grading module %s", moduleName)
		
		if err = Short.EndModule(module, *config); err != nil {
			return fmt.Errorf("error ending module %s: %v", moduleName, err)
		}

		startTime = endTime
	}

	return nil
}
