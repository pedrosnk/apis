# -*- coding: utf-8 -*-

import tornadoredis
from tornado.gen import Task


class RedisStorage():

    def __init__(self):
        self._client = tornadoredis.Client()

    def ping(self):
        return Task(self._client.ping)
