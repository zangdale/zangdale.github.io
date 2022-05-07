

run:
	hugo -D server

build:
	hugo -D

push:
	mv public/ docs/
	git add -A && git commit -m "add doc" && git push