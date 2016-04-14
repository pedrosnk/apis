# -*- coding: utf-8 -*-

from setuptools import setup

setup(
    name='api_tor',
    install_requires=[
        "tornado>=4.3.0,<5.0.0",
        "motor>=0.6.2",
        "tornado-redis>=2.4.18",
        "torstomp>=0.1.7",
        "jsonschema>=2.5.1",
    ],
)
