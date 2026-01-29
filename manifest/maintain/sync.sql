TRUNCATE TABLE need_sync_table;
INSERT INTO need_sync_table
SELECT * FROM mysql('host:3306', 'schema', 'need_sync_table', 'user', 'password');
