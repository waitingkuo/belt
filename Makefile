all:
	mkdir	-p	bin
	GOOS=linux	go	build	-o	bin/belt-linux
	GOOS=darwin	go	build	-o	bin/belt-darwin
gsutil cp bin/belt-darwin gs://waitingkuo-belt/belt-darwin

upload:
	gsutil cp bin/belt-linux gs://waitingkuo-belt/belt-linux
	gsutil cp bin/belt-darwin gs://waitingkuo-belt/belt-darwin

