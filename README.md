# Sobre o Projeto   
- É uma solução para simularmos transferências de valores entre banco fictícios através de chaves (email, cpf).
- Simularemos diversos bancos e contas bancárias que possuem uma chave Pix atrbuída.
- Cada conta bancária poderá cadastrar suas chaves Pix.
- Uma conta bancária poderá realizar uma transferência para outra conta em outro banco utilizando a chave Pix da conta destino.
- Uma transação não pode ser perdida mesmo que o sistema do CodePix ou de um dos bancos estejam fora do ar.

# Sobre os Bancos
- O banco será um microserviço limitado a fazer tranferências, cadastrar contas e chaves Pix.
- Utilizaremos a mesma aplicação par a fazer vários bancos mudando somente as cores, nome e código identificador.
- Nest.js (Banckend)
- Next.js (Frontend)

# Sobre o CodePix
- O CodePix será um microsserviço responsável por intermediar as transferências bancárias.
- Receberá a transação de transferência.
- Encaminhar a transação para o banco de destino (status: "pending").
- Recebe a confirmação do banco de destino (status: "confirmed").
- Envia a confirmação para o banco de origem informando que o banco de destino processou.
- Recebe a confirmação do banco de origem de que ele processou (status: "completed")
- Marca a transação como completa (status: "completed")

# Cadastro e consulta de chave Pix 
![](https://github.com/PedroGuilhermeSilv/codepix-go/blob/main/documentation/img/cadastro-consulta-pix.png)

# Dinâmica do processo
1. Registra transação.
2. Muda status para: "confirmed".
3. Informa banco de origem que a transação foi confirmada pelo banco de destino com status: "confirmed".
4. Finaliza a transação mudando o status para: "completed".

![](https://github.com/PedroGuilhermeSilv/codepix-go/blob/main/documentation/img/dinamica-processo.png)

# Principais desafios

- Comunicação rápida e eficiente.
- Criação de consulta instantânea das chaves (síncrona)
    - gRPC: Framework que utiliza o http2 e protocol buffers transicionando os dados de forma binária.   
- Garantia de que nenhuma transação seja perdida, mesmo que um sistema esteja fora do ar(assíncrona)
    - Apache Kafka: Stream de dados

# Sobre o CodePix 
- Será capaz de atuar como um servidor gRPC
- Consumir e publicar mensagens no Apache kafka
- Ambas operações devem ser realizadas de forma simultânea ao executar o serviço 
- Trabalhar com um design focado em solucionar o problema de domínio

# Estrutura e camadas do CodePix (clean architecture)
![](https://github.com/PedroGuilhermeSilv/codepix-go/blob/main/documentation/img/estrutura-camada.png)


#  gRPC
## Em quais casos podemos utilizar?
- Idel para microservicos 
- Mobile, Browsers e Backend
- Geraco das bibliotecas de forma automatica.
- Streaming bidirecional utilizando HTTP/2

# RPC -  Remote Procedure Call
![](https://github.com/PedroGuilhermeSilv/codepix-go/blob/main/documentation/grpc.png)

# Protocol Burffers
![](https://github.com/PedroGuilhermeSilv/codepix-go/blob/main/documentation/img/protocol-buffers.png)

- Primeiro é definido como deseja que os dados sejam estruturados — em um arquivo de extensão .proto.
- Em seguida, esta definição é compilada e o resultado é um código-fonte automaticamente gerado na linguagem desejada — no momento que escrevo este post, as linguagens compatíveis são C++, C#, Go, Java e Python.
- Finalmente, código-fonte gerado é utilizado para gravar e ler os dados estruturados.
Sempre que houver mudança na estrutura dos dados, o ciclo se repetirá.

# Protocol buffers vs JSON
- Arquivos binarios < JSON  \
- Processo de serializacao e mais leve (CPU) do que o JSON
- Gasta menos recursos de rede
- Processo e mais veloz

# HTTP/2
- Dados trafegados sao em formatos binarios
Utiliza a mesma conexao TCP para enviar e receber dados do cliente e do servidor (Multiplex).
- Headers sao comprimidos

# gRPC - API "unary"
![](https://github.com/PedroGuilhermeSilv/codepix-go/blob/main/documentation/img/grpc-unary.png)

# gRPC - API "Server streaming"
![](https://github.com/PedroGuilhermeSilv/codepix-go/blob/main/documentation/img/grpc-streaming.png)

# gRPC - API "Client streaming"
![](https://github.com/PedroGuilhermeSilv/codepix-go/blob/main/documentation/img/grpc-client-streaming.png)

# REST vs gRPC
- gRPC para comunicao entre servidor e browsers nao e eficiente. Porem, e muito bom para comunicacao de microservicos
- Texto / JSONj vs Protocol Buffers
- Unidirecional vs Bidirecional
- Alta latencia vs Baixa latencia
- Sem contrato vs Contrato definido(.proto)
- Sem suporte a streaming (Request / Response) vs Suporte a streaming

# Apache kafka
- Event-driven
- Tempo real
- Historico dos dados 
- Caracteristicas:
    - Plataforma
    - Trabalha de forma distribuida
    - Banco de dados
    - Extremamente rapido
    - Utiliza o disco ao inves de memoria para processar os dados  
- Topic
    - Stream de dados que atua como um banco de dados
    - Todos os dados ficam armazenados
    - O topico possui diversas particoes
    - Cada particao e definido por numero
    - Voce e obrigado a definir a quantidade de porticoes de um topico
![](https://github.com/PedroGuilhermeSilv/codepix-go/blob/main/documentation/img/topic.png)

- Producers: mandam mensagens
- Consumer: lêem mensagens
![](https://github.com/PedroGuilhermeSilv/codepix-go/blob/main/documentation/img/producer-consumer.png)

# Kafka Cluster
- Conjunto de Brokers
- Cada Broker e um server
- Cada Broker e responsavel por armazenar os dados de uma particao
- Cada particao de Topic esta distribuido em diferentes brokes
![](https://github.com/PedroGuilhermeSilv/codepix-go/blob/main/documentation/img/kafka-cluster.png)

- Com replication factor o Kafka faz diversas copias de particoes de um Topic em diferentes brokes fazendo com que mesmo que uma maquina esteja fora do ar ele tenha a informacao completa.
![](https://github.com/PedroGuilhermeSilv/codepix-go/blob/main/documentation/img/kafka-replication1.png)

![](https://github.com/PedroGuilhermeSilv/codepix-go/blob/main/documentation/img/kafka-replication1.png)

![](https://github.com/PedroGuilhermeSilv/codepix-go/blob/main/documentation/img/kafka-replication1.png)

# Ecossitema
- Kafka Connect
    - Connectors
- Confluent Schema Registry
- Rest Proxy
- ksqlDB
- Streams

Dúvidas: Aula 02: Preload, minuto 1:25:04 comando no cmd não executa