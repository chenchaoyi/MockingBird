# Initial support for three APIs
# 1) For server to server traking call, between ad-tracking server to ad-network
# 2) Get a list of received event (in JSON array), for automation framework to use
# 3) Reset call to clear tracking history

# For development purpose, run shutgun dummy_server.rb so we don't have to restart
# server everytime we make a change

require 'sinatra'
require 'uri'
require 'json'
require 'pp'

# Example: MdotM request URL
# http://ads.mdotm.com/ads/trackback.php?deviceid=6f1ed70dd7df021ac227fe7cc4d2d9ca405d7579&advid=jennifer.lin%40gree.net&appid=505706305&odin=354289cc87f0b525e7f217ff56a605db63aa48b

# Example: Freakout request URL
# http://conv-us.fout.jp/app/AbCdE123/?ios_udid=b64430510ce2772e212f831b48dce47912cb298a&ios_udid_raw=F595C311-FCE0-4F2D-90E1-3A23A7D4DAC7&ifa=FCE0-4F2D-90E1&timestamp=1354742391580&IpAddress=10.0.2.2

# Example: GAP request URL
# http://gii66.gree-dev.net/md.install.php?appid=31415&deviceid=F595C311-FCE0-4F2D-90E1-3A23A7D4DAC7&ifa=FCE0-4F2D-90E1

# Host address for the dummy server
$dummy_host = "localhost"

get %r{/receive/(.*)} do
    file = File.open("events.txt", 'a')
    pp "url from ad tracking is " + request.url
    if file.nil?
        status 500
        content_type 'application/json'
        body({"error"=>"an error occurs, please try again later"}.to_json)
    else
        ad_network_url = request.url
        if params.has_key?("url")  
            # Valid URL string
            ad_network_url.slice!("http://#{$dummy_host}:4567/receive/?url=")
            response_body = {:error=>0, :description=>"success", :response=>nil}
            response_code = 200
            status 200
            content_type 'application/json'
            body({"error"=>0, "description"=>"success","response"=>nil}.to_json)
        else
           ad_network_url.slice!("http://#{$dummy_host}:4567/receive/")
           response_body = {:error=>1, :description=>"bad URL string", :response=>nil}
           response_code = 400
           status 400
           content_type 'application/json'
           body({"error"=>1, "description"=>"bad URL string", "response"=>nil}.to_json)
        end
        event = {:url => ad_network_url,
                 :timestamp => Time.now.to_s,
                 :response_body => response_body,
                 :response_code => response_code}.to_json
        file.puts(event)
        file.close
    end
end

# Automation framework call to retrieve events (return JSON array)
get '/retrieve' do
    event_array_json = []
    if !File.exist?("events.txt")
        status 404
        body({"error"=>"no data found"}.to_json)
    else
        file = File.open("events.txt",'r')
        while line = file.gets
            event_array_json << line.to_json
        end
        file.close
        status 200
        content_type 'application/json'
        body(event_array_json.to_json)
    end
end

# Reset to clear all tracking history
delete '/' do
    File.delete("events.txt") if File.exist?("events.txt")
    status 200
end
