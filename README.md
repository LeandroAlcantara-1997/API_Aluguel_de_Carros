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

## Constrói as imagens da aplicação e sobe os containers
### Necessário esperar até que o MySQL suba para fazer requisições
~~~docker
docker-compose up --build
~~~

