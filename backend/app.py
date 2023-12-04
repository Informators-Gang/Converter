from aiohttp import web
from .views import handle_upload


app = web.Application()
app.router.add_post('/upload', handle_upload)


if __name__ == "__main__":
    web.run_app(app, port=8080)
