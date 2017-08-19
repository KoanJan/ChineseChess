package logic

import (
	"encoding/json"
	"errors"
	"fmt"
)

var api map[string]func([]byte) ([]byte, error) = map[string]func([]byte) ([]byte, error){

	GameLogicFuncPlay: func(data []byte) ([]byte, error) {

		form := new(PlayForm)
		if err := json.Unmarshal(data, form); err != nil {
			return nil, err
		}
		return nil, Play(form.X1, form.Y1, form.X2, form.Y2, form.BoardID, form.UserID)
	},
}

// GameLogicFunc returns a logic func if exists,
// or else a func that always returns a error says that the function doesn't exist.
func GameLogicFunc(funcName string) func([]byte) ([]byte, error) {

	if f, existed := api[funcName]; existed {
		return f
	}
	return noSuchFunc(funcName)
}

func noSuchFunc(funcName string) func([]byte) ([]byte, error) {

	return func([]byte) ([]byte, error) {

		return nil, errors.New(fmt.Sprintf("no such func called '%s'", string(funcName)))
	}
}
