# Transjakarta Fleet Management System

A backend fleet management service for Transjakarta to receive GPS telemetry via MQTT, store location logs into PostgreSQL, trigger geofence entry events to RabbitMQ, and provide REST APIs for real-time tracking.

---

## Getting Started

Ensure **Docker** and **Docker Desktop** are installed and running on your system.

### 1. Clone Repository
```bash
git clone https://github.com/xiaojam/transjakarta.git
cd transjakarta
```

### 2. Run with Docker Compose
Run the following command in the project root directory to automatically build and launch all containers (PostgreSQL, Eclipse Mosquitto, RabbitMQ, Backend APP, and Mock Publisher):

```bash
docker compose up --build -d
```

### 3. Stream Live Logs
To monitor real-time logs from all active services:

```bash
docker compose logs -f
```

---

## REST API Endpoints

Base URL: `http://localhost:8080`

### 1. Get Latest Vehicle Location
Retrieves the most recent GPS telemetry record for a specified vehicle.

- **Endpoint:** `GET /vehicles/:vehicle_id/location`
- **Example:** `http://localhost:8080/vehicles/B1234XYZ/location`

### 2. Get Vehicle Location History
Retrieves historical location records for a vehicle within a specific UNIX timestamp range.

- **Endpoint:** `GET /vehicles/:vehicle_id/history?start={start_timestamp}&end={end_timestamp}`
- **Example:** `http://localhost:8080/vehicles/B1234XYZ/history?start=1700000000&end=2000000000`

---

## Stopping the Application

To temporarily pause all running containerswithout destroying data:

```bash
docker compose stop
```

To shut down and remove all containers, networks, and persistent data volumes:

```bash
docker compose down -v
```