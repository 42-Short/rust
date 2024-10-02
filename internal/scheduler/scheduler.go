package scheduler

import (
	"sort"
	"time"

	Module "github.com/42-Short/shortinette/pkg/interfaces/module"
	Short "github.com/42-Short/shortinette/pkg/short"
)

func schedule(short Short.Short, config Short.Config, module Module.Module, startTime time.Time, endTime time.Time) {
	now := time.Now()
	if now.Before(startTime) {
		time.Sleep(time.Until(startTime))
	}

	short.StartModule(module.Name)

	if now.Before(endTime) {
		time.Sleep(time.Until(endTime))
	}

	Short.EndModule(module, config)
}

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
		go schedule(short, *config, module, startTime, startTime.Add(moduleDuration))
		startTime = startTime.Add(moduleDuration)
	}

	select {}
}
