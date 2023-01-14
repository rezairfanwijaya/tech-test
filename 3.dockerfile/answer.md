#### FROM golang
#### ADD . /go/src/github.com/telkomdev/indihome/backend
#### WORKDIR /go/src/github.com/telkomdev/indihome
#### RUN go get github.com/tools/godep
#### RUN godep restore
#### RUN go install github.com/telkomdev/indihome
#### ENTRYPOINT /go/bin/indihome
#### LISTEN 80

#### Menurut analisa saya, Dockerfile diatas memiliki kesalahan berupa tidak adanya proses build code golang untuk dijadikan sebagai entrypoint. Karena entrypoint akan dipanggil ketika container di run (bukan ketika image dibuild) dan entrypoint harus berisi binary file dari program golang yang sudah dibuild sebelumnya