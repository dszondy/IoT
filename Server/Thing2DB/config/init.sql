
CREATE TABLE lot_data(
id  SERIAL PRIMARY KEY,
device_id VARCHAR,
lot_id    VARCHAR,
UNIQUE (device_id, lot_id)
);

CREATE INDEX on lot_data(
device_id, lot_id
);

INSERT INTO lot_data (device_id, lot_id) VALUES
(1, 1),
(1, 2),
(1, 3);

CREATE TABLE raw_data (
id        SERIAL PRIMARY KEY,
device_id VARCHAR,
lot_id    VARCHAR,
lot_timestamp TIMESTAMP(3),
is_clear  BOOLEAN,
FOREIGN KEY (device_id, lot_id) REFERENCES lot_data (device_id, lot_id),
UNIQUE (  device_id, lot_id, lot_timestamp)
);

CREATE INDEX on raw_data(
device_id, lot_id, lot_timestamp
);

CREATE VIEW lot_states (id, state) AS (
SELECT
	lot_data.id,
	latest_state.is_clear
FROM lot_data,
--      latest_state
	(SELECT
		raw_data.device_id,
		raw_data.lot_id,
		raw_data.is_clear
	FROM raw_data,
--      latest_data
		(SELECT
			device_id,
			lot_id,
			MAX(lot_timestamp) as lot_timestamp
		FROM raw_data
		GROUP BY device_id, lot_id)
		AS latest_data
	WHERE(
		raw_data.device_id LIKE latest_data.device_id AND
		raw_data.lot_id LIKE latest_data.lot_id AND
		raw_data.lot_timestamp = latest_data.lot_timestamp)
	)AS latest_state
WHERE (
	latest_state.device_id LIKE lot_data.device_id AND
	latest_state.lot_id LIKE lot_data.lot_id)
)