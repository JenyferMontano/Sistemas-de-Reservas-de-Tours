-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema reservas
-- -----------------------------------------------------

-- -----------------------------------------------------
-- Schema reservas
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `reservas` DEFAULT CHARACTER SET utf8 ;
USE `reservas` ;

-- -----------------------------------------------------
-- Table `reservas`.`Tour`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `reservas`.`Tour` (
  `idTour` INT NOT NULL,
  `nombre` VARCHAR(45) NOT NULL,
  `descripcion` VARCHAR(45) NOT NULL,
  `tipo` VARCHAR(45) NOT NULL,
  `disponibilidad` TINYINT NOT NULL,
  `precioBase` DECIMAL NOT NULL,
  `ubicacion` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`idTour`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `reservas`.`Persona`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `reservas`.`Persona` (
  `idPersona` INT NOT NULL,
  `nombre` VARCHAR(45) NOT NULL,
  `apellido_1` VARCHAR(45) NOT NULL,
  `apellido_2` VARCHAR(45) NOT NULL,
  `fechaNac` DATE NOT NULL,
  `telefono` VARCHAR(15) NOT NULL,
  `correo` VARCHAR(20) NOT NULL,
  PRIMARY KEY (`idPersona`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `reservas`.`Usuario`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `reservas`.`Usuario` (
  `userName` VARCHAR(15) NOT NULL,
  `password` VARCHAR(30) NOT NULL,
  `rol` VARCHAR(20) NOT NULL,
  `idPersona` INT NOT NULL,
  PRIMARY KEY (`userName`),
  INDEX `fk_Usuario_Persona1_idx` (`idPersona` ASC),
  CONSTRAINT `fk_Usuario_Persona1`
    FOREIGN KEY (`idPersona`)
    REFERENCES `reservas`.`Persona` (`idPersona`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `reservas`.`Transfer`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `reservas`.`Transfer` (
  `idTransfer` INT NOT NULL,
  `tipo` VARCHAR(15) NOT NULL,
  `capacidad` INT NOT NULL,
  PRIMARY KEY (`idTransfer`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `reservas`.`Reserva`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `reservas`.`Reserva` (
  `numReserva` INT NOT NULL,
  `fechaReserva` DATE NOT NULL,
  `horaReserva` TIME NOT NULL,
  `cantidadPersonas` INT NOT NULL,
  `tour` INT NOT NULL,
  `usuario` VARCHAR(15) NOT NULL,
  `persona` INT NOT NULL,
  `transfer` INT NOT NULL,
  PRIMARY KEY (`numReserva`),
  INDEX `fk_Reserva_Tour_idx` (`tour` ASC),
  INDEX `fk_Reserva_Usuario1_idx` (`usuario` ASC),
  INDEX `fk_Reserva_Persona1_idx` (`persona` ASC),
  INDEX `fk_Reserva_Transfer1_idx` (`transfer` ASC),
  CONSTRAINT `fk_Reserva_Tour`
    FOREIGN KEY (`tour`)
    REFERENCES `reservas`.`Tour` (`idTour`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_Reserva_Usuario1`
    FOREIGN KEY (`usuario`)
    REFERENCES `reservas`.`Usuario` (`userName`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_Reserva_Persona1`
    FOREIGN KEY (`persona`)
    REFERENCES `reservas`.`Persona` (`idPersona`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_Reserva_Transfer1`
    FOREIGN KEY (`transfer`)
    REFERENCES `reservas`.`Transfer` (`idTransfer`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `reservas`.`Factura`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `reservas`.`Factura` (
  `idFactura` INT NOT NULL,
  `fechaFact` DATE NOT NULL,
  `metodoPago` VARCHAR(25) NOT NULL,
  `iva` DECIMAL NOT NULL,
  `descuento` DECIMAL NOT NULL,
  `subtotal` DECIMAL NOT NULL,
  `total` DECIMAL NOT NULL,
  `reserva` INT NOT NULL,
  PRIMARY KEY (`idFactura`),
  INDEX `fk_Factura_Reserva1_idx` (`reserva` ASC),
  CONSTRAINT `fk_Factura_Reserva1`
    FOREIGN KEY (`reserva`)
    REFERENCES `reservas`.`Reserva` (`numReserva`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;