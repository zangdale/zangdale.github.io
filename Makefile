

run:
	hugo -D server

build:
	rm -rf docs/ && hugo -D

push:
	mv public/ docs/
	git add -A && git commit -m "add doc" && git push