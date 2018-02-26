# PHP-FPM-CGI

Simple GOLANG wrapper for getting PHP-FPM status page info in JSON format.

This package feature is one useful thing - you can use UNIX or TCP wrapper for connect your status page.  

## Setting UP:

**1) Prepare your PHP server**

Uncomment `pm.status_path = /status` block in your PHP-FPM pool config on your server, like:

````Bash
nano /etc/php/7.2/fpm/pool.d/www.conf
````
And run:

````Bash
service php7.2-fpm restart
````

**2) Download & build package with your OS & ARCH vars:**

````Bash
env GOOS={OS} GOARCH={ARCH} go build -v github.com/globalmac/php-fpm-cgi
````
Where:

{OS} - OS type:

* Mac os - darwin
* Windows - windows
* Linux - linux
* FreeBSD - freebsd

{ARCH} - arhitecture:

* x86_64 - amd64
* x86 - 386
* ARM - arm  (linux only)

Example for Linux x86_64:

```Bash
$ env GOOS=linux GOARCH=amd64 go build -v github.com/globalmac/php-fpm-cgi
```

**3) Switch your connection type UNIX or TCP**

By default in project I use `PHP-FPM via UNIX socket`. If you use `fastcgi_pass 127.0.0.1:9000` you need to set `var viaSocket = false` in `main.go` file. After youn need rebuild your app.
