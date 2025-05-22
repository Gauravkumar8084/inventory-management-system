# Inventory Management System - WebApp

A simple, robust Inventory Management System built using Go (backend), MySQL (database), and GORM (ORM). This system allows full CRUD operations on Products, Purchase Orders, Sales Orders, and provides real-time Inventory tracking and updates.

---

## Tech Stack

- **Backend:** Go (Golang)
- **ORM:** GORM
- **Database:** MySQL
- **Web Framework:** Gin
- **API Testing Tools:** Postman / curl

---

## Table of Contents

- [Getting Started](#getting-started)  
- [Prerequisites](#prerequisites)  
- [Installation](#installation)  
- [Configuration](#configuration)  
- [Database Setup](#database-setup)  
- [Running the Application](#running-the-application)  
- [API Endpoints](#api-endpoints)  
- [Usage Examples](#usage-examples)  
- [Error Handling](#error-handling)  
- [Deployment](#deployment)  
- [Authors](#authors)  
- [Sample Demo - Video Link](#sample-demo---video-link)  

---

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing.

### Step-by-step Overview

1. **Clone the repository** from GitHub  
2. **Set up your environment variables** for MySQL connection  
3. **Run the database schema script** to create tables  
4. **Start the Go backend server**  
5. **Test the APIs** using Postman or curl  

---

## Prerequisites

Before you begin, ensure you have met the following requirements:

- **Go** (version 1.18 or above) installed. You can download it from [golang.org](https://golang.org/dl/)  
- **MySQL Server** installed and running on your machine  
- **Postman** or `curl` for API testing  
- Basic knowledge of terminal/command prompt usage

---

## Installation

### 1. Clone the Project

```bash
git clone https://github.com/manas284/Inventory-management-System.git
cd Inventory-management-System
2. Configuration
Create a .env file in the root directory (you may copy .env.example if available):

env
Copy
Edit
DB_USER=your_mysql_username
DB_PASSWORD=your_mysql_password
DB_HOST=localhost
DB_PORT=3306
DB_NAME=inventory
Note: Replace your_mysql_username and your_mysql_password with your MySQL credentials.

Database Setup
Open your MySQL client (CLI or Workbench).

Create the database if it doesn't exist:

sql
Copy
Edit
CREATE DATABASE inventory;
USE inventory;
Run the schema script to create tables:

sql
Copy
Edit
SOURCE schema.sql;
The schema.sql file contains all the necessary commands to create the tables and constraints.

Running the Application
1. Install dependencies
bash
Copy
Edit
go mod tidy
2. Start the backend server
bash
Copy
Edit
go run main.go
You should see the server start on port 8080 by default.

API Endpoints
Products
Method	Endpoint	Description	Request Body Example
GET	/api/products	Get all products	N/A
POST	/api/products	Create a new product	{ "name": "Laptop", "description": "15 inch", "price": 1200, "quantity": 10 }
GET	/api/products/:id	Get a product by ID	N/A
PUT	/api/products/:id	Update product details	{ "name": "Laptop Pro", "price": 1500 }
DELETE	/api/products/:id	Delete a product by ID	N/A

Sales Orders
Method	Endpoint	Description	Request Body Example
POST	/api/sales-orders	Create a sales order (reduces stock)	{ "product_id": 1, "quantity": 2 }

Purchase Orders
Method	Endpoint	Description	Request Body Example
POST	/api/purchase-orders	Create a purchase order (adds stock)	{ "product_id": 1, "quantity": 5, "supplier": "ABC Supplies" }

Inventory
Method	Endpoint	Description	Request Body Example
GET	/api/inventory	Get current stock levels	N/A

Usage Examples
Create Product
bash
Copy
Edit
curl -X POST http://localhost:8080/api/products \
-H "Content-Type: application/json" \
-d '{"name":"Monitor","description":"24 inch HD","price":200,"quantity":15}'
Get All Products
bash
Copy
Edit
curl http://localhost:8080/api/products
Create Sales Order (reduce stock)
bash
Copy
Edit
curl -X POST http://localhost:8080/api/sales-orders \
-H "Content-Type: application/json" \
-d '{"product_id":1,"quantity":3}'
Get Inventory
bash
Copy
Edit
curl http://localhost:8080/api/inventory
Error Handling
All API responses return appropriate HTTP status codes:

200 OK for successful requests

201 Created for successful POST operations

400 Bad Request for validation errors or malformed inputs

404 Not Found when requested resources don’t exist

500 Internal Server Error for unexpected failures

Error responses have a JSON body with a message, e.g.:

json
Copy
Edit
{
  "error": "Product not found"
}

Deployment
This project is primarily set up for local development.

You can deploy it to any cloud platform that supports Go apps, such as:

Railway

Render

DigitalOcean

Ensure environment variables and database connections are properly configured in the deployment environment.

Authors
Manas Tiwari

GitHub: https://github.com/manas284

