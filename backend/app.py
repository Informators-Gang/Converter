from aiohttp import web
from views import convert_file


app = web.Application()
app.router.add_post('/convert', convert_file)


if __name__ == "__main__":
    web.run_app(app, port=8080)
