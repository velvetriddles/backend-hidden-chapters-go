// // project_structure.txt
// ecommerce/
// ├── cmd/
// │   └── api/
// │       └── main.go
// ├── internal/
// │   ├── domain/
// │   │   ├── product.go
// │   │   ├── user.go
// │   │   ├── order.go
// │   │   └── category.go
// │   ├── usecase/
// │   │   ├── product.go
// │   │   ├── user.go
// │   │   ├── order.go
// │   │   ├── category.go
// │   │   └── interfaces.go
// │   ├── repository/
// │   │   ├── postgres/
// │   │   │   ├── product.go
// │   │   │   ├── user.go
// │   │   │   ├── order.go
// │   │   │   └── category.go
// │   │   └── interfaces.go
// │   ├── delivery/
// │   │   └── http/
// │   │       ├── handler/
// │   │       │   ├── product.go
// │   │       │   ├── user.go
// │   │       │   ├── order.go
// │   │       │   └── category.go
// │   │       ├── middleware/
// │   │       │   ├── auth.go
// │   │       │   └── logging.go
// │   │       └── router.go
// │   ├── app/
// │   │   ├── app.go
// │   │   └── dependencies.go
// │   └── service/
// │       ├── payment.go
// │       └── email.go
// ├── pkg/
// │   ├── config/
// │   │   └── config.go
// │   ├── logger/
// │   │   └── logger.go
// │   ├── database/
// │   │   └── postgres.go
// │   └── auth/
// │       └── jwt.go
// ├── migrations/
// │   ├── 001_create_users.up.sql
// │   ├── 001_create_users.down.sql
// │   ├── 002_create_products.up.sql
// │   ├── 002_create_products.down.sql
// │   ├── 003_create_orders.up.sql
// │   ├── 003_create_orders.down.sql
// │   ├── 004_create_categories.up.sql
// │   └── 004_create_categories.down.sql
// ├── configs/
// │   ├── config.yaml
// │   └── config.production.yaml
// ├── go.mod
// └── README.md

// // internal/domain/product.go
// package domain

// type Product struct {
//     ID          int64   `json:"id"`
//     Name        string  `json:"name"`
//     Description string  `json:"description"`
//     Price       float64 `json:"price"`
//     CategoryID  int64   `json:"category_id"`
//     Stock       int     `json:"stock"`
//     CreatedAt   string  `json:"created_at"`
//     UpdatedAt   string  `json:"updated_at"`
// }

// // internal/domain/user.go
// package domain

// type User struct {
//     ID        int64  `json:"id"`
//     Email     string `json:"email"`
//     Password  string `json:"-"`
//     FirstName string `json:"first_name"`
//     LastName  string `json:"last_name"`
//     CreatedAt string `json:"created_at"`
//     UpdatedAt string `json:"updated_at"`
// }

// // internal/domain/order.go
// package domain

// type Order struct {
//     ID         int64   `json:"id"`
//     UserID     int64   `json:"user_id"`
//     TotalPrice float64 `json:"total_price"`
//     Status     string  `json:"status"`
//     CreatedAt  string  `json:"created_at"`
//     UpdatedAt  string  `json:"updated_at"`
// }

// type OrderItem struct {
//     ID        int64   `json:"id"`
//     OrderID   int64   `json:"order_id"`
//     ProductID int64   `json:"product_id"`
//     Quantity  int     `json:"quantity"`
//     Price     float64 `json:"price"`
// }

// // internal/domain/category.go
// package domain

// type Category struct {
//     ID        int64  `json:"id"`
//     Name      string `json:"name"`
//     ParentID  int64  `json:"parent_id"`
//     CreatedAt string `json:"created_at"`
//     UpdatedAt string `json:"updated_at"`
// }

// // internal/repository/interfaces.go
// package repository

// import "ecommerce/internal/domain"

// type ProductRepository interface {
//     Create(product *domain.Product) error
//     GetByID(id int64) (*domain.Product, error)
//     Update(product *domain.Product) error
//     Delete(id int64) error
//     List(offset, limit int) ([]*domain.Product, error)
// }

