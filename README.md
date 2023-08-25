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
    - [Testing](#testing)
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

- First of all to see this project's graphical interface make sure you run the [front-end](https://github.com/Thinus01/Resort_Booking_Front-end) part
- Before running this repo on your local machine make sure you have installed <a name="install"></a> [Ruby](https://www.ruby-lang.org/), [Rails](https://rubyonrails.org/), and [PostgreSQL](https://www.postgresql.org/)

### Setup

- Clone this repository to your desired folder:

```sh
  cd your-folder
  https://github.com/Thinus01/Resort_Booking_Back-end
```

- Before running the project please make sure you create a `.env` file to your project root directory and add `DATABASE_USER_NAME`, `DATABASE_PASSWORD`, and `JWT_SECRET_KEY` environment variables to the file. For example:
```
DATABASE_USER_NAME=place_your_postgres_database_username
DATABASE_PASSWORD=place_your_database_password

JWT_SECRET_KEY=place_your_jwt_secret_key
```
![back-end-env](https://github.com/Thinus01/Resort_Booking_Back-end/assets/35267447/3fd60673-53a2-4e6c-a386-ff894546c850)

- After that make sure you change cors. To change the cors open the project with your favorite editor then click config folder then click initializers folder and open cors.rb file and after that change the `origins` to your front-end URL to solve the cross-origins request error. For example:
```
origins 'http://localhost:3000'
```
![cors](https://github.com/Thinus01/Resort_Booking_Back-end/assets/35267447/44eb387a-c8d5-40c8-b910-860ed9bf0195)

- To generate the `JWT_SECRET_KEY` as a secure and random hash using Ruby's `SecureRandom module`. You can do this in a Ruby script or in your Rails console:
```ruby
require 'securerandom'

jwt_secret_key = SecureRandom.hex(64)
puts jwt_secret_key
```
This will generate a random string of 128 characters (64 bytes in hexadecimal representation) and print it to the console. Then you can copy it from the console and paste it into your .env file as a value of `JWT_SECRET_KEY`.


### Install

Install this project with:

- Install the required gems with:

```sh
bundle install
npm install
```

### Database

- Create the databases and run migrations with:

```sh
rails db:create
rails db:migrate
rails db:seed
```

### Usage

- To run the development server, execute the following command:

```sh
rails server
```

`Note:` For more information about this back-end API end-points run the server and visit `project_base_url/api-docs` For example: `http://localhost:4000/api-docs`.

### Testing

- To run tests, run the following command:

```sh
rspec spec 
```


### Deployment

You can visit the app using [Render](https://www.render.com/)

<p align="right">(<a href="#readme-top">back to top</a>)</p>


## 👥 Authors <a name="author"></a>

👤 **Abu Raihan**

- GitHub: [@raihan2bd](https://github.com/raihan2bd)
- Twitter: [@raihan2bd](https://twitter.com/raihan2bd)
- LinkedIn: [raihan2bd](https://linkedin.com/in/raihan2bd)

### 👤 **Amaka Konwea**:
- GitHub: [@lawrah_on_GitHub](https://github.com/lawrahkonwea)
- Twitter: [@lawrah_on_Twitter](https://twitter.com/lawrah_xo)
- LinkedIn: [@laura_on_LinkedIn](https://www.linkedin.com/in/amakalaurakonwea/)

### 👤 **Thinus Van De Venter**

- GitHub: [@Thinus01](https://github.com/Thinus01)
- Twitter: [@thinus_v_d_v](https://twitter.com/thinus_v_d_v)
- LinkedIn: [Thinus Van De Venter](https://www.linkedin.com/in/thinus-van-de-venter-99aa26203)

👤 **Winnie Edube**

- GitHub: [@edubew](https://github.com/edubew)
- Twitter: [@edube_winne](https://twitter.com/edube_winne)
- LinkedIn: [Winfred Edube](https://www.linkedin.com/in/winfred-edube/)

👤 **Michale Kithinji**
- GitHub: [@githubhandle](https://github.com/MICHAELKITH)
- Twitter: [@twitterhandle](https://twitter.com/DevMichael11)
- LinkedIn: [LinkedIn](https://www.linkedin.com/in/michaelkithinji/)

<p align="right">(<a href="#readme-top">back to top</a>)</p>


## 🔭 Future Features <a name="future-features"></a>

- [ ] **Improve user experience**

<p align="right">(<a href="#readme-top">back to top</a>)</p>


## 🤝 Contributing <a name="contributing"></a>

Contributions, issues, and feature requests are welcome!

Feel free to check the [issues page](https://github.com/lawrahkonwea/Rails_blog_app/issues).

<p align="right">(<a href="#readme-top">back to top</a>)</p>


## ⭐️ Show your support <a name="support"></a>

If you like this project, please leave a ⭐️

<p align="right">(<a href="#readme-top">back to top</a>)</p>


## 🙏 Acknowledgments <a name="acknowledgements"></a>

We want to give a big thanks to Microverse for giving us the chance to achieve this milestone
We also want to thank [Murat Korkmaz](https://www.behance.net/muratk) for his [design](https://www.behance.net/gallery/26425031/Vespa-Responsive-Redesign), that our project is based on.

<p align="right">(<a href="#readme-top">back to top</a>)</p>


## 📝 License <a name="license"></a>

This project is [MIT](https://github.com/Thinus01/Resort_Booking_Back-end/blob/dev/LICENSE) licensed.

<p align="right">(<a href="#readme-top">back to top</a>)</p>
