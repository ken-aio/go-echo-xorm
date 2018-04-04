# go-echo-sqlboiler
Golang sample project using Echo and XORM.  

## WebFW
### Echo
https://github.com/labstack/echo

## ORM
### XORM
https://github.com/go-xorm/xorm
```
# generate new orm codes
$ make gen

# create new db schema rs ... replace schema
$ make rs
```

## Application Config
### Viper
https://github.com/spf13/viper

see: config dir

## DB Migration
### DBFlute
https://github.com/dbflute/dbflute-core

## Swagger
https://github.com/swaggo/swag  
https://github.com/swaggo/echo-swagger

## Hot reload
https://github.com/codegangsta/gin
```
# run echo web server via gin
$ make run
$ curl http://localhost:1314
```

## Test watch
Use looper for test watching.  
looper  
https://github.com/nathany/looper  

Starting test watching when you entered blow command.
```
$ make watch
```