// type UserRepository interface {
//     Create(user *domain.User) error
//     GetByID(id int64) (*domain.User, error)
//     GetByEmail(email string) (*domain.User, error)
//     Update(user *domain.User) error
//     Delete(id int64) error
// }

// type OrderRepository interface {
//     Create(order *domain.Order) error
//     GetByID(id int64) (*domain.Order, error)
//     Update(order *domain.Order) error
//     Delete(id int64) error
//     ListByUserID(userID int64) ([]*domain.Order, error)
// }

// type CategoryRepository interface {
//     Create(category *domain.Category) error
//     GetByID(id int64) (*domain.Category, error)
//     Update(category *domain.Category) error
//     Delete(id int64) error
//     List() ([]*domain.Category, error)
// }

// // internal/usecase/interfaces.go
// package usecase

// import "ecommerce/internal/domain"

// type ProductUseCase interface {
//     CreateProduct(product *domain.Product) error
//     GetProduct(id int64) (*domain.Product, error)
//     UpdateProduct(product *domain.Product) error
//     DeleteProduct(id int64) error
//     ListProducts(page, pageSize int) ([]*domain.Product, error)
// }

// type UserUseCase interface {
//     Register(user *domain.User) error
//     Login(email, password string) (string, error)
//     GetUser(id int64) (*domain.User, error)
//     UpdateUser(user *domain.User) error
//     DeleteUser(id int64) error
// }

// type OrderUseCase interface {
//     CreateOrder(order *domain.Order) error
//     GetOrder(id int64) (*domain.Order, error)
//     UpdateOrderStatus(id int64, status string) error
//     CancelOrder(id int64) error
//     ListUserOrders(userID int64) ([]*domain.Order, error)
// }

// type CategoryUseCase interface {
//     CreateCategory(category *domain.Category) error
//     GetCategory(id int64) (*domain.Category, error)
//     UpdateCategory(category *domain.Category) error
//     DeleteCategory(id int64) error
//     ListCategories() ([]*domain.Category, error)
// }

// // internal/delivery/http/handler/product.go
// package handler

// import (
//     "ecommerce/internal/usecase"
//     "ecommerce/pkg/logger"
//     "encoding/json"
//     "net/http"
//     "strconv"

//     "github.com/gorilla/mux"
// )

// type ProductHandler struct {
//     useCase usecase.ProductUseCase
//     logger  *logger.Logger
// }

// func NewProductHandler(useCase usecase.ProductUseCase, logger *logger.Logger) *ProductHandler {
//     return &ProductHandler{
//         useCase: useCase,
//         logger:  logger,
//     }
// }

// func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
//     var product domain.Product
//     if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
//         h.logger.Error("Failed to decode product", "error", err.Error())
//         http.Error(w, "Invalid request payload", http.StatusBadRequest)
//         return
//     }

//     if err := h.useCase.CreateProduct(&product); err != nil {
//         h.logger.Error("Failed to create product", "error", err.Error())
//         http.Error(w, "Failed to create product", http.StatusInternalServerError)
//         return
//     }

//     w.WriteHeader(http.StatusCreated)
//     json.NewEncoder(w).Encode(product)
// }

// func (h *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
//     vars := mux.Vars(r)
//     id, err := strconv.ParseInt(vars["id"], 10, 64)
//     if err != nil {
//         h.logger.Error("Invalid product ID", "error", err.Error())
//         http.Error(w, "Invalid product ID", http.StatusBadRequest)
//         return
//     }

//     product, err := h.useCase.GetProduct(id)
//     if err != nil {
//         h.logger.Error("Failed to get product", "error", err.Error())
//         http.Error(w, "Product not found", http.StatusNotFound)
//         return
//     }

//     json.NewEncoder(w).Encode(product)
// }

// // Implement other handler methods (UpdateProduct, DeleteProduct, ListProducts)...

// // internal/delivery/http/middleware/auth.go
// package middleware

// import (
//     "ecommerce/pkg/auth"
//     "net/http"
//     "strings"
// )

// func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
//     return func(w http.ResponseWriter, r *http.Request) {
//         authHeader := r.Header.Get("Authorization")
//         if authHeader == "" {
//             http.Error(w, "Missing auth token", http.StatusUnauthorized)
//             return
//         }

