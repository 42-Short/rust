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

	desiredSwitchTime := time.Date(startTime.Year(), startTime.Month(), startTime.Day(), 9, 42, 0, 0, startTime.Location())

	for _, moduleName := range moduleList {
		module := short.Modules[moduleName]

		if err = short.StartModule(moduleName); err != nil {
			return fmt.Errorf("error starting module %s: %v", moduleName, err)
		}

		now := time.Now()
		if now.Before(desiredSwitchTime) {
			time.Sleep(time.Until(desiredSwitchTime))
		} else {
			desiredSwitchTime = desiredSwitchTime.Add(24 * time.Hour)
			time.Sleep(time.Until(desiredSwitchTime))
		}

		logger.Info.Printf("Grading module %s", moduleName)

		if err = Short.EndModule(module, *config); err != nil {
			return fmt.Errorf("error ending module %s: %v", moduleName, err)
		}

		desiredSwitchTime = desiredSwitchTime.Add(24 * time.Hour)
	}

	return nil
}
