ALTER TABLE driver RENAME TO drivers;

ALTER INDEX idx_driver_user_id RENAME TO idx_drivers_user_id;
ALTER INDEX idx_driver_sppg_id RENAME TO idx_drivers_sppg_id;