//         bearerToken := strings.Split(authHeader, " ")
//         if len(bearerToken) != 2 {
//             http.Error(w, "Invalid auth token", http.StatusUnauthorized)
//             return
//         }

//         token := bearerToken[1]
//         if err := auth.ValidateToken(token); err != nil {
//             http.Error(w, "Invalid auth token", http.StatusUnauthorized)
//             return
//         }

//         next.ServeHTTP(w, r)
//     }
// }

// // internal/delivery/http/router.go
// package http

// import (
//     "ecommerce/internal/delivery/http/handler"
//     "ecommerce/internal/delivery/http/middleware"
//     "github.com/gorilla/mux"
// )

// func NewRouter(
//     productHandler *handler.ProductHandler,
//     userHandler *handler.UserHandler,
//     orderHandler *handler.OrderHandler,
//     categoryHandler *handler.CategoryHandler,
// ) *mux.Router {
//     r := mux.NewRouter()

//     // Public routes
//     r.HandleFunc("/products", productHandler.ListProducts).Methods("GET")
//     r.HandleFunc("/products/{id}", productHandler.GetProduct).Methods("GET")
//     r.HandleFunc("/categories", categoryHandler.ListCategories).Methods("GET")
//     r.HandleFunc("/register", userHandler.Register).Methods("POST")
//     r.HandleFunc("/login", userHandler.Login).Methods("POST")

//     // Protected routes
//     r.HandleFunc("/products", middleware.AuthMiddleware(productHandler.CreateProduct)).Methods("POST")
//     r.HandleFunc("/products/{id}", middleware.AuthMiddleware(productHandler.UpdateProduct)).Methods("PUT")
//     r.HandleFunc("/products/{id}", middleware.AuthMiddleware(productHandler.DeleteProduct)).Methods("DELETE")

//     r.HandleFunc("/orders", middleware.AuthMiddleware(orderHandler.CreateOrder)).Methods("POST")
//     r.HandleFunc("/orders/{id}", middleware.AuthMiddleware(orderHandler.GetOrder)).Methods("GET")
//     r.HandleFunc("/orders/{id}/status", middleware.AuthMiddleware(orderHandler.UpdateOrderStatus)).Methods("PUT")
//     r.HandleFunc("/orders/{id}/cancel", middleware.AuthMiddleware(orderHandler.CancelOrder)).Methods("POST")
//     r.HandleFunc("/users/{id}/orders", middleware.AuthMiddleware(orderHandler.ListUserOrders)).Methods("GET")

//     return r
// }

// // cmd/api/main.go
// package main

// import (
//     "ecommerce/internal/app"
//     "ecommerce/pkg/config"
//     "ecommerce/pkg/logger"
//     "log"
//     "os"
//     "path/filepath"
// )

// func main() {
//     configPath := os.Getenv("CONFIG_PATH")
//     if configPath == "" {
//         configPath = filepath.Join("configs")
//     }

//     cfg, err := config.LoadConfig(configPath)
//     if err != nil {
//         log.Fatalf("Failed to load config: %v", err)
//     }

//     logger, err := logger.NewLogger(cfg.Logger)
//     if err != nil {
//         log.Fatalf("Failed to initialize logger: %v", err)
//     }

//     app := app.NewApp(cfg, logger)
//     if err := app.Run(); err != nil {
//         logger.Error("Failed to run app", "error", err.Error())
//         os.Exit(1)
//     }
// }

// internal/usecase/product.go
// package usecase

// import (
//     "context"
//     "ecommerce/internal/domain"
//     "ecommerce/internal/repository"
//     "errors"
// )

// type productUseCase struct {
//     repo repository.ProductRepository
// }

// func NewProductUseCase(repo repository.ProductRepository) ProductUseCase {
//     return &productUseCase{repo: repo}
// }

// func (uc *productUseCase) CreateProduct(ctx context.Context, product *domain.Product) error {
//     // Здесь можно добавить бизнес-логику, валидацию и т.д.
//     if product.Price <= 0 {
//         return errors.New("price must be positive")
//     }
//     return uc.repo.Create(ctx, product)
// }

// func (uc *productUseCase) GetProduct(ctx context.Context, id int64) (*domain.Product, error) {
//     return uc.repo.GetByID(ctx, id)
// }

