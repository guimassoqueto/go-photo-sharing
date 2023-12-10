kill-port:
	sudo kill -9 `sudo lsof -t -i:3000`

dev:
	modd