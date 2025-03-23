package modules

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// SetupSwagger sets up the Swagger UI endpoint and serves the OpenAPI spec
func SetupSwagger(e *echo.Echo, spec string) {
	// Serve the Swagger UI index page
	e.GET("/swagger", func(c echo.Context) error {
		return c.HTML(http.StatusOK, swaggerIndexHTML)
	})

	// Serve the OpenAPI specification
	e.GET("/swagger/openapi.json", func(c echo.Context) error {
		return c.String(http.StatusOK, spec)
	})
}

// swaggerIndexHTML is the Swagger UI HTML page
const swaggerIndexHTML = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Comparative Market Analysis (CMA) API - Swagger UI</title>
    <link rel="stylesheet" type="text/css" href="https://unpkg.com/swagger-ui-dist@5.1.0/swagger-ui.css">
    <style>
        html { box-sizing: border-box; overflow: -moz-scrollbars-vertical; overflow-y: scroll; }
        *, *:before, *:after { box-sizing: inherit; }
        body { margin: 0; padding: 0; }
        .topbar { display: none; }
    </style>
</head>
<body>
    <div id="swagger-ui"></div>

    <script src="https://unpkg.com/swagger-ui-dist@5.1.0/swagger-ui-bundle.js"></script>
    <script src="https://unpkg.com/swagger-ui-dist@5.1.0/swagger-ui-standalone-preset.js"></script>
    <script>
    window.onload = function() {
        const ui = SwaggerUIBundle({
            url: "/swagger/openapi.json",
            dom_id: '#swagger-ui',
            deepLinking: true,
            presets: [
                SwaggerUIBundle.presets.apis,
                SwaggerUIStandalonePreset
            ],
            plugins: [
                SwaggerUIBundle.plugins.DownloadUrl
            ],
            layout: "StandaloneLayout"
        });
        window.ui = ui;
    };
    </script>
</body>
</html>
`
