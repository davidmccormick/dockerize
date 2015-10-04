template
=============

Cut down version of Jason WIlder's 'dockerize' go program that only includes the template writing functionality.
See [A Simple Way To Dockerize Applications](http://jasonwilder.com/blog/2014/10/13/a-simple-way-to-template-applications/)
Please explore dockerize, and only use this version if you decide, like me, that you want simpler functionality.

The point is not to force the running of an executable and therefore all of exec functionality and log tailing
has been removed.


The typical use case for template is when you have an application that has one or more
configuration files and you would like to control some of the values using environment variables.

For example, a Python application using Sqlalchemy may be able to use environment variables directly.
It may require that the database URL be read from a python settings file with a variable named
`SQLALCHEMY_DATABASE_URI`.  template allows you to set an environment variable such as
`DATABASE_URL` and update the python file when the container starts.

###Usage

It is intended that this be run from a shell script setting up a container before executing the main service.

template /etc/nginx/nginx.tmpl:/etc/nginx/nginx.conf /data/anotherfile.tmpl:/data/anotherfile

## Using Templates

Templates use Golang [text/template](http://golang.org/pkg/text/template/). You can access environment
variables within a template with `.Env`.

```
{{ .Env.PATH }} is my path
```

There are a few built in functions from the original 'dockerize':

  * `default` - Returns a default value for one that does not exist
  * `contains` - Returns true if a string is within another string
  * `exists` - Determines if a file path exists or not
  * `split` - Splits a string into an array using a separator string
  * `replace` - Replaces all occurences of a string within another string
  * `parseUrl`- Parses a URL into it's protocol, scheme, host, etc. parts.

## License

MIT
