--Get nearest parkhouse
SELECT * from park_house
ORDER BY (y_coords-($1))*(y_coords-($1))+(x_coords-($2))*(x_coords-($2))
LIMIT 1;
--Get nearest lot from entry
SELECT * from park_house_entry
WHERE level = ($3) AND house_id = ($4)
ORDER BY (y_coords-($1))*(y_coords-($1))+(x_coords-($2))*(x_coords-($2))
LIMIT 1;
--Get entries
SELECT * from park_house_entry
WHERE level = ($1) AND house_id = ($2);
--Get lots on level
SELECT * from park_house_entry
WHERE level = ($1) AND house_id = ($2);
--Get lots on level with state
