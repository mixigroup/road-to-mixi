require 'json'
require 'sinatra'

class App < Sinatra::Base
  set :protection, :except => [:ip_spoofing]

  get '/' do
    "I am live!"
  end

  get '/test' do
    env.to_json
  end
end