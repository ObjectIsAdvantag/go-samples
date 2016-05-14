# URL redirector in #golang

Test if cross-scheme redirect can be easily implemented as I found no cloud offering to do that.

Takes an URL as query parameter, and redirects to this URL

Syntax : http://domain:port/?url=XXXXXXX 
=> redirects to XXXXXXXX 

Example : http://localhost:8080/?url=spark://rooms/4131c373-81db-3731-837a-3d1eedcf76a3
=> redirects to spark://rooms/4131c373-81db-3731-837a-3d1eedcf76a3
 
Good to know : this is a ONE hour hack, do not expect too much from this code but to do exactly what it was meant for, 
ie., to test that an HTTP URL shortener could open a Cisco Spark room.


### How to run 

```
> git clone https://github.com/ObjectIsAdvantag/go-samples
> cd go-samples/go-shortener
> go run *.go
2016/05/14 00:24:50 Starting shortener, listening at :8080
```

### How to test

Open a Web client and hit: http://localhost:8080/ping

```
{
  "name": "URL Shortener",
  "version": "v0.1",
  "port": "8080",
  "started": "2016-05-14T00:24:50+02:00"
}
```


### How to invoke

Open a Web client and hit: http://localhost:8080/?url=spark://rooms/4131c373-81db-3731-837a-3d1eedcf76a3

You get redirected to spark://rooms/4131c373-81db-3731-837a-3d1eedcf76a3

You should see similar logs

```
2016/05/14 00:24:58 hit healthcheck endpoint
2016/05/14 00:25:28 Parsing URL :/?url=spark://rooms/4131c373-81db-3731-837a-3d1eedcf76a3
2016/05/14 00:25:28 Parsed
2016/05/14 00:25:28 Url: spark://rooms/4131c373-81db-3731-837a-3d1eedcf76a3
2016/05/14 00:25:28 redirecting to spark://rooms/4131c373-81db-3731-837a-3d1eedcf76a3
```