// func (uc *productUseCase) UpdateProduct(ctx context.Context, product *domain.Product) error {
//     // Проверка существования продукта и другая логика
//     existingProduct, err := uc.repo.GetByID(ctx, product.ID)
//     if err != nil {
//         return err
//     }
//     if existingProduct == nil {
//         return errors.New("product not found")
//     }
//     return uc.repo.Update(ctx, product)
// }

// func (uc *productUseCase) DeleteProduct(ctx context.Context, id int64) error {
//     return uc.repo.Delete(ctx, id)
// }

// func (uc *productUseCase) ListProducts(ctx context.Context, page, pageSize int) ([]*domain.Product, error) {
//     offset := (page - 1) * pageSize
//     return uc.repo.List(ctx, offset, pageSize)
// }

// // internal/repository/postgres/product.go
// package postgres

// import (
//     "context"
//     "database/sql"
//     "ecommerce/internal/domain"
//     "fmt"
// )

// type productRepository struct {
//     db *sql.DB
// }

// func NewProductRepository(db *sql.DB) *productRepository {
//     return &productRepository{db: db}
// }

// func (r *productRepository) Create(ctx context.Context, product *domain.Product) error {
//     query := `
//         INSERT INTO products (name, description, price, category_id, stock)
//         VALUES ($1, $2, $3, $4, $5)
//         RETURNING id, created_at, updated_at`

//     err := r.db.QueryRowContext(ctx, query,
//         product.Name, product.Description, product.Price, product.CategoryID, product.Stock,
//     ).Scan(&product.ID, &product.CreatedAt, &product.UpdatedAt)

//     if err != nil {
//         return fmt.Errorf("failed to create product: %w", err)
//     }

//     return nil
// }

// func (r *productRepository) GetByID(ctx context.Context, id int64) (*domain.Product, error) {
//     query := `
//         SELECT id, name, description, price, category_id, stock, created_at, updated_at
//         FROM products
//         WHERE id = $1`

//     var product domain.Product
//     err := r.db.QueryRowContext(ctx, query, id).Scan(
//         &product.ID, &product.Name, &product.Description, &product.Price,
//         &product.CategoryID, &product.Stock, &product.CreatedAt, &product.UpdatedAt,
//     )

//     if err != nil {
//         if err == sql.ErrNoRows {
//             return nil, nil // Продукт не найден
//         }
//         return nil, fmt.Errorf("failed to get product: %w", err)
//     }

//     return &product, nil
// }

// func (r *productRepository) Update(ctx context.Context, product *domain.Product) error {
//     query := `
//         UPDATE products
//         SET name = $2, description = $3, price = $4, category_id = $5, stock = $6, updated_at = NOW()
//         WHERE id = $1`

//     _, err := r.db.ExecContext(ctx, query,
//         product.ID, product.Name, product.Description, product.Price, product.CategoryID, product.Stock,
//     )

//     if err != nil {
//         return fmt.Errorf("failed to update product: %w", err)
//     }

//     return nil
// }

// func (r *productRepository) Delete(ctx context.Context, id int64) error {
//     query := "DELETE FROM products WHERE id = $1"

//     _, err := r.db.ExecContext(ctx, query, id)
//     if err != nil {
//         return fmt.Errorf("failed to delete product: %w", err)
//     }

//     return nil
// }

// func (r *productRepository) List(ctx context.Context, offset, limit int) ([]*domain.Product, error) {
//     query := `
//         SELECT id, name, description, price, category_id, stock, created_at, updated_at
//         FROM products
//         ORDER BY id
//         LIMIT $1 OFFSET $2`

//     rows, err := r.db.QueryContext(ctx, query, limit, offset)
//     if err != nil {
//         return nil, fmt.Errorf("failed to list products: %w", err)
//     }
//     defer rows.Close()

//     var products []*domain.Product
//     for rows.Next() {
//         var p domain.Product
//         err := rows.Scan(
//             &p.ID, &p.Name, &p.Description, &p.Price,
//             &p.CategoryID, &p.Stock, &p.CreatedAt, &p.UpdatedAt,
//         )
//         if err != nil {
//             return nil, fmt.Errorf("failed to scan product: %w", err)
//         }
//         products = append(products, &p)
//     }

