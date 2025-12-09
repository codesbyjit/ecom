# 🛒 **Ecom Monorepo**

### **Nx + Next.js (Web) + Go API | Modern Full-Stack E-Commerce Starter**

A clean, scalable, and developer-friendly monorepo built using **Nx 22**, featuring:

🚀 **Next.js Frontend** (`apps/web`)
⚡ **Go Backend API** (`apps/api`)
📦 **Monorepo Architecture**
🧪 Ready for future features like auth, payments, cart, admin dashboard

---

## 📁 **Project Structure**

```
ecom/
├── web/        # Next.js 14 app
├── api/        # Go backend API
├── node_modules/
├── nx.json
├── package.json
└── README.md
```

---

## ⚙️ **Technologies Used**

### **Frontend (web)**

- Next.js 14
- React
- Tailwind CSS
- API Routes (optional)
- Nx optimized build system

### **Backend (api)**

- Go 1.22+
- net/http
- Future-ready for:

  - Fiber / Gin
  - PostgreSQL / MongoDB
  - JWT auth

### **Monorepo**

- Nx 22
- Fast builds
- Task orchestration
- Cached pipelines

---

## 🚀 **Running the Project**

### **Start Web App**

```
nx run web:dev
```

Open:
👉 [http://localhost:3000](http://localhost:3000)

---

### **Start Go API**

```
nx run api:serve
```

API runs on:
👉 [http://localhost:4000](http://localhost:4000)

---

## 🔧 **API Overview**

Located in:

```
api/
```

# 🎨 **Frontend Overview**

Locat in

```
web/
```

Includes:

- Clean Next.js starter
- Tailwind setup
- Pages & components structure

---

## 📦 **Nx Commands**

| Action            | Command            |
| ----------------- | ------------------ |
| Serve web         | `nx run web:dev`   |
| Serve API         | `nx run api:serve` |
| Build web         | `nx run web:build` |
| Show all projects | `nx show projects` |

---

## 🤝 **Contributing**

PRs, suggestions, and contributions are welcome.

---

## 📜 **License**

MIT License — free to use for any project.
