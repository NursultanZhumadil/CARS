AwesomeProject12 - Car Advertisement Site
AwesomeProject12 — бұл көлік жарнамаларын басқару үшін жасалған веб-қосымша. Бұл қосымша пайдаланушыларға көліктерді қосуға, көлік ақпараттарын көруге, жаңартуға және жоюға мүмкіндік береді.

Құралдар мен технологиялар
Go (Golang) - Бэкенд бағдарламалау тілі
Gin - Go үшін жеңіл және жоғары өнімді веб-фреймворк
GORM - Go үшін ORM (Object-Relational Mapping) кітапханасы
PostgreSQL - Дерекқор басқару жүйесі
Git - Нұсқаларды басқару жүйесі
Орнату
1. Go орнату
Егер сіздің жүйеңізде Go орнатылмаған болса, ресми сайттан Go тілін жүктеп, орнатыңыз.

2. Жоба репозиторийін клондау
Жобаны өз компьютеріңізге клондау үшін келесі команданы орындаңыз:

git clone https://github.com/yourusername/awesomeProject12.git cd awesomeProject12

Тәуелділіктерді орнату

go mod tidy

PostgreSQL дерекқорын орнату
CREATE DATABASE car_db; CREATE USER caradmin WITH PASSWORD 'sultan05'; ALTER ROLE caradmin SET client_encoding TO 'utf8'; ALTER ROLE caradmin SET timezone TO 'UTC'; GRANT ALL PRIVILEGES ON DATABASE car_db TO yourusername;

.env файлын жасау

DB_HOST=localhost DB_USER=postgres DB_PASSWORD=sultan05 DB_NAME=car_db DB_PORT=5432

Жобаны іске қосу
go run main.go

API-лер
[ { "id": 1, "brand": "Toyota", "model": "Camry", "year": 2020, "price": 25000, "mileage": 15000 }, { "id": 2, "brand": "BMW", "model": "X5", "year": 2018, "price": 35000, "mileage": 30000 } ]

GET /car/:id
GET /car/1

{
"brand": "Audi", "model": "A4", "year": 2021, "price": 45000, "mileage": 5000 }

POST /car
{ "brand": "Audi", "model": "A4", "year": 2021, "price": 45000, "mileage": 5000 }

PUT /car/:id
{ "brand": "Audi", "model": "A6", "year": 2022, "price": 55000, "mileage": 3000 }

DELETE /car/:id
