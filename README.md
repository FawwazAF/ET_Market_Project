<a href="https://github.com/FawwazAF/ET_Market_Project"><img height="70" src="https://image.flaticon.com/icons/png/512/862/862819.png"></a>


# ET Market
API Electronic Traditional Market 

[![Go.Dev reference](https://img.shields.io/badge/gorm-reference-blue?logo=go&logoColor=blue)](https://pkg.go.dev/gorm.io/gorm?tab=doc)
[![Go.Dev reference](https://img.shields.io/badge/echo-reference-blue?logo=go&logoColor=blue)](https://github.com/labstack/echo)
[![Codecov](https://img.shields.io/badge/coverage-80-blue?)](https://github.com/FawwazAF/ET_Market_Project)

# Table of Content

- [Introduction](#introduction)
- [Feature Overview](#feature-overview)
- [High Level Architecture](#high-level-architecture)
- [Flowchart](#flowchart)
- [Entity Relations Diagram](#entity-relations-diagram)
- [How to use](#how-to-use)
- [API usage examples](#api-usage-examples)
- [Contribute](#contribute)
- [Credits](#credits)
- [Image source](#image-source)

# Introduction
This project is an RESTFul API that is intended to be used in mobile aplication, let's say like Grab Market or Gojek. This project is written in Golang with API framework of Echo an ORM with Gorm, Dgri-jalva for jwt auth and MySQL for the database. 

# Feature Overview

There are 3 main user in this project that consume the API, the Seller, Consumer, and Driver. In this image below is a use case for every user:

<img src="https://i.imgur.com/YXi6A6X.png">

# High Level Architecture
Here is the High Level Architecture.

<img src="https://i.imgur.com/ln6nvt3.jpg">

# Flowchart
This is the flow of the application that we build.

<img src="https://i.imgur.com/z2MPGUv.png">

# Entity Relations Diagram
This is the database entity that we build.

<img src="https://i.imgur.com/LvFLLnz.jpg">

# How to use
Make sure to install Go and MySQL in order to use this API, you can use Google Cloud Platform or AWS to deploy this API.

Clone this repository in your $PATH:
```
$ git clone https://github.com/FawwazAF/ET_Market_Project.git
```
Run the main.go.
```
$ go run main.go
```
# API usage examples
You can go to this website to test the api:

https://ihsan-null.github.io/ETMarket_OpenAPI/


## Contribute

**Use issues for everything**

- For a small change, just send a PR.
- For bigger changes open an issue for discussion before sending a PR.
- PR should have:
  - Test case
  - Documentation
  - Example (If it makes sense)
- You can also contribute by:
  - Reporting issues
  - Suggesting new features or enhancements
  - Improve/fix documentation

## Credits

- [Fawwaz Amjad Fuadi](https://github.com/FawwazAF) (Author and maintainer)
- [Riska Kurnia Dewi](https://github.com/riskakrndw) (Author and maintainer)
- [Patmiza Nopiani](https://github.com/Patmiza) (Author and maintainer)
- [Ihsanul Ilham Akbar](https://github.com/ihsan-null) (Author and maintainer)

## Image Source
- "https://www.flaticon.com/authors/mynamepong"
