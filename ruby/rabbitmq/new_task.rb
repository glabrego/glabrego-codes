#!/usr/bin/env ruby
# encoding: utf-8

require 'bunny'

connection = Bunny.new
connection.start

channel = connection.create_channel

queue = channel.queue('hello')

message = ARGV.empty? ? "Hello World!" : ARGV.join

queue.publish(message, routing_key: queue.name)

puts "|> Sent '#{message}'"

connection.close
