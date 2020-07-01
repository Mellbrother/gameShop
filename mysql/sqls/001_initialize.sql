-- ==================================
-- Crean up old tables
-- ==================================
DROP TABLE IF EXISTS items;

-- ==================================
-- CREATE TABLES:
--   items
-- ==================================
CREATE TABLE IF NOT EXISTS `items` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(255) DEFAULT NULL,
  `price` VARCHAR(255) DEFAULT NULL, 
  `exhibitor_address` VARCHAR(255) NOT NULL,
  `description` TEXT DEFAULT NULL,
  `is_sold` TINYINT(1) DEFAULT NULL,
  `updated_at` DATETIME NULL DEFAULT CURRENT_TIMESTAMP,
  `created_at` DATETIME NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8
AUTO_INCREMENT=1
COMMENT = 'master table of host items';