# Decorator Pattern Example in Go

Este projeto demonstra a implementação do padrão de projeto Decorator em Go (Golang). O exemplo usa o conceito de um café, onde adicionamos dinamicamente funcionalidades como leite e açúcar.

## Descrição

O padrão Decorator é um padrão de design estrutural que permite a adição de comportamentos a objetos individualmente, em tempo de execução, sem afetar o comportamento de outros objetos da mesma classe. Este projeto implementa o padrão Decorator para decorar um objeto `Coffee` com diferentes ingredientes como `Milk` e `Sugar`.

## Estrutura do Projeto

- `Coffee`: Interface principal que define a estrutura do café.
- `SimpleCoffee`: Implementação concreta da interface `Coffee`.
- `CoffeeDecorator`: Decorator base que implementa a interface `Coffee`.
- `MilkDecorator`: Decorador concreto que adiciona leite ao café.
- `SugarDecorator`: Decorador concreto que adiciona açúcar ao café.

## Como Usar

1. Clone o repositório para sua máquina local.
2. Navegue até a pasta do projeto.
3. Execute o programa com o comando `go run main.go`.

## Exemplo de Saída

- Cost of Simple Coffee: $10.00
- Cost of Milk Coffee: $12.00
- Cost of Sugar Milk Coffee: $13.00