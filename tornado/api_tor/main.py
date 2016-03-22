# -*- coding: utf-8 -*-

import tornado.ioloop
from tornado.web import RequestHandler

from handlers.healthcheck import HealthcheckHandler
from storage.mongo import MongoStorage
from storage.redis import RedisStorage


class MainHandler(RequestHandler):
    def get(self):
        self.write('Hello To api')


def make_app():
    mongo_db = MongoStorage()
    redis_db = RedisStorage()
    return tornado.web.Application([
        ('/', MainHandler),
        ('/healthcheck', HealthcheckHandler),
    ],
        mongo_db=mongo_db,
        redis_db=redis_db,
    )


if __name__ == '__main__':
    app = make_app()
    app.listen(8888)
    tornado.ioloop.IOLoop.current().start()
