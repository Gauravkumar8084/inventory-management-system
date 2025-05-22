# Inventory Management System - WebApp

A simple Inventory Management System built using Go (backend), MySQL (database), and GORM (ORM).  
The system supports full CRUD operations on products, purchase orders, sales orders, and real-time inventory updates.

---

## Tech Stack

- **Backend:** GoLang  
- **ORM:** GORM  
- **Database:** MySQL  
- **Web Framework:** Gin  
- **API Testing:** Postman or curl

---

## Table of Contents

- [Getting Started](#getting-started)  
- [Prerequisites](#prerequisites)  
- [Installation](#installation)  
- [Usage](#usage)  
- [API Endpoints](#api-endpoints)  
- [Testing](#testing)  
- [Deployment](#deployment)  
- [Authors](#authors)  

---

## Getting Started

These instructions will help you set up the Inventory Management System on your local machine for development and testing purposes.

### Steps:

1. Clone the repository  
2. Configure `.env` file with your MySQL credentials  
3. Run the `schema.sql` script on your MySQL server  
4. Run the backend server using Go  

---

## Prerequisites

- Go (version 1.18 or above)  
- MySQL Server  
- Postman or curl (for API testing)  

---

## Installation

### Clone the Project

```bash
git clone https://github.com/manas284/Inventory-management-System.git
cd Inventory-management-System
Configure Environment
Create or update the .env file with your MySQL credentials:

ini
Copy
Edit
DB_USER=your_mysql_username
DB_PASSWORD=your_mysql_password
DB_HOST=localhost
DB_PORT=3306
DB_NAME=inventory
Run SQL Script
Open MySQL CLI or MySQL Workbench and run:

sql
Copy
Edit
SOURCE schema.sql;
Run the Go Server
bash
Copy
Edit
go mod tidy
go run main.go
The server should be live at: http://localhost:8080

Usage
Test the APIs using:

curl commands from the terminal

Postman (a collection is provided with the project for ease of testing)

API Endpoints
Products
Method	Endpoint	Description
GET	/api/products	List all products
POST	/api/products	Create a new product
GET	/api/products/:id	View a product by ID
PUT	/api/products/:id	Update a product
DELETE	/api/products/:id	Delete a product

Sales Orders
Method	Endpoint	Description
POST	/api/sales-orders	Create a sales order (reduces stock)

Purchase Orders
Method	Endpoint	Description
POST	/api/purchase-orders	Create a purchase order (adds stock)

Inventory
Method	Endpoint	Description
GET	/api/inventory	View current stock levels

Testing
Currently, unit tests are not included in this project.
Unit testing will be added in future releases to ensure robustness of CRUD operations and inventory updates.

Deployment
This project is set up for local development. You can deploy it to cloud platforms like:

Railway

Render

DigitalOcean

Authors
Manas Tiwari
GitHub: https://github.com/manas284

