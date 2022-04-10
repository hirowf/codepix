# Case Study - CodePix

## About

- A solution to simulate value transfers between fictional banks through keys (email, cpf);
- We’ll simulate several banks and bank accounts that have an Pix key assigned;
- Each bank account will can to register your Pix key
- An bank account will can realise a transfer to the other account at other bank using an Pix key from the destination account;
- An transfer cannot be lost even that: CodePix system is down;
- An transfer cannot be lost even that: An destinantion Bank system is down;

## About Banks

- A Bank will be a microservice with limited functions to register new account and Pix keys, as well as transfers of values;
- We’ll use the same application to simulate several banks, changing only colors, name and code.
- Nest.js in the backend
- Next.js as frontend

## Idea

![image](https://user-images.githubusercontent.com/17819811/162598398-af8c5387-9982-4e48-9206-bf380755b2c9.png)


## About CodePix

- The microservice CodePix will be responsible for intermediate the banks transfers
- Will receive a transaction of transfer
- Will forward to transaction from destination bank (Status: “pending”)
- Receive a confirmation from destionation bank (Status: “confirmed”)
- Send confirmation to the origin bank informing when destination bank processed
- Receive confirmation from origin bank that transaction was processed (Status: “completed”)
- Mark the transaction as complete (Status: “completed”)

## Register and consulting a key Pix

- Consulte if key Pix exists and return if exists or not
- If not exists, we can create on sending a request to CodePix
- Key are storage in the bank & CodePix

![image](https://user-images.githubusercontent.com/17819811/162598370-f8421a0c-d616-41eb-b096-5f5539000d61.png)

## Process dynamic

![image](https://user-images.githubusercontent.com/17819811/162598380-dbe5f75a-90a2-4459-a294-bc391ddf802e.png)

**CodePix:**

1. Register a transaction
2. Change status to: ‘Confirmed’
3. Informs origin bank that a transaction was confirmed by the destination bank with status: “confirmed
4. Finalize the transaction changing status to: “completed”

## Main Challenges

- Fast and efficient communication
- Instant creation and query of keys (Synchronous)
- Guarantee that no transaction is lost even if any of the 3 systems are down (Asynchronous)

For this will use **gRPC,**  a frameworkd that use HTTP2 and protocol buffers

For asynchronous comunication will use **APACHE Kafka,** to process data

## CodePix

- Will be able to act as a gRPC server
- Consume and publish messages in Apache Kafka
- Both operations must be performed simultaneously when executing the service
- Work with a design focused on solving the domain problem
- Leave the technical complexity to the “application layer”, responsible for the gRPC and Kafka server
- Flexible for the implementation of other communication formats such as API Rest, CLI clients, etc. WITHOUT changing any other application components or the domain model.

## CodePix structure and layers

![image](https://user-images.githubusercontent.com/17819811/162598407-31c28b9a-3762-4458-b946-cbc7f14c4c52.png)

## Resources

- Docker
- Golang
- Apache Kafka
- Postgres
