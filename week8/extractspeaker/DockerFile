FROM alpine


RUN apk update
RUN apk upgrade
RUN apk add --no-cache ffmpeg



WORKDIR /indata


ENTRYPOINT ["/indata/extractspeaker.sh"]