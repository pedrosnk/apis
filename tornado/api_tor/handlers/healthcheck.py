# -*- coding: utf-8 -*-

from tornado.web import RequestHandler

class HealthcheckHandler(RequestHandler):
    def get(self):
        self.write({'status': 'Working'})

    def post(self):
        self.write({
            'status': {
                'mongodb': 'ok'
            }
        })
