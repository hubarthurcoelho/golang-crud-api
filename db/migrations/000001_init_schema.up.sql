CREATE TABLE suppliers (
	token CHAR(36) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL,
    `name` VARCHAR(64) NOT NULL,
    `type` ENUM('COMPANY_NEGOTIATED', 'CONSOLIDATOR'),
	created_at DATETIME NOT NULL,
	updated_at DATETIME NOT NULL,
    PRIMARY KEY (token)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;
CREATE TABLE credentials (
	token CHAR(36) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL,
	`code` CHAR(36) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL,
    `active` tinyint(1) NOT NULL DEFAULT 1,
    agency ENUM('SMARTRIPS'),
    `description` VARCHAR(256) DEFAULT NULL,
    created_at DATETIME NOT NULL,
	updated_at DATETIME NOT NULL,
    PRIMARY KEY (token)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;
CREATE TABLE credential_suppliers (
	token CHAR(36) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL,
    credential_token CHAR(36) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL,
    supplier_token CHAR(36) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL,
    supplier_auth_credentials JSON NOT NULL,
	created_at DATETIME NOT NULL,
	updated_at DATETIME NOT NULL,
    PRIMARY KEY (token),
    KEY credential_token (credential_token),
    CONSTRAINT credential_token FOREIGN KEY (credential_token) REFERENCES credentials (token) ON DELETE CASCADE ON UPDATE CASCADE,
	KEY supplier_token (supplier_token),
    CONSTRAINT supplier_token_cred FOREIGN KEY (supplier_token) REFERENCES suppliers (token) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;
CREATE TABLE base_hotels (
	token CHAR(36) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL,
    interpec_id INT NOT NULL,
    city_id INT NOT NULL,
    `chain` VARCHAR(255) DEFAULT NULL,
    `name` VARCHAR(256) NOT NULL,
    geohash VARCHAR(8) DEFAULT NULL,
    zipcode VARCHAR(255) DEFAULT NULL,
    country_code VARCHAR(7) DEFAULT NULL,
	address VARCHAR(255) DEFAULT NULL,
	phone VARCHAR(31) DEFAULT NULL,
	stars SMALLINT DEFAULT '0',
    hotel_main_picture VARCHAR(256) NOT NULL,
    latitude DECIMAL(17,15) DEFAULT NULL,
    longitude DECIMAL(18,15) DEFAULT NULL,
    PRIMARY KEY (token),
    KEY geohash (geohash)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;
CREATE TABLE negotiated_hotels (
	token CHAR(36) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL,
    supplier_token CHAR(36) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL,
    base_hotel_token CHAR(36) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL,
    rooms JSON NOT NULL,
    hotel_pictures JSON NOT NULL,
    `description` TEXT,
    email varchar(127) NOT NULL,
    amenities JSON DEFAULT NULL,
	`active` TINYINT(1) NOT NULL DEFAULT '1',
	created_at DATETIME NOT NULL,
	updated_at DATETIME NOT NULL,
    PRIMARY KEY (token),
    KEY supplier_token (supplier_token),
    CONSTRAINT supplier_token_hot FOREIGN KEY (supplier_token) REFERENCES suppliers (token) ON DELETE CASCADE ON UPDATE CASCADE,
    KEY base_hotel_token (base_hotel_token),
    CONSTRAINT base_hotel_token FOREIGN KEY (base_hotel_token) REFERENCES base_hotels (token) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;