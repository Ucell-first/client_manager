# Client Manager - Foydalanuvchilarni Boshqarish Tizimi

Bu loyiha Go tilida yozilgan web-asoslangan CRUD (Create, Read, Update, Delete) tizimi bo'lib, foydalanuvchilarni boshqarish uchun mo'ljallangan.

## 🚀 Texnologiyalar

- **Backend**: Go 1.24.4
- **Database**: PostgreSQL 13
- **Frontend**: HTML Templates + Bulma CSS Framework
- **Deployment**: Docker & Docker Compose

## 📁 Loyiha Strukturasi

```
client_manager/
├── main.go                              # Asosiy fayl
├── go.mod                              # Go dependencies
├── go.sum                              # Dependencies checksums
├── .env                                # Environment variables
├── Dockerfile                          # Docker build file
├── docker-compose.yml                  # Docker Compose konfiguratsiya
├── README.md                           # Bu fayl
│
├── configuration/                      # Konfiguratsiya paketlari
│   ├── configuration.go
│   ├── errors.go
│   ├── loader.go
│   ├── postgres.go
│   └── server.go
│
├── internal/delivery/http_cms/         # HTTP handlers va templates
│   ├── handler.go                      # Asosiy handler
│   ├── users.go                        # User CRUD operations
│   ├── assets/
│   │   └── style.css                   # CSS stillari
│   └── templates/                      # HTML templates
│       ├── base.html                   # Asosiy template
│       ├── list.html                   # Foydalanuvchilar ro'yxati
│       ├── view.html                   # Foydalanuvchini ko'rish
│       ├── new.html                    # Yangi foydalanuvchi
│       └── edit.html                   # Tahrirlash
│
├── storage/                            # Ma'lumotlar bazasi qatlami
│   ├── storage.go                      # Storage interface
│   ├── repo/
│   │   └── repo.go                     # Repository interfaces
│   └── postgres/
│       ├── postgres.go                 # PostgreSQL ulanish
│       └── user_cruds.go               # User CRUD operations
│
└── migrations/                         # Database migrations
    ├── 000001_user_tables.up.sql      # Users jadvali yaratish
    └── 000001_user_tables.down.sql    # Users jadvali o'chirish
```

## 📋 Talablar

### Lokal ishga tushirish uchun:
- Go 1.24+
- PostgreSQL 13+

### Docker bilan ishga tushirish uchun:
- Docker
- Docker Compose

## 🛠 O'rnatish va Ishga Tushirish

### 1-usul: Docker bilan ishga tushirish (Tavsiya etiladi)

#### 1. Loyihani yuklab oling:
```bash
git clone <repository-url>
cd client_manager
```

#### 2. Environment variables o'rnating:
`.env` faylini tekshiring va kerak bo'lsa o'zgartiring:
```env
# PostgreSQL
PDB_NAME=client_manager
PDB_PORT=5432
PDB_PASSWORD=12345678
PDB_USER=postgres
PDB_HOST=localhost

# Server
SERVER_PORT=:8080
```

#### 3. Docker Compose bilan ishga tushiring:
```bash
docker-compose up --build
```

Bu buyruq quyidagi tartibda ishlaydi:
1. PostgreSQL konteynerini ishga tushiradi
2. Database tayyorlashni kutadi (health check)
3. Migration servisini ishga tushiradi (tabllar yaratadi)
4. Migration tugaganini kutadi
5. Go ilovasini build qiladi va ishga tushiradi
6. Loyihani `http://localhost:8080` da ochadi

#### 4. Faqat rebuild qilish uchun:
```bash
docker-compose up --build app
```

#### 5. Konteynerlarni to'xtatish:
```bash
docker-compose down
```

#### 6. Ma'lumotlar bilan birga to'xtatish:
```bash
docker-compose down -v
```

### 2-usul: Lokal development

#### 1. Loyihani yuklab oling:
```bash
git clone <repository-url>
cd client_manager
```

#### 2. Dependencies yuklab oling:
```bash
go mod tidy
```

#### 3. PostgreSQL o'rnating va database yarating:
```bash
# MacOS
brew install postgresql
brew services start postgresql
psql -U postgres -c "CREATE DATABASE client_manager;"

# Ubuntu/Debian
sudo apt install postgresql postgresql-contrib
sudo systemctl start postgresql
sudo -u postgres psql -c "CREATE DATABASE client_manager;"
```

#### 4. Migration bajaring:
```bash
# Docker migrate bilan
make migrate-up

# Yoki migrate tool o'rnatib
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
migrate -path ./migrations -database "postgres://postgres:12345678@localhost:5432/client_manager?sslmode=disable" up
```

