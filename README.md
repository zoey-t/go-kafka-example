# go-kafka-example
Data processing with Apache Kafka, REST API and Redis.

This project involves processing on-chain blockchain data using Apache Kafka, REST API, and Redis.
There are three topics: "address", "label" and "transaction". These topics are handled by a single producer and consumer.

The producer, operating on port 8080, receives REST API requests from users, sending the respective messages to the consumer (port 8081). Redis is used as the in-memory database.


## Prerequisites
- docker ([install](https://docs.docker.com/engine/install/))
## stack
- [bitnami/kafka](https://hub.docker.com/r/bitnami/kafka): kafka docker image
- [ibm/sarama](https://github.com/IBM/sarama): Go library for Apache Kafka
- [echo](https://echo.labstack.com/): Go web framework used to build REST APIs
- [redis](https://redis.uptrace.dev/guide/go-redis.html): for in-memory database
## Run the application
1. Start docker: 
    ```bash
    make run/docker
    ```
2. Start the producer in a new terminal
    ```bash
    make run/producer
    ```
3. Start the consumer in a new terminal
    ```bash
    make run/consumer
    ```
4. Send API requests to `localhost:8080`
   
    For example:
   
    - create an address:
        send POST request to `localhost:8080/addresses/0x87631B45877794f9cdd50a70c827403e3C36d072`
        with body
        ```json
        {
        "address": "0x87631B45877794f9cdd50a70c827403e3C36d072",
        "labels": ["eoa"]
        
        }
        ```

    - get an address:
        send GET request to `localhost:8080/addresses/0x87631B45877794f9cdd50a70c827403e3C36d072`
     - create a transaction: 
       send POST request to `localhost:8080/transactions/0x6ffa912cc7da2b5ec51a2cc1152ab39a54f0c72f1b3f32072c9bba154b585780` with body
       ```json
        {
        "hash":"0x6ffa912cc7da2b5ec51a2cc1152ab39a54f0c72f1b3f32072c9bba154b585780",
        "chainid": 2,
        "from": "0x4838B106FCe9647Bdf1E7877BF73cE8B0BAD5f97",
        "to":"0x388C818CA8B9251b393131C08a736A67ccB19297",
        "status":"success"
        }
       ```
## REST API endpoints
- port: `localhost:8080`
- `/addresses`: `GET`
- `/addresses/:address`: `GET`, `POST`, `PUT`, `DELETE`
- `/labels`: `GET`
- `/labels/:label`: `GET`, `POST`, `PUT`, `DELETE`
- `/transactions`: `GET`
- `/transactions/:hash`: `GET`, `POST`, `PUT`, `DELETE`

## Project structure
- `/cmd`
main applications for this project
- `/cmd/producer`
kafka producer. it sends messages to kafka.
- `/cmd/consumer`
kafka consumer. It process the received 
- `/config`
configuration
- `/controllers`
controllers that handles api endpoints request
- `/models`
data structures and models used in REST API
- `/utils`
util/helper functions
