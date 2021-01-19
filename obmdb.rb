require 'rest-client'
require 'json'

class Execution
  attr_accessor(:request_url, :response, :header, :raw_response)

  def initialize
    api_key = "2a1b7137"
    @request_url = "https://www.omdbapi.com/query_string&apikey=#{api_key}"
    @header = {"content-type" => "charset=utf-8"}
  end

  def find_movie(title)
    @request_url = request_url.gsub("query_string", "?t=" + title)
    @raw_response = RestClient::Request.execute(method: :get, url: request_url, timeout: 10, headers: header)
    @response = JSON.parse(@raw_response)
    if @response['Error'] != nil
      output_hash = {"Error" => @response['Error']}
    else
      output_hash = {"Year" => @response['Year']}
    end
    JSON.generate(output_hash)
  end
end

execution = Execution.new
print execution.find_movie(ARGV[0])