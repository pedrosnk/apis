# -*- coding: utf-8 -*-

from tornado.web import RequestHandler
from tornado import gen


class HealthcheckHandler(RequestHandler):

    @gen.coroutine
    def get(self):
        mongo_ping = yield self.settings['mongo_storage'].ping()
        self.write({
            'status': {
                'mongodb': mongo_ping,
            }
        })
