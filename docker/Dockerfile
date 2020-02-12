FROM golang:1.13
WORKDIR /var/app
COPY ./start.server.sh /bin/start.server.sh
RUN go get -u github.com/lbernardo/lambda-local github.com/ghodss/yaml github.com/spf13/cobra github.com/docker/docker/client github.com/lbernardo/gomon
RUN go install github.com/lbernardo/lambda-local
RUN go install github.com/lbernardo/gomon
RUN curl -o- -L https://slss.io/install | bash
ENTRYPOINT [ "sh","/bin/start.server.sh" ]