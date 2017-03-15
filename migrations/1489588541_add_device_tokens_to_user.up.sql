ALTER TABLE users ADD COLUMN device_tokens TEXT [] NOT NULL DEFAULT array[]::TEXT[];
