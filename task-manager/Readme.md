# Nome do Projeto

Uma breve descrição do projeto, seus objetivos e funcionalidades principais.

## Índice

- [Sobre](#sobre)
- [Instalação](#instalação)
- [Uso](#uso)
- [Estrutura do Projeto](#estrutura-do-projeto)
- [Contribuição](#contribuição)
- [Licença](#licença)
- [Contato](#contato)

## Sobre

Este projeto tem como objetivo desenvolver uma API RESTful para gerenciamento de tarefas, proporcionando funcionalidades de CRUD (Create, Read, Update, Delete) de tarefas, autenticação de usuários e armazenamento dos dados em um banco de dados.


## Tecnologias
- Golang v1.22.4  (Liguagem de Programação)
- SQLite (Banco de Dados)
- Swagger (Documentação da API)
- DB Designer (Modelagem do BD)
- Git e Github (Versionamento e armazenamento de código fonte)

## Aplicação e Docs
- Swagger:
- Servidor:

## Instalação
Caso não tenha instalado ainda o Golang e SQLite em sua máquina, segue as instruções. Caso contrário, pule para o passo seguinte.

1. Instale o Golang em sua máquina:
    1.1. Acesse https://go.dev/doc/install e siga as instruções para instalação conforme OS.
    1.2. Verique a versão instalada:
    ```bash
    go version
    ```
2. Instale o SQLite em sua máquina:
    2.1. Acesse https://www.sqlite.org/ e siga as instruções da documentação para instalação conforme OS.
    2.2. Verique a versão instalada:
    ```bash
    sqlite3 --version
    ```


Passos para instalar e configurar o ambiente de desenvolvimento.

1. Clone o repositório:
    ```bash
    git clone git@github.com:rodolfoalvesg/TrilhaBackEndJR-JUN15.git
    ```
2. Vá para o diretório do projeto:
    ```bash
    cd TrilhaBackEndJR-JUN15/task-manager
    ```
3. Instale as dependências:
    ```bash
    make deps
    ```
4. Gere a documentação
     ```bash
    make generate
    ```

## Uso

Instruções sobre como usar o projeto:

1. Acesso ao banco de dados SQLite:
    ```bash
    sqlite3 app/gateway/http/db/sqlite3/database/db.db
    ```
2. Execute a aplicação:
    ```bash
    make start
    ```

## Estrutura do Projeto

Descrição da estrutura do diretório e uma breve explicação sobre o propósito de cada pasta/arquivo.

project-root/
│
├── app/
├──├── config/
├──├── domain/
├──├──├── entities/
├──├──└── usecases/
├──├── gateway/
├──├──└── http
├──├──├──├── middleware/
├──├──├──├── db/
├──├──├──├── rest/
├──├──├──├── router.go
├──└──└──└── server.go
│
├── cmd/
├──├── api/
├──└──└── main.go
│
├── docs/
│
├── .env
├── Readme.md
├── Makefile
├── go.mod
└── go.sum

Essa estrutura de projeto em Golang segue uma organização modular e bem definida, facilitando a manutenção e a escalabilidade do código. Vamos analisar cada diretório e arquivo:

#### Diretórios e Arquivos
`project-root/`: Diretório raiz do projeto.
`app/`: Diretório principal onde reside a lógica de aplicação.
`config/`: Contém configurações da aplicação, como parâmetros de inicialização, variáveis de ambiente, etc.
`domain/`: Contém a lógica de negócios e definições principais.
`entities/`: Define as entidades do domínio, ou seja, as estruturas de dados que representam os objetos do negócio.
`usecases/`: Define os casos de uso ou interações que podem ser feitas com as entidades.
`gateway/`: Define a camada de acesso a dados e serviços externos.
`http/`: Contém lógica relacionada à comunicação HTTP.
`middleware/`: Implementações de middleware para manipulação de requisições HTTP, como autenticação, autorização, etc.
`db/`: Contém a lógica de acesso ao banco de dados.
`rest/`: Define os controladores REST para expor APIs.
`router.go`: Define as rotas HTTP, associando endpoints aos controladores correspondentes.
`server.go`: Inicializa e configura o servidor HTTP.
`cmd/`: Contém os pontos de entrada da aplicação.
`api/`: Diretório específico para o comando API.
`main.go`: Arquivo principal que inicializa a aplicação API.
`docs/`: Armazena as informações de documentações. Como por exemplo esquema de banco de dados e informações da API utilizando ferramentas de terceiros, a exemplo do swagger.
`.env`: Arquivo que armazena variáveis de ambiente.
`Readme.md`: Arquivo de documentação do projeto.
`Makefile`: Arquivo que contém scripts de automação para tarefas comuns, como build, test, etc.
`go.mod`: Arquivo de definição de módulos do Go, gerencia dependências.
`go.sum`: Arquivo que garante a integridade das dependências listadas em go.mod.

#### Funcionalidade dos Componentes
`app/config:` Gerencia configurações que podem variar entre ambientes (desenvolvimento, produção).
`app/domain/entities:` Define os modelos de dados usados na aplicação, garantindo uma representação clara e consistente das entidades do domínio.
`app/domain/usecases:` Implementa a lógica de negócio e as regras que operam sobre as entidades.
`app/gateway/http/middleware:` Manipula requisições e respostas HTTP, aplicando lógica transversal como a autenticação.
`app/gateway/http/db:` Gerencia a persistência dos dados.
`app/gateway/http/rest:` Define padrões de responses e requests.
`app/gateway/http/router.go:` Configura o roteamento HTTP, mapeando URLs para os manipuladores apropriados.
`app/server.go:` Inicializa o servidor web, configurando portas, handlers e middlewares.
`cmd/api/main.go:` Ponto de entrada para execução da API, configurando e iniciando a aplicação.

## Processos e Etapas de Desenvolvimento
**1 - Definição Inicial**

O desenvolvimento da aplicação começou com duas decisões fundamentais:

**Linguagem de Programação:** Escolhi Go (Golang) devido à sua performance, simplicidade e forte suporte a concorrência, o que é ideal para desenvolvimento de aplicações web robustas(E que não necessitamos desse último caso).
**Banco de Dados:** Optei pelo SQLite devido a restrição do projeto, porém o mesmo é um banco de dados leve e autossuficiente, perfeito para nosso caso de uso. Além disso, adotei a prática de migrations para gerenciar e versionar alterações no esquema do banco de dados de forma controlada e reproduzível.

**2. Definição dos Payloads**
Antes de começar a codificação, defini os payloads da API. Isso envolveu a criação de contratos claros para as requisições e respostas que a aplicação iria manejar, garantindo consistência e clareza na comunicação entre cliente e servidor.

**3. Configuração Inicial do Projeto**
Iniciei a codificação com a configuração básica do projeto. A estrutura preferida se assemelha a uma arquitetura Hexagonal. Esta arquitetura promove a separação de preocupações, permitindo que a lógica de negócio seja isolada das interfaces externas, como interfaces de usuário, APIs, banco de dados, etc.

`app/config`: Defini as configurações da aplicação, como variáveis de ambiente, parâmetros de inicialização e outros ajustes necessários para diferentes ambientes.


**4. Criação do Servidor e Configuração de Rotas**
Em seguida, criei o servidor HTTP e configurei as rotas:

`app/gateway/http/server.go:` Preparei o servidor HTTP, configurando.
`app/gateway/http/router.go`: Configurei o roteamento HTTP, mapeando URLs para os manipuladores apropriados. Aqui, associei os endpoints definidos com os controladores que lidariam com cada requisição.

**5. Definição das Entidades e Casos de Uso**
Com o servidor e as rotas configuradas, defini as entidades do domínio e os casos de uso:

`app/domain/entities:` Criei as estruturas de dados que representam as entidades do nosso domínio, garantindo uma representação clara e consistente dos objetos do negócio.
`app/domain/usecases:` Implementei a lógica de negócio e as regras que operam sobre as entidades. Esses casos de uso definem as interações possíveis com as entidades, encapsulando a lógica de aplicação.

**6. Configuração dos Repositórios**
Por fim, configurei os repositórios que gerenciam a persistência dos dados:

`app/gateway/http/db:` Implementei a lógica de acesso ao banco de dados, utilizando SQLite e aplicando as migrations para versionar as alterações no esquema de dados.
`app/gateway/http/rest:` Criei os controladores REST que expõem os endpoints da API, organizando as operações CRUD e outras interações com o cliente.

**7. Implementação das Autorização**
Para garantir a segurança da aplicação, implementei um middleware de autorização que valida as rotas das tarefas:

`app/gateway/http/middleware`: Implementei middleware de autorização para verificar e validar as credenciais dos usuários antes de permitir o acesso aos recursos protegidos. Esse middleware intercepta as requisições HTTP, verifica o token jwt, de autenticação e valida permissões, assegurando que apenas usuários autorizados possam acessar determinadas rotas.

**7. Processo de Integração Contínua**
Cada etapa do desenvolvimento foi submetida através de Pull Requests (PRs). Este processo garantiu que cada mudança fosse avaliada de forma independente, mantendo a qualidade e a integridade do código.


## Licença

Informações sobre a licença do projeto.

Este projeto está licenciado sob a Licença MIT - veja o arquivo [LICENSE](LICENSE) para mais detalhes.

## Contato

Informações de contato para dúvidas ou sugestões.

- Nome: Rodolfo Alves Gonçalves
- Email: rodolfoalves.inf@gmail.com
- LinkedIn: [Seu LinkedIn](https://www.linkedin.com/in/rodolfoalvesg)
- GitHub: [Seu GitHub](https://github.com/rodolfoalvesg)