//     if err = rows.Err(); err != nil {
//         return nil, fmt.Errorf("error iterating product rows: %w", err)
//     }

//     return products, nil
// }

// // internal/delivery/http/handler/product.go
// package handler

// import (
//     "ecommerce/internal/domain"
//     "ecommerce/internal/usecase"
//     "ecommerce/pkg/logger"
//     "encoding/json"
//     "net/http"
//     "strconv"

//     "github.com/gorilla/mux"
// )

// type ProductHandler struct {
//     useCase usecase.ProductUseCase
//     logger  logger.Logger
// }

// func NewProductHandler(useCase usecase.ProductUseCase, logger logger.Logger) *ProductHandler {
//     return &ProductHandler{
//         useCase: useCase,
//         logger:  logger,
//     }
// }

// func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
//     var product domain.Product
//     if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
//         h.logger.Error("Failed to decode product", "error", err.Error())
//         http.Error(w, "Invalid request payload", http.StatusBadRequest)
//         return
//     }

//     if err := h.useCase.CreateProduct(r.Context(), &product); err != nil {
//         h.logger.Error("Failed to create product", "error", err.Error())
//         http.Error(w, "Failed to create product", http.StatusInternalServerError)
//         return
//     }

//     w.WriteHeader(http.StatusCreated)
//     json.NewEncoder(w).Encode(product)
// }

// func (h *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
//     vars := mux.Vars(r)
//     id, err := strconv.ParseInt(vars["id"], 10, 64)
//     if err != nil {
//         h.logger.Error("Invalid product ID", "error", err.Error())
//         http.Error(w, "Invalid product ID", http.StatusBadRequest)
//         return
//     }

//     product, err := h.useCase.GetProduct(r.Context(), id)
//     if err != nil {
//         h.logger.Error("Failed to get product", "error", err.Error())
//         http.Error(w, "Product not found", http.StatusNotFound)
//         return
//     }

//     json.NewEncoder(w).Encode(product)
// }

// func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
//     vars := mux.Vars(r)
//     id, err := strconv.ParseInt(vars["id"], 10, 64)
//     if err != nil {
//         h.logger.Error("Invalid product ID", "error", err.Error())
//         http.Error(w, "Invalid product ID", http.StatusBadRequest)
//         return
//     }

//     var product domain.Product
//     if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
//         h.logger.Error("Failed to decode product", "error", err.Error())
//         http.Error(w, "Invalid request payload", http.StatusBadRequest)
//         return
//     }
//     product.ID = id

//     if err := h.useCase.UpdateProduct(r.Context(), &product); err != nil {
//         h.logger.Error("Failed to update product", "error", err.Error())
//         http.Error(w, "Failed to update product", http.StatusInternalServerError)
//         return
//     }

//     w.WriteHeader(http.StatusOK)
// }

// func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
//     vars := mux.Vars(r)
//     id, err := strconv.ParseInt(vars["id"], 10, 64)
//     if err != nil {
//         h.logger.Error("Invalid product ID", "error", err.Error())
//         http.Error(w, "Invalid product ID", http.StatusBadRequest)
//         return
//     }

//     if err := h.useCase.DeleteProduct(r.Context(), id); err != nil {
//         h.logger.Error("Failed to delete product", "error", err.Error())
//         http.Error(w, "Failed to delete product", http.StatusInternalServerError)
//         return
//     }

//     w.WriteHeader(http.StatusNoContent)
// }

// func (h *ProductHandler) ListProducts(w http.ResponseWriter, r *http.Request) {
//     page, _ := strconv.Atoi(r.URL.Query().Get("page"))
//     pageSize, _ := strconv.Atoi(r.URL.Query().Get("page_size"))

//     if page <= 0 {
//         page = 1
//     }
//     if pageSize <= 0 {
//         pageSize = 10
//     }

//     products, err := h.useCase.ListProducts(r.Context(), page, pageSize)
//     if err != nil {
//         h.logger.Error("Failed to list products", "error", err.Error())
//         http.Error(w, "Failed to list products", http.StatusInternalServerError)
//         return
//     }

//     json.NewEncoder(w).Encode(products)
//