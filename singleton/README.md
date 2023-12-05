# Singleton Pattern Example in Go

## Sobre o Projeto
Este repositório contém um exemplo de implementação do padrão de design Singleton em Go. O padrão Singleton é usado para garantir que uma classe tenha apenas uma instância e forneça um ponto de acesso global a ela.

## Funcionalidades
Configuração Única: Garante que apenas uma instância da configuração seja criada.
Concorrência Segura: Utiliza o pacote sync para garantir que a instância única seja segura para acesso concorrente. (*Importante para ambiente multithreaded)

O código define uma estrutura Config com dois campos: DatabaseURL e Port. A função GetConfig é usada para obter a instância singleton da configuração.

## Vantagens

- Você pode ter certeza que uma classe só terá uma única instância.
- Você ganha um ponto de acesso global para aquela instância.
- O objeto singleton é inicializado somente quando for pedido pela primeira vez.

## Desvantagens

- Viola o princípio de responsabilidade única. O padrão resolve dois problemas de uma só vez.
- O padrão Singleton pode mascarar um design ruim, por exemplo, quando os componentes do programa sabem muito sobre cada um.
- O padrão requer tratamento especial em um ambiente multithreaded para que múltiplas threads não possam criar um objeto singleton várias vezes.
- Pode ser difícil realizar testes unitários do código cliente do Singleton porque muitos frameworks de teste dependem de herança quando produzem objetos simulados. Já que o construtor da classe singleton é privado e sobrescrever métodos estáticos é impossível na maioria das linguagem, você terá que pensar em uma maneira criativa de simular o singleton. Ou apenas não escreva os testes. Ou não use o padrão Singleton.

## Executando o Projeto
Com o Go devidamente instalado, clone o repositório e execute o seguinte comando no diretório do projeto:

```bash
go run main.go

