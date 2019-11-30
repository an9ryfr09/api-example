a6-api
=
```
This is a restful style API for a6
```


### Directory

* `ca:` certificate files
* `conf:` app configs
* `handlers:` controllers
* * `v1:` version 1
* * `v2:` version 2
* `middleware:` all middleware
* * `cors:` cors domain control
* * `jwt:` authentication
* * `verification:` params Verification
* `models:` data model
* * `v1:` version 1
* * * `cms:` model cms
* * * `photo:` model photo
* * `v2:` version 2
* `router:` request routers
* `utils:` tools packages
* * `helper:` helper functions
* * `loader:` load configure file
* * `verification:` validate handles
* `vendor:` the third party packages




### Description
The server supported https protocol, Need edit config file: conf/app.yml.

#### step 1
```
$ cd /path/a6-api-golang
$ cp conf/app.yml.example conf/app.yml
$ vim conf/app.yml
```
#### step 2
```
$ go build
$ go install
```

#### step 3
```
$ /path/a6-api-golang
```

