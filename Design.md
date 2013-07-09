# design

A series of notes about how to present the data collected from users.

## Raw Numbers

Page views, visits, unique visitors and returning visitors should be shown as
raw/plain numbers.

## Bar graphs

A bar graph can be a clear way to present data archived over time.

## Expandable Pie Chart

Browser usage could be presented as a pie chart, with all version numbers
included. When a piece of the pie chart is clicked on, the section explands and
divides based on the browser version.

## Storage

A table called `Page` with the following fields:

* Id: __String__ a web site's unique id
* PageView: __Number__ always incrementing
* Visit: __Number__ always incrementing
* UniqueVisitor: __Number__ always incrementing
* ReturnVisitor: __Number__ always incrementing

A table called `Sums` with the following fields:

* Id: __String__ a website's unique id
* Timestamp: __Number__ time of the save in epoch format
* Name: __String__ the field name thats being saved
* Count: __Number__ the sum of elements at the time of saving
* Type: __String__ the type of data being saved, eg: topk, features, resolution, os, view duration, browser usage
