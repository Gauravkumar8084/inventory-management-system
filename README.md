# Inventory Management System - WebApp

A simple Inventory Management System built using Go (backend), MySQL (database), and GORM (ORM). The system allows full CRUD operations on products, purchase orders, sales orders, and real-time inventory updates.

---

## Tech Stack

### Backend

* GoLang
* GORM (ORM)
* MySQL (Database)
* Gin (Web Framework)

### API Testing

* Postman or curl

---

## Table of Contents

* [Getting Started](#getting-started)
* [Prerequisites](#prerequisites)
* [Installation](#installation)
* [Usage](#usage)
* [API Endpoints](#api-endpoints)
* [Deployment](#deployment)
* [Authors](#authors)
* [Sample Demo - Video Link](#sample-demo---video-link)

---

## Getting Started

These instructions will help you set up the Inventory Management System on your local machine.

### Steps:

1. Clone the Repository
2. Configure `.env` for MySQL credentials
3. Run `schema.sql` in your MySQL instance
4. Run the backend server using Go

---

## Prerequisites

* **Go (1.18 or above)**
* **MySQL Server**
* **Postman or curl** (for API testing)

---

## Installation

### Clone Project

```bash
git clone https://github.com/your-username/inventory-management-go.git
cd inventory-management-go
```

### Configure Environment

Update the `.env` file:

```env
DB_USER=your_mysql_username
DB_PASSWORD=your_mysql_password
DB_HOST=localhost
DB_PORT=3306
DB_NAME=inventory
```

### Run SQL Script

Open MySQL CLI or Workbench and run:

```sql
SOURCE schema.sql;
```

### Run Go Server

```bash
go mod tidy
go run main.go
```

Your server should be live at: `http://localhost:8080`

---

## Usage

Use the following tools:

* **curl** commands from terminal
* **Postman** to test all APIs (collection provided)

---

## API Endpoints

### Products

* `GET /api/products` – List all products
* `POST /api/products` – Create product
* `GET /api/products/:id` – View product by ID
* `PUT /api/products/:id` – Update product
* `DELETE /api/products/:id` – Delete product

### Sales Orders

* `POST /api/sales-orders` – Create sales order (reduces stock)

### Purchase Orders

* `POST /api/purchase-orders` – Create purchase order (adds stock)

### Inventory

* `GET /api/inventory` – View current stock levels

---

## Deployment

This project is currently set up for local development. Deployment can be done to platforms like:

* **Railway**
* **Render**
* **DigitalOcean**

---

## Authors

**Manas Tiwari** – [GitHub: manas284]((https://github.com/manas284))

---

