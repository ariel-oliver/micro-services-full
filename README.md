# micro-services-full

# Como Executar um Docker Compose

Este guia irá orientá-lo sobre como executar um projeto usando Docker Compose. Docker Compose é uma ferramenta para definir e executar aplicações Docker multi-contêineres. Com ele, você pode usar um arquivo YAML para configurar os serviços da sua aplicação e, em seguida, iniciar todos os serviços com um único comando.

## Pré-requisitos

- Ter o Docker instalado na sua máquina. Se você ainda não tem o Docker instalado, você pode baixá-lo e instalá-lo a partir do [site oficial do Docker](https://www.docker.com/products/docker-desktop).
- Ter um arquivo `docker-compose.yml` válido no diretório do seu projeto. Este arquivo define os serviços, redes e volumes da sua aplicação.

## Passos para Executar o Docker Compose

1. **Abra o Terminal**: Abra o terminal ou prompt de comando no seu computador.

2. **Navegue até o Diretório do Projeto**: Use o comando `cd` para navegar até o diretório que contém o arquivo `docker-compose.yml` do seu projeto.



3. **Execute o Docker Compose**: No diretório do projeto, execute o seguinte comando para iniciar os serviços definidos no seu arquivo `docker-compose.yml`:
    ```bash
    docker-compose up
    ```


    Este comando irá construir, (re)criar, iniciar e anexar aos contêineres definidos no arquivo `docker-compose.yml`. Se você quiser executar os contêineres em segundo plano, você pode adicionar a opção `-d`:
    ```bash
    docker-compose up -d
    ```


