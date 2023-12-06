from views import convert_file
import aiohttp_cors
from aiohttp import web

app = web.Application()
app.router.add_post('/convert', convert_file)

cors = aiohttp_cors.setup(app, defaults={
    "*": aiohttp_cors.ResourceOptions(
        allow_credentials=True,
        expose_headers="*",
        allow_headers="*"
    )
})

for route in list(app.router.routes()):
    cors.add(route)

if __name__ == "__main__":
    web.run_app(app, port=8080)
