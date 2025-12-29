# Project Based Internship Rakamin x Evermos - Backend Web Developer

## Description
This project is part of the **Project Based Internship (PBI) Rakamin x Evermos** for the Backend Web Developer role. It simulates the backend architecture of a social commerce platform like Evermos. The system is built using **Golang**, **Fiber** (Web Framework), and **MySQL**. It follows **Clean Architecture** principles to ensure scalability, maintainability, and testability.

The platform allows users to register (which automatically creates a store), manage their profile, manage addresses, view products, and make transactions. Admins have special privileges to manage categories and products.

## 1. Auth Service
Handles user registration and authentication. When a user registers, a corresponding Store (Toko) is automatically created for them.

**Files:**
- Router: `app/router/auth_router/index.go`
- Controller: `app/controller/auth_controller/auth_controller.go`
- Request: `app/request/auth_request.go`
- Service: `app/service/auth_service/`
- Repository: `app/repository/auth_repository/`

**Routes:**

| Method | Endpoint | Description | Auth Required |
| :--- | :--- | :--- | :--- |
| `POST` | `/auth/register` | Register a new user and create a store. | No |
| `POST` | `/auth/Login` | Login and retrieve JWT token. | No |

**Example Request (Register):**
```json
{
  "nama": "John Doe",
  "email": "john@example.com",
  "notelp": "081234567890",
  "password": "password123",
  "is_admin": false
}
```

**Example Request (Login):**
```json
{
  "email": "john@example.com",
  "password": "password123"
}
```

---

## 2. Account Management Service
Allows users to view and update their personal profile information.

**Files:**
- Router: `app/router/user_router/index.go`
- Controller: `app/controller/user_controller/index.go`
- Request: `app/request/user_request.go`
- Service: `app/service/user_service/`
- Repository: `app/repository/user_repository/`

**Routes:**

| Method | Endpoint | Description | Auth Required |
| :--- | :--- | :--- | :--- |
| `GET` | `/user/` | Get current user profile. | Yes (JWT) |
| `PUT` | `/user/update-user` | Update user profile details. | Yes (JWT) |

**Example Request (Update User):**
```json
{
  "nama": "John Doe Updated",
  "password": "newpassword123",
  "notelp": "081234567899",
  "tanggal_lahir": "1990-01-01T00:00:00Z",
  "jenis_kelamin": "Laki-laki",
  "tentang": "Reseller of Muslim products",
  "pekerjaan": "Entrepreneur"
}
```

---

## 3. Store (Toko) Service
Manages the user's store. The store is created automatically upon registration. Users can update their store's photo.

**Files:**
- Router: `app/router/toko_router/index.go`
- Controller: `app/controller/toko_controller/toko_controller.go`
- Request: `app/request/toko_request.go`
- Service: `app/service/toko_service/`
- Repository: `app/repository/toko_repository/`

**Routes:**

| Method | Endpoint | Description | Auth Required |
| :--- | :--- | :--- | :--- |
| `GET` | `/toko/` | Get current user's store info. | Yes (JWT) |
| `PUT` | `/toko/photo` | Update store photo (Upload file). | Yes (JWT) |

**Example Request (Update Photo):**
- **Content-Type**: `multipart/form-data`
- **Body**: `url_foto` (File)

---

## 4. Address (Alamat) Service
Manages shipping addresses for the user.

**Files:**
- Router: `app/router/alamat_router/index.go`
- Controller: `app/controller/alamat_controller/index.go`
- Request: `app/request/alamat_request.go`
- Service: `app/service/alamat_service/`
- Repository: `app/repository/alamat_repository/`

**Routes:**

| Method | Endpoint | Description | Auth Required |
| :--- | :--- | :--- | :--- |
| `POST` | `/alamat/add-alamat` | Add a new address. | Yes (JWT) |
| `GET` | `/alamat/` | Get all addresses for the user. | Yes (JWT) |
| `GET` | `/alamat/:id` | Get a specific address by ID. | Yes (JWT) |
| `PUT` | `/alamat/update-alamat/:id` | Update an address. | Yes (JWT) |
| `DELETE` | `/alamat/delete-alamat/:id` | Delete an address. | Yes (JWT) |

**Example Request (Add/Update Address):**
```json
{
  "judul_alamat": "Home",
  "nama_penerima": "John Doe",
  "notelp": "081234567890",
  "detail_alamat": "Jl. Sudirman No. 1, Jakarta"
}
```

---

## 5. Category Service
Manages product categories. **Only Admins** can create, update, or delete categories.

**Files:**
- Router: `app/router/category_router/index.go`
- Controller: `app/controller/category_controller/index.go`
- Request: `app/request/category_request.go`
- Service: `app/service/category_service/`
- Repository: `app/repository/category_repository/`

**Routes:**

