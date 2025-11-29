# 🧩 **Online Test Platform — Microservices Architecture**

### 🎯 **Project Overview**

This project is a fully containerized **microservices-based online testing platform** built for managing tests, sections, students, and automated result processing.
Each service is independently deployable and communicates via **gRPC**, while the system ensures scalability, isolation, and maintainability across all components.

The goal is to provide a modern, high-performance backend for education platforms — enabling admins to create, manage, and analyze student test results in real-time.

---

## ⚙️ **Architecture Overview**

* **Static Gateway** serves as the single entry point for all external requests.
* **Auth Service** manages authentication and JWT token generation.
* **Admin & User Services** store and validate entity data.
* **Test Services** handle question banks, section management, and dynamic test generation.
* **Queue Job Service** automates scoring and result processing.
* **Filter Service** provides flexible filtering and reporting across users and tests.

All services are **Dockerized**, follow **Clean Architecture principles**, and integrate using **gRPC**, **Redis**, and **RabbitMQ** for asynchronous communication.

---

## 🧱 **Microservices**

### 1️⃣ **Static Gateway (Entry Point)**

* Entry point for all HTTP requests.
* Verifies JWT tokens for authentication.
* Implements Redis-based rate limiting.
* Routes incoming traffic to appropriate gRPC services.

🧰 **Tech Stack:** `Go (Gin)` · `gRPC Client` · `Redis` · `JWT` · `Docker`

---

### 2️⃣ **Auth Service (Admin)**

* Handles admin registration and login.
* Generates and validates JWT tokens.
* Verifies admin credentials via **Admin Service**.

🧰 **Tech Stack:** `Go (Gin)` · `MySQL` · `gRPC` · `bcrypt` · `JWT` · `Redis (Session)`

---

### 3️⃣ **Admin Service**

* Stores admin credentials (username, password).
* Provides gRPC endpoints for user verification.
* Used exclusively by **Auth Service** for validation.

🧰 **Tech Stack:** `Go` · `MySQL` · `gRPC Server` · `bcrypt`

---

### 4️⃣ **User Service**

* Manages student data and profile information.
* Supports full CRUD operations.
* Shares user data with **Filter Service**.

🧰 **Tech Stack:** `Go` · `MySQL` · `GORM` · `gRPC Server` · `Validator.v10`

---

### 5️⃣ **Test & Section Service**

* Handles test creation, sections, and question management.
* Manages multiple-choice options and correct answers.
* Provides `GetFullSectionStructure` gRPC method to retrieve nested test structures.

🧰 **Tech Stack:** `Go` · `MySQL` · `GORM` · `Validator.v10` · `gRPC Server`

---

### 6️⃣ **Test Link Generator**

* Generates unique links for specific test sections.
* Temporarily stores test data in Redis.
* Allows users to access tests using a shared link.

**Endpoints:**

* `POST /api/v1/test/generate` → create test & return access link
* `GET /api/v1/test/:key` → fetch test data from Redis

🧰 **Tech Stack:** `Go (Gin)` · `Redis` · `UUID` · `gRPC Client`

---

### 7️⃣ **Queue Job Service**

* Processes test submissions asynchronously.
* Automatically calculates scores and sends results to **User Service**.
* Uses RabbitMQ + asynq for background job handling.

🧰 **Tech Stack:** `Go` · `RabbitMQ` · `asynq` · `gRPC Client` · `JSON Schema`

---

### 8️⃣ **Filter Service**

* Filters users and tests by date, year, or section.
* Provides analytics and reporting features.
* Fetches user data via gRPC from **User Service**.

🧰 **Tech Stack:** `Go (Gin)` · `REST API` · `MySQL` · `gRPC Client (User Service)`

---

## 🐳 **Deployment & Infrastructure**

* Each service is containerized with **Docker Compose**.
* Uses **Redis** for caching, rate limiting, and session storage.
* **RabbitMQ** handles asynchronous communication between services.
* Future plan: integrate **Consul** for service discovery and health checks.


---