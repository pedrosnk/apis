# -*- coding: utf-8 -*-

import tornado.ioloop
from tornado.web import RequestHandler

from handlers.healthcheck import HealthcheckHandler
from storage.mongo import MongoStorage


class MainHandler(RequestHandler):
    def get(self):
        self.write('Hello To api')


def make_app():
    mongo_storage = MongoStorage()
    return tornado.web.Application([
        ('/', MainHandler),
        ('/healthcheck', HealthcheckHandler),
    ], mongo_storage=mongo_storage)


if __name__ == '__main__':
    app = make_app()
    app.listen(8888)
    tornado.ioloop.IOLoop.current().start()