| Method | Endpoint | Description | Auth Required |
| :--- | :--- | :--- | :--- |
| `GET` | `/category/` | Get all categories. | Yes (JWT) |
| `GET` | `/category/:id` | Get category by ID. | Yes (JWT) |
| `POST` | `/category/add-category` | Add a new category. | Yes (JWT, Admin) |
| `PUT` | `/category/update-category/:id` | Update a category. | Yes (JWT, Admin) |
| `DELETE` | `/category/delete-category/:id` | Delete a category. | Yes (JWT, Admin) |

**Example Request (Add/Update Category):**
```json
{
  "nama_category": "Fashion Muslim"
}
```

---

## 6. Product Service
Manages products in the system. Admins manage the product listings.

**Files:**
- Router: `app/router/product_router/index.go`, `app/router/foto_router/index.go`
- Controller: `app/controller/product_controller/index.go`, `app/controller/foto_controller/index.go`
- Request: `app/request/product_request.go`
- Service: `app/service/product_service/`, `app/service/foto_service/`
- Repository: `app/repository/product_repository/`, `app/repository/foto_repository/`

**Routes:**

| Method | Endpoint | Description | Auth Required |
| :--- | :--- | :--- | :--- |
| `GET` | `/product/` | Get all products. | Yes (JWT) |
| `GET` | `/product/:id` | Get product by ID. | Yes (JWT) |
| `GET` | `/product/search` | Search products by category. | Yes (JWT) |
| `POST` | `/product/add-product` | Add a new product. | Yes (JWT, Admin) |
| `PUT` | `/product/update-product/:id` | Update a product. | Yes (JWT, Admin) |
| `DELETE` | `/product/delete-product/:id` | Delete a product. | Yes (JWT, Admin) |
| `POST` | `/foto/add-foto/:id` | Add a photo to a product. | Yes (JWT, Admin) |

**Example Request (Add/Update Product):**
```json
{
  "nama_produk": "Gamis Syari",
  "harga_reseller": "150000",
  "harga_konsumen": "200000",
  "stok": 100,
  "deskripsi": "Gamis bahan katun jepang",
  "IDCategory": "uuid-category-id"
}
```

---

## 7. Transaction Service
Handles purchasing of products.

**Files:**
- Router: `app/router/trx_router/index.go`
- Controller: `app/controller/trx_controller/index.go`
- Request: `app/request/trx_request.go`
- Service: `app/service/trx_service/`
- Repository: `app/repository/trx_repository/`

**Routes:**

| Method | Endpoint | Description | Auth Required |
| :--- | :--- | :--- | :--- |
| `POST` | `/trx/add-trx` | Create a new transaction. | Yes (JWT) |
| `GET` | `/trx/:id` | Get transaction details. | Yes (JWT, Admin) |

**Example Request (Create Transaction):**
```json
{
  "id_alamat": "uuid-alamat-id",
  "method_bayar": "transfer",
  "items": [
    {
      "id_product": "uuid-product-id-1",
      "id_toko": "uuid-toko-id-1",
      "kuantitas": 2
    },
    {
      "id_product": "uuid-product-id-2",
      "id_toko": "uuid-toko-id-2",
      "kuantitas": 1
    }
  ]
}
```

---

## 8. Beyond the Code — Performance Benchmark & Trade-off Analysis

### Performance Benchmark
To ensure the system handles high traffic efficiently, benchmarking tools like **k6** or **Apache Benchmark (ab)** can be used. Key metrics to monitor include:
- **Requests Per Second (RPS)**: How many concurrent requests the API can handle.
- **Latency**: The time taken for the server to respond (P95 and P99 latency are critical).
- **Resource Usage**: CPU and Memory consumption during load tests.

*Example Scenario*: Simulating 1000 concurrent users performing login and product search operations to identify bottlenecks in database queries or application logic.

### Trade-off Analysis
- **Monolithic vs Microservices**: This project uses a **Monolithic** architecture (Clean Architecture).
  - *Trade-off*: Monoliths are easier to develop, deploy, and debug initially. However, scaling specific components (like just the Product Service) is harder compared to Microservices.
- **SQL (MySQL) vs NoSQL**:
  - *Trade-off*: MySQL provides strong ACID compliance, which is crucial for transactional data (orders, payments). However, it might be less flexible for unstructured product data compared to NoSQL (like MongoDB), and horizontal scaling is more complex.
- **ORM vs Raw SQL**:
  - *Trade-off*: Using an ORM (like Gorm) speeds up development and prevents SQL injection but may introduce performance overhead for complex queries compared to optimized Raw SQL.

---

## 9. Conculsion

