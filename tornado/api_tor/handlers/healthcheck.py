# -*- coding: utf-8 -*-

from tornado.web import RequestHandler
from tornado import gen


class HealthcheckHandler(RequestHandler):

    @gen.coroutine
    def get(self):
        status = yield [
            self.settings['mongo_db'].ping(),
            self.settings['redis_db'].ping()
        ]
        # stomp_ping = self.settings['stomp_db'].ping()
        self.write({
            'status': {
                'mongodb': status[0],
                'redis': status[1],
                # 'stomp': stomp_ping,
            }
        })
