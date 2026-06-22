ALTER TABLE drivers RENAME TO driver;

ALTER INDEX idx_drivers_user_id RENAME TO idx_driver_user_id;
ALTER INDEX idx_drivers_sppg_id RENAME TO idx_driver_sppg_id;