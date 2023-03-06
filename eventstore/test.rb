require 'net/http'
require 'json'

$eventstore_host = 'localhost'
$eventstore_port = 3000
$stream_name = "stream_1_unique_id"

def append_event(event)
  uri = URI("http://#{$eventstore_host}:#{$eventstore_port}/#{$stream_name}")
  pp uri
  req = Net::HTTP::Post.new(uri, 'Content-Type' => 'application/json')
  req.body = JSON.pretty_generate(event)
  res = Net::HTTP.start(uri.hostname, uri.port) do |http|
    http.request(req)
  end
  if res.is_a?(Net::HTTPSuccess)
  else
    raise "bad: #{res}"
  end
end

event = {
  id: SecureRandom.uuid,
  aggregate_id: SecureRandom.uuid,
  version: "1",
  data: { type: :added, title: "heya", content: "woah", filename: "indeded", "newfield": "test" }
}

append_event(event)
