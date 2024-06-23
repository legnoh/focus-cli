package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var (
	formatJson             bool
	assertions             Assertions
	modeConfigurations     ModeConfigurations
	homePath               = os.Getenv("HOME")
	modeConfigurationsPath = homePath + "/Library/DoNotDisturb/DB/ModeConfigurations.json"
	assertionsPath         = homePath + "/Library/DoNotDisturb/DB/Assertions.json"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get Focus mode",
	Long:  `Get Focus mode via Focus configration files`,
	Run:   getFocusConfig,
}

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.PersistentFlags().BoolVarP(&formatJson, "json", "j", false, "output result by print modeConfig json")
}

func getFocusConfig(cmd *cobra.Command, args []string) {

	var (
		now           = time.Now()
		hour          = now.Hour()
		minute        = now.Minute()
		weekday       = now.Weekday()
		latestStarted = 1440
		latestFocused = ModeInfo{Name: "Off"}
	)

	if err := readJson(modeConfigurationsPath, &modeConfigurations); err != nil {
		log.Fatalf("failed to unmarshal ModeConfigurations.json: %s", err)
	}
	if err := readJson(assertionsPath, &assertions); err != nil {
		log.Fatalf("failed to unmarshal Assertions.json: %s", err)
	}

	// focus set manually
	if len(assertions.Data) > 0 && len(assertions.Data[0].StoreAssertionRecords) > 0 {
		modeIdentifier := assertions.Data[0].StoreAssertionRecords[0].AssertionDetails.AssertionDetailsModeIdentifier
		if modeIdentifier != "" {
			latestFocused = modeConfigurations.Data[0].ModeConfigurations[modeIdentifier].Mode
		} else {
			log.Error("StoreAssertionRecords was existed, but AssertionDetailsModeIdentifier was not found.")
		}
	} else {
		// focus set by schedule trigger
		for _, modeConfig := range modeConfigurations.Data[0].ModeConfigurations {
			for _, t := range modeConfig.Triggers.Triggers {
				nowhm := hour*60 + minute
				log.Debugf("mode: %s, trigger: %+v", modeConfig.Mode.Name, t)

				if t.EnabledSetting == TriggerDisabled {
					continue
				}
				if t.Class != "DNDModeConfigurationScheduleTrigger" {
					continue
				}
				isToday := isWeekday(t.TimePeriodWeekdays, weekday)
				if !isToday {
					continue
				}

				start := t.TimePeriodStartTimeHour*60 + t.TimePeriodStartTimeMinute
				end := t.TimePeriodEndTimeHour*60 + t.TimePeriodEndTimeMinute

				// includes midnight
				if start > end {
					end += 1440
					if nowhm < start {
						nowhm += 1440
					}
				}

				if nowhm >= start && nowhm < end {

					// Multiple focus modes at the same time were First-wins
					// https://discussions.apple.com/thread/253168601?answerId=255941986022&sortBy=best#255941986022
					if latestStarted > start {
						latestStarted = start
						latestFocused = modeConfig.Mode
						log.Debugf("triggered mode: %s(%d)", latestFocused.Name, latestStarted)
						break
					} else {
						log.Debugf("mode not activated: latestStarted(%d) < start(%d)", latestStarted, start)
					}
				} else {
					log.Debugf("mode not activated: nowhm(%d) start(%d) end(%d)", nowhm, start, end)
				}
			}
		}
	}
	if formatJson {
		bytesJson, _ := json.Marshal(latestFocused)
		fmt.Print(string(bytesJson))
	} else {
		fmt.Print(latestFocused.Name)
	}
}
