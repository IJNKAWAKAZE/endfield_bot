package player

import (
	"endfield_bot/plugins/commandoperation"
)

var (
	playerOperationMap = map[string]commandoperation.OperationI{
		"state": PlayerOperationState{},
		"box":   PlayerOperationBox{},
		"card":  PlayerOperationCard{},
	}
)

func initFactory() {
	for k, f := range playerOperationMap {
		commandoperation.OperationTypeMaps[k] = f
	}
}
func playerOperationFactory(operation string) *commandoperation.OperationI {
	result, ok := playerOperationMap[operation]
	if !ok {
		return nil
	} else {
		return &result
	}
}
