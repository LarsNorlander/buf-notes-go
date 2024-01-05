buf-generate:
	rm -rf gen/
	buf generate https://github.com/LarsNorlander/buf-notes-api.git
