ALL: out/build out/fenci

out/build: build/build.go
	go build -o out/build github.com/jianingy/fenci/build 

out/fenci: fenci/fenci.go
	go build -o out/fenci github.com/jianingy/fenci/fenci


