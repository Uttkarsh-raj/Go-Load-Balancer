# Go-Load-Balancer

**Go-Load-Balancer** is a lightweight load-balancer built in Golang. It includes implementations for round-robin static routing and least connection dynamic routing, offering a simple yet powerful solution for distributing network traffic efficiently.

## Table of Contents
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li><a href="#about-the-project">About The Project</a></li>
    <li><a href="#built-with">Built With</a></li>
    <li><a href="#getting-started">Getting Started</a></li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgements">Acknowledgements</a></li>
  </ol>
</details>

## About The Project

**Go-Load-Balancer** is a lightweight load-balancer built in Golang. It supports round-robin static routing and least connection dynamic routing, making it an ideal tool for evenly distributing traffic across multiple servers.
### Load Balancer

A load balancer is a network device or software that distributes incoming network traffic across multiple servers. This helps ensure that no single server becomes overwhelmed, improving reliability, availability, and scalability.

### Static Routing

Static routing refers to a predetermined, fixed path that network traffic follows. It doesn't change unless manually adjusted, making it simple but inflexible.

#### Round Robin

Round Robin is a static routing method where requests are distributed in a cyclic manner. Each server in the pool receives an equal number of requests in turn, regardless of server load or capacity.

https://github.com/user-attachments/assets/4caf8a24-25fe-4309-8ef9-656eebea0c3c



### Dynamic Routing

Dynamic routing adjusts the path of network traffic based on real-time conditions, such as server load or network congestion. This helps optimize performance and resource utilization.

#### Least Connection

The Least Connection method is a dynamic routing strategy that directs new requests to the server with the fewest active connections. This helps balance the load more evenly, especially in environments with varying server capacities or request processing times.

> **_NOTE:_** Here, I have set up 3 demo servers where the sleep time simulates computation time, varying for each server based on its processing power and input data. As seen in the response, the 2nd server, with the shortest sleep time (representing faster computation), handles more requests due to its quicker processing.

https://github.com/user-attachments/assets/b43712a0-e4cb-4dc9-a7f7-274a80b87d46

### Built With
- **Golang** - An open-source programming language.

<img height="100px" src="https://upload.wikimedia.org/wikipedia/commons/0/05/Go_Logo_Blue.svg"/>

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## Getting Started

To get started with **go-load-balancer**, follow these steps:

1. **Install Golang**: Download and install Golang from the [official website](https://golang.org/dl/).
2. **Set Up Your Workspace**: Create a new directory for your project and set your `GOPATH` environment variable to point to this directory.
3. **Initialize Your Project**: Inside your project directory, run:
   ```
   go mod init github.com/your-username/go-load-balancer
   ```
4. **Run without Debugging**: 
   ```
   go run main.go
   ```

### Installation

1. Clone the repo:
   ```
   git clone https://github.com/Uttkarsh-raj/go-load-balancer
   ```
2. Install the packages:
   ```
   go mod tidy
   ```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## Routes

- **POST "/addServer"**
  * Request as:
    - body
      ```
      {
        "address":"https://server_address.com"
      }
      ```
  * Response as:
    - Success: Respective output.
      ```
      {
        "message": "Server added successfully",
        "success": true
      }
      ```
    - Failure: "message" : Error.
      ```
      {
        "message": "error: Unable to add the server",
        "success": false
      }
      ```
- **DELETE "/delete"**
  * Request as:
    - body
      ```
      {
        "address":"https://server_address.com"
      }
      ```
  * Response as:
    - Success: Respective output.
      ```
      {
        "message": "Server deleted successfully",
        "success": true
      }
      ```
    - Failure: "message" : Error.
      ```
      {
        "message": "error: Server does not exists",
        "success": false
      }
      ```
- **ANY "/path_for_the_servers"**
  * Request as:
    - body as required by the servers being handled by the load-balancer
  * Response as:
    - Success: Respective output.
    - Failure: "message" : Error.

## Usage
The **Go-Load-Balancer** project offers a flexible solution for load balancing across multiple servers. It supports:

1. **Web Application Scaling**: Distribute user requests across multiple web servers to handle increased traffic and ensure smooth performance.
2. **API Gateway**: Balance incoming API requests among backend services, improving reliability and response time.
3. **Database Load Management**: Distribute queries across multiple database instances to prevent any single instance from being overwhelmed.

## Screenshots
<center>
<br> 
<img width="1000" src="https://github.com/user-attachments/assets/3068c64b-1303-4779-9bb8-5647efc5f0bf"></img>
<img width="1000" src="https://github.com/user-attachments/assets/aa9c5525-db19-4412-bc22-f7dba121e52a"></img>
<br>
</center>

## Contributing

Contributions make the open-source community an amazing place to learn, inspire, and create. Any contributions you make are *greatly appreciated*.

1. Fork the Project.
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`).
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`).
4. Push to the Branch (`git push origin feature/AmazingFeature`).
5. Open a Pull Request.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## License
This project is licensed under the MIT License, a permissive open-source license that allows for the use, modification, and distribution of this software. It grants users the freedom to use the software for any purpose, including commercial purposes, as long as the original copyright and license notice is included. See the [LICENSE](LICENSE) file for more information.
(<a href="#readme-top">back to top</a>)</p>

## Contact
Uttkarsh Raj - [GitHub](https://github.com/Uttkarsh-raj)

Project Link: https://github.com/Uttkarsh-raj/go-load-balancer

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## Acknowledgments

- [Choose an Open Source License](https://choosealicense.com)

<p align="right">(<a href="#readme-top">back to top</a>)</p>
