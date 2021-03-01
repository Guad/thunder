package proto

//go:generate sh -c "docker run -v `pwd`:/defs namely/protoc-all:1.11 -l gogo -d . && mv gen/pb-gogo/github.com/guad/thunder/internal/proto/* . && rm -rf gen"
