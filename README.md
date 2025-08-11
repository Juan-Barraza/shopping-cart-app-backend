# Shopping Cart App (Prueba Técnica)

Pequeña app de carrito de compras: **API** (sin BD, datos estáticos) + **Frontend** en React.
La API expone `/products` y `/cart` (GET/POST). El front consume la API, muestra productos y permite agregarlos al carrito.
Incluye una función que calcula la mejor combinación de productos sin exceder un presupuesto.

## Estructura
- `backend/` — API en Go, carrito en memoria, productos desde JSON estático.
- `frontend/` — React, consume la API y renderiza lista y carrito.

## Requisitos
- Go 1.20+
  
##Endpoints
- GET /products:
 Response ([
  { "id": 1, "name": "Producto 1", "price": 100 },
  { "id": 2, "name": "Producto 2", "price": 200 }
])
- POST /cart:
Request (
{ "id": 1 } )
Response 
   ( "items": [
        {
            "product": {
                "id": 1,
                "name": "Producto 1",
                "price": 60
            },
            "quantity": 2
        }
    "total": 60
})
- GET /cart:
Response 
   ( "items": [
        {
            "product": {
                "id": 1,
                "name": "Producto 1",
                "price": 60
            },
            "quantity": 2
        },
        {
            "product": {
                "id": 4,
                "name": "Producto 4",
                "price": 70
            },
            "quantity": 1
        }
    ],
    "total": 190
})

## Cómo correr
### Backend
```bash
go mod tidy
go run main.go
# API en: http://localhost:8080
