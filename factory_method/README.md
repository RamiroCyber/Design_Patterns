<<<<<<< HEAD:factory_method/README.md
# Factory Method Pattern Example in Go
=======
# Go Factory method notifier
>>>>>>> 0286e25170a50bcb15a9577b4ad741912499db6b:README.md

## Sobre o Projeto
Este projeto em Go exemplifica a implementação do padrão de projeto Factory Method para construir um sistema de notificação. O Factory Method é um padrão de design de criação que oferece uma interface para criar objetos em uma superclasse, permitindo que subclasses modifiquem o tipo de objetos que serão criados.

## Estrutura do Projeto
O sistema é composto por uma interface `Notifier` e suas implementações concretas `EmailNotifier` e `SmsNotifier`, representando diferentes mecanismos de entrega de notificação.

### Interface Notifier
A interface `Notifier` define o método `Send`, que deve ser implementado por todos os notificadores concretos para enviar mensagens.

### Implementações Concretas
- `EmailNotifier`: Implementa o método `Send`, simulando o envio de uma mensagem por email.
- `SmsNotifier`: Similar ao `EmailNotifier`, mas para envio de mensagens SMS.

### Factory Interface e Classes
- `NotifierFactory`: Interface que declara o método `CreateNotifier`.
- `EmailNotifierFactory`: Cria uma instância de `EmailNotifier`.
- `SmsNotifierFactory`: Cria uma instância de `SmsNotifier`.

### Funções Factory
As funções `CreateNotifierEmail` e `CreateNotifierSms` instanciam notificadores usando as respectivas fábricas.

## Como Usar
O ponto de entrada principal (`main`) não instancia diretamente os notificadores. Em vez disso, utiliza as funções `CreateNotifierEmail` e `CreateNotifierSms` para abstrair a criação dos objetos notificadores.

## Vantagens
- **Desacoplamento**: O código do cliente opera com uma interface abstrata, promovendo a flexibilidade e reduzindo o acoplamento.
- **Extensibilidade**: Novos tipos de notificação podem ser adicionados com mínimas ou nenhuma alteração no código cliente.
- **Responsabilidade Única**: Cada fábrica é responsável por criar um único tipo de notificador.

## Executando o Projeto
Com o Go devidamente instalado, clone o repositório e execute o seguinte comando no diretório do projeto:

```bash
go run main.go
