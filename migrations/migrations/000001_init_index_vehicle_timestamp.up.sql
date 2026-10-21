CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE INDEX IF NOT EXISTS idx_vehicle_timestamp_desc ON vehicle_locations (vehicle_id, timestamp DESC);