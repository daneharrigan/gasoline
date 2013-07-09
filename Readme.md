# gasoline

The project tracks the following user data:

* Page Views: when a user reaches a web page
* Visits: when a user reaches a web site. This could mean multiple page views
* Unique Visitors: a user that has reached the web site more than once in a year
* Returning Visitor: a user that has viewed the web site more than once in a day
* TopK: the top 10 most frequented web pages and their view counts
* Features: the count of what users support cookies, quicktime, flash, google talk, java, silverlight or a retina display
* Resolutions: display resolutions used to view the web page and their counts
* OS: operating systems used to view the web page and their counts
* View Durations: the length of time a user spent on a web page broken down by p50, p95 and p99
* Browser Usage: the count of the browsers and their versions used to view the web page

## Endpoints

A list of all endpoints used in the project

```
GET /example
```

This endpoint is an example of how to embed the user tracker JavaScript into a
page.

```
GET /assets/tracker.js
```

This is the JavaScript that gets loaded onto a page to tracker user information.
It sends data back to the `/tracker` endpoint.

```
GET /tracker
```

This endpoint returns headers of `Content-Length: 0` and `Content-Type:
image/gif`. It also records all information sent by `/assets/tracker.js`.

```
GET /dashboard
```

This is a view to all of the aggregated user data. All aggregated data is
delivered through the `/stream` endpoint.

```
GET /stream
```

This endpoint streams data in the form of Server-Sent Events. Aggregated data is
sent once every second.

## Setup

The project uses Go 1.1. At this time there are no other dependencies. All data
is kept in-memory. The project can be started by running: `go run main.go` or
with Foreman:

```bash
$ go build -o bin/gasoline
$ foreman start
```

#### Running on Heroku

GitHub repository access is required for the following steps:

```bash
$ git clone git@github.com:daneharrigan/gasoline.git
$ cd gasoline
$ heroku create --buildpack https://github.com/kr/heroku-buildpack-go.git
$ git push heroku master
```
