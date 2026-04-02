# 🚍 Fleet Management System (Backend Engineer Test)

## 📌 Deskripsi

Aplikasi ini merupakan sistem backend untuk manajemen armada kendaraan berbasis **microservices**.

Sistem ini terdiri dari:

* **Publisher** → mengirim data lokasi kendaraan setiap 2 detik
* **MQTT (Mosquitto)** → message broker untuk ingest data
* **Subscriber** → validasi + geofence check
* **RabbitMQ** → event broker
* **Worker** → consume event dan simpan ke database
* **API (Gin)** → menyediakan endpoint untuk membaca data

---

## 🛠️ Teknologi

* Golang (Gin Framework)
* MQTT (Eclipse Mosquitto)
* RabbitMQ
* PostgreSQL
* Docker & Docker Compose

---

## 📁 Struktur Project

```
fleet-system/
├── go-app        # REST API (Gin)
├── go-subs       # MQTT Subscriber + Geofence + Publisher RabbitMQ
├── go-worker     # RabbitMQ Consumer → PostgreSQL
├── go-publish    # MQTT Publisher (simulator)
├── docker-compose.yml
├── mosquitto.conf
```

---

## Cara Menjalankan Aplikasi

### 1. Jalankan Docker

```bash
docker-compose up --build
```

---

### 2. Tunggu semua service berjalan

Cek:

```bash
docker ps
```

Pastikan semua status:

```
healthy ✅
```

---

## 🌐 Akses Service

| Service     | URL                    |
| ----------- | ---------------------- |
| API         | http://localhost:8080  |
| RabbitMQ UI | http://localhost:15672 |
| MQTT        | tcp://localhost:1883   |

---

## 🔑 RabbitMQ Login

```
username: guest
password: guest
```

---

## 🧪 Testing

---

### 1. Health Check

#### API

```http
GET http://localhost:8080/health
```

#### Subscriber / Worker

```http
GET http://localhost:8081/health
```

---

### 2. Data Otomatis (Publisher)

Publisher akan otomatis mengirim data setiap **2 detik** ke MQTT:

```json
{
  "vehicle_id": "B1234XYZ",
  "latitude": -6.2088,
  "longitude": 106.8456,
  "timestamp": 1715003456
}
```

---

### 3. Cek Database

Data akan otomatis masuk ke PostgreSQL melalui Worker.

---

### 4. Test API

#### 🔹 Get Last Location

```http
GET http://localhost:8080/vehicles/B1234XYZ/location
```

---

#### 🔹 Get History

```http
GET http://localhost:8080/vehicles/B1234XYZ/history?start=1715000000&end=1715009999
```

---

## 📡 Geofence

Jika kendaraan masuk radius **50 meter** dari titik:

```
Lat: -6.2088
Lon: 106.8456
```

Maka event akan dikirim ke RabbitMQ:

```json
{
  "vehicle_id": "B1234XYZ",
  "event": "geofence_entry",
  "location": {
    "latitude": -6.2088,
    "longitude": 106.8456
  },
  "timestamp": 1715003456
}
```

---

## 🐳 Docker Services

* postgres
* mqtt (mosquitto)
* rabbitmq
* app (API)
* subs (subscriber)
* worker
* publish
