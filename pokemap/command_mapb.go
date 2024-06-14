package pokemap

import (
	"errors"
	"github.com/ds1242/pokedex-go/config"
)

func commandMapB(conf *config.Config) error{
	if conf.PreviousURL == "" {
		return errors.New("no more previous data to query, please use map to check forward")
	}
	GetLocation(conf.PreviousURL, conf)
	return nil
}