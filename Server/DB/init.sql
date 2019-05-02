CREATE TABLE parking_lots
(
    id INTEGER PRIMARY KEY,
    park_house_id INTEGER REFERENCES park_house(id),
    INTEGER dev_id NOT NULL,
    INTEGER sensor_id NOT NULL,
    UNIQUE (dev_id, sensor_id)
);
    CREATE UNIQUE INDEX ON devices(real_Id);
    CREATE UNIQUE INDEX ON devices(dev_id, sensor_id);

CREATE TABLE raw_data (
  SERAL id PRIMARY KEY,
  INTEGER dev_id NOT NULL,
  INTEGER sensor_id NOT NULL,
  TIMESTAMP ts NOT NULL,
  BOOLEAN is_empty NOT NULL,
  FOREIGN KEY (dev_id, sensor_id) REFERENCES devices (dev_id, sensor_id)
);

CREATE TABLE actual_state (
  INTEGER real_id PRIMARY KEY ,
  BOOLEAN is_empty NOT NULL,
  BOOLEAN is_locked DEFAULT FALSE,
  TIMESTAMP ts,
  FOREIGN KEY (real_id) REFERENCES devices(real_id)
);
CREATE UNIQUE INDEX ON actual_state (real_Id);

CREATE FUNCTION update_row () RETURNS TRIGGER LANGUAGE plpgsql
SECURITY DEFINER AS $$
BEGIN
    UPDATE actual_state
    SET state   = NEW.state,
        ts      = NEW.ts,
        locked  = FALSE
    WHERE real_Id IN (SELECT real_id FROM devices
                      WHERE NEW.dev_id = devices.dev_id AND NEW.sensor_id = devices.sensor_id)
          AND (ts ISNULL  OR ts < _ts);
    RETURN TRUE;
END;
$$;

CREATE TRIGGER update_state AFTER INSERT ON raw_data
      FOR EACH ROW
      EXECUTE PROCEDURE update_row();

CREATE TABLE park_house(
  SERIAL id PRIMARY KEY,
  VARCHAR name,
  DOUBLE  latitude,
  DOUBLE  longitude
);

CREATE TABLE park_house_entry(
  SERIAL id,
  INTEGER house_id REFERENCES  park_house(id),
  DOUBLE x_coords,
  DOUBLE y_coords,
  DOUBLE width,
  DOUBLE direction,
  INTEGER level
);

CREATE TABLE local_data(
  id INTEGER PRIMARY KEY REFERENCES parking_lots(id),
  DOUBLE x_coords,
  DOUBLE y_coords,
  DOUBLE width,
  DOUBLE length,
  DOUBLE direction,
  INTEGER level
);

