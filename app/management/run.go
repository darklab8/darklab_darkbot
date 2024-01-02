/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package management

import (
	"darkbot/app/configurator"
	"darkbot/app/forumer"
	"darkbot/app/listener"
	"darkbot/app/scrappy"
	"darkbot/app/settings"
	"darkbot/app/settings/darkbot_logus"
	"darkbot/app/viewer"
	"runtime"
	"time"

	"github.com/darklab8/darklab_goutils/goutils/utils"

	"net/http"
	_ "net/http/pprof"

	"github.com/pkg/profile"
	"github.com/spf13/cobra"
)

// runCmd represents the check command
var runCmd = &cobra.Command{
	Use: "run",
	Run: func(cmd *cobra.Command, args []string) {
		darkbot_logus.Log.Info("run called")

		// migrate db
		configurator.NewConfigurator(settings.Dbpath).AutoMigrateSchema()
		forumenacer := forumer.NewForumer(settings.Dbpath)

		var scrappy_storage *scrappy.ScrappyStorage
		if settings.Config.DevEnvMockApi == "true" {
			scrappy_storage = scrappy.FixtureMockedStorage()
		} else {
			scrappy_storage = scrappy.NewScrappyWithApis()
		}

		scrappy_storage.Update()
		go scrappy_storage.Run()
		go listener.Run()
		go viewer.NewViewer(settings.Dbpath, scrappy_storage).Run()
		go forumenacer.Run()

		// profiler
		if settings.Config.ProfilingEnabled == settings.EnvTrue {
			p := profile.Start(profile.MemProfile, profile.ProfilePath("."), profile.NoShutdownHook)
			defer p.Stop()

			go func() {
				http.ListenAndServe(":8080", nil)
			}()
		}

		// May be garbage collector will help
		go func() {
			for {
				time.Sleep(time.Second * 10)
				runtime.GC()
			}
		}()

		utils.SleepAwaitCtrlC()
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
