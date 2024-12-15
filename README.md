# E-commerce Product Management System

## Overview
This system facilitates the management of products for an e-commerce platform, allowing users to add, update, delete, and view products, along with additional functionality for processing product data asynchronously.

## Architectural Choices
- **Go**: Selected for its high performance, simplicity, and strong concurrency handling, making it ideal for building scalable services.
- **RabbitMQ**: Employed for decoupling microservices and enabling asynchronous messaging between services, ensuring better scalability.
- **PostgreSQL**: Chosen for its powerful relational database capabilities, ACID compliance, and support for complex queries.
- **Redis**: Implemented for caching frequently accessed data, improving read performance and reducing database load.

## Installation & Setup

1. Clone the repository:
   ```bash
   git clone https://github.com/ajay3750/Product-Management.git
