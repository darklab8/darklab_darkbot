from fastapi import APIRouter
from fastapi import Query

from fastapi import Depends
from scrappy.core.databases import DatabaseFactory, Database
from . import actions as base_actions
from .schemas import BaseQueryParams

router = APIRouter(
    prefix="/bases",
    tags=["items"],
)

query_default_values = BaseQueryParams()


@router.get("/")
async def get_bases(
    database: Database = Depends(DatabaseFactory.get_default_database),
    page: int = Query(default=query_default_values.page),
    name: list[str] = Query(default=query_default_values.name_tags),
):

    bases = base_actions.ActionGetFilteredBases(
        database=database,
        page=page,
        name_tags=name,
    )
    return bases