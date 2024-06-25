package cmd

import (
	"encoding/json"
	"os"
	"time"
)

const (
	TriggerDisabled = 1
	TriggerEnabled  = 2
)

type ModeInfo struct {
	Identifier      string `json:"identifier,omitempty"`
	ModeIdentifier  string `json:"modeIdentifier,omitempty"`
	Name            string `json:"name,omitempty"`
	SemanticType    int    `json:"semanticType,omitempty"`
	SymbolImageName string `json:"symbolImageName,omitempty"`
	TintColorName   string `json:"tintColorName,omitempty"`
	Visibility      int    `json:"visibility,omitempty"`
}

type Assertions struct {
	Data []struct {
		StoreAssertionRecords []struct {
			AssertionUUID   string `json:"assertionUUID,omitempty"`
			AssertionSource struct {
				AssertionClientIdentifier string `json:"assertionClientIdentifier,omitempty"`
			} `json:"assertionSource,omitempty"`
			AssertionStartDateTimestamp float64 `json:"assertionStartDateTimestamp,omitempty"`
			AssertionDetails            struct {
				AssertionDetailsIdentifier     string `json:"assertionDetailsIdentifier,omitempty"`
				AssertionDetailsModeIdentifier string `json:"assertionDetailsModeIdentifier,omitempty"`
				AssertionDetailsLifetime       struct {
					AssertionDetailsScheduleLifetimeScheduleIdentifier string `json:"assertionDetailsScheduleLifetimeScheduleIdentifier,omitempty"`
					AssertionDetailsLifetimeType                       string `json:"assertionDetailsLifetimeType,omitempty"`
					AssertionDetailsScheduleLifetimeBehavior           string `json:"assertionDetailsScheduleLifetimeBehavior,omitempty"`
				} `json:"assertionDetailsLifetime,omitempty"`
				AssertionDetailsReason string `json:"assertionDetailsReason,omitempty"`
				Dummy                  struct {
					Dummy string `json:"dummy,omitempty"`
				} `json:"dummy,omitempty"`
			} `json:"assertionDetails,omitempty"`
		} `json:"storeAssertionRecords,omitempty"`
		StoreInvalidationRecords []struct {
			InvalidationAssertion struct {
				AssertionUUID               string          `json:"assertionUUID,omitempty"`
				AssertionSource             AssertionSource `json:"assertionSource,omitempty"`
				AssertionStartDateTimestamp float64         `json:"assertionStartDateTimestamp,omitempty"`
				AssertionDetails            struct {
					AssertionDetailsIdentifier     string `json:"assertionDetailsIdentifier,omitempty"`
					AssertionDetailsModeIdentifier string `json:"assertionDetailsModeIdentifier,omitempty"`
					AssertionDetailsReason         string `json:"assertionDetailsReason,omitempty"`
				} `json:"assertionDetails,omitempty"`
			} `json:"invalidationAssertion,omitempty"`
			InvalidationDateTimestamp float64         `json:"invalidationDateTimestamp,omitempty"`
			InvalidationReason        string          `json:"invalidationReason,omitempty"`
			InvalidationSource        AssertionSource `json:"invalidationSource,omitempty"`
		} `json:"storeInvalidationRecords,omitempty"`
		StoreInvalidationRequestRecords []struct {
			InvalidationRequestDateTimestamp float64 `json:"invalidationRequestDateTimestamp,omitempty"`
			InvalidationRequestPredicate     struct {
				InvalidationPredicateType string `json:"invalidationPredicateType,omitempty"`
			} `json:"invalidationRequestPredicate,omitempty"`
			InvalidationRequestReason string          `json:"invalidationRequestReason,omitempty"`
			InvalidationRequestSource AssertionSource `json:"invalidationRequestSource,omitempty"`
			InvalidationRequestUUID   string          `json:"invalidationRequestUUID,omitempty"`
		} `json:"storeInvalidationRequestRecords,omitempty"`
	} `json:"data,omitempty"`
	Header struct {
		Version   int     `json:"version,omitempty"`
		Timestamp float64 `json:"timestamp,omitempty"`
	} `json:"header,omitempty"`
}

