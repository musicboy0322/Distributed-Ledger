# Distributed-Ledger

* Motivation：My friends and I are developing a blockchain-related system. During the process, I realized that blockchain is built on a distributed system, a field I had never explored before.
  This realization inspired me to actively seek out new knowledge and challenge myself to implement and build a distributed system, expanding my skills and understanding in this area.

* Purpose：Using distributed ledger concepts to build an application that simulates blockchain-like functionality.

* Feature：
    * Implementing multi-threading to increase system throughput and concurrency, allowing multiple node connections to be handled simultaneously
    * Using TCP Socket for node-to-node communication
    * Using Docker to stimulate mutiple nodes
      
      <img width="1187" alt="截圖 2024-08-26 下午4 02 42" src="https://github.com/user-attachments/assets/e63cb1a1-9c37-41a6-9c21-41cb8c1621fb">
      
* Command：
    * Check Money：Verify the balance of target wallet
    * Check Log：Review the history transition of target wallet
    * Transition：Transfer money from one wallet to another
    * Check Chain：Check if the local block has been altered
    * Check All Chains：Check if any blocks across all nodes have been altered

* Technique：
    * Programing Language：Golang
    * Network Programing：Socket、TCP
    * Containerization Tool：Docker
      
* Future Expectation：
    * Adding a heartbeat mechanism to monitor the status of long connections and prevent redundant or stale connections
    * Implementing a load balancing mechanism to distribute incoming transitions, preventing overload on any one server and improving system performance and stability

* Demo：https://youtu.be/h7gLIvKN9pk
