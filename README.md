<a name="readme-top"></a>
<h1 align='center'>FilmWise</h1>


# 📗 Table of Contents

- [📗 Table of Contents](#-table-of-contents)
- [ FilmWise ](#-about-project-)
  - [🛠 Built With ](#-built-with-)
    - [Tech Stack ](#tech-stack-)
    - [Key Features ](#key-features-)
  - [💻 Getting Started ](#-getting-started-)
    - [Prerequisites](#prerequisites)
    - [Setup](#setup)
    - [Install](#install)
    - [Database](#database)
    - [Usage](#usage)
    - [Build](#build)
    - [Deployment](#deployment)
  - [👥 Authors ](#-authors-)
  - [🔭 Future Features ](#-future-features-)
  - [🤝 Contributing ](#-contributing-)
  - [⭐️ Show your support ](#️-show-your-support-)
  - [🙏 Acknowledgments ](#-acknowledgments-)
  - [📝 License ](#-license-)


# FilmWise <a name="about-project"></a>
FilmWise: Your Ultimate Movie Hub. Experience movies like never before with FilmWise, the web application that lets you dive into detailed movie information, share your thoughts through comments and reviews, and connect with fellow cinephiles. Discover, discuss, and decide on your favorite films all in one place. if you want to see the front-end part, you can click [here](https://github.com/raihan2bd/filmwise-front)

## 🛠 Built With <a name="built-with"></a>
### Tech Stack <a name="tech-stack"></a>

<details>
  <summary>Front End</summary>
  <ul>
    <li>Nextjs</li>
    <li>React</li>
    <li>Redux</li>
    <li>JAVASCRIPT</li>
    <li>Html</li>
    <li>CSS</li>
  </ul>
</details>
<details>
  <summary>Back End</summary>
  <ul>
    <li>Golang</li>
    <li>PostgreSQL</li>
  </ul>
</details>


<!-- LIVE DEMO -->

## 🚀 Live Demo <a name="live-demo"></a>

> Live demo will update soon!.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

### Key Features <a name="key-features"></a>

- Discover detailed movie information, from plots to cast.
- Engage in vibrant discussions by commenting on movies.
- Add your personal reviews to contribute to the community.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## 💻 Getting Started <a name="getting-started"></a> 

To get a local copy up and running, follow these steps.

### Prerequisites

In order to run this project you need:
- Then Make sure you have installed [Go (golang)](https://go.dev/dl/) version 1.20.4 or the latest stable version.
- Then make sure you have installed [PostgreSQL](https://www.postgresql.org/) on your local machine if you want to use this project locally.
- Then Create a database called `filmwise` inside the database and create tables using this `project>database` `schema` SQL query.

- First of all to see this project's graphical interface make sure you run the [front-end](https://github.com/raihan2bd/filmwise-front) part

### Setup

- Clone this repository to your desired folder:

```sh
  cd your-folder
  https://github.com/raihan2bd/filmwise.git
```

- Before running the project please make sure you create a `.env` file to your project root directory and add `DATABASE_URI`, and `JWT_SECRET_KEY` environment variables to the file. For example:
```
DATABASE_URI="host=localhost port=5432 dbname=filmwise user=postgres password=your password sslmode=disable"
JWT_SECRET="your jwt secret key"
CLD_URI="Secreat Key of cloudinary"
CLOUD_NAME="Name of cloudinary account"
```

### Getting JWT Secret Key

To obtain the JWT secret key, please <a href="https://go.dev/play/" target="_blank" style="color: blue; font-size: 12px; text-decoration: underline;">click here</a>.

An open Go terminal is required. In the terminal, please copy the following code and paste it into your <span style="color: lightblue;">.env</span> file.

Like this

JWT_SECRET="*******"

```

package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func main() {
	// Define the desired length of the secret key in bytes
	keyLength := 64 // Adjust the length as needed

	// Create a byte slice to hold the random bytes
	key := make([]byte, keyLength)

	// Generate random bytes using crypto/rand
	_, err := rand.Read(key)
	if err != nil {
		fmt.Println("Error generating random bytes:", err)
		return
	}

	// Encode the random bytes in base64 to create a string
	secretKey := base64.StdEncoding.EncodeToString(key)

	fmt.Println("Generated JWT secret key:", secretKey)
}

```

### Getting Cloudinary Serect key and Name

Cloudinary is a cloud-based media management platform that helps businesses and developers efficiently store, manage, and deliver images and videos for websites and applications. It provides features like image and video uploading, storage, transformation, optimization, and content delivery via a content delivery network (CDN), making it easier to handle media assets in web and mobile applications. Cloudinary's services can enhance website performance, user experience, and streamline media asset workflows.

For using this you need to <a href="https://cloudinary.com/users/register_free" target="_blank" style="hover:text-decoration: underline; font-size:1.1rem">Create an account</a> Or if you have an Account you need to <a href="https://cloudinary.com/users/login" target="_blank" style="hover:text-decoration: underline; font-size:1.2rem;">Sign In </a>

After that go into the <span style="color: #3448c5;">Dashborad</span>

Copy the <span style="color: #3448c5;">Cloud Name</span> and
<span style="color: #3448c5;">API Environment variable</span>

```
CLD_URI="cloudinary://******"

CLOUD_NAME="****"

```



### Install

Install this project with:

- Install the required gems with:

```sh
go mod tidy
```

### Database

- Create the databases properly, You need to open an SQL editor and run the `/database/schema.sql` file script. Make sure you run the script block by block.

### Usage

- To run the development server, execute the following command:

```sh
go run ./cmd/api/ .
```

### Build

- To build the project for production-ready run the following command:

```sh
go build -o main ./cmd/api/*.go
```




### Deployment

To deploy your project online You can visit [Render](https://www.render.com/)

<p align="right">(<a href="#readme-top">back to top</a>)</p>


## 👥 Author <a name="author"></a>

👤 **Abu Raihan**

- GitHub: [@raihan2bd](https://github.com/raihan2bd)
- Twitter: [@raihan2bd](https://twitter.com/raihan2bd)
- LinkedIn: [raihan2bd](https://linkedin.com/in/raihan2bd)
  
<p align="right">(<a href="#readme-top">back to top</a>)</p>


## 🔭 Future Features <a name="future-features"></a>

- [ ] **Improve user experience**

<p align="right">(<a href="#readme-top">back to top</a>)</p>


## 🤝 Contributing <a name="contributing"></a>

Contributions, issues, and feature requests are welcome!

Feel free to check the [issues page](https://github.com/raihan2bd/filmwise/issues).

<p align="right">(<a href="#readme-top">back to top</a>)</p>


## ⭐️ Show your support <a name="support"></a>

If you like this project, please leave a ⭐️

<p align="right">(<a href="#readme-top">back to top</a>)</p>


## 🙏 Acknowledgments <a name="acknowledgements"></a>

I would like to thank [Trevor Sawler](https://www.gocode.ca/) Who helped me a lot to learn Golang.

<p align="right">(<a href="#readme-top">back to top</a>)</p>


## 📝 License <a name="license"></a>

This project is [MIT](./LICENSE) licensed.

<p align="right">(<a href="#readme-top">back to top</a>)</p>