The Rakamin x Evermos Project (Backend Web Developer) simulates the backend architecture of an Evermos-style social e-commerce platform. This system is built with Golang, Fiber, and MySQL, and implements the Clean Architecture principle to ensure scalability, maintainability, and ease of testing. Its main functionalities include user registration (automatically creating stores), profile and address management, product search, and transaction processing; while administrators have special rights to manage categories and products. This analysis will thoroughly discuss fundamental aspects such as system architecture (Clean Architecture and monolithic design), software engineering principles (SOLID, separation of concerns), technology usage (Golang, Fiber, MySQL), and relevance to the social e-commerce context in Indonesia. In addition, the technological trade-offs (ORM vs raw SQL, SQL vs NoSQL), potential bottlenecks, performance testing strategies (benchmarking), and the maintainability, extensibility, and production readiness of this project will be examined.

### System Architecture: Clean Architecture and Monolithic

The architecture of this project adopts the Clean Architecture pattern, which separates code into structured layers (router, controller, service, repository) for each service domain (e.g. Auth, User, Shop). This approach is in line with the Separation of Concerns (SoC) principle, where each part handles one business responsibility separately. This allows independent modules to be tested and developed without affecting each other, facilitating the application of SOLID principles such as Single Responsibility and Dependency Inversion. For example, the Auth service directory structure shows separate file groupings for routers, controllers, services, repositories, and requests, illustrating efforts to comply with SRP and SoC.

### Technical Analysis: Golang, Fibre, and Technology Trade-offs

At the technological level, the choice of Golang and Fibre reflects a focus on performance and the ability to handle high traffic. In fact, studies show that Go web applications can achieve high throughput with efficient CPU and memory usage. Thus, the Golang+Fiber stack generally supports low latency and relatively concise resource usage. However, at extreme traffic levels (e.g., 100,000 QPS requests), CPU, memory, and I/O aspects can become major bottlenecks. 

MySQL was chosen as the database due to its reliability and ACID compliance, which are important for transactional data (orders, payments). MySQL is also widely used in the local industry, providing a mature ecosystem for e-commerce needs. However, MySQL is vertically scalable (increasing server capacity) rather than horizontally scalable (adding nodes). Thus, if the data volume increases dramatically, replication and partitioning settings become a challenge. The project README recognises this as a trade-off: MySQL is strong on transactional data consistency, but less flexible for unstructured data (products with rich attributes) and more difficult to scale horizontally. NoSQL alternatives such as MongoDB can handle semi-structured data more flexibly, but at the cost of sacrificing the essential ACID guarantees in the transaction module.

The choice between ORM (Gorm) and raw SQL queries is also recognised as a critical trade-off. ORM simplifies development by abstracting SQL, handling schema migrations and automatically preventing SQL injection. On the other hand, ORMs can cause performance overhead, especially when performing complex queries or handling large data. Studies show that Gorm adds computational cost to bulk insert/update operations or complex multi-table queries. Therefore, for latency-sensitive queries (e.g., product searches with category joins), caution is needed: developers may have to write raw SQL or tune indexes for efficiency. Additionally, other potential inefficiencies include product image uploads (processed at the /foto/add-foto endpoint), which can be an I/O-intensive step, as well as caching requirements (Redis/memcached) that are not mentioned in the documentation. Overall, the chosen technology stack supports rapid development and good performance, but special attention to database optimisation and ORM usage is crucial to avoid performance bottlenecks.

### Relevance to Social E-commerce in Indonesia
From an industry perspective, the architecture and features of this system are highly relevant to Evermos' social commerce business model. Evermos itself is a social e-commerce platform focused on halal products in Indonesia, which facilitates resellers (users) to open online stores without capital. The fact that 90% of Indonesia's population is Muslim has enabled Evermos to grow rapidly by targeting this market. This project's system simplifies the registration process so that new users automatically have a shop (the shop is created during registration), fulfilling the concept of resellers without capital. Users (resellers) can manage their profiles, addresses, view products, and make transactions – mimicking the active role of resellers who sell to their social networks. Meanwhile, the admin feature that allows adding/editing/deleting categories and products reflects the platform's curation and oversight functions. Thus, this system balances two key roles in social e-commerce: resellers/users who interact directly with commercial content, and admins as infrastructure providers (categories, products, transaction reports). This is in line with Evermos' strategy of relying on resellers to ‘help customers choose products more efficiently’, as well as its multi-seller marketplace model that includes local MSMEs. Overall, the specifications in the README (e.g. JWT authentication, profile management, shop photos) are in line with the basic requirements of a commercial social platform in Indonesia.

### Maintainability, Extensibility, and Production Readiness
From a maintainability perspective, the Clean Architecture-based code structure is highly supportive. Each separate layer (controller, service, repository) facilitates continuous development and the addition of new features. Software engineering theory explains that systems that adhere to SOLID principles tend to be easier to maintain, more flexible, extensible, and easier to test. For example, because each service module has a single responsibility and depends on an interface, adding new features should not break the old code. Complete project documentation (e.g., API routes and JSON payload examples) also speeds up the understanding of the next team. Clean Architecture promises the ability to test the business logic layer without being tied to a specific database or framework, which is excellent for setting up unit/integration testing. 


