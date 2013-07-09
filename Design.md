# design

A series of notes about how to present the data collected from users.

### Raw Numbers

Page views, visits, unique visitors and returning visitors should be shown as
raw/plain numbers.

### Bar graphs

A bar graph can be a clear way to present data archived over time.

### Expandable Pie Chart

Browser usage could be presented as a pie chart, with all version numbers
included. When a piece of the pie chart is clicked on, the section explands and
divides based on the browser version.

### Storage

A table called `Page` with the following fields:

* Id: _String_ a web site's unique id
* PageView: _Number_ always incrementing
* Visit: _Number_ always incrementing
* UniqueVisitor: _Number_ always incrementing
* ReturnVisitor: _Number_ always incrementing

A table called `Sums` with the following fields:

* Id: _String_ a website's unique id
* Timestamp: _Number_ time of the save in epoch format
* Name: _String_ the field name thats being saved
* Count: _Number_ the sum of elements at the time of saving
* Type: _String_ the type of data being saved, eg: topk, features, resolution, os, view duration, browser usage
