# WalletCore

- [WalletCore](#walletcore)
  - [Descrição](#descrição)
  - [Componentes](#componentes)
    - [Serviços](#serviços)
      - [wallet-core](#wallet-core)
      - [balance-app](#balance-app)
    - [Banco de Dados](#banco-de-dados)
      - [mysql-wallet](#mysql-wallet)
      - [mysql-balance](#mysql-balance)
    - [Mensageria](#mensageria)
      - [Apache Kafka e Zookeeper](#apache-kafka-e-zookeeper)
      - [Confluent Control Center](#confluent-control-center)
  - [Tecnologias](#tecnologias)
    - [wallet-core](#wallet-core-1)
      - [Linguagem](#linguagem)
      - [Bibliotecas](#bibliotecas)
      - [Migração de Banco de Dados](#migração-de-banco-de-dados)
      - [Gerenciamento de Dependências](#gerenciamento-de-dependências)
    - [balance-app](#balance-app-1)
      - [Linguagem](#linguagem-1)
      - [Frameworks e Bibliotecas](#frameworks-e-bibliotecas)
      - [Migração de Banco de Dados](#migração-de-banco-de-dados-1)
      - [Gerenciamento de Dependências](#gerenciamento-de-dependências-1)
    - [Banco de Dados](#banco-de-dados-1)
    - [Mensageria](#mensageria-1)
    - [Contêineres e Orquestração](#contêineres-e-orquestração)
  - [Pré-requisitos](#pré-requisitos)
  - [Executando e Testando o projeto](#executando-e-testando-o-projeto)
    - [Como executar o projeto](#como-executar-o-projeto)
    - [Verificando logs e status dos serviços](#verificando-logs-e-status-dos-serviços)
    - [Portas dos Serviços](#portas-dos-serviços)
    - [Observações Adicionais](#observações-adicionais)
  - [Documentação da API](#documentação-da-api)
    - [wallet-core](#wallet-core-2)
      - [Cria um cliente](#cria-um-cliente)
        - [Corpo da Requisição](#corpo-da-requisição)
        - [Resposta](#resposta)
      - [Cria uma conta](#cria-uma-conta)
        - [Corpo da Requisição](#corpo-da-requisição-1)
        - [Resposta](#resposta-1)
      - [Realiza uma transação](#realiza-uma-transação)
        - [Corpo da Requisição](#corpo-da-requisição-2)
        - [Resposta](#resposta-2)
    - [balance-app](#balance-app-2)
      - [Consulta o saldo de uma conta](#consulta-o-saldo-de-uma-conta)
        - [Parâmetros da Requisição](#parâmetros-da-requisição)
        - [Resposta](#resposta-3)
      - [Consulta a lista de contas](#consulta-a-lista-de-contas)
        - [Resposta](#resposta-4)
  - [Licença](#licença)
  - [Autores](#autores)
  - [Agradecimentos](#agradecimentos)

## Descrição

Este projeto é um sistema de gerenciamento de carteiras e saldos financeiros desenvolvido no contexto do curso Event Driven Architecture oferecido pela FullCycle. O objetivo é aplicar os conceitos de arquitetura orientada a eventos em uma aplicação prática, utilizando tecnologias modernas e boas práticas de desenvolvimento.

O sistema é composto por dois serviços principais:

- wallet-core: Implementado em Go durante as aulas do curso, este serviço é responsável pelo gerenciamento de contas e transações financeiras da carteira digital. Ele lida com as operações fundamentais que garantem a integridade e a segurança das transações dos usuários.
- balance-app: Desenvolvido em Java como parte do desafio proposto, este serviço é dedicado ao cálculo e atualização dos saldos. Ele assegura que todas as movimentações financeiras refletem corretamente nos balanços, mantendo a consistência dos dados.

A arquitetura do projeto foi concebida seguindo o paradigma de microserviços, o que promove escalabilidade, flexibilidade e facilidade de manutenção. A utilização de contêineres Docker permite que os serviços sejam implantados de maneira consistente em qualquer ambiente, facilitando o processo de desenvolvimento e implantação contínua.

Para comunicação assíncrona entre os serviços, o projeto emprega o Apache Kafka junto com o Zookeeper, implementando os princípios da arquitetura orientada a eventos. Essa abordagem garante que os serviços possam interagir de forma desacoplada, aumentando a robustez e a resiliência do sistema como um todo.

---

## Componentes

### Serviços

O projeto é composto por dois serviços principais:

#### wallet-core

- **Descrição**: Serviço desenvolvido em **Go** responsável pelo gerenciamento de contas e transações da carteira digital. Ele lida com as operações fundamentais que garantem a integridade e a segurança das transações dos usuários.

#### balance-app

- **Descrição**: Aplicação desenvolvida em **Java** responsável pelo gerenciamento de balanços financeiros. Ela assegura que todas as movimentações financeiras sejam refletidas corretamente nos balanços, mantendo a consistência e integridade dos dados.

### Banco de Dados

#### mysql-wallet

- **Descrição**: Instância do MySQL utilizada pelo serviço `wallet-core` para armazenamento de dados relacionados a clientes, contas e transações.
- **Dados Persistentes**: Armazenados no diretório `mysql-wallet`.
- **Configuração**: Detalhes de configuração podem ser encontrados no arquivo `docker-compose.yaml`.

#### mysql-balance

- **Descrição**: Instância do MySQL utilizada pelo serviço `balance-app` para armazenamento de dados relacionados aos balanços das contas.
- **Dados Persistentes**: Armazenados no diretório `mysql-balance`.
- **Configuração**: Detalhes de configuração podem ser encontrados no arquivo `docker-compose.yaml`.

### Mensageria

#### Apache Kafka e Zookeeper

- **Descrição**: Utilizados para comunicação assíncrona entre os serviços, implementando uma arquitetura orientada a eventos que permite o desacoplamento e escalabilidade dos componentes.
- **Configuração**: As configurações dos serviços Kafka e Zookeeper estão definidas no arquivo `docker-compose.yaml`.

#### Confluent Control Center

- **Descrição**: Ferramenta de monitoramento e gerenciamento para plataformas baseadas em Kafka. Permite visualizar em tempo real o fluxo de dados, tópicos, consumidores, produtores e outras métricas relevantes.
- **Características**:
  - Facilita o monitoramento das operações de mensageria entre os serviços.
  - Acessível através de uma interface web intuitiva.
- **Configuração**: Detalhes sobre a integração com o Confluent Control Center estão disponíveis no arquivo `docker-compose.yaml`.

---

## Tecnologias

O projeto utiliza as seguintes tecnologias e ferramentas:

### wallet-core

#### Linguagem

- **[Go](https://golang.org/)**: Linguagem de programação utilizada no `wallet-core` para desenvolvimento de serviços e aplicações.

#### Bibliotecas

- **[Confluent Kafka Go](https://github.com/confluentinc/confluent-kafka-go)**: Cliente Go para Apache Kafka, utilizado no `wallet-core` para comunicação assíncrona.
- **[Go-Chi](https://github.com/go-chi/chi)**: Framework leve para criação de APIs em Go, utilizado no `wallet-core` para roteamento e manipulação de requisições HTTP.
- **[Go-SQL-Driver](https://github.com/go-sql-driver/mysql)**: Driver MySQL para Go, utilizado no `wallet-core` para persistência de dados em ambiente de produção.
- **[Godotenv](https://github.com/joho/godotenv)**: Carrega variáveis de ambiente a partir de arquivos `.env` no `wallet-core`.
- **[Go-SQLite3](https://github.com/mattn/go-sqlite3)**: Driver SQLite para Go, utilizado no `wallet-core` para persistência de dados em ambiente de desenvolvimento.
- **[Ulid](https://github.com/oklog/ulid/v2)**: Biblioteca Go para geração de identificadores únicos lexicograficamente ordenados, utilizada no `wallet-core`.
- **[Testify](https://github.com/stretchr/testify)**: Pacote de utilidades para testes em Go, utilizado no `wallet-core` para asserções e mocks.

#### Migração de Banco de Dados

- **[Migrate](https://github.com/golang-migrate/migrate)**: Ferramenta de migração de banco de dados utilizada no `wallet-core` para gerenciar mudanças na estrutura do banco de dados de forma controlada.

#### Gerenciamento de Dependências

- **[Go Modules](https://blog.golang.org/using-go-modules)**: Sistema oficial de gerenciamento de dependências em Go.
- **[Makefile](https://www.gnu.org/software/make/manual/make.html)**: Utilizado para automação de tarefas comuns, como build, execução e testes.

### balance-app

#### Linguagem

- **[Java](https://www.java.com/)**: Linguagem de programação utilizada no `balance-app` para desenvolvimento de serviços e aplicações.

#### Frameworks e Bibliotecas

- **[Lombok](https://projectlombok.org/)**: Biblioteca Java que reduz o código boilerplate, fornecendo anotações para geração automática de getters, setters e outros métodos comuns.
- **[Spring Boot](https://spring.io/projects/spring-boot)**: Framework Java utilizado no `balance-app` para facilitar o desenvolvimento de aplicações web ao fornecer configurações padrão e integração com outros serviços.
- **[Spring Web](https://spring.io/projects/spring-web)**: Projeto Spring para desenvolvimento de aplicações web, utilizado no `balance-app` para criação de APIs RESTful.
- **[Spring Data JPA](https://spring.io/projects/spring-data-jpa)**: Projeto Spring para integração com JPA, utilizado no `balance-app` para persistência de dados.
- **[Spring Kafka](https://spring.io/projects/spring-kafka)**: Projeto Spring para integração com Apache Kafka, utilizado no `balance-app` para comunicação assíncrona.

#### Migração de Banco de Dados

- **[Liquibase](https://www.liquibase.org/)**: Ferramenta de migração de banco de dados utilizada no `balance-app` para gerenciar mudanças na estrutura do banco de dados de forma controlada.

#### Gerenciamento de Dependências

- **[Gradle](https://gradle.org/)**: Ferramenta de automação de build e gerenciamento de dependências Java, utilizada no `balance-app`.

### Banco de Dados

- **[MySQL](https://www.mysql.com/)**: Sistema de gerenciamento de banco de dados relacional utilizado pelos serviços `wallet-core` e `balance-app

### Mensageria

- **[Apache Kafka](https://kafka.apache.org/)**: Plataforma de streaming de eventos utilizada para comunicação assíncrona entre os serviços.
- **[Confluent Control Center](https://www.confluent.io/confluent-control-center/)**: Ferramenta de monitoramento e gerenciamento para plataformas baseadas em Kafka.

### Contêineres e Orquestração

- **[Docker](https://www.docker.com/)**: Utilizado para containerização dos serviços e dependências.
- **[Docker Compose](https://docs.docker.com/compose/)**: Utilizado para orquestração dos contêineres e serviços.

---

## Pré-requisitos

- **Docker**: Instalado e configurado na máquina local.
- **Docker Compose**: Instalado e configurado na máquina local.
- **Git**: Instalado e configurado na máquina local.
- **Make**: Opcional, mas recomendado para facilitar a execução de tarefas comuns.

Caso queira desenvolver localmente, é necessário ter as seguintes ferramentas instaladas:

- **Go**: Versão 1.23 ou superior.
- **Java**: Versão 23 ou superior.
- **VS Code**: Editor de código recomendado.
- **IntelliJ IDEA**: IDE recomendada para desenvolvimento Java.

As outras dependências são gerenciadas pelos contêineres Docker e Docker Compose, ou pelas ferramentas de gerenciamento de dependências específicas de cada linguagem.

---

## Executando e Testando o projeto

### Como executar o projeto

Após clonar o repositório, siga as instruções abaixo para executar o projeto localmente.

1. Construir e iniciar os contêineres:

    Execute o comando abaixo na raiz do projeto para construir e iniciar todos os serviços em segundo plano:

    ```bash
    make up
    ```

    Ou, se preferir iniciar diretamente com o Docker Compose:

    ```bash
    docker-compose up -d --build
    ```

2. Parar e remover os contêineres:

    Para parar todos os serviços e remover os contêineres e volumes associados, execute:

    ```bash
    make down
    ```

    Ou, se preferir parar diretamente com o Docker Compose:

    ```bash
    docker-compose down
    sudo rm -rf ./.docker/mysql*
    ```

### Verificando logs e status dos serviços

1. Para visualizar os logs de um serviço específico, utilize:

    ```bash
    docker compose logs -f balanceapp
    ```

    ou

    ```bash
    docker compose logs -f goapp
    ```

Para verificar o status e os logs dos contêineres de maneira mais fácil, recomendo o uso do [lazydocker](https://github.com/jesseduffield/lazydocker), uma interface de terminal interativa para gerenciamento de contêineres Docker.

### Portas dos Serviços

- **wallet-core**: [http://localhost:8080](http://localhost:8080)
- **balance-app**: [http://localhost:8081](http://localhost:3003)
- **Confluent Control Center**: [http://localhost:9021](http://localhost:9021)
- **MySQL Wallet**: Porta 3306
- **MySQL Balance**: Porta 3307
- **Zookeeper**: Porta 2181
- **Kafka**: Porta 9092

Mais detalhes sobre as portas e configurações de rede podem ser encontrados no arquivo `docker-compose.yaml` e nos arquivos de configuração de cada serviço.

### Observações Adicionais

- Certifique-se de que as portas definidas nos arquivos docker-compose.yaml não estejam em uso por outros serviços em sua máquina.
- Os volumes persistentes são armazenados na pasta .docker. Se necessário, você pode remover dados antigos excluindo os diretórios correspondentes.

---

## Documentação da API

A API do projeto é composta por dois serviços principais: `wallet-core` e `balance-app`. Cada serviço fornece endpoints específicos para a criação de clientes, contas e transações, bem como a consulta dos saldos das contas.

Para facilitar o uso da API, o projeto contem dois arquivos .http que podem ser utilizados com a extensão REST Client do VS Code. Basta clicar no botão "Send Request" para executar as requisições e visualizar as respostas.

- **[[requests/wallet-core.http]]**: Contém as requisições para criação de clientes, contas e transações.
- **[[requests/balance-app.http]]**: Contém as requisições para consulta dos saldos das contas.

A documentação abaixo descreve os endpoints disponíveis, os parâmetros necessários e as respostas esperadas.

### wallet-core

A API do wallet-core fornece endpoints para a criação de clientes, contas e transações em uma carteira digital.

#### Cria um cliente

Essa rota permite criar um novo cliente na carteira digital, fornecendo o nome e o e-mail do cliente.

```http
POST http://localhost:8080/clients HTTP/1.1
Content-Type: application/json

{
  "name": "Zé Galinha",
  "email": "ze@galinha.com"
}
```

##### Corpo da Requisição

| Parâmetro | Tipo     | Descrição                            |
| :-------- | :------- | :----------------------------------- |
| `name`    | `string` | **Obrigatório**. O nome do cliente   |
| `email`   | `string` | **Obrigatório**. O e-mail do cliente |

##### Resposta

A resposta contém os dados do cliente criado:

```json
{
  "ID": "01JCMDQ6W0RCBZ4ZHCVQG0SAZ8",
  "CreatedAt": "2024-11-14T04:08:21.120816294Z",
  "UpdateAt": "2024-11-14T04:08:21.120816334Z",
  "Name": "Zé Galinha",
  "Email": "ze@galinha.com"
}
```

#### Cria uma conta

Essa rota permite criar uma nova conta para um cliente existente, fornecendo o ID do cliente.

```http
POST http://localhost:8080/accounts HTTP/1.1'
Content-Type: application/json

{
  "client_id": "01JCBJNZFP2G6QHBX9MJHDAGR3"
}
```

##### Corpo da Requisição

| Parâmetro   | Tipo     | Descrição                        |
| :---------- | :------- | :------------------------------- |
| `client_id` | `string` | **Obrigatório**. O ID do cliente |

##### Resposta

A resposta contém o ID da conta ecém criada.

```json
{
  "ID": "01JCMDZN7V1QHKPK45NMXGYTN7"
}
```

#### Realiza uma transação

Essa rota permite realizar uma transação financeira entre duas contas, fornecendo os IDs das contas de origem e destino, bem como o valor da transação.

```http
POST http://localhost:8080/transactions HTTP/1.1
Content-Type: application/json

{
  "from_account_id": "01JCBJMXSSV0EGRRJ69ND23NEV",
  "to_account_id": "01JCBJPGHCQA0DJRFAXWE32C68",
  "amount": 1
}
```

##### Corpo da Requisição

| Parâmetro         | Tipo     | Descrição                                 |
| :---------------- | :------- | :---------------------------------------- |
| `from_account_id` | `string` | **Obrigatório**. O ID da conta de origem  |
| `to_account_id`   | `string` | **Obrigatório**. O ID da conta de destino |
| `amount`          | `float`  | **Obrigatório**. O valor da transação     |

##### Resposta

A resposta contém o id da trasnsação, a conta de origem, a conta de destino e o valor da transação.

```json
{
  "id": "01JCME4624TBSM22901YHMCR7K",
  "from_account_id": "01JCBJMXSSV0EGRRJ69ND23NEV",
  "to_account_id": "01JCBJPGHCQA0DJRFAXWE32C68",
  "amount": 1
}
```

### balance-app

A API do balance-app fornece endpoints para consulta dos saldos das contas.

#### Consulta o saldo de uma conta

Essa rota permite consultar o saldo de uma conta específica, fornecendo o ID da conta.

```http
GET http://localhost:3003/balances/01JCBJMXSSV0EGRRJ69ND23NEV HTTP/1.1
```

##### Parâmetros da Requisição

| Parâmetro | Tipo     | Descrição                      |
| :-------- | :------- | :----------------------------- |
| `id`      | `string` | **Obrigatório**. O ID da conta |

##### Resposta

A resposta contém o saldo atual da conta consultada, assim como o ID da conta.

```json
{
  "id": "01JCBJMXSSV0EGRRJ69ND23NEV",
  "createdAt": "2024-11-11T00:00:00",
  "updatedAt": "2024-11-14T04:15:26",
  "clientId": "01JCBJNZFP2G6QHBX9MJHDAGR3",
  "balance": 998.00
}
```

#### Consulta a lista de contas

Essa rota permite consultar a lista de contas cadastradas no sistema.

```http
GET http://localhost:3003/balances HTTP/1.1
```

##### Resposta

A resposta contém uma lista de contas, cada uma com seu ID e saldo atual.

```json
[
  {
    "id": "01JCBJMXSSV0EGRRJ69ND23NEV",
    "createdAt": "2024-11-11T00:00:00",
    "updatedAt": "2024-11-14T04:15:26",
    "clientId": "01JCBJNZFP2G6QHBX9MJHDAGR3",
    "balance": 998.00
  },
  {
    "id": "01JCBJPGHCQA0DJRFAXWE32C68",
    "createdAt": "2024-11-11T00:00:00",
    "updatedAt": "2024-11-14T04:15:26",
    "clientId": "01JCBJPMXY1H52YGGG5FD328MT",
    "balance": 1002.00
  }
]
```

---

## Licença

Distribuído sob a licença [MIT](LICENSE).

---

## Autores

- [Josenaldo Matos](https://www.github.com/josenaldo)

## Agradecimentos

Este projeto foi desenvolvido com base no curso Event Driven Architecture, da FullCycle, que foi ministrado por Wesley Willians. Agradeço a toda a equipe da [FullCycle](https://fullcycle.com.br/) por disponibilizar conteúdos de qualidade e por promover a disseminação de conhecimento na comunidade de desenvolvimento de software.

- [Wesley Willians](https://www.github.com/wesleywillians)
- [Luiz Carlos](https://www.github.com/argentinaluiz)
- [Full Cycle](https://fullcycle.com.br/)
