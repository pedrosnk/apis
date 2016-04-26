# -*- coding: utf-8 -*-

import json

from tornado.web import RequestHandler
from jsonschema import Draft3Validator


class ItemSchemaPostHandler(RequestHandler):
    def post(self):
        validator = Draft3Validator(json.loads(self.request.body))
        any_error = False

        for x in validator.iter_errors({}):
            any_error = True
            break

        self.write({"status": "WORKING", "errors": any_error})
