FROM golang

ARG app_env
ENV APP_ENV $app_env

COPY ./app /go/src/github.com/robrotheram/baldrick_engine/app
WORKDIR /go/src/github.com/robrotheram/baldrick_engine/app

RUN go get ./
RUN go build
RUN go get -u github.com/go-swagger/go-swagger/cmd/swagger


CMD if [ ${APP_ENV} = production ]; \
	then \
	app; \
	else \
	swagger generate spec -o ./static/swagger.json; \
	go get github.com/pilu/fresh && app \
	fresh; \
	fi

EXPOSE 8080