# -*- coding: utf-8 -*-

import tornado.ioloop

from handlers.healthcheck import HealthcheckHandler
from handlers.main import MainHandler
from handlers.item_schema import ItemSchemaPostHandler
from storage.mongo import MongoStorage
from storage.redis import RedisStorage


def make_app():
    mongo_db = MongoStorage()
    redis_db = RedisStorage()
    return tornado.web.Application([
        ('/', MainHandler),
        ('/healthcheck', HealthcheckHandler),
        (r'/item-schemas', ItemSchemaPostHandler),
    ],
        mongo_db=mongo_db,
        redis_db=redis_db,
    )


if __name__ == '__main__':
    app = make_app()
    app.listen(8888)
    tornado.ioloop.IOLoop.current().start()
