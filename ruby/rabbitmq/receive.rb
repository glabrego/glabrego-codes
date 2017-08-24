#!/usr/bin/env ruby
# encoding: utf-8

require 'bunny'

connection = Bunny.new # cria a conexão com localhost:5672 por default ou podemos indicar um host passando o atributo hostname: host
connection.start # inicia a conexão com o host indicado

channel = connection.create_channel
queue = channel.queue('hello')

begin
  puts "|= Waiting for messages. To exit press CTRL-C"
  queue.subscribe(block: true) do |delivery_info, properties, body|
    puts "|< Received #{body}"
  end
rescue Interrupt => _
  connection.close
  exit(0)
end
