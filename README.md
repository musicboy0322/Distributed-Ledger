# Distributed-Ledger

* Motivation: After participating in a blockchain project, I realized blockchain is fundametally a large-scale distributed system which is an area I had never deeply explored before. This insight motivated me to actively seek new knowledge and challenge myself by implementing a distributed ledger system similar to blockchain. Through this project, I aim to expand my skills and deepen my understanding of distributed systems.
  
* Purpose: To develop an application that simulates blockchain-like functionality using distributed ledger concepts.

* Command:
    * Check Money: Verify the balance of target wallet.
    * Check Log: Review the history transition of target wallet.
    * Transition: Transfer money from one wallet to another.

* System Feature:
    * Implementing multi-threading in each node using Goroutines to increase connection throughput and concurrency.
    * Using Socket TCP to handle long connections among nodes and short conncetions on the client side.
    * Using Kubernetes to manage Docker for stimulating mutiple nodes.
      
      <img width="853" alt="Screenshot 2025-03-29 at 9 25 21 AM" src="https://github.com/user-attachments/assets/57ab9dcc-af9f-4369-b80d-c397eff49b01" />
      <img width="661" alt="Screenshot 2025-03-29 at 9 26 30 AM" src="https://github.com/user-attachments/assets/04cbce50-629d-4331-ac48-bab9c41d48b2" />

* Blockchain Feature:
    *  Each node acts as both a data store and an entry point for transactions.
    *  Each node follows the principle of eventual consistency when storing transaction information.
    *  Every transaction information is stored on all nodes.
    *  Clients do not store any transaction information.
    *  Clients randomly connect to a node for transaction submission and transaction log retrieval.
      
* Technique:
    * Programming Language: Golang
    * Network Programming: Socket, TCP
    * DevOps Tools: Docker, Kubernetes(Minikube)

* Demo: [https://youtu.be/h7gLIvKN9pk](https://youtu.be/NwrnfNtub5o)
