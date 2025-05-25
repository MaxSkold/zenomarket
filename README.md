# Scalable E-commerce Platform — Microservices Architecture (Go + Docker)

This repository contains my solution to the [Scalable E-commerce Platform Project](https://roadmap.sh/projects/scalable-ecommerce-platform) from roadmap.sh.

✅ [Project URL](https://roadmap.sh/projects/scalable-ecommerce-platform)

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
