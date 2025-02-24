# Product Management Application

This repository contains the implementation of a full-stack application for product management with a frontend built using Vue.js and a backend powered by Go (Golang) and the Gin framework.
The directories product-management-fe and product-management-be are both managed by **Git**

### Tech stack:
- **Frontend**: JavaScript, Vue.js, Vue Router, Vuex, ElementUI
- **Backend**: Golang, Gin, Gorm
- **Database**: PostgreSQL
- **Local Deployment**: Docker, DockerCompose
- **Version Control**: Git

## Features Implemented
1. Product Creation
2. Product Update
3. Product Deletion
4. Product List (including filtering and sorting)

The full API documentation is in `./product_management_be/readme.md`

   
## Setup and Installation

### Prerequisites
Docker and Docker Compose must be installed on your machine.

### Steps to Run Locally

1. **Start Services**:
``` bash
docker-compose up --build
```
2. **Create Database and Index (if it's the first time you run the app)**
``` bash
docker exec -i product_management_db psql -U postgres -d postgres -f /custom-scripts/init.sql
```
3. **Visit Our Frontend**

visit http://localhost:8080
4. **Stop the Application**
```bash
docker-compose down
```

## Feature Demonstration
### initial empty list
![img.png](readme_images/img.png)
### product creation page
![img_1.png](readme_images/img_1.png)
#### form validation
![img_2.png](readme_images/img_2.png)
#### form complete
![img_3.png](readme_images/img_3.png)
#### product creation success
![img_4.png](readme_images/img_4.png)
#### image preview
![img_5.png](readme_images/img_5.png)
### click "update" on one item
#### update page (reuse the creation page, but have a query parameter id)
![img_6.png](readme_images/img_6.png)
#### change the price, description and image
![img_7.png](readme_images/img_7.png)
#### update result
![img_8.png](readme_images/img_8.png)
### show full description by hovering on it
![img_9.png](readme_images/img_9.png)
### click "delete" on one item
#### deletion confirmation
![img_10.png](readme_images/img_10.png)
### delete result
![img_11.png](readme_images/img_11.png)
### more products in the list
![img_12.png](readme_images/img_12.png)
#### change from 10/page to 5/page
![img_13.png](readme_images/img_13.png)
#### change page number
![img_14.png](readme_images/img_14.png)
### search products by filters
![img_15.png](readme_images/img_15.png)
![img_16.png](readme_images/img_16.png)
![img_17.png](readme_images/img_17.png)
### order by price desc
![img_18.png](readme_images/img_18.png)
### order by name asc
![img_19.png](readme_images/img_19.png)![img.png](order_name.png)
### enter the url with filters and pagination directly
enter `http://localhost:8080/#/?name=a&type=Beauty&priceMin=&priceMax=&page=1&page_size=5` in a new page

and it will show the below page and fill the search form and pagination automatically
![img_20.png](readme_images/img_20.png)![img.png](readme_images/enter_link.png)
