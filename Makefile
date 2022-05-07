

run:
	hugo -D server

build:
	hugo -D

push:
	git add -A && git commit -m "add doc" && git push