-- Adminer 4.8.1 MySQL 11.5.2-MariaDB-ubu2404 dump

SET NAMES utf8;
SET time_zone = '+00:00';
SET foreign_key_checks = 0;

SET NAMES utf8mb4;

CREATE DATABASE `BTS` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci */;
USE `BTS`;

CREATE TABLE `Checklists` (
  `ChecklistId` bigint(20) NOT NULL AUTO_INCREMENT,
  `UserId` bigint(20) DEFAULT NULL,
  `Title` varchar(255) DEFAULT NULL,
  `Description` text DEFAULT NULL,
  `IsDeleted` tinyint(4) DEFAULT NULL COMMENT '1. True, 0. false',
  `CreatedAt` datetime DEFAULT NULL,
  `UpdatedAt` datetime DEFAULT NULL,
  PRIMARY KEY (`ChecklistId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;


CREATE TABLE `Checklist_Items` (
  `ItemId` bigint(20) NOT NULL AUTO_INCREMENT,
  `ChecklistId` bigint(20) DEFAULT NULL,
  `Description` text DEFAULT NULL,
  `Status` tinyint(4) DEFAULT NULL COMMENT '1. Checked, 0. Unchecked',
  `IsDeleted` tinyint(4) DEFAULT NULL COMMENT '1. True, 0. False',
  `CreatedAt` datetime DEFAULT NULL,
  `UpdatedAt` datetime DEFAULT NULL,
  PRIMARY KEY (`ItemId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;


CREATE TABLE `Users` (
  `UserId` bigint(20) NOT NULL AUTO_INCREMENT,
  `Name` varchar(255) DEFAULT NULL,
  `Email` varchar(100) DEFAULT NULL,
  `Password` varchar(50) DEFAULT NULL,
  `Status` tinyint(4) DEFAULT NULL COMMENT '0. Logout, 1. Login',
  `CreatedAt` datetime DEFAULT NULL,
  `UpdatedAt` datetime DEFAULT NULL,
  PRIMARY KEY (`UserId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;


-- 2024-12-11 09:15:49
