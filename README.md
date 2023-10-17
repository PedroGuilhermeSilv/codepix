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
![Screenshot of a comment on a GitHub issue showing an image, added in the Markdown, of an Octocat smiling and raising a tentacle.](https://myoctocat.com/assets/images/base-octocat.svg)

# Dinâmica do processo
1. Registra transação.
2. Muda status para: "confirmed".
3. Informa banco de origem que a transação foi confirmada pelo banco de destino com status: "confirmed".
4. Finaliza a transação mudando o status para: "completed".

![Screenshot of a comment on a GitHub issue showing an image, added in the Markdown, of an Octocat smiling and raising a tentacle.](https://myoctocat.com/assets/images/base-octocat.svg)

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
![Screenshot of a comment on a GitHub issue showing an image, added in the Markdown, of an Octocat smiling and raising a tentacle.](https://myoctocat.com/assets/images/base-octocat.svg)
