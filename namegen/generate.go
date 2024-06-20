package namegen

import (
	"math/rand"

	"github.com/mtariq99/namegen/consts"
)

type nameGenImpl interface {
	GetUniqueName() string
}

type nameGen struct {
	random *rand.Rand
}

func (ng *nameGen) GetUniqueName() string {
	listName := consts.ListNamesArr()
	randomIndex := ng.random.Intn(consts.LastIndex-consts.StartIndex+1) + consts.StartIndex
	return listName[randomIndex]
}

func NewGenerator(seed int64) nameGenImpl {
	ng := &nameGen{
		random: rand.New(rand.NewSource(seed)),
	}
	return ng
}