type ModeConfiguration struct {
	AutomaticallyGenerated bool `json:"automaticallyGenerated,omitempty"`
	CompatibilityVersion   int  `json:"compatibilityVersion,omitempty"`
	Configuration          struct {
		ApplicationConfigurationType int `json:"applicationConfigurationType,omitempty"`
		CompatibilityVersion         int `json:"compatibilityVersion,omitempty"`
		HideApplicationBadges        int `json:"hideApplicationBadges,omitempty"`
		MinimumBreakthroughUrgency   int `json:"minimumBreakthroughUrgency,omitempty"`
		SenderConfigurationType      int `json:"senderConfigurationType,omitempty"`
		SuppressionMode              int `json:"suppressionMode,omitempty"`
		SuppressionType              int `json:"suppressionType,omitempty"`
	} `json:"configuration,omitempty"`
	Created                      float64  `json:"created,omitempty"`
	DimsLockScreen               int      `json:"dimsLockScreen,omitempty"`
	HasSecureData                bool     `json:"hasSecureData,omitempty"`
	ImpactsAvailability          int      `json:"impactsAvailability,omitempty"`
	LastModified                 float64  `json:"lastModified,omitempty"`
	LastModifiedByDeviceID       string   `json:"lastModifiedByDeviceID,omitempty"`
	LastModifiedByVersion        string   `json:"lastModifiedByVersion,omitempty"`
	Mode                         ModeInfo `json:"mode,omitempty"`
	ResolvedCompatibilityVersion int      `json:"resolvedCompatibilityVersion,omitempty"`
	Triggers                     struct {
		Triggers []struct {
			Class                     string  `json:"class,omitempty"`
			CreationDate              float64 `json:"creationDate,omitempty"`
			EnabledSetting            int     `json:"enabledSetting,omitempty"`
			TimePeriodEndTimeHour     int     `json:"timePeriodEndTimeHour,omitempty"`
			TimePeriodEndTimeMinute   int     `json:"timePeriodEndTimeMinute,omitempty"`
			TimePeriodStartTimeHour   int     `json:"timePeriodStartTimeHour,omitempty"`
			TimePeriodStartTimeMinute int     `json:"timePeriodStartTimeMinute,omitempty"`
			TimePeriodWeekdays        int     `json:"timePeriodWeekdays,omitempty"` // 7bitを日土金木水火月に準えた2進数で表現
		} `json:"triggers,omitempty"`
	} `json:"triggers,omitempty"`
}

type ModeConfigurations struct {
	Data []struct {
		ModeConfigurations map[string]ModeConfiguration `json:"modeConfigurations,omitempty"`
	} `json:"data,omitempty"`
	Header struct {
		Version   int     `json:"version,omitempty"`
		Timestamp float64 `json:"timestamp,omitempty"`
	} `json:"header,omitempty"`
}

type AssertionSource struct {
	AssertionClientIdentifier       string `json:"assertionClientIdentifier,omitempty"`
	AssertionSourceDeviceIdentifier string `json:"assertionSourceDeviceIdentifier,omitempty"`
}

func readJson(filePath string, configs any) error {
	log.Debugf("reading file: %s", filePath)
	jsonBytes, err := os.ReadFile(filePath)
	if err != nil {
		log.Warnf("failed to read file: %s", err)
		return err
	}
	if json.Unmarshal(jsonBytes, &configs) != nil {
		log.Warnf("failed to unmarshal json file: %s", err)
		return err
	}
	return nil
}

func getTargetBit(weekDay time.Weekday) int {
	if weekDay == time.Sunday {
		return 6
	} else {
		return int(weekDay) - 1
	}
}

func isWeekday(jsonInt int, weekDay time.Weekday) bool {
	targetBit := getTargetBit(weekDay)
	judge := (jsonInt>>targetBit)&1 == 1
	log.Debugf("[isWeekday] bit: %07b, target: %d, judge: %t", jsonInt, targetBit, judge)
	return judge
}
