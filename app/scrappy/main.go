package scrappy

import (
	"darkbot/app/scrappy/base"
	"darkbot/app/scrappy/baseattack"
	"darkbot/app/scrappy/player"
	"darkbot/app/settings"
	"darkbot/app/settings/logus"
	"time"
)

type ScrappyStorage struct {
	baseStorage       *base.BaseStorage
	playerStorage     *player.PlayerStorage
	baseAttackStorage *baseattack.BaseAttackStorage
}

func NewScrapyStorage(base_api base.IbaseAPI, player_api player.IPlayerAPI, base_attack baseattack.IbaseAttackAPI) *ScrappyStorage {
	s := &ScrappyStorage{}
	s.baseStorage = base.NewBaseStorage(base_api)
	s.playerStorage = player.NewPlayerStorage(player_api)
	s.baseAttackStorage = baseattack.NewBaseAttackStorage(base_attack)
	return s
}

func (s *ScrappyStorage) Update() {
	s.baseStorage.Update()
	s.playerStorage.Update()
	s.baseAttackStorage.Update()
}

func (s *ScrappyStorage) GetBaseStorage() *base.BaseStorage {
	return s.baseStorage
}
func (s *ScrappyStorage) GetPlayerStorage() *player.PlayerStorage {
	return s.playerStorage
}

func (s *ScrappyStorage) GetBaseAttackStorage() *baseattack.BaseAttackStorage {
	return s.baseAttackStorage
}

var Storage *ScrappyStorage

func init() {
	logus.Info("initialized scrappy")
	Storage = NewScrapyStorage(base.NewBaseApi(), player.NewPlayerAPI(), baseattack.NewBaseAttackAPI())
}

func Run() {
	logus.Info("starting scrappy infinity update loop")
	for {
		Storage.Update()
		time.Sleep(time.Duration(settings.LoopDelay) * time.Second)
	}
	logus.Info("gracefully shutdown scrappy infinity loop")
}