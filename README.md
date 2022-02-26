# API de Aluguel de Carros

## Tecnologias utilizadas no Backend

* MySQL 5.7
* Golang 1.17
* Docker 20.10.7
* Postman

## Executando build com CompileDaemon

~~~Linux
$HOME/go/bin/CompileDaemon -command="./LeandroAlcantara-1997"
~~~

~~~Windows
C:\Users\leand\go\bin\CompileDaemon -command="./LeandroAlcantara-1997"
~~~

## Foi utilizado no projeto:

* Arquitetura MVC (Model, View, Controller)
* Persistência de dados com DAO (Data access object) 
* Conteinerização com Docker


# Instruções básicas para executar os containers

## Constrói as imagens da aplicação e do MySQL
~~~docker
docker-compose up -d --build
~~~

## Necessário acessar o bash da aplicação
~~~docker 
docker-compose exec app bash
~~~

## Executar dentro do bash o comando para rodar a aplicação
~~~golang
go run main.go
~~~