
# Golang, Templ, Tailwind, and HTMX Starter Repository

This repository is a starter template for building modern web applications using **Golang**, **Templ**, **Tailwind CSS**, and **HTMX**. It is designed to provide a clean and scalable structure to kickstart your development.

## Features

- **Golang**: Backend logic and routing using [Echo](https://echo.labstack.com/) (or your preferred framework).
- **Templ**: Efficient and type-safe HTML templating.
- **Tailwind CSS**: Utility-first CSS framework for rapid UI development.
- **HTMX**: Lightweight framework for modern front-end interactivity without heavy JavaScript.

## Prerequisites

Ensure you have the following installed:

- [Go](https://golang.org/dl/) (version 1.19 or higher)
- [Node.js](https://nodejs.org/) and npm (for Tailwind CSS)
- A code editor like [VS Code](https://code.visualstudio.com/)

---

## Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/AdamZeman/gotth-starter.git
cd gotth-starter
```

### 2. Install Dependencies

#### Install Tailwind CSS dependencies:

```bash
npm install
```

```bash
npx tailwindcss init
```

- Configure the `tailwind.config.js` file
- Add the necessary filetypes as shown `.templ` example

```javascript
/** @type {import('tailwindcss').Config} */
module.exports = {
    content: [
        './**/*.templ',
        './**/*.go',
    ],
    theme: {
        extend: {
            colors: {
            },
        },
    },
    plugins: [],
};
```

### 3. Run the Development Server

#### Start the Go backend:

```bash
go run main.go
```

#### Start Tailwind CSS in watch mode:

```bash
npx tailwindcss -i ./css/input.css -o ./css/output.css --watch
```

### 4. Access the Application

Open your browser and navigate to `http://localhost:4000`.

### 5. Air

- install
```bash
go install github.com/air-verse/air@latest
```

- Create .air.toml config file
```bash
air init
```

- Config the .toml file
- `exclude_regex = ["_test.go", "._templ.go"]`
- `include_ext = ["go", "tpl", "tmpl", "html", "templ"]`

#### Start

```bash
air
```

Open your browser and navigate to `http://localhost:4000`.

## Project Structure
```
golang-templ-tailwind-htmx-starter/
├── css/                   # Tailwind input/output CSS files
│   ├── input.css          # Tailwind input file
│   ├── output.css         # Compiled CSS (ignored by Git)
├── app/                   # Main App folder
│   ├── handlers/          # Handlers + utils (.go files)
│   ├── views/             # Example handler for the home page
│   │   ├── components/    # Components for the layout and pages (.templ files)
│   │   ├── layouts/       # Layouts (.templ files)
│   │   ├── pages/         # Pages (.templ files)
├── go.mod                 # Go packages
├── main.go                # Main application entry point
├── package.json           # Tailwind configuration
├── .gitignore             # Git ignore file
├── README.md              # Project documentation
```

## Customization

1. **Templating**:
   - Add new `.templ` files in the `templates/` directory.
   - Update your handlers to use the new templates.

2. **Tailwind CSS**:
   - Customize your styles in the `input.css` file.
   - Extend the configuration in `tailwind.config.js`.

3. **Routing**:
   - Add new routes in `handlers/utils.go` and corresponding handlers in the `handlers/` folder.

## Scripts

- **Start Tailwind in Watch Mode**:
  ```bash
  npm run dev
  ```
- **Build Tailwind for Production**:
  ```bash
  npx tailwindcss -i ./css/input.css -o ./css/output.css --minify
  npx tailwindcss -i css/input.css -o css/output.css --watch
  ```

## Database
### Migrations

```bash
go get -u github.com/golang-migrate/migrate/v4
go get -u github.com/golang-migrate/migrate/v4/database/postgres

migrate -path ./db/migrations -database "postgres://localhost:5432/ecomm?sslmode=disable" up
```

## Contributing

Contributions are welcome! Feel free to open an issue or submit a pull request for any improvements or bug fixes.


## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Acknowledgments

- [Golang](https://golang.org/)
- [Templ](https://templ.dev/)
- [Tailwind CSS](https://tailwindcss.com/)
- [HTMX](https://htmx.org/)

---
