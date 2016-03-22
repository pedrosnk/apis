# -*- coding: utf-8 -*-

from motor import motor_tornado


class MongoStorage():

    def __init__(self):
        self._client = motor_tornado.MotorClient('localhost', 27017)

    def ping(self):
        return self._client.admin.command('ping')
