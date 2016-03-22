# -*- coding: utf-8 -*-

from torstomp import TorStomp


class StompStorage():

    def __init__(self):
        self._client = TorStomp('localhost', 61613, connect_headers={
            'heart-beat': '1000,1000'
        })
        self._status = 'WAITING'
        self._client.connect()

    def ping(self):
        return self._client.connected
