# API de Aluguel de Carros


Essa API tem a finalidade de simular um sistema de aluguel de veiculos. 

## Tecnologias utilizadas:

* MySQL 5.7
* Golang 1.17
* Docker 20.10.7
* Postman

## Foi utilizado no projeto:

* Arquitetura MVC (Model, View, Controller)
* Persistência de dados com DAO (Data access object) 
* Conteinerização com Docker


## Instruções básicas para executar os containers

### Constrói as imagens da aplicação e sobe os containers
#### * Necessário esperar até que o MySQL suba para fazer requisições
~~~docker
docker-compose up --build
~~~

## Para entrar no bash do serviço 

~~~docker
docker-compose exec <servico> bash
~~~

## Parar os serviços

~~~docker
docker-compose stop
~~~

## Importe os testes feitos no Postman:
~~~postman
https://www.getpostman.com/collections/4109e9233a3533e109e0
~~~

# Autor

## Linkedin
https://www.linkedin.com/in/leandro-alcantara-pro

## Email
leandro1997silva@gmail.com