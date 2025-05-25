# Scalable E-commerce Platform — Microservices Architecture (Go + Docker)

## About the Project  
ZenoMarket is a scalable e-commerce platform developed as a prototype for small-scale online retail.  
The project is built using a microservices architecture with domain-specific logic separated into individual services: users, products, cart, orders, and payments.

---

## 🛠 Stack

- Go (Golang)
- PostgreSQL
- Redis
- Docker + Docker Compose
- REST API (net/http)
- JWT Authentication
- Clean Architecture (modular structure)

---

## 📦 Microservices

- `user-service`: registration, login, profile
- `product-service`: catalog, inventory
- `cart-service`: add/remove products
- `order-service`: checkout, order history
- `payment-service`: mock payment handler
- `notification-service`: email/SMS mock

---

## 🚀 Run Project

```bash
docker-compose up --build
