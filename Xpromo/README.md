## Intro ##

This is a 3rd party dummy ad network server that aggregates install events to each of the real ad network. It will expose a set of APIs for Xpromo Automation Framework to verify install events for end to end scenario

## Prerequisites ##

* Ruby version >= 1.9.3
* RubyGems version >= 1.3.6
* Install [Bundle](http://gembundler.com/)
    * (sudo) gem install bundler
* Install all required gems
    * bundle install

## Usage ##

* To start the server: ruby dummy_server.rb 

## API ##

Currently, there are three APIs supported

* GET http://localhost:4567/receive/?url=http://ad.example.com
  * this is what the ad-tracking server will call to dummy server to pass the url information

* GET http://localhost:4567/retrieve
  * this will return an array of JSON where each network event will look like this:
     {"url" => http://ad.example.com,
      "timestamp" => "2012-12-10 16:26:33 -0800",
      "response code" => 200,
      "response_body" => {"error":0,"description":"success","response":null}
      }

* DELETE http://localhost:4567/
  * remove all tracking history 
