FROM golang:1.13.8

RUN \
    apt-get update \
      && apt-get install -y --no-install-recommends \
         netcat \
         rsync \
         python3 \
         python3-pip \
         python3-venv \
      && rm -rf /var/lib/apt/lists/*

ENV PYTHON_ENV=/tmp/python-env

RUN pip3 install --upgrade pip
RUN pip3 install --upgrade setuptools
RUN pip3 install --upgrade docker-compose==1.23.2

# Setup work environment
ENV LIBBEAT_PATH /go/src/github.com/elastic/beats/x-pack/libbeat

RUN mkdir -p $LIBBEAT_PATH/build/coverage
WORKDIR $LIBBEAT_PATH
