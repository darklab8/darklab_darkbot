from utils.porto import AbstractAction
from .storage import PlayerStorage
from . import schemas as player_schemas
import scrappy.core.settings as settings
from scrappy.commons.subtasks import SubTaskGetItemsData, SubTaskSaveItemsToStorage
from scrappy.core.logger import base_logger

logger = base_logger.getChild(__name__)


class SubTaskGetPlayerData(SubTaskGetItemsData):
    @property
    def _url(self):
        return settings.API_PLAYER_URL


class SubTaskParsePlayers(AbstractAction):
    def __init__(self, data: dict):
        self._data = data

    def run(self) -> list[player_schemas.PlayerIn]:
        players = [
            player_schemas.PlayerIn(**player, timestamp=self._data["timestamp"])
            for player in self._data["players"]
        ]
        logger.debug(f"{self.__class__.__name__} is done")
        return players


class SubTaskSavePlayersToStorage(SubTaskSaveItemsToStorage):
    storage = PlayerStorage

    def __init__(self, players: list[player_schemas.PlayerIn], database):
        super().__init__(items=players, database=database)