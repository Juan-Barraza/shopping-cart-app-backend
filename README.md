# Shopping Cart API

API REST desarrollada en Go para un sistema de carrito de compras. Maneja productos estáticos desde JSON y carrito en memoria con algoritmo de optimización de presupuesto.


## 📋 Descripción de la solución

Esta API REST proporciona un sistema de carrito de compras completo con las siguientes características:

- **Gestión de productos**: Catálogo estático cargado desde archivo JSON
- **Carrito en memoria**: Sistema de carrito temporal por sesión
- **Algoritmo de optimización**: Calcula la mejor combinación de productos dentro de un presupuesto determinado
- **CORS habilitado**: Permite requests desde cualquier origen para facilitar el desarrollo

## 📋 Endpoints

### GET /products
Obtiene todos los productos disponibles

**Response:**
```json
[
  { "id": 1, "name": "Producto 1", "price": 100 },
  { "id": 2, "name": "Producto 2", "price": 200 },
  { "id": 3, "name": "Producto 3", "price": 50 }
]
```
### POST /cart
Agrega un producto al carrito
**Request**:
``` json
{ "id": 1 }
```
**Response:**
``` json
{
  "items": [
    {
      "product": { "id": 1, "name": "Producto 1", "price": 100 },
      "quantity": 1
    }
  ],
  "total": 100
}
```

### GET /cart
Obtiene el carrito actual con todos los productos y total

**Response:**

``` json

{
  "items": [
    {
      "product": { "id": 1, "name": "Producto 1", "price": 100 },
      "quantity": 2
    },
    {
      "product": { "id": 3, "name": "Producto 3", "price": 50 },
      "quantity": 1
    }
  ],
  "total": 250
}
```

### 🛠️ Instalación y ejecución
Requisitos
- Go 1.20 o superior

Pasos para ejecutar localmente

``` bash# Clonar el repositorio
git clone https://github.com/Juan-Barraza/shopping-cart-app-backend.git
cd shopping-cart-backend

# Instalar dependencias
go mod tidy

# Ejecutar el servidor
go run main.go

# La API estará disponible en: http://localhost:8080
```
### 🌐 Frontend relacionado
El frontend de esta aplicación está disponible en: https://shopping-cart-app-frontend.vercel.app/

### 🚀 Despliegue
Esta API está desplegada en Railway. Para desplegar tu propia instancia:

1. Fork este repositorio
2. Conecta tu cuenta de Railway con GitHub
3. Selecciona este repositorio
4. Railway detectará automáticamente que es un proyecto Go
5. El despliegue se realizará automáticamente

### 📝 Notas adicionales
- Los datos del carrito se almacenan en memoria, por lo que se pierden al reiniciar el servidor
- Los productos se cargan desde products.json al inicializar la aplicación


