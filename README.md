## Middleware

Client <-> Python-RSocket <-> NATS-Stream <-> Golang-CQRS

## Python Enviroment

virtualenv --python=/Users/tjikaljedy/.pyenv/versions/3.9.1/bin/python ~/melody-io-env
source ~/melody-io-env/bin/activate

## Python RSocket

The main reason not using golang in middleware, golang version RSocket still far from perfect

#### Module

pip install rsocket
pip install pyjwt
pip install nats-py
pip install transitions
