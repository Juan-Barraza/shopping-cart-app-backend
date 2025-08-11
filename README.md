# Shopping Cart App (Prueba Técnica)

Pequeña app de carrito de compras: **API** (sin BD, datos estáticos) + **Frontend** en React.
La API expone `/products` y `/cart` (GET/POST). El front consume la API, muestra productos y permite agregarlos al carrito.
Incluye una función que calcula la mejor combinación de productos sin exceder un presupuesto.

## Estructura
- `backend/` — API en Go, carrito en memoria, productos desde JSON estático.
- `frontend/` — React, consume la API y renderiza lista y carrito.

## Requisitos
- Go 1.20+

## Cómo correr
### Backend
```bash
go mod tidy
go run main.go
# API en: http://localhost:8080