#### 5. Ilovani ishga tushiring:
```bash
go run main.go
```

Brauzeringizda `http://localhost:8080` ni oching.

## 🌐 Foydalanish

### Asosiy sahifalar:

1. **Bosh sahifa**: `http://localhost:8080` yoki `http://localhost:8080/users`
   - Barcha foydalanuvchilar ro'yxati
   - Yangi foydalanuvchi qo'shish tugmasi

2. **Yangi foydalanuvchi qo'shish**: `http://localhost:8080/user/new`
   - MSISDN (telefon raqam) - majburiy
   - Ism - majburiy
   - Holati (faol/nofaol) - ixtiyoriy

3. **Foydalanuvchini ko'rish**: `http://localhost:8080/user/view?id=USER_ID`
   - Foydalanuvchi ma'lumotlarini batafsil ko'rish

4. **Foydalanuvchini tahrirlash**: `http://localhost:8080/user/edit?id=USER_ID`
   - Mavjud ma'lumotlarni o'zgartirish

5. **Foydalanuvchini o'chirish**: `http://localhost:8080/user/delete?id=USER_ID`
   - Foydalanuvchini butunlay o'chirish (tasdiqlash bilan)

### Interfeys xususiyatlari:
- **Responsive design** - mobil qurilmalarda ham ishlaydi
- **Bulma CSS Framework** - zamonaviy dizayn
- **Font Awesome icons** - chiroyli ikonkalar
- **O'zbek tilida** - barcha matnlar o'zbek tilida

## 🗄 Database Schema

```sql
CREATE TABLE users (
    user_id SERIAL PRIMARY KEY,
    msisdn VARCHAR(15) NOT NULL UNIQUE,
    name VARCHAR(100) NOT NULL,
    is_active BOOLEAN DEFAULT true,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

## 🔧 Konfiguratsiya

### Environment Variables:

| Variable | Tavsif | Default |
|----------|--------|---------|
| `PDB_HOST` | PostgreSQL server manzili | `localhost` |
| `PDB_PORT` | PostgreSQL porti | `5432` |
| `PDB_NAME` | Database nomi | `client_manager` |
| `PDB_USER` | Database foydalanuvchisi | `postgres` |
| `PDB_PASSWORD` | Database paroli | `12345678` |
| `SERVER_PORT` | Server porti | `:8080` |

## 📝 Development

### Migration boshqaruvi:

```bash
# Migration yaratish
make migrate-create NAME=add_users_table

# Migration up
make migrate-up
docker-compose run --rm migrate

# Migration down
make migrate-down

# Migration version ko'rish
make migrate-version

# Migration force (xatolik bo'lsa)
make migrate-force VERSION=1
```

### Yangi route qo'shish:
1. `internal/delivery/http_cms/handler.go` da route qo'shing
2. `internal/delivery/http_cms/users.go` da handler funksiya yozing
3. Kerak bo'lsa yangi template yarating

### Template o'zgartirish:
- `internal/delivery/http_cms/templates/` papkasidagi HTML fayllarni tahrirlang
- Barcha templatelar `base.html` dan foydalanadi

## 🐛 Troubleshooting

### 1. "Sahifa topilmadi" xatoligi:
```bash
# Serverni to'g'ri portda ishga tushirganingizni tekshiring
curl http://localhost:8080
```

### 2. Database ulanish xatoligi:
```bash
# PostgreSQL ishlab turganini tekshiring
ps aux | grep postgres

# Database mavjudligini tekshiring
psql -U postgres -l | grep client_manager
```

### 3. Template xatoliklari:
```bash
# Template fayllari mavjudligini tekshiring
ls -la internal/delivery/http_cms/templates/
```

### 4. Docker xatoliklari:
```bash
# Konteyner loglarini ko'ring
docker-compose logs app
docker-compose logs postgres

# Konteynerlarni qayta ishga tushiring
docker-compose restart
```

## 🚀 Production ga deploy qilish

### 1. Environment variables o'rnating:
- Real database credentials
- SSL konfiguratsiya
- Production server porti

### 2. Dockerfile optimizatsiya qiling:
- Multi-stage build (allaqachon mavjud)
- Health check qo'shing
- Security best practices

### 3. Nginx yoki reverse proxy qo'shing:
```nginx
server {
    listen 80;
    server_name your-domain.com;

    location / {
        proxy_pass http://localhost:8080;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
    }
}
```

## 📞 Support

Muammolar yoki takliflar uchun GitHub Issues ochishingiz mumkin yoki loyiha maintainer bilan bog'lanishingiz mumkin.

## 📄 License

Bu loyiha MIT litsenziyasi ostida tarqatiladi.
