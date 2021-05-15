CREATE DATABASE IF NOT EXISTS onlineshopdb

CREATE TABLE onlineshopdb.`product` (
  `sku` varchar(50) NOT NULL,
  `name` varchar(255) NOT NULL,
  `price` decimal(17,2) unsigned NOT NULL,
  `qty` decimal unsigned NOT NULL DEFAULT(0),
  `created_at` datetime NOT NULL DEFAULT current_timestamp(),
  `created_by` varchar(255) NOT NULL,
  `updated_at` datetime DEFAULT NULL ON UPDATE current_timestamp(),
  `updated_by` varchar(255) NULL,
  `deleted_at` datetime NULL,
  `deleted_by` varchar(255) NULL,
  `status` smallint(1) NOT NULL DEFAULT 1,
  PRIMARY KEY (`sku`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8

insert into onlineshopdb.`product` (sku, name, price, qty, created_by)
values ('120P90', 'Google Home', 49.99, 10, 'admin'),
('43N23P', 'MacBook Pro', 5399.99, 5, 'admin'),
('A304SD', 'Alexa Speaker', 109.50, 10, 'admin'),
('234234', 'Raspberry Pi B', 30.00, 5, 'admin')

CREATE TABLE onlineshopdb.`promotion` (
  `id` int(8) unsigned NOT NULL AUTO_INCREMENT,
  `apply_to_sku` varchar(50) NOT NULL,
  `min_required_qty` decimal NOT NULL,
  `description` varchar(255) NULL,
  `free_item_sku` varchar(50) NULL,
  `new_price_qty` decimal unsigned NULL,
  `discount_in_percent` decimal unsigned NULL,
  `created_at` datetime NOT NULL DEFAULT current_timestamp(),
  `created_by` varchar(255) NOT NULL,
  `updated_at` datetime DEFAULT NULL ON UPDATE current_timestamp(),
  `updated_by` varchar(255) NULL,
  `deleted_at` datetime NULL,
  `deleted_by` varchar(255) NULL,
  `status` smallint(1) NOT NULL DEFAULT 1,
  PRIMARY KEY (`id`),
  FOREIGN KEY (`apply_to_sku`) REFERENCES product(`sku`),
  FOREIGN KEY (`free_item_sku`) REFERENCES product(`sku`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8

insert into onlineshopdb.`promotion` (apply_to_sku, min_required_qty, description, free_item_sku, new_price_qty, discount_in_percent, created_by)
values ('43N23P', 1, 'Each sale of a MacBook Pro comes with a free Raspberry Pi B', '234234', NULL, NULL, 'admin'),
('120P90', 3, 'Buy 3 Google Homes for the price of 2', NULL, 2, NULL, 'admin'),
('A304SD', 4, 'Buy more than 3 Alexa speakers will have a 10% discount on all Alexa speakers', NULL, NULL, 10, 'admin')
