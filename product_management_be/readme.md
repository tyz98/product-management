
# local development
## run postgresql on docker
```
docker pull postgres:17
docker volume create postgre-data
docker run -id --name=postgresql -v postgre-data:/var/lib/postgresql/data -p 5432:5432 -e POSTGRES_PASSWORD=<your_password_here> postgres:17
```
## install go packages
go version: 1.20
```
go get github.com/gin-gonic/gin
go get github.com/jinzhu/gorm
go get github.com/jinzhu/gorm/dialects/postgres
go get github.com/spf13/viper
```
## run the backend
```
go build -o ./product-backend ./cmd/main.go
./product-backend
```

# API Documentation
## Base URL
http://localhost:9000/api
## Endpoints
### **1. Get all products**
- **URL**: `/api/product`
- **Method**: `GET`
- **Description**: Retrieves a list of products with pagination, filtering, and sorting.
- **Query Parameters**:
    - `page` (default: `1`).
    - `page_size` (default: `10`).
    - `name`: Filter by product name.
    - `type`: Filter by product type.
    - `price_min`: Filter by minimum price.
    - `price_max`: Filter by maximum price.
    - `sort_by`: The field to sort by (default: `id`).
    - `sort_order`: The sorting order (`asc` or `desc`, default: `asc`).
- **Response**:
    - `products`: A list of products matching the criteria.
    - `total`: The total number of products available based on the filtering criteria.

    ```json
    {
      "products": [
        {
          "id": 1,
          "name": "Product A",
          "type": "Electronics",
          "price": 100,
          "description": "Description of Product A",
          "image_path": "uploads/20250222_150405_imageA.png"
        },
        {
          "id": 2,
          "name": "Product B",
          "type": "Clothing",
          "price": 50.12,
          "description": "Description of Product B",
          "image_path": "uploads/20250222_150405_imageB.png"
        }
      ],
      "total": 2
    }
    ```
### **2. Create a product**
- **URL**: `/api/product`
- **Method**: `POST`
- **Description**: Creates a new product.
- **Request Body** (JSON):
    - `name`: The name of the product.
    - `type`: The type of the product.
    - `price`: The price of the product.
    - `description`: A description of the product.
    - `image_path`: The path to the product's image.(returned by api 6)
  - **Response**: The created product.

    ```json
    {
        "id": 3,
        "name": "Product C",
        "type": "Furniture",
        "price": 150.8,
        "description": "Description of Product C",
        "image_path": "uploads/20250222_150405_imageC.png"
    }
    ```
### **3. Get a product by ID**
- **URL**: `/api/product/:id`
- **Method**: `GET`
- **Description**: Retrieves a product by its ID.
- **URL Parameters**:
    - `id`: The ID of the product.
- **Response**: The requested product.

    ```json
    {
      "id": 1,
      "name": "Product A",
      "type": "Electronics",
      "price": 100,
      "description": "Description of Product A",
      "image_path": "uploads/20250222_150405_imageA.png"
    }
    ```
### **4. Update a product by ID**
- **URL**: `/api/product/:id`
- **Method**: `PUT`
- **Description**: Updates a product by its ID.
- **URL Parameters**:
    - `id`: The ID of the product.
- **Request Body** (JSON):
    - `name`: The updated name of the product.
    - `type`: The updated type of the product.
    - `price`: The updated price of the product.
    - `description`: The updated description of the product.
    - `image_path`: The updated image_path of the product.
- **Response**: The updated product.

    ```json
    {
      "id": 1,
      "name": "Updated Product A",
      "type": "Electronics",
      "price": 120,
      "description": "Updated description of Product A",
      "image_path": "uploads/20250222_150405_updated_imageA.png"
    }
    ```
### **5. Delete a product by ID**
- **URL**: `/api/product/:id`
- **Method**: `DELETE`
- **Description**: Deletes a product by its ID.
- **URL Parameters**:
    - `id`: The ID of the product.
- **Response**: `null`
### **6. Upload an image for a product**
- **URL**: `/api/product/upload`
- **Method**: `POST`
- **Description**: Uploads an image and returns the image URL.
- **Request Body** (Form-data):
    - `file`: The image file to be uploaded.
- **Response**:
    - `url`: The URL of the uploaded image.

    ```json
    {
      "url": "uploads/20250222_150405_imageA.png"
    }
    ```
### **7. Get all product types**
- **URL**: `/api/product/type`
- **Method**: `GET`
- **Description**: Retrieves a list of allowed product types.
- **Response**:
  - `types`: A list of allowed product types.

  ```json
  {
    "types": [
      "Electronics",
      "Clothing",
      "Beauty",
      "Books",
      "Jewelry",
      "Furniture",
      "Grocery"
    ]
  }
  ```
