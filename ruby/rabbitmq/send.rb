#!/usr/bin/env ruby
# encoding: utf-8

require 'bunny'

connection = Bunny.new # cria a conexão com localhost:5672 por default ou podemos indicar um host passando o atributo hostname: host
connection.start # inicia a conexão com o host indicado

channel = connection.create_channel # cria um canal, que é por onde a maioria das APIs para se trabalhar fica

queue = channel.queue('hello') # declara a fila de nome 'hello' a qual queremos enviar mensagens

channel.default_exchange.publish('Hello World!', routing_key: queue.name) # envia a mensagem 'Hello World!' para a fila 'hello'

puts "|> Sent 'Hello World!'"

connection.close # fecha a conexão
