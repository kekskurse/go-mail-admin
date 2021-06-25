package main

import "github.com/rs/zerolog/log"

func resetDBForTest() {

	log.Warn().Msgf("Reset Database")
	db.Exec("DROP TABLE accounts;")
	db.Exec("DROP TABLE aliases;")
	db.Exec("DROP TABLE tlspolicies;")
	db.Exec("DROP TABLE domains;")
	db.Exec("CREATE TABLE `domains` (\n                           `id` int unsigned NOT NULL AUTO_INCREMENT,\n                           `domain` varchar(255) NOT NULL,\n                           PRIMARY KEY (`id`),\n                           UNIQUE KEY (`domain`)\n);")
	db.Exec("CREATE TABLE `accounts` (\n                            `id` int unsigned NOT NULL AUTO_INCREMENT,\n                            `username` varchar(64) NOT NULL,\n                            `domain` varchar(255) NOT NULL,\n                            `password` varchar(255) NOT NULL,\n                            `quota` int unsigned DEFAULT '0',\n                            `enabled` boolean DEFAULT '0',\n                            `sendonly` boolean DEFAULT '0',\n                            PRIMARY KEY (id),\n                            UNIQUE KEY (`username`, `domain`),\n                            FOREIGN KEY (`domain`) REFERENCES `domains` (`domain`)\n);")
	db.Exec("CREATE TABLE `aliases` (\n                           `id` int unsigned NOT NULL AUTO_INCREMENT,\n                           `source_username` varchar(64),\n                           `source_domain` varchar(255) NOT NULL,\n                           `destination_username` varchar(64) NOT NULL,\n                           `destination_domain` varchar(255) NOT NULL,\n                           `enabled` boolean DEFAULT '0',\n                           PRIMARY KEY (`id`),\n                           UNIQUE KEY (`source_username`, `source_domain`, `destination_username`, `destination_domain`),\n                           FOREIGN KEY (`source_domain`) REFERENCES `domains` (`domain`)\n);\n")
	db.Exec("CREATE TABLE `tlspolicies` (\n                               `id` int unsigned NOT NULL AUTO_INCREMENT,\n                               `domain` varchar(255) NOT NULL,\n                               `policy` enum('none', 'may', 'encrypt', 'dane', 'dane-only', 'fingerprint', 'verify', 'secure') NOT NULL,\n                               `params` varchar(255),\n                               PRIMARY KEY (`id`),\n                               UNIQUE KEY (`domain`)\n);\n")
}