# About Project

Hi everyone! The project is a restaurant management system. The purpose here is to perform invoice/menu/food/order transactions with the manager's account.
## Description
The Restaurant Management System provides functionalities to manage restaurant menus, orders, tables, and invoices. It allows users to create, read, update, and delete (CRUD) information about the restaurant's offerings and operations. The system ensures data integrity and provides an easy-to-use interface for managing restaurant activities.


## API Endpoints
The system exposes the following API endpoints:

Menu

GET /menus - Get all menus <br>
GET /menus/:id - Get a menu by ID<br>
POST /menus - Create a new menu<br>
PATCH /menus/:id - Update a menu by ID Food
<br><br>
GET /foods - Get all foods<br>
GET /foods/:id - Get a food by ID<br>
POST /foods - Create a new food<br>
PATCH /foods/:id - Update a food by ID Order
<br><br>
GET /orders - Get all orders <br>
GET /orders/:id - Get an order by ID <br>
POST /orders - Create a new order <br>
PATCH /orders/:id - Update an order by ID Order_Item 
<br><br>
GET /orderItems - Get all order items <br>
GET /orderItems/:id - Get an order item by ID <br>
POST /orderItems - Create a new order item <br>
PATCH /orderItems/:id - Update an order item by ID Table
<br><br>
GET /tables  -  Get all tables<br>
GET /tables/:id  -  Get a table by ID<br>
POST /tables  -  Create a new table<br>
PATCH /tables/:id  -  Update a table by ID Invoice
<br><br>
GET /invoices - Get all invoices<br>
GET /invoices/:id - Get an invoice by ID<br>
POST /invoices - Create a new invoice<br>
PATCH /invoices/:id - Update an invoice by ID<br><br>

## Project diagram is here:

![Project Scheme](https://github.com/Furkanturan8/restaurant-management/blob/main/project_structure.png)
