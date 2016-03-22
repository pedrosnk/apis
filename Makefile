
setup:
	@.virtualenv/bin/pypy setup.py install

run:
	@.virtualenv/bin/pypy api_tor/main.py
